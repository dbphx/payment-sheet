package authenhandler

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURI = os.Getenv("GOOGLE_REDIRECT_URI")
	oauthState = os.Getenv("GOOGLE_OAUTH_STATE")
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
