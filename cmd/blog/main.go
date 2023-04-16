// Copyright qppHuster ;1587299799@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package main

import (
	"blog/internal/blog"
	_ "go.uber.org/automaxprocs"
	"os"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := blog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
