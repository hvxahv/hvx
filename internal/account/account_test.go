package account

import (
	"fmt"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

func TestAccount_IsExist(t *testing.T) {
	// Output: false if Exist in database.
	d := &pb.IsExistRequest{
		Username: "hvturingga",
	}
	s := &server{}
	a, err := s.IsExist(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a.IsExist)

	// Output: true if not found in database.
	d2 := &pb.IsExistRequest{
		Username: "isNotExist",
	}
	s2 := &server{}
	a2, err := s2.IsExist(context.Background(), d2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a2.IsExist)
}

func TestAccount_Create(t *testing.T) {
	d := &pb.CreateAccountRequest{
		Username:  "hvxahv",
		Mail:      "hvxahv@halfmemories.com",
		Password:  "hvxahv123",
		PublicKey: "public_key",
	}
	s := &server{}
	a, err := s.Create(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestAccount_GetAccountByUsername(t *testing.T) {
	d := &pb.GetAccountByUsernameRequest{
		Username: "hvturingga",
	}
	s := &server{}
	a, err := s.GetAccountByUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestAccount_Delete(t *testing.T) {
	d := &pb.DeleteAccountRequest{
		Username: "hvxahv2",
		Password: "hvxahv123",
	}
	s := &server{}
	a, err := s.DeleteAccount(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestAccount_EditUsername(t *testing.T) {
	d := &pb.EditUsernameRequest{
		Id:       "737973421798785025",
		Username: "hvxahv2",
	}
	s := &server{}
	a, err := s.EditUsername(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestAccount_EditPassword(t *testing.T) {
	d := &pb.EditPasswordRequest{
		Username: "hvxahv2",
		Password: "hvxahv123",
		New:      "hvxahv1234",
	}
	s := &server{}
	a, err := s.EditPassword(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestAccount_EditEmail(t *testing.T) {
	d := &pb.EditEmailRequest{
		Id:   "737973421798785025",
		Mail: "hvxahv2@halfmemories.com",
	}

	s := &server{}

	a, err := s.EditEmail(context.Background(), d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}
