package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//Key is used to disambiguate context keys
type Key int

const (
	// DBKey is a context key for pulling the DB reference
	DBKey Key = 0
	// JWTKey is a context key for pulling valid authenticated struct containing player information
	JWTKey Key = 1
)

// normally I only want to pass the db to my handlers, the other keys are seldom used... which is why I split them up

// PassDB middleware takes a reference to db and inserts it into context for subsequent handlers through a closure
func PassDB(db *gorm.DB) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), DBKey, db)
			r = r.WithContext(ctx)
			f(w, r)
		}
	}
}

// JWTIdentity middleware is used for checking user authentication and passing user identifier to handler
func JWTIdentity(jwtSecret string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			reqToken = strings.Split(reqToken, "Bearer ")[1]

			token, err := jwt.ParseWithClaims(reqToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				// interested in passing config information through middleware as well... not a big fan of everything hanging off server type
				return []byte(jwtSecret), nil
			})

			claims, ok := token.Claims.(*jwt.StandardClaims)
			if !(ok && token.Valid) {
				fmt.Printf("error with Parse with Claims: %s", err)
				w.Write([]byte("{\"failure\":true}"))
				return
			}

			accountID, err := strconv.Atoi(claims.Id)
			if err != nil {
				fmt.Printf("Atoi call failed: %s", err)
				w.Write([]byte("{\"failure\":true}"))
				return
			}

			ctx := context.WithValue(r.Context(), JWTKey, accountID)
			r = r.WithContext(ctx)
			f(w, r)
		}
	}
}
