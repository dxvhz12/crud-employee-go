package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if auth == "" {
			http.Error(w, "missing token", 401)
			return
		}

		tokenString := strings.TrimPrefix(
			auth,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, jwt.ErrSignatureInvalid

				}

				return []byte(
					os.Getenv("JWT_SECRET"),
				), nil
			},
		)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				401,
			)

			return
		}

		if !token.Valid {

			http.Error(
				w,
				"invalid token",
				401,
			)

			return
		}

		next(w, r)

	}

}
