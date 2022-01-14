package article

import (
	"log"
	"testing"
)

func TestConversations_Create(t *testing.T) {
	c := NewConversations("https://halfmemories.com/status/727770128571465729", 727491252124745729, "https://halfmemories.com/status/727770128571465729", "XXS", 727491252124745729)
	if err := c.Create(); err != nil {
		log.Println(err)
		return
	}
}
