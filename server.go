package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var counter = 1

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Rating   int    `json:"rating"`
	Club     string `json:"club"`
}

func newPlayer(id int, name string, position string, rating int, club string) Player {
	return Player{Id: id, Name: name, Position: position, Rating: rating, Club: club}
}

var players []Player

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(players)

}

func CreatePlayers(w http.ResponseWriter, r *http.Request) {
	var player Player
	json.NewDecoder(r.Body).Decode(&player)
	player.Id = counter
	counter++
	players = append(players, player)
	json.NewEncoder(w).Encode(player)
}

func GetPlayerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for _, item := range players {

		if item.Id == int(id) {
			player = item
			break
		}
	}
	json.NewEncoder(w).Encode(player)
}

func DeletePlayerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for index, item := range players {

		if item.Id == int(id) {
			players = append(players[:index], players[index+1:]...)
			break
		}
	}
	fmt.Fprintf(w, "Deleted successfully")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/player", GetPlayers).Methods("GET")
	router.HandleFunc("/player", CreatePlayers).Methods("POST")
	router.HandleFunc("/player/{id}", GetPlayerById).Methods("GET")
	router.HandleFunc("/player/{id}", DeletePlayerById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8090", router))

}
