package mysqldriver

type I interface {
	Query(query string, args ...interface{}) ([]map[string][]byte, error)
}
