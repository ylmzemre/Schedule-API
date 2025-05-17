package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ylmzemre/Schelude-API/config"
	"github.com/ylmzemre/Schelude-API/models"
	"github.com/ylmzemre/Schelude-API/service"
	"time"
)

type AuthHandler struct {
	svc service.StudentService
}

func NewAuthHandler(s service.StudentService) *AuthHandler { return &AuthHandler{svc: s} }

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	var r req
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	st, err := h.svc.Register(&models.Student{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(st)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var r req
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	st, err := h.svc.Login(r.Email, r.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}
	// Token
	claims := jwt.MapClaims{
		"sub": st.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(config.Cfg.JWTSecret))
	return c.JSON(fiber.Map{"token": ss})
}
