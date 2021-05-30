package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/u/:user", ActorHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Actor struct {
	Context           []interface{} `json:"@context"`
	Types             string        `json:"type"`
	Id                string        `json:"id"`
	Following         string        `json:"following"`
	Followers         string        `json:"followers"`
	Liked             string        `json:"liked"`
	Inbox             string        `json:"inbox"`
	Outbox            string        `json:"outbox"`
	PreferredUsername string        `json:"preferredUsername"`
	Name              string        `json:"name"`
	Summary           string        `json:"summary"`
	Icon              []string      `json:"icon"`
}

type Lang struct {
	Language string `json:"@language"`
}

func NewCtx(lang string) []interface{} {
	link := "https://www.w3.org/ns/activitystreams"
	//
	//l := Lang{Language: lang}
	//la := make([]interface{}, 2)
	//la[1] = link
	//la[2] = l

	c := []interface{}{
		link,
		map[string]interface{}{"@language": lang},
	}
	return c

}
func NewActor(context []interface{}, id string, following string, followers string, liked string, inbox string, outbox string, preferredUsername string, name string, summary string, icon []string) *Actor {
	return &Actor{
		Context:           context,
		Types:             "Person",
		Id:                id,
		Following:         following,
		Followers:         followers,
		Liked:             liked,
		Inbox:             inbox,
		Outbox:            outbox,
		PreferredUsername: preferredUsername,
		Name:              name,
		Summary:           summary,
		Icon:              icon,
	}
}

func ActorHandler(c *gin.Context) {
	icon := []string{"https://avatars.githubusercontent.com/u/35920389?v=4"}
	ctx := NewCtx("jp")
	fmt.Println(ctx)

	na := NewActor(ctx, "id",
		"following",
		"followers",
		"liked",
		"inbox",
		"outbox",
		"hvturingga",
		"幸せ",
		"光を進め",
		icon,
	)

	c.JSON(200, gin.H{
		"response": na,
	})

}
