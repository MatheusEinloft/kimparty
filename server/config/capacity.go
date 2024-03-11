package config

import (
	"log"
	"os"
	"strconv"
)

func GetPartyMaxCapacity() uint8 {
    envCapacity := os.Getenv("KIMPARTY_PARTY_MAX_CAPACITY")
    capacity, err := strconv.ParseUint(envCapacity, 10, 8)

    if err != nil {
        log.Println("KIMPARTY_PARTY_MAX_CAPACITY is not set, using default party max capacity 2")
        return 2
    }

    return uint8(capacity)
}
