// Copyright qppHuster &lt;1587299799@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package middleware

import (
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/known"
	"blog/internal/pkg/log"
	"blog/pkg/token"

	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(c)
		log.Infow("auth filter", "username", username)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()

			return
		}

		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
