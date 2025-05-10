// Lógica de negocio 
package services

import (
	"go-microservices/db"
	"go-microservices/models"
	"go-microservices/utils"

	"github.com/google/uuid"
)

func CreateUser(email, password string, role models.Role, orgID uuid.UUID) (*models.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:          email,
		PasswordHash:   hashed,
		Role:           role,
		OrganizationID: orgID,
	}
	err = db.DB.Create(user).Error
	return user, err
}

func CreateOrganization(name string) (*models.Organization, error) {
	org := &models.Organization{Name: name}
	err := db.DB.Create(org).Error
	return org, err
}
