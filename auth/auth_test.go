package auth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
	"time"
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
	if err := cockroach.NewRoach().Dial(); err != nil {
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

func TestClaims_JWTTokenGenerator(t *testing.T) {
	var (
		issuer = viper.GetString("domain")
		expir  = time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
		signer = viper.GetString("authentication.token.signed")
	)
	u := NewUserdata(uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String())
	c := NewRegisteredClaims(issuer, "", "", expir)
	generator, err := NewClaims(u, c).JWTTokenGenerator(signer)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(generator)
}
