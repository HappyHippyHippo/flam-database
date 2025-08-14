package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
)

// Connection is a mock of DatabaseConnection interface.
type Connection struct {
	ctrl     *gomock.Controller
	recorder *ConnectionRecorder
}

// ConnectionRecorder is the mock recorder for Connection.
type ConnectionRecorder struct {
	mock *Connection
}

// NewConnection creates a new mock instance.
func NewConnection(ctrl *gomock.Controller) *Connection {
	mock := &Connection{ctrl: ctrl}
	mock.recorder = &ConnectionRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Connection) EXPECT() *ConnectionRecorder {
	return m.recorder
}

// AddError mocks base method.
func (m *Connection) AddError(e error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddError", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddError indicates an expected call of AddError.
func (mr *ConnectionRecorder) AddError(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddError", reflect.TypeOf((*Connection)(nil).AddError), e)
}

// Assign mocks base method.
func (m *Connection) Assign(attrs ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range attrs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Assign", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Assign indicates an expected call of Assign.
func (mr *ConnectionRecorder) Assign(attrs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*Connection)(nil).Assign), attrs...)
}

// Association mocks base method.
func (m *Connection) Association(column string) *gorm.Association {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Association", column)
	ret0, _ := ret[0].(*gorm.Association)
	return ret0
}

// Association indicates an expected call of Association.
func (mr *ConnectionRecorder) Association(column interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Association", reflect.TypeOf((*Connection)(nil).Association), column)
}

// Attrs mocks base method.
func (m *Connection) Attrs(attrs ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range attrs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Attrs", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Attrs indicates an expected call of Attrs.
func (mr *ConnectionRecorder) Attrs(attrs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*Connection)(nil).Attrs), attrs...)
}

// AutoMigrate mocks base method.
func (m *Connection) AutoMigrate(dst ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range dst {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *ConnectionRecorder) AutoMigrate(dst ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*Connection)(nil).AutoMigrate), dst...)
}

// Begin mocks base method.
func (m *Connection) Begin(opts ...*sql.TxOptions) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Begin", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *ConnectionRecorder) Begin(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*Connection)(nil).Begin), opts...)
}

// Clauses mocks base method.
func (m *Connection) Clauses(conds ...clause.Expression) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Clauses", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Clauses indicates an expected call of Clauses.
func (mr *ConnectionRecorder) Clauses(conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clauses", reflect.TypeOf((*Connection)(nil).Clauses), conds...)
}

// Close mocks base method.
func (m *Connection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *ConnectionRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Connection)(nil).Close))
}

// Commit mocks base method.
func (m *Connection) Commit() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *ConnectionRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*Connection)(nil).Commit))
}

// Connection mocks base method.
func (m *Connection) Connection(fc func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connection", fc)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connection indicates an expected call of Connection.
func (mr *ConnectionRecorder) Connection(fc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connection", reflect.TypeOf((*Connection)(nil).Connection), fc)
}

// Count mocks base method.
func (m *Connection) Count(count *int64) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", count)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *ConnectionRecorder) Count(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*Connection)(nil).Count), count)
}

// Create mocks base method.
func (m *Connection) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *ConnectionRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Connection)(nil).Create), value)
}

// CreateInBatches mocks base method.
func (m *Connection) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatches", value, batchSize)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// CreateInBatches indicates an expected call of CreateInBatches.
func (mr *ConnectionRecorder) CreateInBatches(value, batchSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatches", reflect.TypeOf((*Connection)(nil).CreateInBatches), value, batchSize)
}

// DB mocks base method.
func (m *Connection) DB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DB indicates an expected call of DB.
func (mr *ConnectionRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*Connection)(nil).DB))
}

// Debug mocks base method.
func (m *Connection) Debug() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Debug indicates an expected call of Debug.
func (mr *ConnectionRecorder) Debug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*Connection)(nil).Debug))
}

// Delete mocks base method.
func (m *Connection) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{value}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *ConnectionRecorder) Delete(value interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{value}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Connection)(nil).Delete), varargs...)
}

