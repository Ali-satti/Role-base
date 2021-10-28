package Service

import (
	"RolesBased/models"
	"strings"
)

func CheckRole(role *models.UserRole )(int,string){
	CheckDesignation := strings.ToLower(role.Designation)
	CheckapiName := strings.ToLower(role.ApiName)
 	if CheckDesignation=="admin"{
		 if CheckapiName=="add"{
			 return 200,"Role found successfully"
		 }else if CheckapiName=="update"{
			 return 200,"Role found successfully"
		}else if CheckapiName=="showteam"{
			return 200,"Role found successfully"
		}else if CheckapiName =="delete"{
			return 200,"Role found successfully"
		 }else if CheckapiName == "showself" {
			 return 200, "Role found successfully"
		 }else {
			 return 400,"Invalid apiName"
		 }
	}else if CheckDesignation=="team lead"{
		if CheckapiName == "showself"{
			return 200,"Role found successfully"
		}else if CheckapiName == "showteam"{
			return 200,"Role found successfully"
		}else if CheckapiName == "add"{
			return 403,"Forbidden"
		}else if CheckapiName == "update"{
			return 403,"Forbidden"
		}else if CheckapiName == "delete"{
			return 403,"Forbidden"
		}else {
			return 400,"Invalid apiName"
		}
	}else if CheckDesignation=="employee" {
		if CheckapiName == "showself" {
			return 200,"Role found successfully"
		}else if CheckapiName == "showteam"{
			return 200,"Role found successfully"
		}else if CheckapiName == "add"{
			return 403,"Forbidden"
		}else if CheckapiName == "update"{
			return 403,"Forbidden"
		}else if CheckapiName == "delete"{
			return 403,"Forbidden"
		}else {
			return 400,"Invalid apiName"
		}
	}else {
		return 400,"Invalid designation"
	}
}

