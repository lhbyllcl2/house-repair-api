package handler

import (
	"house-repair-api/pkg/response"
	"net/http"

	"house-repair-api/admin/internal/logic/sys/role"
	"house-repair-api/admin/internal/svc"

	logic "house-repair-api/admin/admin/internal/logic/sys/role"
)

func RoleDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewRoleDeleteLogic(r.Context(), ctx)
		err := l.RoleDelete()
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, nil)
	}
}
