package misc

import (
	"log"
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		log.Printf("set key=[%v] environment variable=[%v]", key, value)
		return value
	}
	log.Printf("WARNING: invoked fallback for key=[%v] environment variable=[%v]",key, fallback)
	return fallback
}

func Int64Env(key string, fallback int64) int64 {
	env := GetEnv(key, strconv.FormatInt(fallback, 10))
	parsed, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		log.Panic("overflow int64 for env")
	}
	return parsed
}