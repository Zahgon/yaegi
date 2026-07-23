//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"context"
	"database/sql/driver"
	"reflect"
)

func init() {
	Symbols["database/sql/driver/driver"] = map[string]reflect.Value{

		"Bool":                      reflect.ValueOf(&driver.Bool).Elem(),
		"DefaultParameterConverter": reflect.ValueOf(&driver.DefaultParameterConverter).Elem(),
		"ErrBadConn":                reflect.ValueOf(&driver.ErrBadConn).Elem(),
		"ErrRemoveArgument":         reflect.ValueOf(&driver.ErrRemoveArgument).Elem(),
		"ErrSkip":                   reflect.ValueOf(&driver.ErrSkip).Elem(),
		"Int32":                     reflect.ValueOf(&driver.Int32).Elem(),
		"IsScanValue":               reflect.ValueOf(driver.IsScanValue),
		"IsValue":                   reflect.ValueOf(driver.IsValue),
		"ResultNoRows":              reflect.ValueOf(&driver.ResultNoRows).Elem(),
		"String":                    reflect.ValueOf(&driver.String).Elem(),

		"ColumnConverter":                reflect.ValueOf((*driver.ColumnConverter)(nil)),
		"Conn":                           reflect.ValueOf((*driver.Conn)(nil)),
		"ConnBeginTx":                    reflect.ValueOf((*driver.ConnBeginTx)(nil)),
		"ConnPrepareContext":             reflect.ValueOf((*driver.ConnPrepareContext)(nil)),
		"Connector":                      reflect.ValueOf((*driver.Connector)(nil)),
		"Driver":                         reflect.ValueOf((*driver.Driver)(nil)),
		"DriverContext":                  reflect.ValueOf((*driver.DriverContext)(nil)),
		"Execer":                         reflect.ValueOf((*driver.Execer)(nil)),
		"ExecerContext":                  reflect.ValueOf((*driver.ExecerContext)(nil)),
		"IsolationLevel":                 reflect.ValueOf((*driver.IsolationLevel)(nil)),
		"NamedValue":                     reflect.ValueOf((*driver.NamedValue)(nil)),
		"NamedValueChecker":              reflect.ValueOf((*driver.NamedValueChecker)(nil)),
		"NotNull":                        reflect.ValueOf((*driver.NotNull)(nil)),
		"Null":                           reflect.ValueOf((*driver.Null)(nil)),
		"Pinger":                         reflect.ValueOf((*driver.Pinger)(nil)),
		"Queryer":                        reflect.ValueOf((*driver.Queryer)(nil)),
		"QueryerContext":                 reflect.ValueOf((*driver.QueryerContext)(nil)),
		"Result":                         reflect.ValueOf((*driver.Result)(nil)),
		"Rows":                           reflect.ValueOf((*driver.Rows)(nil)),
		"RowsAffected":                   reflect.ValueOf((*driver.RowsAffected)(nil)),
		"RowsColumnTypeDatabaseTypeName": reflect.ValueOf((*driver.RowsColumnTypeDatabaseTypeName)(nil)),
		"RowsColumnTypeLength":           reflect.ValueOf((*driver.RowsColumnTypeLength)(nil)),
		"RowsColumnTypeNullable":         reflect.ValueOf((*driver.RowsColumnTypeNullable)(nil)),
		"RowsColumnTypePrecisionScale":   reflect.ValueOf((*driver.RowsColumnTypePrecisionScale)(nil)),
		"RowsColumnTypeScanType":         reflect.ValueOf((*driver.RowsColumnTypeScanType)(nil)),
		"RowsNextResultSet":              reflect.ValueOf((*driver.RowsNextResultSet)(nil)),
		"SessionResetter":                reflect.ValueOf((*driver.SessionResetter)(nil)),
		"Stmt":                           reflect.ValueOf((*driver.Stmt)(nil)),
		"StmtExecContext":                reflect.ValueOf((*driver.StmtExecContext)(nil)),
		"StmtQueryContext":               reflect.ValueOf((*driver.StmtQueryContext)(nil)),
		"Tx":                             reflect.ValueOf((*driver.Tx)(nil)),
		"TxOptions":                      reflect.ValueOf((*driver.TxOptions)(nil)),
		"Validator":                      reflect.ValueOf((*driver.Validator)(nil)),
		"Value":                          reflect.ValueOf((*driver.Value)(nil)),
		"ValueConverter":                 reflect.ValueOf((*driver.ValueConverter)(nil)),
		"Valuer":                         reflect.ValueOf((*driver.Valuer)(nil)),

		"_ColumnConverter":                reflect.ValueOf((*_database_sql_driver_ColumnConverter)(nil)),
		"_Conn":                           reflect.ValueOf((*_database_sql_driver_Conn)(nil)),
		"_ConnBeginTx":                    reflect.ValueOf((*_database_sql_driver_ConnBeginTx)(nil)),
		"_ConnPrepareContext":             reflect.ValueOf((*_database_sql_driver_ConnPrepareContext)(nil)),
		"_Connector":                      reflect.ValueOf((*_database_sql_driver_Connector)(nil)),
		"_Driver":                         reflect.ValueOf((*_database_sql_driver_Driver)(nil)),
		"_DriverContext":                  reflect.ValueOf((*_database_sql_driver_DriverContext)(nil)),
		"_Execer":                         reflect.ValueOf((*_database_sql_driver_Execer)(nil)),
		"_ExecerContext":                  reflect.ValueOf((*_database_sql_driver_ExecerContext)(nil)),
		"_NamedValueChecker":              reflect.ValueOf((*_database_sql_driver_NamedValueChecker)(nil)),
		"_Pinger":                         reflect.ValueOf((*_database_sql_driver_Pinger)(nil)),
		"_Queryer":                        reflect.ValueOf((*_database_sql_driver_Queryer)(nil)),
		"_QueryerContext":                 reflect.ValueOf((*_database_sql_driver_QueryerContext)(nil)),
		"_Result":                         reflect.ValueOf((*_database_sql_driver_Result)(nil)),
		"_Rows":                           reflect.ValueOf((*_database_sql_driver_Rows)(nil)),
		"_RowsColumnTypeDatabaseTypeName": reflect.ValueOf((*_database_sql_driver_RowsColumnTypeDatabaseTypeName)(nil)),
		"_RowsColumnTypeLength":           reflect.ValueOf((*_database_sql_driver_RowsColumnTypeLength)(nil)),
		"_RowsColumnTypeNullable":         reflect.ValueOf((*_database_sql_driver_RowsColumnTypeNullable)(nil)),
		"_RowsColumnTypePrecisionScale":   reflect.ValueOf((*_database_sql_driver_RowsColumnTypePrecisionScale)(nil)),
		"_RowsColumnTypeScanType":         reflect.ValueOf((*_database_sql_driver_RowsColumnTypeScanType)(nil)),
		"_RowsNextResultSet":              reflect.ValueOf((*_database_sql_driver_RowsNextResultSet)(nil)),
		"_SessionResetter":                reflect.ValueOf((*_database_sql_driver_SessionResetter)(nil)),
		"_Stmt":                           reflect.ValueOf((*_database_sql_driver_Stmt)(nil)),
		"_StmtExecContext":                reflect.ValueOf((*_database_sql_driver_StmtExecContext)(nil)),
		"_StmtQueryContext":               reflect.ValueOf((*_database_sql_driver_StmtQueryContext)(nil)),
		"_Tx":                             reflect.ValueOf((*_database_sql_driver_Tx)(nil)),
		"_Validator":                      reflect.ValueOf((*_database_sql_driver_Validator)(nil)),
		"_Value":                          reflect.ValueOf((*_database_sql_driver_Value)(nil)),
		"_ValueConverter":                 reflect.ValueOf((*_database_sql_driver_ValueConverter)(nil)),
		"_Valuer":                         reflect.ValueOf((*_database_sql_driver_Valuer)(nil)),
	}
}

