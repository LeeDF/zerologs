package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"github.com/LeeDF/zerologs/api/internal/logic/foo"
	"github.com/LeeDF/zerologs/api/internal/svc"
	"github.com/LeeDF/zerologs/api/internal/types"
)

func HelloHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FooReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), ctx)
		resp, err := l.Hello(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
