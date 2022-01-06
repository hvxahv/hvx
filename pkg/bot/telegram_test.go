package bot

import "testing"

func TestBot_Send(t *testing.T) {
	b := NewBot(1, "HEELO")
	if err := b.Send(); err != nil {
		return
	}
}

func TestBot_GetUpdateId(t *testing.T) {
	b := NewBot(1, "HEELO")
	err := b.GetUpdateId()
	if err != nil {
		return
	}
}
