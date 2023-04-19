// Copyright qppHuster &lt;1587299799@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qppHust/blog.

package model

import "time"

type PostM struct {
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	Email     string    `gorm:"column:email"`          //
	ID        int64     `gorm:"column:id;primary_key"` //
	Nickname  string    `gorm:"column:nickname"`       //
	Password  string    `gorm:"column:password"`       //
	Phone     string    `gorm:"column:phone"`          //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
	Username  string    `gorm:"column:username"`       //
}

// TableName sets the insert table name for this struct type
func (p *PostM) TableName() string {
	return "user"
}