// Distinct mocks base method.
func (m *Connection) Distinct(args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Distinct", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Distinct indicates an expected call of Distinct.
func (mr *ConnectionRecorder) Distinct(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distinct", reflect.TypeOf((*Connection)(nil).Distinct), args...)
}

// Exec mocks base method.
func (m *Connection) Exec(sql string, values ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *ConnectionRecorder) Exec(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*Connection)(nil).Exec), varargs...)
}

// Find mocks base method.
func (m *Connection) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *ConnectionRecorder) Find(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*Connection)(nil).Find), varargs...)
}

// FindInBatches mocks base method.
func (m *Connection) FindInBatches(dest interface{}, batchSize int, fc func(*gorm.DB, int) error) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInBatches", dest, batchSize, fc)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// FindInBatches indicates an expected call of FindInBatches.
func (mr *ConnectionRecorder) FindInBatches(dest, batchSize, fc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInBatches", reflect.TypeOf((*Connection)(nil).FindInBatches), dest, batchSize, fc)
}

// First mocks base method.
func (m *Connection) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *ConnectionRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*Connection)(nil).First), varargs...)
}

// FirstOrCreate mocks base method.
func (m *Connection) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstOrCreate", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// FirstOrCreate indicates an expected call of FirstOrCreate.
func (mr *ConnectionRecorder) FirstOrCreate(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrCreate", reflect.TypeOf((*Connection)(nil).FirstOrCreate), varargs...)
}

// FirstOrInit mocks base method.
func (m *Connection) FirstOrInit(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstOrInit", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// FirstOrInit indicates an expected call of FirstOrInit.
func (mr *ConnectionRecorder) FirstOrInit(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrInit", reflect.TypeOf((*Connection)(nil).FirstOrInit), varargs...)
}

// Get mocks base method.
func (m *Connection) Get(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *ConnectionRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Connection)(nil).Get), key)
}

// Group mocks base method.
func (m *Connection) Group(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *ConnectionRecorder) Group(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*Connection)(nil).Group), name)
}

// Having mocks base method.
func (m *Connection) Having(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Having", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Having indicates an expected call of Having.
func (mr *ConnectionRecorder) Having(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Having", reflect.TypeOf((*Connection)(nil).Having), varargs...)
}

// InnerJoins mocks base method.
func (m *Connection) InnerJoins(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InnerJoins", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// InnerJoins indicates an expected call of InnerJoins.
func (mr *ConnectionRecorder) InnerJoins(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerJoins", reflect.TypeOf((*Connection)(nil).InnerJoins), varargs...)
}

// InstanceGet mocks base method.
func (m *Connection) InstanceGet(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceGet", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// InstanceGet indicates an expected call of InstanceGet.
func (mr *ConnectionRecorder) InstanceGet(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceGet", reflect.TypeOf((*Connection)(nil).InstanceGet), key)
}

// InstanceSet mocks base method.
func (m *Connection) InstanceSet(key string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceSet", key, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// InstanceSet indicates an expected call of InstanceSet.
func (mr *ConnectionRecorder) InstanceSet(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceSet", reflect.TypeOf((*Connection)(nil).InstanceSet), key, value)
}

// Joins mocks base method.
func (m *Connection) Joins(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Joins", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Joins indicates an expected call of Joins.
func (mr *ConnectionRecorder) Joins(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Joins", reflect.TypeOf((*Connection)(nil).Joins), varargs...)
}

// Last mocks base method.
func (m *Connection) Last(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Last", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *ConnectionRecorder) Last(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*Connection)(nil).Last), varargs...)
}

// Limit mocks base method.
func (m *Connection) Limit(limit int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", limit)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Limit indicates an expected call of Limit.
func (mr *ConnectionRecorder) Limit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*Connection)(nil).Limit), limit)
}

// MapColumns mocks base method.
func (m_2 *Connection) MapColumns(m map[string]string) *gorm.DB {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "MapColumns", m)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// MapColumns indicates an expected call of MapColumns.
func (mr *ConnectionRecorder) MapColumns(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapColumns", reflect.TypeOf((*Connection)(nil).MapColumns), m)
}

// Migrator mocks base method.
func (m *Connection) Migrator() gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator")
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

// Migrator indicates an expected call of Migrator.
func (mr *ConnectionRecorder) Migrator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*Connection)(nil).Migrator))
}

