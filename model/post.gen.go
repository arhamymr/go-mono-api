// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePost = "Post"

// Post mapped from table <Post>
type Post struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Content   string    `gorm:"column:content;not null" json:"content"`
	Published bool      `gorm:"column:published;not null" json:"published"`
	Deleted   time.Time `gorm:"column:deleted" json:"deleted"`
	AuthorID  int32     `gorm:"column:author_id;not null" json:"author_id"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}