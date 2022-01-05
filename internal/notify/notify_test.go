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
	if err := NewNotifies(725154968683708417,
		"https://fcm.googleapis.com/fcm/send/fFXw2lgfA-U:APA91bH4UHigzqa2vpIHJ446QIJepIIn4u_Y53cLABwUTsEFrhDR1BAvN_t0-ApOM5CzQTVFUfNvEY0TzCSgx7xedsR-eFizI4OucFv-E1k24KISzRCQsQOVRN7nWYurxlrauS_h5sC-",
		"BLA7yzeyG9aaqRgURqHbRPcYVX4LDyii6VkOFl8f80Wpjxt69INhqd0xflC8LZDFNZP0ze1F-4VLTEGm84bI9nc",
		"gJwr2dSCZqORdWLajmhg1g",
	).Create(); err != nil {
		fmt.Println(err)
	}
}

func TestNotify_Get(t *testing.T) {
	r, err := NewNotifiesByDeviceID(725154968683708417).Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
