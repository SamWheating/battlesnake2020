package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/SamWheating/battlesnake2020/pkg/simple_moves"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
	"net/http"
)

// These are used to determine the appearance of the snake
var COLOUR = "#517550"
var HEADTYPE = "beluga"
var TAILTYPE = "round-bum"

func Start(w http.ResponseWriter, r *http.Request) {

	resp := structs.StartResponse{}
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

func End(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK.")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK.")
}

func Index(w http.ResponseWriter, r *http.Request) {
	url := "https://www.google.com/search?q=snake&source=lnms&tbm=isch"
	http.Redirect(w, r, url, 302)
}

func Move(w http.ResponseWriter, r *http.Request) {

	var body structs.MoveRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call external move function (swap this out for different algs)
	move := simple_moves.PlayItSafe(body)
	//move := simple_moves.Greedy(body)
	//move := simple_moves.FollowTail(body)

	resp := structs.MoveResponse{}
	resp.Move = move
	resp.Shout = "I am snek"

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
