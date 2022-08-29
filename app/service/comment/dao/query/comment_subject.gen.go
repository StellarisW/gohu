// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"main/app/service/comment/dao/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCommentSubject(db *gorm.DB) commentSubject {
	_commentSubject := commentSubject{}

	_commentSubject.commentSubjectDo.UseDB(db)
	_commentSubject.commentSubjectDo.UseModel(&model.CommentSubject{})

	tableName := _commentSubject.commentSubjectDo.TableName()
	_commentSubject.ALL = field.NewField(tableName, "*")
	_commentSubject.ID = field.NewInt64(tableName, "id")
	_commentSubject.ObjType = field.NewInt32(tableName, "obj_type")
	_commentSubject.ObjID = field.NewInt64(tableName, "obj_id")
	_commentSubject.Count = field.NewInt32(tableName, "count")
	_commentSubject.RootCount = field.NewInt32(tableName, "root_count")
	_commentSubject.State = field.NewInt32(tableName, "state")
	_commentSubject.Attrs = field.NewInt32(tableName, "attrs")
	_commentSubject.CreateTime = field.NewTime(tableName, "create_time")
	_commentSubject.UpdateTime = field.NewTime(tableName, "update_time")

	_commentSubject.fillFieldMap()

	return _commentSubject
}

type commentSubject struct {
	commentSubjectDo commentSubjectDo

	ALL        field.Field
	ID         field.Int64
	ObjType    field.Int32
	ObjID      field.Int64
	Count      field.Int32
	RootCount  field.Int32
	State      field.Int32
	Attrs      field.Int32
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (c commentSubject) Table(newTableName string) *commentSubject {
	c.commentSubjectDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentSubject) As(alias string) *commentSubject {
	c.commentSubjectDo.DO = *(c.commentSubjectDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentSubject) updateTableName(table string) *commentSubject {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt64(table, "id")
	c.ObjType = field.NewInt32(table, "obj_type")
	c.ObjID = field.NewInt64(table, "obj_id")
	c.Count = field.NewInt32(table, "count")
	c.RootCount = field.NewInt32(table, "root_count")
	c.State = field.NewInt32(table, "state")
	c.Attrs = field.NewInt32(table, "attrs")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")

	c.fillFieldMap()

	return c
}

func (c *commentSubject) WithContext(ctx context.Context) *commentSubjectDo {
	return c.commentSubjectDo.WithContext(ctx)
}

func (c commentSubject) TableName() string { return c.commentSubjectDo.TableName() }

func (c commentSubject) Alias() string { return c.commentSubjectDo.Alias() }

func (c *commentSubject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentSubject) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["obj_type"] = c.ObjType
	c.fieldMap["obj_id"] = c.ObjID
	c.fieldMap["count"] = c.Count
	c.fieldMap["root_count"] = c.RootCount
	c.fieldMap["state"] = c.State
	c.fieldMap["attrs"] = c.Attrs
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
}

func (c commentSubject) clone(db *gorm.DB) commentSubject {
	c.commentSubjectDo.ReplaceDB(db)
	return c
}

type commentSubjectDo struct{ gen.DO }

func (c commentSubjectDo) Debug() *commentSubjectDo {
	return c.withDO(c.DO.Debug())
}

func (c commentSubjectDo) WithContext(ctx context.Context) *commentSubjectDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentSubjectDo) ReadDB() *commentSubjectDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentSubjectDo) WriteDB() *commentSubjectDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentSubjectDo) Clauses(conds ...clause.Expression) *commentSubjectDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentSubjectDo) Returning(value interface{}, columns ...string) *commentSubjectDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentSubjectDo) Not(conds ...gen.Condition) *commentSubjectDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentSubjectDo) Or(conds ...gen.Condition) *commentSubjectDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentSubjectDo) Select(conds ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentSubjectDo) Where(conds ...gen.Condition) *commentSubjectDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentSubjectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *commentSubjectDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentSubjectDo) Order(conds ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentSubjectDo) Distinct(cols ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentSubjectDo) Omit(cols ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentSubjectDo) Join(table schema.Tabler, on ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentSubjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentSubjectDo) RightJoin(table schema.Tabler, on ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentSubjectDo) Group(cols ...field.Expr) *commentSubjectDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentSubjectDo) Having(conds ...gen.Condition) *commentSubjectDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentSubjectDo) Limit(limit int) *commentSubjectDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentSubjectDo) Offset(offset int) *commentSubjectDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentSubjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *commentSubjectDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentSubjectDo) Unscoped() *commentSubjectDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentSubjectDo) Create(values ...*model.CommentSubject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentSubjectDo) CreateInBatches(values []*model.CommentSubject, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentSubjectDo) Save(values ...*model.CommentSubject) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentSubjectDo) First() (*model.CommentSubject, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentSubject), nil
	}
}

func (c commentSubjectDo) Take() (*model.CommentSubject, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentSubject), nil
	}
}

func (c commentSubjectDo) Last() (*model.CommentSubject, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentSubject), nil
	}
}

func (c commentSubjectDo) Find() ([]*model.CommentSubject, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommentSubject), err
}

func (c commentSubjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommentSubject, err error) {
	buf := make([]*model.CommentSubject, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentSubjectDo) FindInBatches(result *[]*model.CommentSubject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentSubjectDo) Attrs(attrs ...field.AssignExpr) *commentSubjectDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentSubjectDo) Assign(attrs ...field.AssignExpr) *commentSubjectDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentSubjectDo) Joins(fields ...field.RelationField) *commentSubjectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentSubjectDo) Preload(fields ...field.RelationField) *commentSubjectDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentSubjectDo) FirstOrInit() (*model.CommentSubject, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentSubject), nil
	}
}

func (c commentSubjectDo) FirstOrCreate() (*model.CommentSubject, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentSubject), nil
	}
}

func (c commentSubjectDo) FindByPage(offset int, limit int) (result []*model.CommentSubject, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentSubjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentSubjectDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c *commentSubjectDo) withDO(do gen.Dao) *commentSubjectDo {
	c.DO = *do.(*gen.DO)
	return c
}