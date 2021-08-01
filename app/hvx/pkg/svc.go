package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
)

// NewServices Create a microservices.
func NewServices(name string) {
	// author
	a := viper.GetString("author")

	// package name and directory name
	pn := fmt.Sprintf("github.com/disism/hvxahv/app/%s", name)
	dn := fmt.Sprintf("../%s", name)

	// Execute the command to create the microservices and return the standard output.
	cmd := exec.Command("cobra", "init", "--pkg-name", pn, dn, "-a", a)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf(string(out))
}
