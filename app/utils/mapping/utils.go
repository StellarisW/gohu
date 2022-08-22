package mapping

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/zeromicro/go-zero/core/stringx"
)

const (
	defaultOption      = "default"
	stringOption       = "string"
	optionalOption     = "optional"
	optionsOption      = "options"
	rangeOption        = "range"
	optionSeparator    = "|"
	equalToken         = "="
	escapeChar         = '\\'
	leftBracket        = '('
	rightBracket       = ')'
	leftSquareBracket  = '['
	rightSquareBracket = ']'
	segmentSeparator   = ','
)

var (
	errUnsupportedType  = errors.New("unsupported type on setting field value")
	errNumberRange      = errors.New("wrong number range setting")
	optionsCache        = make(map[string]optionsCacheValue)
	cacheLock           sync.RWMutex
	structRequiredCache = make(map[reflect.Type]requiredCacheValue)
	structCacheLock     sync.RWMutex
)

type (
	optionsCacheValue struct {
		key     string
		options *fieldOptions
		err     error
	}

	requiredCacheValue struct {
		required bool
		err      error
	}
)

// Contains checks if str is in list.
func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
}

// Deref dereferences a type, if pointer type, returns its element type.
func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

// Repr returns the string representation of v.
func Repr(v interface{}) string {
	if v == nil {
		return ""
	}

	// if func (v *Type) String() string, we can't use Elem()
	switch vt := v.(type) {
	case fmt.Stringer:
		return vt.String()
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	return reprOfValue(val)
}

// ValidatePtr validates v if it's a valid pointer.
func ValidatePtr(v *reflect.Value) error {
	// sequence is very important, IsNil must be called after checking Kind() with reflect.Ptr,
	// panic otherwise
	if !v.IsValid() || v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("not a valid pointer: %v", v)
	}

	return nil
}

func convertType(kind reflect.Kind, str string) (interface{}, error) {
	switch kind {
	case reflect.Bool:
		return str == "1" || strings.ToLower(str) == "true", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("the value %q cannot parsed as int", str)
		}

		return intValue, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("the value %q cannot parsed as uint", str)
		}

		return uintValue, nil
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("the value %q cannot parsed as float", str)
		}

		return floatValue, nil
	case reflect.String:
		return str, nil
	default:
		return nil, errUnsupportedType
	}
}

func doParseKeyAndOptions(field reflect.StructField, value string) (string, *fieldOptions, error) {
	segments := parseSegments(value)
	key := strings.TrimSpace(segments[0])
	options := segments[1:]

	if len(options) == 0 {
		return key, nil, nil
	}

	var fieldOpts fieldOptions
	for _, segment := range options {
		option := strings.TrimSpace(segment)
		if err := parseOption(&fieldOpts, field.Name, option); err != nil {
			return "", nil, err
		}
	}

	return key, &fieldOpts, nil
}

// ensureValue ensures nested members not to be nil.
// If pointer value is nil, set to a new value.
func ensureValue(v reflect.Value) reflect.Value {
	for {
		if v.Kind() != reflect.Ptr {
			break
		}

		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}

	return v
}

func implicitValueRequiredStruct(tag string, tp reflect.Type) (bool, error) {
	numFields := tp.NumField()
	for i := 0; i < numFields; i++ {
		childField := tp.Field(i)
		if usingDifferentKeys(tag, childField) {
			return true, nil
		}

		_, opts, err := parseKeyAndOptions(tag, childField)
		if err != nil {
			return false, err
		}

		if opts == nil {
			if childField.Type.Kind() != reflect.Struct {
				return true, nil
			}

			if required, err := implicitValueRequiredStruct(tag, childField.Type); err != nil {
				return false, err
			} else if required {
				return true, nil
			}
		} else if !opts.Optional && len(opts.Default) == 0 {
			return true, nil
		} else if len(opts.OptionalDep) > 0 && opts.OptionalDep[0] == notSymbol {
			return true, nil
		}
	}

	return false, nil
}

func isLeftInclude(b byte) (bool, error) {
	switch b {
	case '[':
		return true, nil
	case '(':
		return false, nil
	default:
		return false, errNumberRange
	}
}

func isRightInclude(b byte) (bool, error) {
	switch b {
	case ']':
		return true, nil
	case ')':
		return false, nil
	default:
		return false, errNumberRange
	}
}

func maybeNewValue(field reflect.StructField, value reflect.Value) {
	if field.Type.Kind() == reflect.Ptr && value.IsNil() {
		value.Set(reflect.New(value.Type().Elem()))
	}
}

func parseGroupedSegments(val string) []string {
	val = strings.TrimLeftFunc(val, func(r rune) bool {
		return r == leftBracket || r == leftSquareBracket
	})
	val = strings.TrimRightFunc(val, func(r rune) bool {
		return r == rightBracket || r == rightSquareBracket
	})
	return parseSegments(val)
}

