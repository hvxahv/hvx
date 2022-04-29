/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
)

// GetActorHandler Get the actor's handler. It will get
// the queried username from Param, then call the gRPC
// service by the username, and return the JsonLD of the
// standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	s, err := account.GetActorClient()
	if err != nil {
		return
	}
	d := &pb.GetActorByAccountUsernameRequest{
		Username: c.Param("actor"),
	}

	a, err := s.GetActorByAccountUsername(c, d)
	if err != nil {
		return
	}
	actor := activitypub.NewActor(a.PreferredUsername, a.Name, a.Summary, a.PublicKey, a.Avatar)
	c.JSON(200, actor)
}

func SearchActorsHandler(c *gin.Context) {
	s, err := account.GetActorClient()
	if err != nil {
		return
	}
	d := &pb.GetActorsByPreferredUsernameRequest{
		PreferredUsername: c.Param("actor"),
	}

	actors, err := s.GetActorsByPreferredUsername(c, d)
	if err != nil {
		return
	}
	c.JSON(200, actors.Actors)
}
