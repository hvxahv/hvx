package notify

import (
	"fmt"
	"os"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	n := cockroach.NewDBAddr()
	if err := n.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}

}

func TestNotify_Create(t *testing.T) {
	if err := NewNotifies(2131312412213312312,
		"https://sg2p.notify.windows.com/w/?token=BQYAAAB6d7igIw1dP7TMQGZTFsdIA9uBWPaiv9od2NtnpO3Oh8wNaFUIC0Yp5yVZFv3VmCea7CTp205cGlaNaHsyb3AAltWOj8SEJ7liLwOm4o2n0bIHn0bsUFdTv3a8i6AWVciis1pIAYAvT%2baEPYkoFWUXbvtpUiSDUy0THeB0QUDMimWZvYfOXkDFbNraKMJAc0JCS0MF8oZRcxMwyH6o6oKOmKFjBp2E8NK%2b2LCUyZU7Xc6wdYTCLAD%2fvQ7K%2bq8za4%2bUAyOmfw9okKHy9%2fcYgYWlypYQ%2fkmqiEfne3Y2dF0z2P6e4YLTcWaxn9lL49cKkdA%3d",
		"BPIKfLEeyz-DpVTLt3oQiqunnBvkbRFr500uEc9hMX52ABKcnDUyQzsnzQjaejXFVVXMhj4wC50WbcDoqVirQJ0",
		"P-8DQm_Gwdit1RApgigAUA",
		"BDWxbKVZjt8vWyO9fcBK8mt-48GFf_iE2wvVrU7ubW5hOHLTkpdhpIud6uUSsvK4H5xqOgoM-lNu9hM5DbN8aAQ",
		"FK7dy8kslbJcOVx4VZ3NH-GgN7LBOOYgj0eXm3r11ho",
	).Create(); err != nil {
		fmt.Println(err)
	}
}

func TestNotify_Get(t *testing.T) {
	r, err := NewNotifiesByDeviceID(2131312412213312312).Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Auth)
}
