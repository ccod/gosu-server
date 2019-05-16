package config

import (
	"fmt"
	"os"

	"github.com/FuzzyStatic/blizzard"
	"github.com/ccod/go-bnet"
	"github.com/jinzhu/gorm"

	// config also creates a db connection, which is why it is here
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// Settings refers to the initial set of values that will configure the web server
type Settings struct {
	Domain       string
	Port         string
	ClientDomain string

	Tools Tools
}

// Tools refers to things that will be passed through middleware as necessary pieces for the handler, like db, blizzard client etc...
type Tools struct {
	Blizz      *blizzard.Client
	DB         *gorm.DB
	AuthClient *oauth2.Config

	OauthSalt string
	JWTSecret string
}

// LoadFromEnv will even start the db connection, but as this should be called from main, you can defer close() in the next line.
func LoadFromEnv() Settings {
	var s Settings

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found, will only pull from environment vars")
	}

	s.Domain = os.Getenv("GOSU_DOMAIN")
	s.Port = os.Getenv("GOSU_PORT")
	s.ClientDomain = os.Getenv("GOSU_CLIENT_DOMAIN")

	blizzClientID := os.Getenv("GOSU_BLIZZ_CLIENT_ID")
	blizzClientSecret := os.Getenv("GOSU_BLIZZ_CLIENT_SECRET")

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=%s",
		os.Getenv("GOSU_DB_HOST"),
		os.Getenv("GOSU_DB_PORT"),
		os.Getenv("GOSU_DB_USER"),
		os.Getenv("GOSU_DB_PASSWORD"),
		os.Getenv("GOSU_DB_DBNAME"),
		os.Getenv("GOSU_DB_SSL"),
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		fmt.Printf("db err: %s\n", err)
		panic("db connection returned with an error")
	}

	// going to fix the redirect url, as the deployment target won't need to incorporate a port in the uri
	s.Tools = Tools{
		DB:    db,
		Blizz: blizzard.NewClient(blizzClientID, blizzClientSecret, blizzard.US, blizzard.Locale("enUS")),
		AuthClient: &oauth2.Config{
			ClientID:     blizzClientID,
			ClientSecret: blizzClientSecret,
			Scopes:       []string{"sc2.profile"},
			RedirectURL:  s.Domain + ":" + s.Port + "/auth/bnet_oauth_cb",
			Endpoint:     bnet.Endpoint("us"),
		},

		OauthSalt: os.Getenv("GOSU_OAUTH_SALT"),
		JWTSecret: os.Getenv("GOSU_JWT_SECRET"),
	}

	if err := s.Tools.Blizz.AccessTokenReq(); err != nil {
		panic("Nope, failed to get blizzard access token")
	}
	fmt.Println("Created Access Token")

	return s
}
