package engine

import (
	"gorestapi/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CSRFToken)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("/", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("/", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentType, echo.HeaderOrigin, echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookieDomain: os.Getenv("API_DOMAIN"),
		CookiePath:   "/",
		// CookieMaxAge: 60,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
	}))

	return e
}
