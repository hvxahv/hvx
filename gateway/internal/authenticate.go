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
// Verify whether the Token is correct,
// Will obtain the username by parsing the Token
// Add the username in the context and set the key to loginUser.
func Auth(c *gin.Context) {
	a := c.Request.Header.Get("Authorization")
	if a == "" {
		c.JSON(500, errors.NewHandler("500", errors.ErrNotAuthorizationTOKEN))
		c.Abort()
		return
	}
	parse, err := ParseAuthorization(a)
	if err != nil {
		c.JSON(401, errors.NewHandler("401", errors.ErrTokenParse))
		c.Abort()
		return
	}
	_ := clientv1.New(c,
		microsvc.NewGRPCAddress("device").Get(),
	)
	if err != nil {
		c.JSON(500, errors.NewHandler("500", errors.ErrConnectDeviceRPCServer))
		c.Abort()
		return
	}
	defer cli.Close()
	devices, err := device.NewDevicesClient(cli.Conn).IsExist(c, &device.IsExistRequest{Id: parse.DeviceID})
	if err != nil {
		c.JSON(501, errors.NewHandler("501", err.Error()))
		c.Abort()
		return
	}

	if devices.IsExist {
		c.JSON(401, errors.NewHandler("401", errors.New(errors.ErrTokenUnauthorized).Error()))
		c.Abort()
		return
	}

	c.Next()
}

// ParseAuthorization Parses the obtained Authorization and returns the Claims data.
// GET AUTHORIZATION EXAMPLE: a := c.Request.Header.Get("Authorization")
func ParseAuthorization(a string) (*auth2.Claims, error) {
	var (
		token  = strings.Split(a, "Bearer ")[1]
		secret = viper.GetString("authentication.token.secret")
	)
	parse, err := auth2.NewParseJWTToken(token, secret).JWTTokenParse()
	if err != nil {
		return nil, err
	}
	return parse, nil
}