type _database_sql_driver_ColumnConverter struct {
	IValue           interface{}
	WColumnConverter func(idx int) driver.ValueConverter
}

func (W _database_sql_driver_ColumnConverter) ColumnConverter(idx int) driver.ValueConverter {
	_ = "STUB: not implemented"
	return *new(driver.ValueConverter)
}

type _database_sql_driver_Conn struct {
	IValue   interface{}
	WBegin   func() (driver.Tx, error)
	WClose   func() error
	WPrepare func(query string) (driver.Stmt, error)
}

func (W _database_sql_driver_Conn) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}
func (W _database_sql_driver_Conn) Close() error { _ = "STUB: not implemented"; return nil }
func (W _database_sql_driver_Conn) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

type _database_sql_driver_ConnBeginTx struct {
	IValue   interface{}
	WBeginTx func(ctx context.Context, opts driver.TxOptions) (driver.Tx, error)
}

func (W _database_sql_driver_ConnBeginTx) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

type _database_sql_driver_ConnPrepareContext struct {
	IValue          interface{}
	WPrepareContext func(ctx context.Context, query string) (driver.Stmt, error)
}

func (W _database_sql_driver_ConnPrepareContext) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

type _database_sql_driver_Connector struct {
	IValue   interface{}
	WConnect func(a0 context.Context) (driver.Conn, error)
	WDriver  func() driver.Driver
}

