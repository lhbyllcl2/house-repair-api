package handler

import (
	"net/http"
	"real-estate/pkg/response"

	"house-repair-api/api/internal/logic/repair/part"
	"house-repair-api/api/internal/svc"
	"house-repair-api/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func PartDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Failed(w, response.ParamBindError, err)
			return
		}

		l := logic.NewPartDeleteLogic(r.Context(), ctx)
		err := l.PartDelete(req)
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, nil)
	}
}
