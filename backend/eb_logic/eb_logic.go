package eb_logic

import "fmt"

type Player struct {
	Name  string
	Score int
}

type Board struct {
	ID      int
	Name    string
	Players []Player
}

func HelloBoard() {
	fmt.Println("Initializing EB logic...")
}

var predefinedPlayers = []Player{
	{Name: "Alice", Score: 0},
	{Name: "Bob", Score: 0},
	{Name: "Charlie", Score: 0},
}

var boards = []Board{
	{Name: "My Board", Players: predefinedPlayers, ID: 1},
}

func GetPlayers(boardID int) []Player {
	for _, board := range boards {
		if board.ID == boardID {
			return board.Players
		}
	}
	fmt.Println("Error: boardID not found")
	return nil
}

func AddBoard(name string, players []Player) {
	newBoard := Board{
		ID:      len(boards) + 1,
		Name:    name,
		Players: players,
	}
	boards = append(boards, newBoard)
}

func GetBoards() []Board {
	return boards
}
