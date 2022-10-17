package rules

import (
	"api-permission/internal/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"api-permission/internal/logic/auth/rules"
	"api-permission/internal/svc"
	"api-permission/internal/types"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRuleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rules.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
