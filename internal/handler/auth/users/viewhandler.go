package users

import (
	"api-permission/internal/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"api-permission/internal/logic/auth/users"
	"api-permission/internal/svc"
	"api-permission/internal/types"
)

func ViewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ViewUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := users.NewViewLogic(r.Context(), svcCtx)
		resp, err := l.View(&req)
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