// don't modify returned fieldOptions, it's cached and shared among different calls.
func parseKeyAndOptions(tagName string, field reflect.StructField) (string, *fieldOptions, error) {
	value := field.Tag.Get(tagName)
	if len(value) == 0 {
		return field.Name, nil, nil
	}

	cacheLock.RLock()
	cache, ok := optionsCache[value]
	cacheLock.RUnlock()
	if ok {
		return stringx.TakeOne(cache.key, field.Name), cache.options, cache.err
	}

	key, options, err := doParseKeyAndOptions(field, value)
	cacheLock.Lock()
	optionsCache[value] = optionsCacheValue{
		key:     key,
		options: options,
		err:     err,
	}
	cacheLock.Unlock()

	return stringx.TakeOne(key, field.Name), options, err
}

// support below notations:
// [:5] (:5] [:5) (:5)
// [1:] [1:) (1:] (1:)
// [1:5] [1:5) (1:5] (1:5)
func parseNumberRange(str string) (*numberRange, error) {
	if len(str) == 0 {
		return nil, errNumberRange
	}

	leftInclude, err := isLeftInclude(str[0])
	if err != nil {
		return nil, err
	}

	str = str[1:]
	if len(str) == 0 {
		return nil, errNumberRange
	}

	rightInclude, err := isRightInclude(str[len(str)-1])
	if err != nil {
		return nil, err
	}

	str = str[:len(str)-1]
	fields := strings.Split(str, ":")
	if len(fields) != 2 {
		return nil, errNumberRange
	}

	if len(fields[0]) == 0 && len(fields[1]) == 0 {
		return nil, errNumberRange
	}

	var left float64
	if len(fields[0]) > 0 {
		var err error
		if left, err = strconv.ParseFloat(fields[0], 64); err != nil {
			return nil, err
		}
	} else {
		left = -math.MaxFloat64
	}

	var right float64
	if len(fields[1]) > 0 {
		var err error
		if right, err = strconv.ParseFloat(fields[1], 64); err != nil {
			return nil, err
		}
	} else {
		right = math.MaxFloat64
	}

	return &numberRange{
		left:         left,
		leftInclude:  leftInclude,
		right:        right,
		rightInclude: rightInclude,
	}, nil
}

func parseOption(fieldOpts *fieldOptions, fieldName, option string) error {
	switch {
	case option == stringOption:
		fieldOpts.FromString = true
	case strings.HasPrefix(option, optionalOption):
		segs := strings.Split(option, equalToken)
		switch len(segs) {
		case 1:
			fieldOpts.Optional = true
		case 2:
			fieldOpts.Optional = true
			fieldOpts.OptionalDep = segs[1]
		default:
			return fmt.Errorf("field %s has wrong optional", fieldName)
		}
	case option == optionalOption:
		fieldOpts.Optional = true
	case strings.HasPrefix(option, optionsOption):
		segs := strings.Split(option, equalToken)
		if len(segs) != 2 {
			return fmt.Errorf("field %s has wrong options", fieldName)
		}

		fieldOpts.Options = parseOptions(segs[1])
	case strings.HasPrefix(option, defaultOption):
		segs := strings.Split(option, equalToken)
		if len(segs) != 2 {
			return fmt.Errorf("field %s has wrong default option", fieldName)
		}

		fieldOpts.Default = strings.TrimSpace(segs[1])
	case strings.HasPrefix(option, rangeOption):
		segs := strings.Split(option, equalToken)
		if len(segs) != 2 {
			return fmt.Errorf("field %s has wrong range", fieldName)
		}

		nr, err := parseNumberRange(segs[1])
		if err != nil {
			return err
		}

		fieldOpts.Range = nr
	}

	return nil
}

// parseOptions parses the given options in tag.
// for example: `json:"name,options=foo|bar"` or `json:"name,options=[foo,bar]"`
func parseOptions(val string) []string {
	if len(val) == 0 {
		return nil
	}

	if val[0] == leftSquareBracket {
		return parseGroupedSegments(val)
	}

	return strings.Split(val, optionSeparator)
}

func parseSegments(val string) []string {
	var segments []string
	var escaped, grouped bool
	var buf strings.Builder

	for _, ch := range val {
		if escaped {
			buf.WriteRune(ch)
			escaped = false
			continue
		}

		switch ch {
		case segmentSeparator:
			if grouped {
				buf.WriteRune(ch)
			} else {
				// need to trim spaces, but we cannot ignore empty string,
				// because the first segment stands for the key might be empty.
				// if ignored, the later tag will be used as the key.
				segments = append(segments, strings.TrimSpace(buf.String()))
				buf.Reset()
			}
		case escapeChar:
			if grouped {
				buf.WriteRune(ch)
			} else {
				escaped = true
			}
		case leftBracket, leftSquareBracket:
			buf.WriteRune(ch)
			grouped = true
		case rightBracket, rightSquareBracket:
			buf.WriteRune(ch)
			grouped = false
		default:
			buf.WriteRune(ch)
		}
	}

	last := strings.TrimSpace(buf.String())
	// ignore last empty string
	if len(last) > 0 {
		segments = append(segments, last)
	}

	return segments
}

