package pkg

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
)

// CreateDB Create a database.
// An error will be returned if the creation fails,
// and a successful creation message will be output if the creation is successful.
func CreateDB(name string) {
	nd :=  cockroach.NewDb()
	if err := nd.Create(name); err != nil {
		fmt.Printf("Failed to initialize cockroach : %s", err)
	} else {
		fmt.Printf("Create %s database successfully.", name)
	}
}


