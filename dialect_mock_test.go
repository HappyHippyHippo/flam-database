package database

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type DialectMock struct {
	ctrl     *gomock.Controller
	recorder *DialectMockRecorder
}

type DialectMockRecorder struct {
	mock *DialectMock
}

func NewDialectMock(ctrl *gomock.Controller) *DialectMock {
	mock := &DialectMock{ctrl: ctrl}
	mock.recorder = &DialectMockRecorder{mock}
	return mock
}

func (m *DialectMock) EXPECT() *DialectMockRecorder {
	return m.recorder
}

func (m *DialectMock) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BindVarTo", writer, stmt, v)
}

func (mr *DialectMockRecorder) BindVarTo(writer, stmt, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindVarTo", reflect.TypeOf((*DialectMock)(nil).BindVarTo), writer, stmt, v)
}

func (m *DialectMock) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *DialectMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*DialectMock)(nil).Close))
}

func (m *DialectMock) DataTypeOf(arg0 *schema.Field) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataTypeOf", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *DialectMockRecorder) DataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataTypeOf", reflect.TypeOf((*DialectMock)(nil).DataTypeOf), arg0)
}

func (m *DialectMock) DefaultValueOf(arg0 *schema.Field) clause.Expression {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultValueOf", arg0)
	ret0, _ := ret[0].(clause.Expression)
	return ret0
}

func (mr *DialectMockRecorder) DefaultValueOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultValueOf", reflect.TypeOf((*DialectMock)(nil).DefaultValueOf), arg0)
}

func (m *DialectMock) Explain(sql string, vars ...interface{}) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	varargs = append(varargs, vars...)
	ret := m.ctrl.Call(m, "Explain", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *DialectMockRecorder) Explain(sql interface{}, vars ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, vars...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Explain", reflect.TypeOf((*DialectMock)(nil).Explain), varargs...)
}

func (m *DialectMock) Initialize(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *DialectMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*DialectMock)(nil).Initialize), arg0)
}

func (m *DialectMock) Migrator(db *gorm.DB) gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator", db)
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

func (mr *DialectMockRecorder) Migrator(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*DialectMock)(nil).Migrator), db)
}

func (m *DialectMock) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *DialectMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*DialectMock)(nil).Name))
}

func (m *DialectMock) QuoteTo(arg0 clause.Writer, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QuoteTo", arg0, arg1)
}

func (mr *DialectMockRecorder) QuoteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuoteTo", reflect.TypeOf((*DialectMock)(nil).QuoteTo), arg0, arg1)
}
