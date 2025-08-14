package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
	schema "gorm.io/gorm/schema"
)

// Dialect is a mock of DatabaseDialect interface.
type Dialect struct {
	ctrl     *gomock.Controller
	recorder *DialectRecorder
}

// DialectRecorder is the mock recorder for Dialect.
type DialectRecorder struct {
	mock *Dialect
}

// NewDialect creates a new mock instance.
func NewDialect(ctrl *gomock.Controller) *Dialect {
	mock := &Dialect{ctrl: ctrl}
	mock.recorder = &DialectRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Dialect) EXPECT() *DialectRecorder {
	return m.recorder
}

// BindVarTo mocks base method.
func (m *Dialect) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BindVarTo", writer, stmt, v)
}

// BindVarTo indicates an expected call of BindVarTo.
func (mr *DialectRecorder) BindVarTo(writer, stmt, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindVarTo", reflect.TypeOf((*Dialect)(nil).BindVarTo), writer, stmt, v)
}

// Close mocks base method.
func (m *Dialect) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *DialectRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Dialect)(nil).Close))
}

// DataTypeOf mocks base method.
func (m *Dialect) DataTypeOf(arg0 *schema.Field) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataTypeOf", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DataTypeOf indicates an expected call of DataTypeOf.
func (mr *DialectRecorder) DataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataTypeOf", reflect.TypeOf((*Dialect)(nil).DataTypeOf), arg0)
}

// DefaultValueOf mocks base method.
func (m *Dialect) DefaultValueOf(arg0 *schema.Field) clause.Expression {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultValueOf", arg0)
	ret0, _ := ret[0].(clause.Expression)
	return ret0
}

// DefaultValueOf indicates an expected call of DefaultValueOf.
func (mr *DialectRecorder) DefaultValueOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultValueOf", reflect.TypeOf((*Dialect)(nil).DefaultValueOf), arg0)
}

// Explain mocks base method.
func (m *Dialect) Explain(sql string, vars ...interface{}) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range vars {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Explain", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Explain indicates an expected call of Explain.
func (mr *DialectRecorder) Explain(sql interface{}, vars ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, vars...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Explain", reflect.TypeOf((*Dialect)(nil).Explain), varargs...)
}

// Initialize mocks base method.
func (m *Dialect) Initialize(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *DialectRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*Dialect)(nil).Initialize), arg0)
}

// Migrator mocks base method.
func (m *Dialect) Migrator(db *gorm.DB) gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator", db)
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

// Migrator indicates an expected call of Migrator.
func (mr *DialectRecorder) Migrator(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*Dialect)(nil).Migrator), db)
}

// Name mocks base method.
func (m *Dialect) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *DialectRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*Dialect)(nil).Name))
}

// QuoteTo mocks base method.
func (m *Dialect) QuoteTo(arg0 clause.Writer, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QuoteTo", arg0, arg1)
}

// QuoteTo indicates an expected call of QuoteTo.
func (mr *DialectRecorder) QuoteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuoteTo", reflect.TypeOf((*Dialect)(nil).QuoteTo), arg0, arg1)
}
