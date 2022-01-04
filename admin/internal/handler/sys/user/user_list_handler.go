package handler

import (
	"house-repair-api/pkg/response"
	"net/http"

	"house-repair-api/admin/internal/logic/sys/user"
	"house-repair-api/admin/internal/svc"
	"house-repair-api/admin/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Failed(w, response.ParamBindError, err)
			return
		}

		l := logic.NewUserListLogic(r.Context(), ctx)
		resp, err := l.UserList(req)
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, resp)
	}
}
