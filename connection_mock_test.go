package database

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ConnectionMock struct {
	ctrl     *gomock.Controller
	recorder *ConnectionMockRecorder
}

type ConnectionMockRecorder struct {
	mock *ConnectionMock
}

func NewConnectionMock(ctrl *gomock.Controller) *ConnectionMock {
	mock := &ConnectionMock{ctrl: ctrl}
	mock.recorder = &ConnectionMockRecorder{mock}
	return mock
}

func (m *ConnectionMock) EXPECT() *ConnectionMockRecorder {
	return m.recorder
}

func (m *ConnectionMock) AddError(e error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddError", e)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) AddError(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddError", reflect.TypeOf((*ConnectionMock)(nil).AddError), e)
}

func (m *ConnectionMock) Assign(attrs ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{}, attrs...)
	ret := m.ctrl.Call(m, "Assign", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Assign(attrs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*ConnectionMock)(nil).Assign), attrs...)
}

func (m *ConnectionMock) Association(column string) *gorm.Association {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Association", column)
	ret0, _ := ret[0].(*gorm.Association)
	return ret0
}

func (mr *ConnectionMockRecorder) Association(column interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Association", reflect.TypeOf((*ConnectionMock)(nil).Association), column)
}

func (m *ConnectionMock) Attrs(attrs ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{}, attrs...)
	ret := m.ctrl.Call(m, "Attrs", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Attrs(attrs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*ConnectionMock)(nil).Attrs), attrs...)
}

func (m *ConnectionMock) AutoMigrate(dst ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{}, dst...)
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) AutoMigrate(dst ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*ConnectionMock)(nil).AutoMigrate), dst...)
}

func (m *ConnectionMock) Begin(opts ...*sql.TxOptions) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Begin", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Begin(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*ConnectionMock)(nil).Begin), opts...)
}

func (m *ConnectionMock) Clauses(conds ...clause.Expression) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Clauses", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Clauses(conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clauses", reflect.TypeOf((*ConnectionMock)(nil).Clauses), conds...)
}

func (m *ConnectionMock) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*ConnectionMock)(nil).Close))
}

func (m *ConnectionMock) Commit() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*ConnectionMock)(nil).Commit))
}

func (m *ConnectionMock) Connection(fc func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionMock", fc)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) Connection(fc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionMock", reflect.TypeOf((*ConnectionMock)(nil).Connection), fc)
}

func (m *ConnectionMock) Count(count *int64) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", count)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Count(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*ConnectionMock)(nil).Count), count)
}

func (m *ConnectionMock) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*ConnectionMock)(nil).Create), value)
}

func (m *ConnectionMock) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatches", value, batchSize)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) CreateInBatches(value, batchSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatches", reflect.TypeOf((*ConnectionMock)(nil).CreateInBatches), value, batchSize)
}

func (m *ConnectionMock) DB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *ConnectionMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*ConnectionMock)(nil).DB))
}

func (m *ConnectionMock) Debug() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Debug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*ConnectionMock)(nil).Debug))
}

func (m *ConnectionMock) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{value}, conds...)
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Delete(value interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{value}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*ConnectionMock)(nil).Delete), varargs...)
}

func (m *ConnectionMock) Distinct(args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{}, args...)
	ret := m.ctrl.Call(m, "Distinct", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Distinct(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distinct", reflect.TypeOf((*ConnectionMock)(nil).Distinct), args...)
}

func (m *ConnectionMock) Exec(sql string, values ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Exec(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*ConnectionMock)(nil).Exec), varargs...)
}

func (m *ConnectionMock) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Find(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*ConnectionMock)(nil).Find), varargs...)
}

func (m *ConnectionMock) FindInBatches(dest interface{}, batchSize int, fc func(*gorm.DB, int) error) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInBatches", dest, batchSize, fc)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) FindInBatches(dest, batchSize, fc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInBatches", reflect.TypeOf((*ConnectionMock)(nil).FindInBatches), dest, batchSize, fc)
}

