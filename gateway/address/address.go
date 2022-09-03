package address

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	Account  string = "account"
	Actor    string = "actor"
	Public   string = "public"
	Auth     string = "auth"
	Device   string = "device"
	Channel  string = "channel"
	Article  string = "article"
	Saved    string = "saved"
	Activity string = "activity"
	Message  string = "message"
)

func GetHTTP(svcName string) string {
	hostname := viper.GetString(fmt.Sprintf("microsvcs.%s.hostname", svcName))
	port := viper.GetString(fmt.Sprintf("microsvcs.%s.ports.http", svcName))
	if viper.GetBool(fmt.Sprintf("microsvcs.%s.ports.useSSL", svcName)) {
		return fmt.Sprintf("https://%s:%s", hostname, port)
	}
	return fmt.Sprintf("http://%s:%s", hostname, port)
}
