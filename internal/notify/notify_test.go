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
	if err := NewNotifies("xxs",
		"https://sg2p.notify.windows.com/w/?token=BQYAAACdRald%2bxWeGHCX0f0vz9EI2J7l6T%2baXY4EejzsoreeDlzmhckR0vdSj%2fP%2fl6KUOocND5Tw6LZ8C5myu5ACPxiwr732KPiyHSFIfrt2TpJuSLFF4QoGPfNB5PuSRSHBcGJAKw7sASDNjbX79L7%2fa8iqDyf3S6vm51RC3WHwBaxTJ4vpBF5u0Uyho8SiBXEiTOnc6chAbKuh0%2bV3LniIsQ0lUrp6kCLCBrSS3sDdwr9Y5CYrnIfcvQ7bER2GBnrMFG7E7yJw3TWzVoFUpd5gyCSi1yPnpLUzBkB2%2fz5FHglCJj5Hqv27JDOPGkR9dZBnYME%3d",
		"BEMYIT01WoBjyPWn035CZS5LA5hOHkXd7I42J77X2-syd9LUtVEfhuW7TbIlfSveLEUNjgWQszUeyBH0cwC_kh8",
		"P-Hl9ObhZFViJ1sEaNyPprkw",
		// "BDWxbKVZjt8vWyO9fcBK8mt-48GFf_iE2wvVrU7ubW5hOHLTkpdhpIud6uUSsvK4H5xqOgoM-lNu9hM5DbN8aAQ",
		// "FK7dy8kslbJcOVx4VZ3NH-GgN7LBOOYgj0eXm3r11ho",
	).Create(); err != nil {
		fmt.Println(err)
	}
}

func TestNotify_Get(t *testing.T) {
	r, err := NewNotifiesByDeviceID("").Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
