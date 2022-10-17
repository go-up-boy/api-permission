package roles

import (
	"api-permission/internal/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"api-permission/internal/logic/auth/roles"
	"api-permission/internal/svc"
	"api-permission/internal/types"
)

func AddRulesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRulesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := roles.NewAddRulesLogic(r.Context(), svcCtx)
		resp, err := l.AddRules(&req)
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
