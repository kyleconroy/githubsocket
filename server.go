package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var port *int = flag.Int("p", 23456, "Port to listen.")
var user *string = flag.String("user", "", "Github username.")
var pass *string = flag.String("pass", "", "Github password.")

var dispatch map[int64]chan int = make(map[int64]chan int)

type Sub struct {
	Mode  string
	Topic string
}

// jsonServer echoes back json string sent from client using websocket.JSON.
func hubServer(ws *websocket.Conn) {
	id := rand.Int63()

	dispatch[id] = make(chan int)
	defer delete(dispatch, id)

	fmt.Printf("listern %d\n", id)

	go func() {
		for {
			// Send send a text message serialized T as JSON.
			event := <-dispatch[id]

			err := websocket.JSON.Send(ws, event)

			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Print(".")
		}
	}()

	for {
		var msg Sub

		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &msg)

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("callback: http://incoming?socket=%d\n", id)
	}

	fmt.Printf("close")
}

func HookListener(w http.ResponseWriter, req *http.Request) {
	socket := req.URL.Query().Get("socket")
	id, err := strconv.ParseInt(socket, 10, 64)

	if err != nil {
		fmt.Fprint(w, "Bad socket id")
		return
	}

	go func() {
		dispatch[id] <- 45
	}()

	fmt.Fprintf(w, "Socket %s contacted", socket)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.Parse()

	http.Handle("/hub", websocket.Handler(hubServer))
	http.Handle("/", http.FileServer(http.Dir("")))
	http.HandleFunc("/incoming", HookListener)

	fmt.Printf(" * http://localhost:%d/\n", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)

	if err != nil {
		panic("ListenANdServe: " + err.Error())
	}
}
