package misc

import (
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