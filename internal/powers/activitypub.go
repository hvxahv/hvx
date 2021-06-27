package powers

import (
	"github.com/gin-gonic/gin"
	"hvxahv/internal/powers/handlers"
)

func activityPubV1(r *gin.Engine) {
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	r.GET("/.well-known/webfinger", handlers.WebFingerHandler)
	// TODO - Actor
	r.GET("", handlers.GetActorHandler)

}
