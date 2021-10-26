package Service

import (
	"RolesBased/models"
	"strings"
)

func CheckRole(role *models.UserRole )(string,error){
	CheckDesignation := strings.ToLower(role.Designation)
	CheckapiName := strings.ToLower(role.ApiName)
 	if CheckDesignation=="admin"{
		 return "200",nil
	}else if CheckDesignation=="team lead"{
		if CheckapiName == "showself"{
			return "200",nil
		}else if CheckapiName == "showteam"{
			return "200",nil
		}else {
			return "400",nil
		}
	}else if CheckDesignation=="employee" {
		if CheckapiName == "showself" {
			return "200", nil
		}else {
			return "400",nil
	}
	}else {
		return "401",nil
	}
}

