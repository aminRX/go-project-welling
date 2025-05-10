// Handlers HTTP 
package controllers

import (
	"net/http"

	"go-microservices/models"
	"go-microservices/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignupInput struct {
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Role     models.Role `json:"role"`
	OrgID    uuid.UUID   `json:"org_id"`
}

type OrgInput struct {
	Name string `json:"name"`
}

func Signup(c *gin.Context) {
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.CreateUser(input.Email, input.Password, input.Role, input.OrgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func CreateOrganization(c *gin.Context) {
	var input OrgInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org, err := services.CreateOrganization(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, org)
}
