// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	foo "github.com/LeeDF/zerologs/api/internal/handler/foo"
	"github.com/LeeDF/zerologs/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hello",
				Handler: foo.HelloHandler(serverCtx),
			},
		},
	)
}
