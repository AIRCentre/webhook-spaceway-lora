package mysqldriver

type driverMock struct {
	queryResult []map[string][]byte
	lastQuery   string
	err         error
}

func (d *driverMock) Query(query string, args ...interface{}) ([]map[string][]byte, error) {
	d.lastQuery = query
	return d.queryResult, d.err
}

func (d *driverMock) SetQueryResult(result []map[string][]byte) {
	d.queryResult = result
}

func (d *driverMock) SetError(err error) {
	d.err = err
}

func (d *driverMock) GetLastQuery() string {
	return d.lastQuery
}

func NewMock() *driverMock {
	return &driverMock{}
}
