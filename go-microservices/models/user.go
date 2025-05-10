// Modelos de GORM 
package models

import (
	"github.com/google/uuid"
)

type Role string

const (
	Admin    Role = "admin"
	Master   Role = "master"
	Personal Role = "personal"
)

type Organization struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"unique;not null"`
	Users []User   `gorm:"foreignKey:OrganizationID"`
}

type User struct {
	ID             uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email          string        `gorm:"uniqueIndex;not null"`
	PasswordHash   string        `gorm:"not null"`
	Role           Role          `gorm:"type:varchar(20);not null"`
	OrganizationID uuid.UUID     `gorm:"type:uuid"`
	Organization   Organization  `gorm:"foreignKey:OrganizationID"`
}
