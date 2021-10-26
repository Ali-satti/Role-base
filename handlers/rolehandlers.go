package handlers

import (
	"RolesBased/Service"
	"RolesBased/gen/restapi/operations/getting_role"
	"github.com/go-openapi/runtime/middleware"
)

type rolebase struct {
}

func RoleBase() getting_role.GetroleHandler {
	return &rolebase{}
}

func (rb *rolebase) Handle(params getting_role.GetroleParams) middleware.Responder {
	rbase,err := Service.CheckRole(toUserDomain(params.Body))
	if err != nil{
		return getting_role.NewGetroleBadRequest().WithPayload("Check you request again")
		panic(err)
	}
	if rbase == "200"{
		return getting_role.NewGetroleOK().WithPayload(toUserGen(rbase))
	}else {
		return getting_role.NewGetroleUnauthorized().WithPayload(toUserGen(rbase))
	}
}

