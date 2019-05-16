package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/FuzzyStatic/blizzard"
	"github.com/ccod/gosu-server/pkg/client"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/ccod/gosu-server/pkg/models"
	re "github.com/ccod/gosu-server/pkg/response"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"

	"github.com/ccod/go-bnet"
)

func login(oauthCfg *oauth2.Config, authSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oauthCfg.AuthCodeURL(authSecret, oauth2.AccessTypeOnline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func bnetCB(oauthCfg *oauth2.Config, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func getUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	accID := r.Context().Value(m.JWTKey).(int)

	var user models.Player
	db.First(&user, accID)

	if user.AccountID != 0 {
		re.RespondJSON(user, w, r)
		return
	}

	// Since Player record does not already exist in DB (first time login), will call on client to do the initial Fetch
	blizz := r.Context().Value(m.BlizzKey).(*blizzard.Client)
	blizz.TokenValidation()

	user, err := client.FetchNewPlayer(blizz, accID)
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	// should move this to player model
	db.Create(&user)
	re.RespondJSON(user, w, r)
}
