package sqlz

import "fmt"

const (
	PostgresqlPassword 	   = "POSTGRESQL_PASSWORD"
	PostgresqlDatabase     = "POSTGRESQL_DATABASE"
	PostgresqlUser         = "POSTGRESQL_USER"
	PostgresqlHost         = "POSTGRESQL_HOST"
	PostgresqlPort         = "POSTGRESQL_PORT"
	PostgresqlSslMode      = "POSTGRESQL_SSL_MODE"
)

//host=%s port=%d user=%s password=%s dbname=%s sslmode=%s
func PostgreSQLURI(user string, pw string, host string, port int, database string, sslmode bool) string {
	var sslmodetext string
	if sslmode {
		sslmodetext = "require"
	} else {
		sslmodetext = "disable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, database, pw, sslmodetext)
}