// Model mocks base method.
func (m *Connection) Model(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *ConnectionRecorder) Model(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*Connection)(nil).Model), value)
}

// Not mocks base method.
func (m *Connection) Not(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Not", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Not indicates an expected call of Not.
func (mr *ConnectionRecorder) Not(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Not", reflect.TypeOf((*Connection)(nil).Not), varargs...)
}

// Offset mocks base method.
func (m *Connection) Offset(offset int) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", offset)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *ConnectionRecorder) Offset(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*Connection)(nil).Offset), offset)
}

// Omit mocks base method.
func (m *Connection) Omit(columns ...string) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Omit", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Omit indicates an expected call of Omit.
func (mr *ConnectionRecorder) Omit(columns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Omit", reflect.TypeOf((*Connection)(nil).Omit), columns...)
}

// Or mocks base method.
func (m *Connection) Or(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Or", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Or indicates an expected call of Or.
func (mr *ConnectionRecorder) Or(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*Connection)(nil).Or), varargs...)
}

// Order mocks base method.
func (m *Connection) Order(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Order indicates an expected call of Order.
func (mr *ConnectionRecorder) Order(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*Connection)(nil).Order), value)
}

// Pluck mocks base method.
func (m *Connection) Pluck(column string, dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pluck", column, dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Pluck indicates an expected call of Pluck.
func (mr *ConnectionRecorder) Pluck(column, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pluck", reflect.TypeOf((*Connection)(nil).Pluck), column, dest)
}

// Preload mocks base method.
func (m *Connection) Preload(query string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Preload", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Preload indicates an expected call of Preload.
func (mr *ConnectionRecorder) Preload(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preload", reflect.TypeOf((*Connection)(nil).Preload), varargs...)
}

// Raw mocks base method.
func (m *Connection) Raw(sql string, values ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Raw", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *ConnectionRecorder) Raw(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*Connection)(nil).Raw), varargs...)
}

// Rollback mocks base method.
func (m *Connection) Rollback() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *ConnectionRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*Connection)(nil).Rollback))
}

// RollbackTo mocks base method.
func (m *Connection) RollbackTo(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTo", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// RollbackTo indicates an expected call of RollbackTo.
func (mr *ConnectionRecorder) RollbackTo(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTo", reflect.TypeOf((*Connection)(nil).RollbackTo), name)
}

// Row mocks base method.
func (m *Connection) Row() *sql.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Row")
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// Row indicates an expected call of Row.
func (mr *ConnectionRecorder) Row() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Row", reflect.TypeOf((*Connection)(nil).Row))
}

// Rows mocks base method.
func (m *Connection) Rows() (*sql.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rows")
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rows indicates an expected call of Rows.
func (mr *ConnectionRecorder) Rows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rows", reflect.TypeOf((*Connection)(nil).Rows))
}

// Save mocks base method.
func (m *Connection) Save(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *ConnectionRecorder) Save(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*Connection)(nil).Save), value)
}

// SavePoint mocks base method.
func (m *Connection) SavePoint(name string) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePoint", name)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// SavePoint indicates an expected call of SavePoint.
func (mr *ConnectionRecorder) SavePoint(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePoint", reflect.TypeOf((*Connection)(nil).SavePoint), name)
}

// Scan mocks base method.
func (m *Connection) Scan(dest interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", dest)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *ConnectionRecorder) Scan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*Connection)(nil).Scan), dest)
}

