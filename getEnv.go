package main

import (
	"fmt"
	"github.com/joho/godotenv"
	stdlog "log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		stdlog.Fatal("env file open fail")
	}

	name := os.Getenv("NAME")
	fmt.Println(name)
}
