package errors

import (
	"log"
)

type E struct {
	commentary string
	error      error
}

// Sends A middle tier designed to send errors to an error server or error collection middleware.
type Sends interface {
	Do() error
}

func Throw(commentary string, error error) {
	x := &E{
		commentary: commentary,
		error:      error,
	}
	if err := x.Do(); err != nil {
		log.Printf("failed to send error message: %v", err)
		return
	}
}

func (e *E) Do() error {
	// TODO - This method will be implemented to send errors to the error collection server.
	log.Println(e.commentary, e.error)
	return nil
}
