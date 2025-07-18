package api

import (
	"log"
	"net/http"
)

func HandleRollDice(w http.ResponseWriter, r *http.Request) {
	log.Println("Rolling dice test...")
	w.Write([]byte("Dice rolled!"))
}
