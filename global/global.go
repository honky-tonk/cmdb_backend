package global

import "github.com/joho/godotenv"

var (
	Env_file map[string]string
)

func Init() {
	Env_file, _ = godotenv.Read(".env")
}
