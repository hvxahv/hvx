package activity

import (
	"fmt"
	"testing"
)

func TestMentions_Create(t *testing.T) {
	IniTestConfig(t)

	m := NewMentions("https://mas.to/users/hvturingga/statuses/107330783584704387", 710444110081654785, 698619813575491585, 713252633895862273)
	if err := m.Create(); err != nil {
		fmt.Println(err)
	}
}