func (m *ConnectionMock) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*ConnectionMock)(nil).First), varargs...)
}

func (m *ConnectionMock) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	ret := m.ctrl.Call(m, "FirstOrCreate", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) FirstOrCreate(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrCreate", reflect.TypeOf((*ConnectionMock)(nil).FirstOrCreate), varargs...)
}

func (m *ConnectionMock) FirstOrInit(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	ret := m.ctrl.Call(m, "FirstOrInit", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) FirstOrInit(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrInit", reflect.TypeOf((*ConnectionMock)(nil).FirstOrInit), varargs...)
}

func (m *ConnectionMock) Get(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0 := ret[0]
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (mr *ConnectionMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*ConnectionMock)(nil).Get), key)
}

func (m *ConnectionMock) Group(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Group(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*ConnectionMock)(nil).Group), name)
}

func (m *ConnectionMock) Having(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Having", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Having(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Having", reflect.TypeOf((*ConnectionMock)(nil).Having), varargs...)
}

func (m *ConnectionMock) InnerJoins(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "InnerJoins", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) InnerJoins(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerJoins", reflect.TypeOf((*ConnectionMock)(nil).InnerJoins), varargs...)
}

func (m *ConnectionMock) InstanceGet(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceGet", key)
	ret0 := ret[0]
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (mr *ConnectionMockRecorder) InstanceGet(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceGet", reflect.TypeOf((*ConnectionMock)(nil).InstanceGet), key)
}

func (m *ConnectionMock) InstanceSet(key string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceSet", key, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) InstanceSet(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceSet", reflect.TypeOf((*ConnectionMock)(nil).InstanceSet), key, value)
}

func (m *ConnectionMock) Joins(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Joins", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Joins(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Joins", reflect.TypeOf((*ConnectionMock)(nil).Joins), varargs...)
}

func (m *ConnectionMock) Last(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	ret := m.ctrl.Call(m, "Last", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Last(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*ConnectionMock)(nil).Last), varargs...)
}

func (m *ConnectionMock) Limit(limit int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", limit)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Limit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*ConnectionMock)(nil).Limit), limit)
}

func (m_2 *ConnectionMock) MapColumns(m map[string]string) *gorm.DB {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "MapColumns", m)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) MapColumns(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapColumns", reflect.TypeOf((*ConnectionMock)(nil).MapColumns), m)
}

func (m *ConnectionMock) Migrator() gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator")
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

func (mr *ConnectionMockRecorder) Migrator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*ConnectionMock)(nil).Migrator))
}

func (m *ConnectionMock) Model(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Model(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*ConnectionMock)(nil).Model), value)
}

func (m *ConnectionMock) Not(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Not", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Not(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Not", reflect.TypeOf((*ConnectionMock)(nil).Not), varargs...)
}

func (m *ConnectionMock) Offset(offset int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", offset)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Offset(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*ConnectionMock)(nil).Offset), offset)
}

func (m *ConnectionMock) Omit(columns ...string) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Omit", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Omit(columns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Omit", reflect.TypeOf((*ConnectionMock)(nil).Omit), columns...)
}

func (m *ConnectionMock) Or(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Or", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Or(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*ConnectionMock)(nil).Or), varargs...)
}

func (m *ConnectionMock) Order(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Order(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*ConnectionMock)(nil).Order), value)
}

func (m *ConnectionMock) Pluck(column string, dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pluck", column, dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Pluck(column, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pluck", reflect.TypeOf((*ConnectionMock)(nil).Pluck), column, dest)
}

func (m *ConnectionMock) Preload(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Preload", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Preload(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preload", reflect.TypeOf((*ConnectionMock)(nil).Preload), varargs...)
}

func (m *ConnectionMock) Raw(sql string, values ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	ret := m.ctrl.Call(m, "Raw", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Raw(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*ConnectionMock)(nil).Raw), varargs...)
}

func (m *ConnectionMock) Rollback() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*ConnectionMock)(nil).Rollback))
}

