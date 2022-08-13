/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/APIs/v1alpha1/device"
	auth2 "github.com/hvxahv/hvx/auth"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/spf13/viper"
)

// Auth authentication middleware for gin network framework,
// Check whether the Token is carried in the request,
// and verify whether the Token is correct,
// Will obtain the username by parsing the Token
// and add the username in the context and set the key to loginUser.
func Auth(c *gin.Context) {
	a := c.Request.Header.Get("Authorization")
	if a == "" {
		c.JSON(500, gin.H{
			"code":    "500",
			"message": "TOKEN IS NOT CARRIED IN THE REQUEST",
		})
		c.Abort()
		return
	}
	var (
		token  = strings.Split(a, "Bearer ")[1]
		secret = viper.GetString("authentication.token.secret")
	)

	parse, err := auth2.NewParseJWTToken(token, secret).JWTTokenParse()
	if err != nil {
		c.JSON(401, gin.H{
			"code":  "401",
			"error": errors.New("TOKEN_PARSING_FAILURE").Error(),
		})
		c.Abort()
		return
	}

	cli, err := clientv1.New(c,
		microsvc.NewGRPCAddress("device").Get(),
	)
	if err != nil {
		return
	}
	defer cli.Close()
	devices, err := device.NewDevicesClient(cli.Conn).IsExist(c, &device.IsExistRequest{Id: parse.DeviceID})
	if err != nil {
		c.JSON(501, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	if devices.IsExist {
		c.JSON(401, gin.H{
			"code":  "401",
			"error": errors.New(errors.ErrTokenUnauthorized).Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
