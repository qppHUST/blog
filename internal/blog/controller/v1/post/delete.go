package post

import (
	"github.com/gin-gonic/gin"
	"github.com/qppHUST/blog/internal/pkg/core"
	"github.com/qppHUST/blog/internal/pkg/known"
	"github.com/qppHUST/blog/internal/pkg/log"
)

// Delete 删除指定的博客.
func (ctrl *PostController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete post function called")

	if err := ctrl.b.Posts().Delete(c, c.GetString(known.XUsernameKey), c.Param("postID")); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
