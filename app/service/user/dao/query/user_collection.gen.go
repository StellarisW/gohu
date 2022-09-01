// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"main/app/service/user/dao/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUserCollection(db *gorm.DB) userCollection {
	_userCollection := userCollection{}

	_userCollection.userCollectionDo.UseDB(db)
	_userCollection.userCollectionDo.UseModel(&model.UserCollection{})

	tableName := _userCollection.userCollectionDo.TableName()
	_userCollection.ALL = field.NewField(tableName, "*")
	_userCollection.ID = field.NewInt64(tableName, "id")
	_userCollection.UserID = field.NewInt64(tableName, "user_id")
	_userCollection.CollectType = field.NewInt32(tableName, "collect_type")
	_userCollection.ObjType = field.NewInt32(tableName, "obj_type")
	_userCollection.ObjID = field.NewInt64(tableName, "obj_id")
	_userCollection.CreateTime = field.NewTime(tableName, "create_time")
	_userCollection.UpdateTime = field.NewTime(tableName, "update_time")

	_userCollection.fillFieldMap()

	return _userCollection
}

type userCollection struct {
	userCollectionDo userCollectionDo

	ALL         field.Field
	ID          field.Int64
	UserID      field.Int64
	CollectType field.Int32
	ObjType     field.Int32
	ObjID       field.Int64
	CreateTime  field.Time
	UpdateTime  field.Time

	fieldMap map[string]field.Expr
}

func (u userCollection) Table(newTableName string) *userCollection {
	u.userCollectionDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCollection) As(alias string) *userCollection {
	u.userCollectionDo.DO = *(u.userCollectionDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCollection) updateTableName(table string) *userCollection {
	u.ALL = field.NewField(table, "*")
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewInt64(table, "user_id")
	u.CollectType = field.NewInt32(table, "collect_type")
	u.ObjType = field.NewInt32(table, "obj_type")
	u.ObjID = field.NewInt64(table, "obj_id")
	u.CreateTime = field.NewTime(table, "create_time")
	u.UpdateTime = field.NewTime(table, "update_time")

	u.fillFieldMap()

	return u
}

func (u *userCollection) WithContext(ctx context.Context) *userCollectionDo {
	return u.userCollectionDo.WithContext(ctx)
}

func (u userCollection) TableName() string { return u.userCollectionDo.TableName() }

func (u userCollection) Alias() string { return u.userCollectionDo.Alias() }

func (u *userCollection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCollection) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 7)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["collect_type"] = u.CollectType
	u.fieldMap["obj_type"] = u.ObjType
	u.fieldMap["obj_id"] = u.ObjID
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["update_time"] = u.UpdateTime
}

func (u userCollection) clone(db *gorm.DB) userCollection {
	u.userCollectionDo.ReplaceDB(db)
	return u
}

type userCollectionDo struct{ gen.DO }

func (u userCollectionDo) Debug() *userCollectionDo {
	return u.withDO(u.DO.Debug())
}

func (u userCollectionDo) WithContext(ctx context.Context) *userCollectionDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCollectionDo) ReadDB() *userCollectionDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCollectionDo) WriteDB() *userCollectionDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCollectionDo) Clauses(conds ...clause.Expression) *userCollectionDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCollectionDo) Returning(value interface{}, columns ...string) *userCollectionDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCollectionDo) Not(conds ...gen.Condition) *userCollectionDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCollectionDo) Or(conds ...gen.Condition) *userCollectionDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCollectionDo) Select(conds ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCollectionDo) Where(conds ...gen.Condition) *userCollectionDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCollectionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userCollectionDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userCollectionDo) Order(conds ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCollectionDo) Distinct(cols ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCollectionDo) Omit(cols ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCollectionDo) Join(table schema.Tabler, on ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCollectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCollectionDo) RightJoin(table schema.Tabler, on ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCollectionDo) Group(cols ...field.Expr) *userCollectionDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCollectionDo) Having(conds ...gen.Condition) *userCollectionDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCollectionDo) Limit(limit int) *userCollectionDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCollectionDo) Offset(offset int) *userCollectionDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCollectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userCollectionDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCollectionDo) Unscoped() *userCollectionDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCollectionDo) Create(values ...*model.UserCollection) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCollectionDo) CreateInBatches(values []*model.UserCollection, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCollectionDo) Save(values ...*model.UserCollection) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCollectionDo) First() (*model.UserCollection, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCollection), nil
	}
}

func (u userCollectionDo) Take() (*model.UserCollection, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCollection), nil
	}
}

func (u userCollectionDo) Last() (*model.UserCollection, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCollection), nil
	}
}

func (u userCollectionDo) Find() ([]*model.UserCollection, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCollection), err
}

func (u userCollectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCollection, err error) {
	buf := make([]*model.UserCollection, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCollectionDo) FindInBatches(result *[]*model.UserCollection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCollectionDo) Attrs(attrs ...field.AssignExpr) *userCollectionDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCollectionDo) Assign(attrs ...field.AssignExpr) *userCollectionDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCollectionDo) Joins(fields ...field.RelationField) *userCollectionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCollectionDo) Preload(fields ...field.RelationField) *userCollectionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCollectionDo) FirstOrInit() (*model.UserCollection, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCollection), nil
	}
}

func (u userCollectionDo) FirstOrCreate() (*model.UserCollection, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCollection), nil
	}
}

func (u userCollectionDo) FindByPage(offset int, limit int) (result []*model.UserCollection, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userCollectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCollectionDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u *userCollectionDo) withDO(do gen.Dao) *userCollectionDo {
	u.DO = *do.(*gen.DO)
	return u
}
