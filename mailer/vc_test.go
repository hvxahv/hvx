package mailer

import (
	"fmt"
	"testing"
)

func TestValidateCode_Generate(t *testing.T) {
	fmt.Println(ValidateCodeGenerator(6))
}
