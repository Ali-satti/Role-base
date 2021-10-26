package handlers

import (
	"RolesBased/gen/models"
	domain "RolesBased/models"
	"github.com/go-openapi/swag"
)
// toUserDomain converts gen to domain model
func toUserDomain(role *models.Role) *domain.UserRole {
	return &domain.UserRole{
		Designation:     swag.StringValue(role.Designation),
		ApiName:            swag.StringValue(role.APIName),
	}
}

// toUserGen converts domain to gen model
func toUserGen(s string) *models.Rolereturn{
	return &models.Rolereturn{
		Response: swag.String(s),
	}
}
