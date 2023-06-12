package config

import "os"

const (
	env_PORT       = "PORT"
	env_USERNAME   = "DB_USERNAME"
	env_PASSWORD   = "DB_PASSWORD"
	env_URL        = "JULODB_URL"
	env_JWT_SECRET = "JWT_SECRET"
)

func getEnv(key string) string {
	val := os.Getenv(key)
	return val
}

func Username() string {
	return getEnv(env_USERNAME)
}
func Password() string {
	return getEnv(env_PASSWORD)
}
func DBURL() string {
	return getEnv(env_URL)
}
func Port() string {
	return getEnv(env_PORT)
}
func JwtSecret() string {
	return getEnv(env_JWT_SECRET)
}
