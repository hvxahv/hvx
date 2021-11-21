package pkg

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
)

// NewDB Create a database.
// An error will be returned if the creation fails,
// and a successful creation message will be output if the creation is successful.
func NewDB(name string) error {
	n := cockroach.NewDBAddr()
	if err := n.New(name); err != nil {
		return errors.Errorf("failed to initialize cockroach : %s", err)
	}

	fmt.Printf("create %s database successfully.", name)
	return nil
}


