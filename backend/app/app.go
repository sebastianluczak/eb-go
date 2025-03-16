package main

import (
	"eb_logic"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

var manager = eb_logic.NewBoardManager()

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}

func GetPlayers(w http.ResponseWriter, r *http.Request, boardID string) {
	players, err := manager.GetPlayers(boardID)
	if err != nil {
		http.Error(w, "Failed to get players", http.StatusInternalServerError)
		return
	}
	playerNames := make([]string, len(players))
	for i, player := range players {
		playerNames[i] = player.Name
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerNames)
}

func AddBoard(w http.ResponseWriter, r *http.Request) {
	randomString := uuid.New().String()[:8]
	manager.AddBoard(randomString, []eb_logic.Player{})
	json.NewEncoder(w).Encode("OK")
}

func GetBoards(w http.ResponseWriter, r *http.Request) {
	boards, _ := manager.GetBoards()
	boardNames := make([]string, len(boards))
	for i, board := range boards {
		boardNames[i] = board.Name
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boards)
}

func main() {
	eb_logic.HelloBoard()

	// Routes definition
	http.HandleFunc("/", getHello)
	http.HandleFunc("/getPlayers", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		boardIDStr := r.URL.Query().Get("board")
		if boardIDStr == "" {
			http.Error(w, "Missing board parameter", http.StatusBadRequest)
			return
		}

		GetPlayers(w, r, boardIDStr)
	})
	http.HandleFunc("/addBoard", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		AddBoard(w, r)
	})
	http.HandleFunc("/getBoards", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		GetBoards(w, r)
	})

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
