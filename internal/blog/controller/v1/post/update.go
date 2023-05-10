// Copyright qppHuster(1587299799@qq.com). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package post

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/qppHUST/blog/internal/pkg/core"
	"github.com/qppHUST/blog/internal/pkg/errno"
	"github.com/qppHUST/blog/internal/pkg/known"
	"github.com/qppHUST/blog/internal/pkg/log"
	v1 "github.com/qppHUST/blog/pkg/api/blog/v1"
)

// Update 更新博客.
func (ctrl *PostController) Update(c *gin.Context) {
	log.C(c).Infow("Update post function called")

	var r v1.UpdatePostRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Posts().Update(c, c.GetString(known.XUsernameKey), c.Param("postID"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
