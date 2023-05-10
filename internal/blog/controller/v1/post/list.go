// Copyright qppHuster(1587299799@qq.com). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/qppHUST/blog/internal/pkg/core"
	"github.com/qppHUST/blog/internal/pkg/errno"
	"github.com/qppHUST/blog/internal/pkg/known"
	"github.com/qppHUST/blog/internal/pkg/log"
	v1 "github.com/qppHUST/blog/pkg/api/blog/v1"
)

// List 返回博客列表.
func (ctrl *PostController) List(c *gin.Context) {
	log.C(c).Infow("List post function called.")

	var r v1.ListPostRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Posts().List(c, c.GetString(known.XUsernameKey), r.Offset, r.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
