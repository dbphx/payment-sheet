package authenhandler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"my-source/sheet-payment/be/repository"
	"net/http"
	"net/url"
	"strings"
)

var (
	clientID     = ""
	clientSecret = ""
	redirectURI  = ""
	oauthState   = "" // Should generate real random string
	authURL      = "https://accounts.google.com/o/oauth2/v2/auth"
)

type SsoHandler struct {
	UserRepo repository.IUserRepository
}

func NewSSOHandler(usr repository.IUserRepository) *SsoHandler {
	return &SsoHandler{
		UserRepo: usr,
	}
}

func (sso *SsoHandler) LoginSSO(c *fiber.Ctx) error {
	urlStr := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=openid%%20email%%20profile&state=%s",
		authURL,
		clientID,
		redirectURI,
		oauthState,
	)

	return c.Redirect(urlStr, fiber.StatusTemporaryRedirect)
}

func (sso *SsoHandler) LoginSSOCallback(c *fiber.Ctx) error {
	if c.Query("state") != oauthState {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid state")
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("No code provided")
	}

	tokenResp, err := http.PostForm("https://oauth2.googleapis.com/token", url.Values{
		"code":          {code},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"redirect_uri":  {redirectURI},
		"grant_type":    {"authorization_code"},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get token")
	}
	defer tokenResp.Body.Close()

	body, _ := io.ReadAll(tokenResp.Body)

	var tokenData struct {
		IDToken string `json:"id_token"`
	}
	if err := json.Unmarshal(body, &tokenData); err != nil {
		return c.Status(500).SendString("Invalid token response")
	}

	// Verify the ID token
	payload, err := verifyGoogleIDToken(tokenData.IDToken)
	if err != nil {
		return c.Status(401).SendString("Token verification failed")
	}

	user := payload["email"].(string)
	dbUser, err := sso.UserRepo.GetByUsername(user)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	jwtToken, err := GenToken(dbUser.Username)
	if err != nil {
		return c.Status(500).SendString("Token generation failed")
	}

	// Return token to frontend (could be redirect or JSON)
	return c.Redirect("http://localhost/login-sso/callback-ui?token=" + jwtToken + "&username=" + user)
}

func verifyGoogleIDToken(idToken string) (map[string]interface{}, error) {
	parts := strings.Split(idToken, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid token")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var claims map[string]interface{}
	err = json.Unmarshal(payload, &claims)
	return claims, err
}
