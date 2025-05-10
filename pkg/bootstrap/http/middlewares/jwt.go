package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/WayDBae/eWallet/internal/entities"
	"github.com/WayDBae/eWallet/pkg/bootstrap/http/misc/response"
	"github.com/WayDBae/eWallet/pkg/utils"
	"github.com/google/uuid"
)

func (m *provider) JWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			data, _ := json.Marshal(response.Build(response.ErrUnauthorized))
			w.Write(data)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			data, _ := json.Marshal(response.Build(response.ErrUnauthorized))
			w.Write(data)
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString, m.config.Server.RefreshSecretKey, r.Context())
		if err != nil {
			data, _ := json.Marshal(response.Build(response.ErrUnauthorized))
			w.Write(data)
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {

		}
		user, err := m.user.Get(entities.User{
			BaseGorm: entities.BaseGorm{
				ID: userID,
			},
		}, r.Context())
		if err != nil || user.DeletedAt.Valid {
			data, _ := json.Marshal(response.Build(response.ErrUnauthorized))
			w.Write(data)
			return
		}

		// можно сохранить user в context
		ctx := context.WithValue(r.Context(), entities.ContextUserIDKey, user.ID)

		next(w, r.WithContext(ctx))
	})
}
