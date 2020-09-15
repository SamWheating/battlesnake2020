package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/SamWheating/battlesnake2020/pkg/lookahead"
	"github.com/SamWheating/battlesnake2020/pkg/structs"
)

// return a random 24-bit hex colour like #A1B514
func getRandomHex() string {
	const hexchars = "1234567890ABCDEF"
	b := make([]byte, 7)
	for i := range b {
		b[i] = hexchars[rand.Intn(len(hexchars))]
	}
	b[0] = '#'
	return string(b)
}

func Start(w http.ResponseWriter, r *http.Request) {
	HEADTYPES := []string{"beluga", "safe"}
	TAILTYPES := []string{"round-bum", "curled"}
	resp := structs.StartResponse{}
	// These are used to determine the appearance of the snake
	resp.Headtype = HEADTYPES[rand.Intn(len(HEADTYPES))]
	resp.Tailtype = TAILTYPES[rand.Intn(len(TAILTYPES))]
	resp.Color = getRandomHex()

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

	depthArg := r.URL.Query().Get("depth")
	countArg := r.URL.Query().Get("count")
	if depthArg == "" {
		depthArg = "10"
	}
	if countArg == "" {
		countArg = "10000"
	}
	depth, _ := strconv.Atoi(depthArg)
	count, _ := strconv.Atoi(countArg)

	// call external move function (swap this out for different algs)
	//move := simple_moves.PlayItSafe(body)
	//move := simple_moves.PlayItSafeFlood(body)
	move := lookahead.Lookahead(body, depth, count)
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
