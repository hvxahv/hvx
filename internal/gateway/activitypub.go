package gateway

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func activityPubV1(r *gin.Engine) {
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	//r.GET("/.well-known/webfinger", handlers.WebFingerHandler)

	// https://www.w3.org/TR/activitypub/#actor-objects
	// Get the actors in the activityPub protocol.
	//r.GET("/u/:actor", handlers.GetActorHandler)

	r.GET("/u/:actor/outbox", handlers.GetActorOutbox)
	//r.POST("/u/:user/inbox", handlers.InboxHandler)

}
