package sqlz

import "testing"

func TestPostgreSQLURI(t *testing.T) {
	println(PostgreSQLURI("root", "root", "localhost", 5432, "jemmadb",false, true))
}