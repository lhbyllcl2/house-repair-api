package handler

import (
	"house-repair-api/pkg/response"
	"net/http"

	"house-repair-api/admin/internal/logic/sys/role"
	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"

	logic "github.com/lhbyllcl2/house-repair-api/admin/admin/internal/logic/sys/role"
)

func RoleAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAddReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Failed(w, response.ParamBindError, err)
			return
		}

		l := logic.NewRoleAddLogic(r.Context(), ctx)
		err := l.RoleAdd(req)
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, nil)
	}
}
