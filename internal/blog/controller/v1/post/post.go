package post

import (
	"github.com/qppHUST/blog/internal/blog/biz"
	"github.com/qppHUST/blog/internal/blog/store"
)

// PostController 是 post 模块在 Controller 层的实现，用来处理博客模块的请求.
type PostController struct {
	b biz.IBiz
}

// New 创建一个 post controller.
func New(ds store.IStore) *PostController {
	return &PostController{b: biz.NewBiz(ds)}
}
