package Service

import (
	"RolesBased/models"
	"strings"
)

func CheckRole(role *models.UserRole )(int,string){
	CheckDesignation := strings.ToLower(role.Designation)
	CheckApiName := strings.ToLower(role.ApiName)
	roles:=[]string{
		"add","update","showteam","delete","showself",
	}
 	if CheckDesignation=="admin"{
		for _,v := range roles{
			if v==CheckApiName{
				return 200,"Roles found successfully"
			}else {
				return 400,"Invalid apiName"
		}
		}
	}else if CheckDesignation=="team lead"{
		if roles[2]== CheckApiName || roles[4]== CheckApiName{
			return 200,"Roles found successfully"
		}else if roles[0]== CheckApiName || roles[1]== CheckApiName || roles[3]== CheckApiName{
			return 403,"Forbidden"
		}else{
			return 400,"Invalid apiName"
		}
	}else if CheckDesignation=="employee"{
		if roles[4]== CheckApiName {
			return 200,"Role found successfully"
		}else if roles[0]== CheckApiName || roles[1]== CheckApiName || roles[2]== CheckApiName || roles[3]== CheckApiName{
			return 403,"Forbidden"
		}else {
			return 400,"Invalid apiName"
		}
	}else {
		return 400, "Invalid designation"
	}
	return 400, "Invalid designation"
}

