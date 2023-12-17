package handlers

import (
	"log"
	"os"
	"quickshare/config"
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler interface {
	SendGoogle(c *fiber.Ctx) error
	SendGoogleCallback(c *fiber.Ctx) error
}

type AuthHandlerImpl struct {
	UtilService services.UtilService
	UserService services.UserService
}

func (h *AuthHandlerImpl) SendGoogle(c *fiber.Ctx) error {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("G_CLIENT"),
		ClientSecret: os.Getenv("G_SECRET"),
		RedirectURL:  os.Getenv("G_REDIRECT"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	// generate random 32-long for state identification
	generated := h.UtilService.GenerateRandomID(32)

	sess, _ := config.StateStore.Get(c)
	sess.Set("session_state", generated)
	sess.Save()

	// create url for auth process.
	// we can pass state as someway to identify
	// and validate the login process.
	URL := conf.AuthCodeURL(generated)

	return c.Redirect(URL)
}

func (h *AuthHandlerImpl) SendGoogleCallback(c *fiber.Ctx) error {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("G_CLIENT"),
		ClientSecret: os.Getenv("G_SECRET"),
		RedirectURL:  os.Getenv("G_REDIRECT"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	// get session store for current context
	sess, sessErr := config.SessionStore.Get(c)
	stateSess, stateErr := config.StateStore.Get(c)
	if sessErr != nil || stateErr != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// state from the session storage
	savedState := stateSess.Get("session_state")

	// get state and  from the google callback
	state := c.Query("state")
	code := c.Query("code")

	// compare the state that is coming from the callback
	// with the one that is stored inside the session storage.
	if state != savedState {
		// Handle the invalid state error
		return c.Status(fiber.StatusBadRequest).SendString("Invalid state")
	}

	// exchange code that retrieved from google via
	// URL query parameter into token, this token
	// can be used later to query information of current user
	// from respective provider.
	token, err := conf.Exchange(c.Context(), code)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	profile, err := h.UtilService.ConvertToken(token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// find current user by provided email,
	// if the user found in the database, then we can just logged in,
	// if not, then register that user.
	isExist, err := h.UserService.FindUserByEmail(profile.Email)
	// this error indicates user not found
	if err != nil {
		// register user and save their data into database
		createdId, err := h.UserService.RegisterUser(profile)
		if err != nil {
			log.Printf("Failed to register user: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
		}

		// Store the user's id in the session
		sess.Set("ID", createdId)

		// Save into memory session and.
		// saving also set a session cookie containing session_id
		if err := sess.Save(); err != nil {
			log.Printf("Failed to save user session: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user session")
		}

		// return immediately
		return c.Status(fiber.StatusOK).Redirect("/")
	}

	// Store the existed user's id in the session
	sess.Set("ID", isExist.ID)

	if err := sess.Save(); err != nil {
		log.Printf("Failed to save user session: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user session")
	}

	return c.Status(fiber.StatusOK).Redirect("/")
}