// ScanRows mocks base method.
func (m *Connection) ScanRows(rows *sql.Rows, dest interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanRows", rows, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanRows indicates an expected call of ScanRows.
func (mr *ConnectionRecorder) ScanRows(rows, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanRows", reflect.TypeOf((*Connection)(nil).ScanRows), rows, dest)
}

// Scopes mocks base method.
func (m *Connection) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range funcs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scopes", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Scopes indicates an expected call of Scopes.
func (mr *ConnectionRecorder) Scopes(funcs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scopes", reflect.TypeOf((*Connection)(nil).Scopes), funcs...)
}

// Select mocks base method.
func (m *Connection) Select(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *ConnectionRecorder) Select(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*Connection)(nil).Select), varargs...)
}

// Session mocks base method.
func (m *Connection) Session(config *gorm.Session) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", config)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Session indicates an expected call of Session.
func (mr *ConnectionRecorder) Session(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*Connection)(nil).Session), config)
}

// Set mocks base method.
func (m *Connection) Set(key string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *ConnectionRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*Connection)(nil).Set), key, value)
}

// SetupJoinTable mocks base method.
func (m *Connection) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupJoinTable", model, field, joinTable)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupJoinTable indicates an expected call of SetupJoinTable.
func (mr *ConnectionRecorder) SetupJoinTable(model, field, joinTable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupJoinTable", reflect.TypeOf((*Connection)(nil).SetupJoinTable), model, field, joinTable)
}

// Table mocks base method.
func (m *Connection) Table(name string, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Table", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Table indicates an expected call of Table.
func (mr *ConnectionRecorder) Table(name interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Table", reflect.TypeOf((*Connection)(nil).Table), varargs...)
}

// Take mocks base method.
func (m *Connection) Take(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Take", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Take indicates an expected call of Take.
func (mr *ConnectionRecorder) Take(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*Connection)(nil).Take), varargs...)
}

// ToSQL mocks base method.
func (m *Connection) ToSQL(queryFn func(*gorm.DB) *gorm.DB) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSQL", queryFn)
	ret0, _ := ret[0].(string)
	return ret0
}

// ToSQL indicates an expected call of ToSQL.
func (mr *ConnectionRecorder) ToSQL(queryFn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSQL", reflect.TypeOf((*Connection)(nil).ToSQL), queryFn)
}

// Transaction mocks base method.
func (m *Connection) Transaction(fc func(*gorm.DB) error, opts ...*sql.TxOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *ConnectionRecorder) Transaction(fc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*Connection)(nil).Transaction), varargs...)
}

// Unscoped mocks base method.
func (m *Connection) Unscoped() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unscoped")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Unscoped indicates an expected call of Unscoped.
func (mr *ConnectionRecorder) Unscoped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unscoped", reflect.TypeOf((*Connection)(nil).Unscoped))
}

// Update mocks base method.
func (m *Connection) Update(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *ConnectionRecorder) Update(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Connection)(nil).Update), column, value)
}

// UpdateColumn mocks base method.
func (m *Connection) UpdateColumn(column string, value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", column, value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *ConnectionRecorder) UpdateColumn(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*Connection)(nil).UpdateColumn), column, value)
}

// UpdateColumns mocks base method.
func (m *Connection) UpdateColumns(values interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumns", values)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// UpdateColumns indicates an expected call of UpdateColumns.
func (mr *ConnectionRecorder) UpdateColumns(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumns", reflect.TypeOf((*Connection)(nil).UpdateColumns), values)
}

// Updates mocks base method.
func (m *Connection) Updates(values interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", values)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *ConnectionRecorder) Updates(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*Connection)(nil).Updates), values)
}

// Use mocks base method.
func (m *Connection) Use(plugin gorm.Plugin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Use", plugin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *ConnectionRecorder) Use(plugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*Connection)(nil).Use), plugin)
}

// Where mocks base method.
func (m *Connection) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *ConnectionRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*Connection)(nil).Where), varargs...)
}

// WithContext mocks base method.
func (m *Connection) WithContext(ctx context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}
