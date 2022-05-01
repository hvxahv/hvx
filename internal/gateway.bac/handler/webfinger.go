/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"github.com/hvxahv/hvx/pkg/activitypub"
)

func GetWebFingerHandler(c *gin.Context) {
	resource := c.Query("resource")

	fmt.Println(resource)
	if ok := activitypub.IsRemote(resource); ok {
		actor := activitypub.GetRemoteWebFinger(resource)
		c.JSON(200, actor)
		return
	}

	name := activitypub.GetActorName(resource)
	s, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.IsExistRequest{
		Username: name,
	}
	exist, err := s.IsExist(c, d)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exist.IsExist {
		c.JSON(200, activitypub.NewWebFinger(name, false))
		return
	}
}
