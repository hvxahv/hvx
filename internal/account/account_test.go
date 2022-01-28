package account

import (
	"fmt"
	"github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"golang.org/x/net/context"
	"os"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/cache"
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

	cache.InitRedis(1)
}

func TestAccount_IsExist(t *testing.T) {
	// Output: false if Exist in database.
	d := &v1alpha1.NewAccountUsername{
		Username: "hvturingga",
	}
	s := &account{}
	a, err := s.IsExist(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a.IsExist)

	// Output: true if not found in database.
	d2 := &v1alpha1.NewAccountUsername{
		Username: "isNotExist",
	}
	s2 := &account{}
	a2, err := s2.IsExist(context.Background(), d2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a2.IsExist)
}

func TestAccount_Create(t *testing.T) {
	d := &v1alpha1.NewAccountCreate{
		Username:  "hvturingga",
		Mail:      "hvturingga@disism.com",
		Password:  "hvxahv123",
		PublicKey: "p",
	}
	s := &account{}
	create, err := s.Create(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(create)
}

func TestAccount_Delete(t *testing.T) {
	d := &v1alpha1.NewAccountDelete{
		Username: "hvxahv",
		Password: "hvxahv123",
	}
	s := &account{}
	r, err := s.Delete(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Code, r.Reply)
}

func TestAccount_EditUsername(t *testing.T) {
	d := &v1alpha1.NewEditAccountUsername{
		Id:       "731607090811043841",
		Username: "hvturingga",
	}
	s := &account{}
	r, err := s.EditUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Code, r.Reply)
}

func TestAccount_EditPassword(t *testing.T) {
	d := &v1alpha1.NewEditAccountPassword{
		Id:       "731607090811043841",
		Password: "Hvxahv123",
	}
	s := &account{}
	r, err := s.EditPassword(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Code, r.Reply)
}

func TestAccount_EditMail(t *testing.T) {
	d := &v1alpha1.NewEditAccountMail{
		Id:   "731607090811043841",
		Mail: "x@disism.com",
	}
	s := &account{}
	r, err := s.EditMail(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r.Code, r.Reply)
}
func TestAccount_GetAccountByUsername(t *testing.T) {
	d := &v1alpha1.NewAccountUsername{
		Username: "hvturingga",
	}
	s := &account{}
	a, err := s.GetAccountByUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}
