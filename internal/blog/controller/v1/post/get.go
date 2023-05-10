package post

import (
	"github.com/gin-gonic/gin"
	"github.com/qppHUST/blog/internal/pkg/core"
	"github.com/qppHUST/blog/internal/pkg/known"
	"github.com/qppHUST/blog/internal/pkg/log"
)

// Get 获取指定的博客.
func (ctrl *PostController) Get(c *gin.Context) {
	log.C(c).Infow("Get post function called")

	post, err := ctrl.b.Posts().Get(c, c.GetString(known.XUsernameKey), c.Param("postID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
