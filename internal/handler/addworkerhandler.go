package handler

import (
	"net/http"

	"github.com/lukaproject/schedulersvr/internal/logic"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddWorkerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddWorkerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddWorkerLogic(r.Context(), svcCtx)
		resp, err := l.AddWorker(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
