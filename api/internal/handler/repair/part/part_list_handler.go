package handler

import (
	"net/http"
	"real-estate/pkg/response"

	"house-repair-api/api/internal/logic/repair/part"
	"house-repair-api/api/internal/svc"
	"house-repair-api/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func PartListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Failed(w, response.ParamBindError, err)
			return
		}

		l := logic.NewPartListLogic(r.Context(), ctx)
		resp, err := l.PartList(req)
		if err != nil {
			response.Failed(w, response.HandleWithFailure, err)
			return
		}
		response.Success(w, resp)
	}
}
