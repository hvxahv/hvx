package maria

import (
	"fmt"
	"testing"
)


func TestGetMaria(t *testing.T) {
	if err := InitMariaDB(); err != nil {
		t.Log(err)
	}
	db := GetMaria()
	fmt.Println(db)
}

