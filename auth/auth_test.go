package auth

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hvxahv/hvx/cockroach"
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
	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}
}

func TestClaims_JWTTokenGenerator(t *testing.T) {
	var (
		issuer = viper.GetString("domain")
		expir  = time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
		secret = viper.GetString("authentication.token.secret")
	)
	u := NewUserdata(uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String())
	c := NewRegisteredClaims(issuer, "", "", expir)
	generator, err := NewClaims(u, c).JWTTokenGenerator(secret)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(generator)
}

func TestClaims_JWTTokenParse(t *testing.T) {
	var (
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyZGF0YSI6eyJhY2NvdW50X2lkIjoiOTJiNGRjMjUtN2FmNi00MDk1LWFjNmYtOTRlYmZiNDc4ZmE5IiwiYWN0b3JfaWQiOiI4MjNhZGM3NS1hMjYxLTRhOTEtYWY3Zi05NzIwYjMyY2Y2MjQiLCJkZXZpY2VfaWQiOiIwMjY4ZTMyNi0xZDI5LTQ2NmUtOTQ0Zi02YzdiMTY4NjQ0ZTgiLCJ1c2VybmFtZSI6IjE0NGY5MDYyLTMxOTAtNDczNS1iMDcxLTVkZDNjNmE3MjJkZiIsIm1haWwiOiI1M2JlMTMzNS1kNWI5LTRiZjUtYmIyYi05OTllZmU5YmZlOWYifSwiand0Xy5fc3RhbmRhcmRfY2xhaW1zIjp7ImlzcyI6Imh2eGFodi5oYWxmbWVtb3JpZXMuY29tIiwiYXVkIjpbIiJdLCJleHAiOjE2NjQ3NTk4NTcsIm5iZiI6MTY1OTU3NTg1NywiaWF0IjoxNjU5NTc1ODU3LCJqdGkiOiIxNDZhNGFiZS02ZGNmLTQ0YmEtOTExZC1kZTQ5MmE1MTViOGMifX0.kMvNzJDK8f3VD7omel6sZ4sI43DZQ1ArbAUZ0TesUyc"
		secret = viper.GetString("authentication.token.secret")
	)
	parse, err := NewParseJWTToken(token, secret).JWTTokenParse()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(parse)
}
