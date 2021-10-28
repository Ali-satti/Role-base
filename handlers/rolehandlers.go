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
	rbase,rmessage := Service.CheckRole(toUserDomain(params.Body))
	if rbase == 200{
		return getting_role.NewGetroleOK().WithPayload(toUserGen(rbase,rmessage))
	}else if rbase == 403{
		return getting_role.NewGetroleForbidden().WithPayload(toUserGen2(rbase,rmessage))
	}else if rbase == 400 {
		return getting_role.NewGetroleBadRequest().WithPayload(toUserGen3(rbase,rmessage))
	}else {
		return getting_role.NewGetroleInternalServerError().WithPayload(toUserGen1(rbase,rmessage))
	}
}

