package eb_logic

import (
	"errors"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type BoardManager struct {
	mu     sync.RWMutex
	boards map[string]*Board
}

func NewBoardManager() *BoardManager {
	bm := &BoardManager{
		boards: make(map[string]*Board),
	}

	initialBoard := &Board{
		ID:   "470b91b9-76ed-48e7-bfbd-0de71395a40a",
		Name: "First Board",
		Players: []Player{
			{Name: "Alice", Score: 0},
			{Name: "Bob", Score: 0},
			{Name: "Charlie", Score: 0},
		},
	}
	bm.boards[initialBoard.ID] = initialBoard

	return bm
}

func (bm *BoardManager) GetBoards() ([]Board, error) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()

	boards := make([]Board, 0, len(bm.boards))
	for _, board := range bm.boards {
		boardCopy := *board
		boards = append(boards, boardCopy)
	}
	sort.Slice(boards, func(i, j int) bool {
		return boards[i].Name < boards[j].Name
	})
	return boards, nil
}

func (bm *BoardManager) GetBoard(id string) (Board, error) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()

	board, exists := bm.boards[id]
	if !exists {
		return Board{}, errors.New("board not found")
	}
	return *board, nil
}

func (bm *BoardManager) AddBoard(name string, players []Player) (Board, error) {
	if name == "" {
		return Board{}, errors.New("board name cannot be empty")
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	playersCopy := make([]Player, len(players))
	copy(playersCopy, players)

	newBoard := &Board{
		ID:      uuid.New().String(),
		Name:    name,
		Players: playersCopy,
	}

	bm.boards[newBoard.ID] = newBoard
	return *newBoard, nil
}

func (bm *BoardManager) GetPlayers(id string) ([]Player, error) {
	board, err := bm.GetBoard(id)
	if err != nil {
		return nil, err
	}

	playersCopy := make([]Player, len(board.Players))
	copy(playersCopy, board.Players)
	return playersCopy, nil
}

func (bm *BoardManager) AddPlayer(id string, name string) ([]Player, error) {
	if name == "" {
		return nil, errors.New("player name cannot be empty")
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	board, exists := bm.boards[id]
	if !exists {
		return nil, errors.New("Board not found")
	}

	board.Players = append(board.Players, Player{Name: name, Score: 0})

	return board.Players, nil
}
