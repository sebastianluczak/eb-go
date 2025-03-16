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

type Server struct {
	manager *eb_logic.BoardManager
}

func NewServer(manager *eb_logic.BoardManager) *Server {
	return &Server{manager: manager}
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *Server) getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}

func (s *Server) GetPlayers(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	boardID := r.URL.Query().Get("board")
	if boardID == "" {
		http.Error(w, "Missing board parameter", http.StatusBadRequest)
		return
	}

	players, err := s.manager.GetPlayers(boardID)
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

func (s *Server) AddBoard(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	boardID := uuid.New().String()[:8]
	s.manager.AddBoard(boardID, []eb_logic.Player{})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func (s *Server) GetBoards(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	boards, err := s.manager.GetBoards()
	if err != nil {
		http.Error(w, "Failed to get boards", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boards)
}

func (s *Server) AddPlayer(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	boardID := r.URL.Query().Get("board")
	if boardID == "" {
		http.Error(w, "Missing board parameter", http.StatusBadRequest)
		return
	}

	playerName := r.URL.Query().Get("player")
	if playerName == "" {
		http.Error(w, "Missing player name parameter", http.StatusBadRequest)
		return
	}

	s.manager.AddPlayer(boardID, playerName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func main() {
	eb_logic.HelloBoard()

	manager := eb_logic.NewBoardManager()
	server := NewServer(manager)

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.getHello)
	mux.HandleFunc("/getPlayers", server.GetPlayers)
	mux.HandleFunc("/addBoard", server.AddBoard)
	mux.HandleFunc("/getBoards", server.GetBoards)
	mux.HandleFunc("/addPlayer", server.AddPlayer)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
