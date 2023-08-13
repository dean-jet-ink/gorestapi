package model

import "time"

type Task struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"userId" gorm:"not null"`
}

type TaskResponse struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" grom:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
