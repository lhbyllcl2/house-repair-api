package handler

import (
	"net/http"

	"house-repair-api/admin/internal/logic/sys/user"
	"house-repair-api/admin/internal/svc"
	"house-repair-api/pkg/response"
)

func UserDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserDeleteLogic(r.Context(), ctx)
		err := l.UserDelete()
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, nil)
	}
}
