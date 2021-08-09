package handlers

import (
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

//// WebFinger and WebFingerLinks form the JSON-LD protocol of the standard Activitypub
//type WebFinger struct {
//	Subject string           `json:"subject"`
//	Links   []WebFingerLinks `json:"links"`
//}
//
//// WebFingerLinks is used by WebFinger
//type WebFingerLinks struct {
//	Rel  string `json:"rel"`
//	Type string `json:"type"`
//	Href string `json:"href"`
//}

func WebFingerHandler(c *gin.Context) {
	res := c.Query("resource")

	// If you are not searching for the user of this instance, go to the remote request.
	//host := activitypub.GetHost(res)
	//if host != viper.GetString("localhost") {
	//	url := fmt.Sprintf("https://%s/.well-known/webfinger?resource=%s", host, res)
	//	method := "GET"
	//	client := &http.Client {}
	//	req, err := http.NewRequest(method, url, nil)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	r, err := client.Do(req)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	defer r.Body.Close()
	//
	//	body, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	var res activitypub.WebFingerData
	//	_ = json.Unmarshal(body,&res)
	//
	//	c.JSON(200, res)
	//	return
	//}


	// Perform some filtering operations from the request to obtain the user name,
	// and then search for the user name to find whether the user exists in the database.
	// Currently only tested mastodon has not supported other ActivityPub implementations.

	name, err := activitypub.IsActorExists(res)
	if err != nil {
		log.Println(err)
		return
	}

	// Use this client to call the remote Accounts gRPC service,
	// and then pass the user name to get the queried data.
	cli, conn,  err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	accounts, err := cli.FindAccount(context.Background(), &pb.AccountByName{Username: name})
	if err != nil {
		return
	}

	c.JSON(200, activitypub.NewWebFinger(accounts.Username))
}

