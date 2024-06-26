package middleware

import (
	"net/http"
	"wechat-backend/common/response"
	"wechat-backend/common/utils"
	"wechat-backend/service/communicate/api/internal/config"
)

type VerifyAuthorityMiddleware struct {
	config config.Config
}

func NewVerifyAuthorityMiddleware(c config.Config) *VerifyAuthorityMiddleware {
	return &VerifyAuthorityMiddleware{config: c}
}

func (v VerifyAuthorityMiddleware) HandleVerifyAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//获取token
		authorization := r.Header.Get("Authorization")

		_, err := utils.ParseToken(authorization, v.config.Auth.AccessSecret)
		if err != nil {
			response.Fail(response.UNAUTHORIZED, response.UNAUTHORIZED_MESSAGE, w, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
