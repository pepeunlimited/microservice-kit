package sqlz

import "fmt"

const (
	MysqlRootPassword = "MYSQL_ROOT_PASSWORD"
	MysqlDatabase     = "MYSQL_DATABASE"
	MysqlUser         = "MYSQL_USER"
	MysqlHost         = "MYSQL_HOST"
	MysqlPort         = "MYSQL_PORT"
)

//db_user:password@tcp(localhost:3306)/my_db
//root:root@(localhost:3306)/shopping_list
//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
func MySQLURI(username string, pw string, host string, port int, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", username, pw, host, port, database)
}
