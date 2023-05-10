// Copyright qppHuster &lt;1587299799@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package middleware

import (
	"github.com/qppHUST/blog/internal/pkg/core"
	"github.com/qppHUST/blog/internal/pkg/errno"
	"github.com/qppHUST/blog/internal/pkg/known"
	"github.com/qppHUST/blog/internal/pkg/log"
	"github.com/qppHUST/blog/pkg/token"

	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token，能解析成功，就意味着认证成功，就意味着成功了
		username, err := token.ParseRequest(c)
		log.Infow("auth filter", "username", username)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()

			return
		}

		// 这是为了授权才存储的，在authz里面
		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
