package handler

import (
	"net/http"

	"house-repair-api/admin/internal/logic/sys/user"
	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"
	"house-repair-api/pkg/response"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Failed(w, response.ParamBindError, err)
			return
		}

		l := logic.NewUserAddLogic(r.Context(), ctx)
		err := l.UserAdd(req)
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, nil)
	}
}
