package errors

import (
	"fmt"
	"strings"
)

func NewDatabaseCreate(tableName string) error {
	return fmt.Errorf("FAILED_TO_AUTOMATICALLY_CREATE_%s_DATABASE", strings.ToUpper(tableName))
}
