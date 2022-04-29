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
	"io/ioutil"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	fmt.Println(name, string(body))
}
