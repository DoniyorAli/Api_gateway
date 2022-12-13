package models

import "time"

type Content struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	AuthorID  string     `json:"author_id" binding:"required"`
	CreateAt  time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type CreateModelArticle struct {
	Content         // Promoted fields
	AuthorID string `json:"author_id" binding:"required"`
}

type UpdateArticleModel struct {
	ID string `json:"id" binding:"required"`
	Content
}

type GetByIDArticleModel struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author  Author     `json:"author_id" binding:"required"`
	CreateAt  time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}