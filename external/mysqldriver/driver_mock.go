package mysqldriver

type driverMock struct {
	queryResult      []map[string][]byte
	lastQuery        string
	lastExec         string
	QueryCallCount   int
	execLastInsert   int64
	execRowsAffected int64
	ExecCallCount    int
	err              error
}

func (d *driverMock) Query(query string, args ...interface{}) ([]map[string][]byte, error) {
	d.lastQuery = query
	d.QueryCallCount++
	return d.queryResult, d.err
}

func (d *driverMock) Exec(query string, args ...interface{}) (lastInsertId int64, rowsAffected int64, err error) {
	d.ExecCallCount++
	d.lastExec = query
	return d.execLastInsert, d.execRowsAffected, d.err
}

func (d *driverMock) SetQueryResult(result []map[string][]byte) {
	d.queryResult = result
}

func (d *driverMock) SetExecResult(lastInsertId int64, rowsAffected int64) {
	d.execLastInsert = lastInsertId
	d.execRowsAffected = rowsAffected
}

func (d *driverMock) SetError(err error) {
	d.err = err
}

func (d *driverMock) GetLastQuery() string {
	return d.lastQuery
}
func (d *driverMock) GetLastExec() string {
	return d.lastExec
}

func NewMock() *driverMock {
	return &driverMock{}
}
