package handler

import (
	"net/http"

	"github.com/lukaproject/schedulersvr/internal/logic"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveWorkerByNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveWorkerByNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveWorkerByNameLogic(r.Context(), svcCtx)
		resp, err := l.RemoveWorkerByName(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
