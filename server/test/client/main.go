package main

import (
	"kimparty/internal/party"
	"sync"

	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/json-iterator/go"
)

var baseUrl string = "://localhost:"
var baseHttpUrl string
var baseWsUrl string

var port int
var parties int
var users int
var logMsg bool

func init() {
	flag.IntVar(&port, "port", 3000, "Port on which server is running")
	flag.IntVar(&parties, "parties", 1, "Parties to create")
	flag.IntVar(&users, "users", 2, "Users to create")
    flag.BoolVar(&logMsg, "log-msg", true, "Log messages received from the connections")

	flag.Parse()

	baseUrl = fmt.Sprintf("%s%d", baseUrl, port)
	baseHttpUrl = "http" + baseUrl
	baseWsUrl = "ws" + baseUrl

	log.Printf("BASE URL: '%s'", baseUrl)
	log.Printf("BASE HTTP URL: '%s'", baseHttpUrl)
	log.Printf("BASE WS URL: '%s'", baseWsUrl)

	log.Printf("Parties: %d", parties)
	log.Printf("Users: %d", users)
    log.Printf("Log messages: %v", logMsg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(parties * users)

	for i := 0; i < parties; i++ {
		pt := createRoom(fmt.Sprintf("%d", i))
		log.Printf("Party [%v - %s] created", i, pt.ID)

		for j := 0; j < users; j++ {
			username := fmt.Sprintf("test-%v-%v",i, j)

			go func(roomID, username string) {
				defer wg.Done()
				connectToParty(roomID, username)
			}(pt.ID, username)
		}

	}

	wg.Wait()
}

func createRoom(partyURL string) *party.Party {
	url := fmt.Sprintf("%s/party?url=%s", baseHttpUrl, partyURL)
	res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))

	if err != nil {
		log.Panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Panic(err)
	}

	var pt party.Party
	err = jsoniter.ConfigFastest.Unmarshal(body, &pt)

	if err != nil {
		log.Panic(err)
	}

	return &pt
}

func connectToParty(ptID string, username string) {
	url := fmt.Sprintf("%s/ws/join?party_id=%s&username=%s", baseWsUrl, ptID, username)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Panicf("User [%s] failed to connect to server:%s", username, err)
		return
	}

	log.Printf("User [%s] connected", username)
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("User [%s] error reading message: %s", username, string(message))
				return
			}
        
            if string(message) == "pong" {
                continue
            }

            if logMsg {
                log.Printf("User [%s] received message: %s", username, string(message))
            }
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				log.Panicf("Error sending message: %v", err)
				return
			}
		}
	}
}
