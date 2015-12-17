package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
	games := RepoFindGames()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(games); err != nil {
		panic(err)
	}
}

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var gameId int
	var err error

	if gameId, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}
	game := RepoFindGame(gameId)

	if game.ID == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
	}

}
