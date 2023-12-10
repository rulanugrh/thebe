package middleware

import (
	"be-project/entity/domain"
	"be-project/entity/web"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var sessions = map[string]session{}

type session struct {
	username string
	expiry time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func GenerateSession(req domain.User) (string, time.Time) {
	sessionToken := uuid.NewString()
	expireAt := time.Now().Add(1 * time.Hour)

	namesUser := fmt.Sprintf("%s %s", req.FName, req.LName)
	sessions[sessionToken] = session {
		username: namesUser,
		expiry: expireAt,
	}

	return sessionToken, expireAt
}

func RefreshSession(req domain.User, w http.ResponseWriter, r *http.Request) (string, time.Time){
	c, err := r.Cookie("session_tokens")
	if err != nil {
		if err == http.ErrNoCookie {
			response := web.ResponseFailure {
				Code: http.StatusUnauthorized,
				Message: "Unauthorized",
			}

			result, _ := json.Marshal(response)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
		}

		w.WriteHeader(http.StatusBadRequest)
	}
	sessionToken := c.Value
	
	newSessionToken := uuid.NewString()
	expireAt := time.Now().Add(1 * time.Hour)

	namesUser := fmt.Sprintf("%s %s", req.FName, req.LName)
	sessions[newSessionToken] = session {
		username: namesUser,
		expiry: expireAt,
	}

	delete(sessions, sessionToken)
	return newSessionToken, time.Now().Add(1 * time.Hour)
}

func SessionVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_tokens")
		if err != nil {
			if err == http.ErrNoCookie {
				response := web.ResponseFailure {
					Code: http.StatusUnauthorized,
					Message: "Unauthorized",
				}

				result, _ := json.Marshal(response)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(result)
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value
		userSession, exist := sessions[sessionToken]
		if !exist {
			response := web.ResponseFailure {
				Code: http.StatusUnauthorized,
				Message: "Unauthorized",
			}

			result, _ := json.Marshal(response)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
			return
		}

		if userSession.isExpired() {
			delete(sessions, sessionToken)
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

	})
}