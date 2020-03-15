package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var COLOUR = "#517550"
var HEADTYPE = "beluga"
var TAILTYPE = "round-bum"

type StartRequest struct {
	Game	string
	Turn	string
	Board	string
	You		string
}

type StartResponse struct {
	Color 		 string
	Headtype     string
	Tailtype     string
}

// This is the request structure from the gameserver -
// We can pass the whole thing to any functions which are
// computing moves.
type MoveRequest struct {
	Game struct {
		ID string `json:"id"`
	} `json:"game"`
	Turn  int `json:"turn"`
	Board struct {
		Height int `json:"height"`
		Width  int `json:"width"`
		Food   []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"food"`
		Snakes []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Health int    `json:"health"`
			Body   []struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"body"`
			Shout string `json:"shout"`
		} `json:"snakes"`
	} `json:"board"`
	You struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Health int    `json:"health"`
		Body   []struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"body"`
		Shout string `json:"shout"`
	} `json:"you"`
}

type MoveResponse struct {
	Move 	string
	Shout 	string
}

type EndRequest struct {
	Game	string
	Turn 	string
	Board 	string
	You 	string
}

func start_handler(w http.ResponseWriter, r *http.Request) {

	resp := StartResponse{}
	resp.Color = COLOUR
	resp.Headtype = HEADTYPE
	resp.Tailtype = TAILTYPE

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)

}

func end_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK.")
}

func ping_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK.")
}

func move_handler(w http.ResponseWriter, r *http.Request) {

	var body MoveRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// call external move function (swap this out for different algs)
	move := goLeft(body)

	resp := MoveResponse{}
	resp.Move = move
	resp.Shout = "I am snek"

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Stats Server started on port %s", port)

	http.HandleFunc("/start", start_handler)
	http.HandleFunc("/end", end_handler)
	http.HandleFunc("/move", move_handler)
	http.HandleFunc("/ping", ping_handler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
