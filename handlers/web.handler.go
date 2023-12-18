package handlers

import (
	"html/template"
	"quickshare/entities"
	"quickshare/services"

	"github.com/gofiber/fiber/v2"
)

type WebHandler interface {
	SendDashboard(c *fiber.Ctx) error
	SendLogin(c *fiber.Ctx) error
	SendPost(c *fiber.Ctx) error
	SendPosts(c *fiber.Ctx) error
}

type WebHandlerImpl struct {
	UtilService services.UtilService
	UserService services.UserService
	PostService services.PostService
}

func (h *WebHandlerImpl) SendDashboard(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	var user *entities.User

	if userID != nil {
		user, _ = h.UserService.FindUserByID(userID.(string))
	}

	return c.Render("pages/index", fiber.Map{
		"User": user,
	})
}

func (h *WebHandlerImpl) SendLogin(c *fiber.Ctx) error {
	return c.Render("pages/login", nil)
}

func (h *WebHandlerImpl) SendPost(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var user *entities.User
	if userID != nil {
		user, _ = h.UserService.FindUserByID(userID.(string))
	}

	postId := c.Params("id")
	if postId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	post, err := h.PostService.FindPostByID(postId)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	parsed := template.HTML(h.UtilService.ParseMD([]byte(post.Content)))
	return c.Render("pages/post", fiber.Map{
		"User":    user,
		"Post":    post,
		"Content": parsed,
	})
}

func (h *WebHandlerImpl) SendPosts(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user, _ := h.UserService.FindUserByID(userID.(string))

	posts, err := h.PostService.FindPostsByAuthor(userID.(string))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Render("pages/current-posts", fiber.Map{
		"User":  user,
		"Posts": posts,
	})
}
