package misc

import (
	"fmt"
	"log"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		log.Printf("using env=[%v]", value)
		return value
	}
	log.Printf("using fallback env=[%v]", fallback)
	return fallback
}

//db_user:password@tcp(localhost:3306)/my_db
//root:root@(localhost:3306)/shopping_list
//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
func MySQLURI(username string, pw string, host string, port int, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", username, pw, host, port, database)
}
