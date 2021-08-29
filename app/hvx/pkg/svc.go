package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os/exec"
)

// NewServices Use this function to create a service.
func NewServices(name string) error {
	a := viper.GetString("author")
	n := viper.GetString("name")

	// package name and directory name
	pn := fmt.Sprintf("%s/app/%s", n, name)
	dn := fmt.Sprintf("../%s", name)

	// Execute the command to create the microservices and return the standard output.
	cmd := exec.Command("cobra", "init", "--pkg-name", pn, dn, "-a", a)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Errorf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf(string(out))
	return nil
}
