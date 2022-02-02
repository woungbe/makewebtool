package configs

import (
	"html/template"
	"io"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type httpConfig struct {
	Service string `json:"service"`
	Host    string `json:"host"`
	Port    string `json:"port"`
}

type CustomValidator struct {
	validator *validator.Validate
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func getCookieStore() *sessions.CookieStore {
	// In real-world applications, use env variables to store the session key.
	sessionKey := "test-session-key"
	return sessions.NewCookieStore([]byte(sessionKey))
}
