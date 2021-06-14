package v1alpha1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/microservice"
	"log"
	"strconv"
)

// NewAccountsHandler ...
func NewAccountsHandler(c *gin.Context) {
	// Username used to log in.
	username := c.PostForm("username")
	// Password for login.
	password := c.PostForm("password")
	// Account avatar.
	avatar := c.PostForm("avatar")
	// User's name, displayed name.
	name := c.PostForm("name")
	// User's email, used to retrieve password.
	email := c.PostForm("email")
	// Choose whether the account is a private account.
	p := c.PostForm("private")
	private, err := strconv.Atoi(p)
	if err != nil {
		log.Println(err)
	}


	addr := "7041"
	nc := microservice.NewClient("Accounts", addr)
	conn, err := nc.NewConn()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	cli := pb.NewAccountsClient(conn)

	r, err := cli.NewAccounts(context.Background(), &pb.NewAccountsData{
		Username: username,
		Password: password,
		Avatar:   avatar,
		Name:     name,
		Email:    email,
		Private: int32(private),
	})
	if err != nil {
		log.Printf("Failed to send message to Accounts server: %v", err)
	}
	fmt.Println(r)
}

// LoginHandler ...
func LoginHandler(c *gin.Context) {
	// Username used to log in.
	username := c.PostForm("username")
	// Password for login.
	password := c.PostForm("password")

	addr := "7041"
	nc := microservice.NewClient("Accounts", addr)
	conn, err := nc.NewConn()
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	cli := pb.NewAccountsClient(conn)

	r, err := cli.LoginAccounts(context.Background(), &pb.AccountsLogin{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Printf("Failed to send message to Accounts server: %v", err)
	}
	fmt.Println(r)
}

