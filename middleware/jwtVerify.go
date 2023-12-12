package middleware

import (
	"be-project/config"
	"be-project/entity/domain"
	"be-project/entity/web"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(user domain.UserLogin) (string, error) {
	conf := config.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &jwtClaims{
		Name: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.Secret))
	if err != nil {
		log.Printf("Cant claim jwt token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(token string) error {
	conf := config.GetConfig()
	tokens, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Secret), nil
	})

	if err != nil {
		log.Printf("Token is not valid: %v", err)
	}

	claims, errClaim := tokens.Claims.(*jwtClaims)
	if !errClaim {
		log.Printf("Cant claim token %v", errClaim)
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		log.Printf("token expired")
	}

	return nil

}

func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token = r.Header.Get("Authorization")
		json.NewEncoder(w).Encode(r)

		token = strings.TrimSpace(token)

		if token == "" {
			res := web.ResponseFailure{
				Message: "Cant login because not have token",
				Code:    http.StatusUnauthorized,
			}

			response, _ := json.Marshal(res)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
		}

		err := ValidateToken(token)
		if err != nil {
			res := web.ResponseFailure{
				Message: "Cant login because token is not valid",
				Code:    http.StatusBadRequest,
			}

			response, _ := json.Marshal(res)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)
		}

		next.ServeHTTP(w, r)
	})
}
