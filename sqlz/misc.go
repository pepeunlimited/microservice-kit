package sqlz

import "fmt"

//db_user:password@tcp(localhost:3306)/my_db
//root:root@(localhost:3306)/shopping_list
//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
func MySQLURI(username string, pw string, host string, port int, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", username, pw, host, port, database)
}
