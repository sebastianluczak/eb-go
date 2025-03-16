package main

import (
	"eb_logic"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayers(t *testing.T) {
	// Mock data
	boardID := "test_board"
	mockPlayers := []eb_logic.Player{
		{Name: "Alice"},
		{Name: "Bob"},
	}
	manager := eb_logic.NewBoardManager()
	actualBoard, _ := manager.AddBoard(boardID, mockPlayers)

	server := NewServer(manager)

	req := httptest.NewRequest("GET", "/getPlayers?board="+actualBoard.ID, nil)
	w := httptest.NewRecorder()

	server.GetPlayers(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
