package errors

import (
	"fmt"
	"strings"
)

func NewFailedToConnect(svcName string) error {
	return fmt.Errorf("FAILED_TO_CONNECT_TO_%s_SERVICE", strings.ToUpper(svcName))
}
