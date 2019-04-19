package auth

import (
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"

	"github.com/ccod/go-bnet"
)

func login(oauthCfg *oauth2.Config, authSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oauthCfg.AuthCodeURL(authSecret, oauth2.AccessTypeOnline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func bnetCB(oauthCfg *oauth2.Config, authSecret string, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		if state != authSecret {
			fmt.Printf("invalid oauth state expected '%s', go '%s'\n", authSecret, state)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		code := r.FormValue("code")
		token, err := oauthCfg.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		oauthClient := oauthCfg.Client(oauth2.NoContext, token)
		client := bnet.NewClient("us", oauthClient)
		user, _, err := client.UserInfo()
		if err != nil {
			fmt.Printf("client.Profile().SC2() failed with '%s'\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		claims := &jwt.StandardClaims{
			Id:     strconv.Itoa(user.ID),
			Issuer: "gosu-beef",
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := jwtToken.SignedString([]byte(jwtSecret))
		if err != nil {
			fmt.Printf("jwt signing failed: %s", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		fmt.Printf("UserInfo is: %v\n", user)
		fmt.Printf("Jwt token: %s", tokenString)
		// will pass client domain as part of tools I think...
		http.Redirect(w, r, "http://localhost:3000/callback#"+tokenString, http.StatusTemporaryRedirect)
	}
}
