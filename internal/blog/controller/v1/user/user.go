// Copyright qppHuster &lt;1587299799@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package user

import (
	"github.com/qppHUST/blog/internal/blog/biz"
	"github.com/qppHUST/blog/internal/blog/store"
	"github.com/qppHUST/blog/pkg/auth"
	pb "github.com/qppHUST/blog/pkg/proto/blog/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedBlogServer
}

// New 创建一个 user controller.
func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
