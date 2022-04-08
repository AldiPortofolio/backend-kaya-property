package models

import "time"

type Blog struct {
	ID   int    `gorm:"column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	Status string `gorm:"column:status" json:"status"`
	BannerPhoto string `gorm:"column:banner_photo" json:"banner_photo"`
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Photo string `gorm:"column:photo" json:"photo"`
	Slug string    `gorm:"column:slug" json:"slug"`
}

type BlogDetail struct {
	ID   int    `gorm:"column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	Status string `gorm:"column:status" json:"status"`
	BannerPhoto string `gorm:"column:banner_photo" json:"banner_photo"`
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Photo string `gorm:"column:photo" json:"photo"`
	Slug string    `gorm:"column:slug" json:"slug"`
	BlogTags []BlogTags `gorm:"foreignKey:BlogId;" json:"blog_tags"`
}

type BlogTags struct {
	ID   int    `gorm:"column:id" json:"id"`
	TagId int    `gorm:"column:tag_id" json:"tagId"`
	BlogId int   `gorm:"column:blog_id" json:"blogId"`
	Blog []Blog  `gorm:"references:BlogId" json:"blog"`
	Tag Tag  `gorm:"references:TagId" json:"tag"`
}

type ResBlogTags struct {
	ID   int    `gorm:"column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
	Status string `gorm:"column:status" json:"status"`
	BannerPhoto string `gorm:"column:banner_photo" json:"banner_photo"`
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Photo string `gorm:"column:photo" json:"photo"`
	TagId int    `gorm:"column:tag_id" json:"tagId"`
	Slug string    `gorm:"column:slug" json:"slug"`
	
}



func (t *Blog) TableName() string {
	return "blogs"
}

func (t *BlogDetail) TableName() string {
	return "blogs"
}

func (t *BlogTags) TableName() string {
	return "blog_tags"
}

