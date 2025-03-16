package eb_logic

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type BoardManager struct {
	mu     sync.Mutex
	boards []Board
}

func NewBoardManager() *BoardManager {
	return &BoardManager{
		boards: []Board{
			{ID: "470b91b9-76ed-48e7-bfbd-0de71395a40a", Name: "First Board", Players: []Player{
				{Name: "Alice", Score: 0},
				{Name: "Bob", Score: 0},
				{Name: "Charlie", Score: 0},
			}},
		},
	}
}

func (bm *BoardManager) GetBoards() ([]Board, error) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	return bm.boards, nil
}

func (bm *BoardManager) GetBoard(id string) (*Board, error) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	for _, board := range bm.boards {
		if board.ID == id {
			return &board, nil
		}
	}
	return nil, errors.New("board not found")
}

func (bm *BoardManager) AddBoard(name string, players []Player) (*Board, error) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	newBoard := Board{
		ID:      uuid.New().String(),
		Name:    name,
		Players: players,
	}
	bm.boards = append(bm.boards, newBoard)
	return &newBoard, nil
}

func (bm *BoardManager) GetPlayers(id string) ([]Player, error) {
	board, err := bm.GetBoard(id)
	if err != nil {
		return nil, err
	}
	return board.Players, nil
}
