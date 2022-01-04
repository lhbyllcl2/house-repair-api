package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"

	"house-repair-api/pkg/response"
)

type AuthMiddleware struct {
	Redis *redis.Redis
}

func NewAuthMiddleware(Redis *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		Redis: Redis,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//判断请求header中是否携带了x-user-id
		userId := r.Context().Value("userId").(json.Number).String()
		if userId == "" {
			logx.Errorf("缺少必要参数x-user-id")
			response.FailedWithCode(w, response.AuthorizationError)
			return
		}
		next(w, r)
	}
}
