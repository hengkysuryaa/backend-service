package main

import (
	"github.com/hengkysuryaa/backend-service/fetch-app/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	godotenv.Load()

	// execute command here
	cmd.RunRest()
}
