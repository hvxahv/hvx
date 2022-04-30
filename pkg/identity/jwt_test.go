package identity

import (
	"fmt"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
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
	if err := cockroach.NewDBAddr().InitDB(); err != nil {
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

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVdWlkIjoiIiwiRW1haWwiOiJ4QGRpc2lzbS5jb20iLCJJRCI6IjczMzEyNDY4MDYzNjU5NjIyNSIsIlVzZXJuYW1lIjoiaHZ0dXJpbmdnYSIsIkRldmljZXNIYXNoIjoiZDlmOGE2ZjctZmY0Yi00ZDU1LWE5NDUtYzViODAxNjk5ZTM4IiwiZXhwIjoxNjUyODY0NDQ3LCJpYXQiOjE2NDc2ODA0NDcsImlzcyI6ImhhbGZtZW1vcmllcy5jb20iLCJzdWIiOiJ0b2tlbiJ9.2oFGQVv3KczKrnsRnJ7sfb5c1XEKXSmBR9wCQBM_zhA"

func TestGenToken(t *testing.T) {
	token, err := GenToken("foo", "bar", "123", "hash")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestVerifyToken(t *testing.T) {
	token, err := VerifyToken(token)
	if err != nil {
		t.Errorf("token parsing error: %v", err)
	}
	t.Log(token)

}

func TestParseToken(t *testing.T) {
	token, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}
