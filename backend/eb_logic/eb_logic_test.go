package eb_logic

import "testing"

func TestHelloBoard(t *testing.T) {
	result := HelloBoard()

	if result != 1 {
		t.Errorf("HelloWorld returned incorrect code. Expected 1, got %d", result)
	}
}