func reprOfValue(val reflect.Value) string {
	switch vt := val.Interface().(type) {
	case bool:
		return strconv.FormatBool(vt)
	case error:
		return vt.Error()
	case float32:
		return strconv.FormatFloat(float64(vt), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(vt, 'f', -1, 64)
	case fmt.Stringer:
		return vt.String()
	case int:
		return strconv.Itoa(vt)
	case int8:
		return strconv.Itoa(int(vt))
	case int16:
		return strconv.Itoa(int(vt))
	case int32:
		return strconv.Itoa(int(vt))
	case int64:
		return strconv.FormatInt(vt, 10)
	case string:
		return vt
	case uint:
		return strconv.FormatUint(uint64(vt), 10)
	case uint8:
		return strconv.FormatUint(uint64(vt), 10)
	case uint16:
		return strconv.FormatUint(uint64(vt), 10)
	case uint32:
		return strconv.FormatUint(uint64(vt), 10)
	case uint64:
		return strconv.FormatUint(vt, 10)
	case []byte:
		return string(vt)
	default:
		return fmt.Sprint(val.Interface())
	}
}

func setMatchedPrimitiveValue(kind reflect.Kind, value reflect.Value, v interface{}) error {
	switch kind {
	case reflect.Bool:
		value.SetBool(v.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(v.(int64))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(v.(uint64))
	case reflect.Float32, reflect.Float64:
		value.SetFloat(v.(float64))
	case reflect.String:
		value.SetString(v.(string))
	default:
		return errUnsupportedType
	}

	return nil
}

func setValue(kind reflect.Kind, value reflect.Value, str string) error {
	if !value.CanSet() {
		return errValueNotSettable
	}

	value = ensureValue(value)
	v, err := convertType(kind, str)
	if err != nil {
		return err
	}

	return setMatchedPrimitiveValue(kind, value, v)
}

func structValueRequired(tag string, tp reflect.Type) (bool, error) {
	structCacheLock.RLock()
	val, ok := structRequiredCache[tp]
	structCacheLock.RUnlock()
	if ok {
		return val.required, val.err
	}

	required, err := implicitValueRequiredStruct(tag, tp)
	structCacheLock.Lock()
	structRequiredCache[tp] = requiredCacheValue{
		required: required,
		err:      err,
	}
	structCacheLock.Unlock()

	return required, err
}

func toFloat64(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case int:
		return float64(val), true
	case int8:
		return float64(val), true
	case int16:
		return float64(val), true
	case int32:
		return float64(val), true
	case int64:
		return float64(val), true
	case uint:
		return float64(val), true
	case uint8:
		return float64(val), true
	case uint16:
		return float64(val), true
	case uint32:
		return float64(val), true
	case uint64:
		return float64(val), true
	case float32:
		return float64(val), true
	case float64:
		return val, true
	default:
		return 0, false
	}
}

func usingDifferentKeys(key string, field reflect.StructField) bool {
	if len(field.Tag) > 0 {
		if _, ok := field.Tag.Lookup(key); !ok {
			return true
		}
	}

	return false
}

func validateAndSetValue(kind reflect.Kind, value reflect.Value, str string, opts *fieldOptionsWithContext) error {
	if !value.CanSet() {
		return errValueNotSettable
	}

	v, err := convertType(kind, str)
	if err != nil {
		return err
	}

	if err := validateValueRange(v, opts); err != nil {
		return err
	}

	return setMatchedPrimitiveValue(kind, value, v)
}

func validateJsonNumberRange(v json.Number, opts *fieldOptionsWithContext) error {
	if opts == nil || opts.Range == nil {
		return nil
	}

	fv, err := v.Float64()
	if err != nil {
		return err
	}

	return validateNumberRange(fv, opts.Range)
}

func validateNumberRange(fv float64, nr *numberRange) error {
	if nr == nil {
		return nil
	}

	if (nr.leftInclude && fv < nr.left) || (!nr.leftInclude && fv <= nr.left) {
		return errNumberRange
	}

	if (nr.rightInclude && fv > nr.right) || (!nr.rightInclude && fv >= nr.right) {
		return errNumberRange
	}

	return nil
}

func validateValueInOptions(val interface{}, options []string) error {
	if len(options) > 0 {
		switch v := val.(type) {
		case string:
			if !stringx.Contains(options, v) {
				return fmt.Errorf(`error: value "%s" is not defined in options "%v"`, v, options)
			}
		default:
			if !stringx.Contains(options, Repr(v)) {
				return fmt.Errorf(`error: value "%v" is not defined in options "%v"`, val, options)
			}
		}
	}

	return nil
}

func validateValueRange(mapValue interface{}, opts *fieldOptionsWithContext) error {
	if opts == nil || opts.Range == nil {
		return nil
	}

	fv, ok := toFloat64(mapValue)
	if !ok {
		return errNumberRange
	}

	return validateNumberRange(fv, opts.Range)
}
