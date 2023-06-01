package mysqldriver

type I interface {
	Query(query string, args ...interface{}) ([]map[string][]byte, error)
	Exec(query string, args ...interface{}) (lastInsertId int64, rowsAffected int64, err error)
}
