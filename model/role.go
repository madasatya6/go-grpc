package model

type Role struct {
	RoleID uint   `json:"role_id" gorm:"primaryKey"`
	Name   string `json:"name"`
}
