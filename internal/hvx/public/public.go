/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"log"
)

func GetPublicAccountCountHandler(c *gin.Context) {
	db := cockroach.GetDB()
	var count int64
	if err := db.Debug().Table("account").Count(&count).Error; err != nil {
		log.Println(err)
		return
	}
	fmt.Println(count)
	c.JSON(200, gin.H{
		"code":          "200",
		"account_count": count,
	})
}

func GetPublicInstanceDetailsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "200",
		"details": gin.H{
			"version":    "0.1.0",
			"build":      "2022-01-01",
			"maintainer": "hvxahv",
		},
	})
}
