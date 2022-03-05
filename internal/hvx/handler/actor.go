/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
)

// GetActorHandler Get the actor's handler. It will get
// the queried username from Param, then call the gRPC
// service by the username, and return the JsonLD of the
// standard activitypub protocol.
func GetActorHandler(c *gin.Context) {

}

func SearchActorsHandler(c *gin.Context) {

}
