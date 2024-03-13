package config

import (
	"log"
	"os"
	"strconv"
)

func GetPort() uint16 {
	envPort := os.Getenv("KIMPARTY_PORT")
	port, err := strconv.ParseInt(envPort, 10, 64)

	if err != nil {
		log.Println("KIMPARTY_PORT is not set, using default port 3000")
		port = 3000
	}

	return uint16(port)
}
