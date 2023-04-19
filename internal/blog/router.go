// Copyright qppHuster &lt;1587299799@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package blog

import (
	"blog/internal/blog/controller/v1/user"
	"blog/internal/blog/store"
	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/log"
	"blog/pkg/auth"

	mw "blog/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S, authz)

	//登录接口
	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get)
			// userv1.PUT(":name",us.UPDATE)
			// userv1.GET("",us.List)
			// userv1.DELETE(":name",us.Delete)
		}
	}

	return nil
}
