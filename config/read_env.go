package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	BACKEND_PORT := os.Getenv("BACKEND_PORT")
	fmt.Println(BACKEND_PORT)
}