func (m *ConnectionMock) RollbackTo(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTo", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) RollbackTo(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTo", reflect.TypeOf((*ConnectionMock)(nil).RollbackTo), name)
}

func (m *ConnectionMock) Row() *sql.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Row")
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

func (mr *ConnectionMockRecorder) Row() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Row", reflect.TypeOf((*ConnectionMock)(nil).Row))
}

func (m *ConnectionMock) Rows() (*sql.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rows")
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *ConnectionMockRecorder) Rows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rows", reflect.TypeOf((*ConnectionMock)(nil).Rows))
}

func (m *ConnectionMock) Save(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Save(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*ConnectionMock)(nil).Save), value)
}

func (m *ConnectionMock) SavePoint(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePoint", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) SavePoint(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePoint", reflect.TypeOf((*ConnectionMock)(nil).SavePoint), name)
}

func (m *ConnectionMock) Scan(dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Scan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*ConnectionMock)(nil).Scan), dest)
}

func (m *ConnectionMock) ScanRows(rows *sql.Rows, dest interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanRows", rows, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) ScanRows(rows, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanRows", reflect.TypeOf((*ConnectionMock)(nil).ScanRows), rows, dest)
}

func (m *ConnectionMock) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range funcs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scopes", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Scopes(funcs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scopes", reflect.TypeOf((*ConnectionMock)(nil).Scopes), funcs...)
}

func (m *ConnectionMock) Select(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Select(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*ConnectionMock)(nil).Select), varargs...)
}

func (m *ConnectionMock) Session(config *gorm.Session) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", config)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Session(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*ConnectionMock)(nil).Session), config)
}

func (m *ConnectionMock) Set(key string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*ConnectionMock)(nil).Set), key, value)
}

func (m *ConnectionMock) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupJoinTable", model, field, joinTable)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) SetupJoinTable(model, field, joinTable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupJoinTable", reflect.TypeOf((*ConnectionMock)(nil).SetupJoinTable), model, field, joinTable)
}

func (m *ConnectionMock) Table(name string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{name}, args...)
	ret := m.ctrl.Call(m, "Table", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Table(name interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Table", reflect.TypeOf((*ConnectionMock)(nil).Table), varargs...)
}

func (m *ConnectionMock) Take(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	ret := m.ctrl.Call(m, "Take", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Take(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*ConnectionMock)(nil).Take), varargs...)
}

func (m *ConnectionMock) ToSQL(queryFn func(*gorm.DB) *gorm.DB) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSQL", queryFn)
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *ConnectionMockRecorder) ToSQL(queryFn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSQL", reflect.TypeOf((*ConnectionMock)(nil).ToSQL), queryFn)
}

func (m *ConnectionMock) Transaction(fc func(*gorm.DB) error, opts ...*sql.TxOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) Transaction(fc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*ConnectionMock)(nil).Transaction), varargs...)
}

func (m *ConnectionMock) Unscoped() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unscoped")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Unscoped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unscoped", reflect.TypeOf((*ConnectionMock)(nil).Unscoped))
}

func (m *ConnectionMock) Update(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Update(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*ConnectionMock)(nil).Update), column, value)
}

func (m *ConnectionMock) UpdateColumn(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) UpdateColumn(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*ConnectionMock)(nil).UpdateColumn), column, value)
}

func (m *ConnectionMock) UpdateColumns(values interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumns", values)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) UpdateColumns(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumns", reflect.TypeOf((*ConnectionMock)(nil).UpdateColumns), values)
}

func (m *ConnectionMock) Updates(values interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", values)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Updates(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*ConnectionMock)(nil).Updates), values)
}

func (m *ConnectionMock) Use(plugin gorm.Plugin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Use", plugin)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *ConnectionMockRecorder) Use(plugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*ConnectionMock)(nil).Use), plugin)
}

func (m *ConnectionMock) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *ConnectionMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*ConnectionMock)(nil).Where), varargs...)
}

func (m *ConnectionMock) WithContext(ctx context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}
