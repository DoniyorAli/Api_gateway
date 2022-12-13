package models

import "time"

type Author struct {
	ID string `json:"id"`
	// Lastname  string `json:"lastname" binding:"required" minLenth:"3" maxLength:"16" example:"John"`
	// Firstname string `json:"firstname" binding:"required" minLenth:"3" maxLength:"16" example:"Doe"`
	Fullname   string     `json:"fullname" binding:"required" minLenth:"3" maxLength:"255" example:"Sherlock Holmes"`
	Middlename string     `json:"middlename" minLenth:"3" maxLength:"15" example:"Sherlock"`
	CreateAt   time.Time  `json:"created_at"`
	UpdateAt   *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

type CreateModelAuthor struct {
	// Lastname  string `json:"lastname" binding:"required" minLenth:"3" maxLength:"16" example:"John"`
	// Firstname string `json:"firstname" binding:"required" minLenth:"3" maxLength:"16" example:"Doe"`
	Fullname   string `json:"fullname" binding:"required" minLenth:"3" maxLength:"255" example:"Sherlock Holmes"`
	Middlename string `json:"middlename" minLenth:"3" maxLength:"15" example:"Sherlock"`
}

type UpdateAuthorResponse struct {
	ID string `json:"id"`
	// Lastname  string `json:"lastname" binding:"required" minLenth:"3" maxLength:"16" example:"John"`
	// Firstname string `json:"firstname" binding:"required" minLenth:"3" maxLength:"16" example:"Doe"`
	Fullname   string `json:"fullname" binding:"required" minLenth:"3" maxLength:"255" example:"Sherlock Holmes"`
	Middlename string `json:"middlename" minLenth:"3" maxLength:"15" example:"Sherlock"`
}