func (W _database_sql_driver_Connector) Connect(a0 context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (W _database_sql_driver_Connector) Driver() driver.Driver {
	_ = "STUB: not implemented"
	return *new(driver.Driver)
}

type _database_sql_driver_Driver struct {
	IValue interface{}
	WOpen  func(name string) (driver.Conn, error)
}

func (W _database_sql_driver_Driver) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

type _database_sql_driver_DriverContext struct {
	IValue         interface{}
	WOpenConnector func(name string) (driver.Connector, error)
}

func (W _database_sql_driver_DriverContext) OpenConnector(name string) (driver.Connector, error) {
	_ = "STUB: not implemented"
	return *new(driver.Connector), nil
}

type _database_sql_driver_Execer struct {
	IValue interface{}
	WExec  func(query string, args []driver.Value) (driver.Result, error)
}

func (W _database_sql_driver_Execer) Exec(query string, args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

type _database_sql_driver_ExecerContext struct {
	IValue       interface{}
	WExecContext func(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error)
}

func (W _database_sql_driver_ExecerContext) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

type _database_sql_driver_NamedValueChecker struct {
	IValue           interface{}
	WCheckNamedValue func(a0 *driver.NamedValue) error
}

func (W _database_sql_driver_NamedValueChecker) CheckNamedValue(a0 *driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_Pinger struct {
	IValue interface{}
	WPing  func(ctx context.Context) error
}

func (W _database_sql_driver_Pinger) Ping(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_Queryer struct {
	IValue interface{}
	WQuery func(query string, args []driver.Value) (driver.Rows, error)
}

func (W _database_sql_driver_Queryer) Query(query string, args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

type _database_sql_driver_QueryerContext struct {
	IValue        interface{}
	WQueryContext func(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error)
}

func (W _database_sql_driver_QueryerContext) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

type _database_sql_driver_Result struct {
	IValue        interface{}
	WLastInsertId func() (int64, error)
	WRowsAffected func() (int64, error)
}

func (W _database_sql_driver_Result) LastInsertId() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
func (W _database_sql_driver_Result) RowsAffected() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type _database_sql_driver_Rows struct {
	IValue   interface{}
	WClose   func() error
	WColumns func() []string
	WNext    func(dest []driver.Value) error
}

func (W _database_sql_driver_Rows) Close() error      { _ = "STUB: not implemented"; return nil }
func (W _database_sql_driver_Rows) Columns() []string { _ = "STUB: not implemented"; return nil }
func (W _database_sql_driver_Rows) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_RowsColumnTypeDatabaseTypeName struct {
	IValue                      interface{}
	WClose                      func() error
	WColumnTypeDatabaseTypeName func(index int) string
	WColumns                    func() []string
	WNext                       func(dest []driver.Value) error
}

func (W _database_sql_driver_RowsColumnTypeDatabaseTypeName) Close() error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeDatabaseTypeName) ColumnTypeDatabaseTypeName(index int) string {
	_ = "STUB: not implemented"
	return ""
}

func (W _database_sql_driver_RowsColumnTypeDatabaseTypeName) Columns() []string {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeDatabaseTypeName) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_RowsColumnTypeLength struct {
	IValue            interface{}
	WClose            func() error
	WColumnTypeLength func(index int) (length int64, ok bool)
	WColumns          func() []string
	WNext             func(dest []driver.Value) error
}

func (W _database_sql_driver_RowsColumnTypeLength) Close() error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeLength) ColumnTypeLength(index int) (length int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (W _database_sql_driver_RowsColumnTypeLength) Columns() []string {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeLength) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_RowsColumnTypeNullable struct {
	IValue              interface{}
	WClose              func() error
	WColumnTypeNullable func(index int) (nullable bool, ok bool)
	WColumns            func() []string
	WNext               func(dest []driver.Value) error
}

func (W _database_sql_driver_RowsColumnTypeNullable) Close() error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeNullable) ColumnTypeNullable(index int) (nullable bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (W _database_sql_driver_RowsColumnTypeNullable) Columns() []string {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeNullable) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_RowsColumnTypePrecisionScale struct {
	IValue                    interface{}
	WClose                    func() error
	WColumnTypePrecisionScale func(index int) (precision int64, scale int64, ok bool)
	WColumns                  func() []string
	WNext                     func(dest []driver.Value) error
}

func (W _database_sql_driver_RowsColumnTypePrecisionScale) Close() error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypePrecisionScale) ColumnTypePrecisionScale(index int) (precision int64, scale int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func (W _database_sql_driver_RowsColumnTypePrecisionScale) Columns() []string {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypePrecisionScale) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_RowsColumnTypeScanType struct {
	IValue              interface{}
	WClose              func() error
	WColumnTypeScanType func(index int) reflect.Type
	WColumns            func() []string
	WNext               func(dest []driver.Value) error
}

func (W _database_sql_driver_RowsColumnTypeScanType) Close() error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeScanType) ColumnTypeScanType(index int) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (W _database_sql_driver_RowsColumnTypeScanType) Columns() []string {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsColumnTypeScanType) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_RowsNextResultSet struct {
	IValue            interface{}
	WClose            func() error
	WColumns          func() []string
	WHasNextResultSet func() bool
	WNext             func(dest []driver.Value) error
	WNextResultSet    func() error
}

func (W _database_sql_driver_RowsNextResultSet) Close() error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsNextResultSet) Columns() []string {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsNextResultSet) HasNextResultSet() bool {
	_ = "STUB: not implemented"
	return false
}
func (W _database_sql_driver_RowsNextResultSet) Next(dest []driver.Value) error {
	_ = "STUB: not implemented"
	return nil
}
func (W _database_sql_driver_RowsNextResultSet) NextResultSet() error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_SessionResetter struct {
	IValue        interface{}
	WResetSession func(ctx context.Context) error
}

func (W _database_sql_driver_SessionResetter) ResetSession(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type _database_sql_driver_Stmt struct {
	IValue    interface{}
	WClose    func() error
	WExec     func(args []driver.Value) (driver.Result, error)
	WNumInput func() int
	WQuery    func(args []driver.Value) (driver.Rows, error)
}

func (W _database_sql_driver_Stmt) Close() error { _ = "STUB: not implemented"; return nil }
func (W _database_sql_driver_Stmt) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (W _database_sql_driver_Stmt) NumInput() int { _ = "STUB: not implemented"; return 0 }
func (W _database_sql_driver_Stmt) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

type _database_sql_driver_StmtExecContext struct {
	IValue       interface{}
	WExecContext func(ctx context.Context, args []driver.NamedValue) (driver.Result, error)
}

func (W _database_sql_driver_StmtExecContext) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

type _database_sql_driver_StmtQueryContext struct {
	IValue        interface{}
	WQueryContext func(ctx context.Context, args []driver.NamedValue) (driver.Rows, error)
}

func (W _database_sql_driver_StmtQueryContext) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

type _database_sql_driver_Tx struct {
	IValue    interface{}
	WCommit   func() error
	WRollback func() error
}

func (W _database_sql_driver_Tx) Commit() error   { _ = "STUB: not implemented"; return nil }
func (W _database_sql_driver_Tx) Rollback() error { _ = "STUB: not implemented"; return nil }

type _database_sql_driver_Validator struct {
	IValue   interface{}
	WIsValid func() bool
}

func (W _database_sql_driver_Validator) IsValid() bool { _ = "STUB: not implemented"; return false }

type _database_sql_driver_Value struct {
	IValue interface{}
}

type _database_sql_driver_ValueConverter struct {
	IValue        interface{}
	WConvertValue func(v any) (driver.Value, error)
}

func (W _database_sql_driver_ValueConverter) ConvertValue(v any) (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

type _database_sql_driver_Valuer struct {
	IValue interface{}
	WValue func() (driver.Value, error)
}

func (W _database_sql_driver_Valuer) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
