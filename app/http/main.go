package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	MateriaEcho "github.com/futurisen-solution/materia/echo"
	MateriaMiddleware "github.com/futurisen-solution/materia/middleware"
	"github.com/futurisen-solution/symphonic-skeleton/bootstrap"
	"github.com/futurisen-solution/symphonic-skeleton/bootstrap/http"
	"github.com/futurisen-solution/symphonic-skeleton/shared/log"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	bootstrap.Boot()

	e := echo.New()

	e.Use(MateriaMiddleware.Logger(log.Logger()))
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:               nil,
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: true,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src *; font-src 'self'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'",
		CSPReportOnly:         false,
		HSTSPreloadEnabled:    true,
		ReferrerPolicy:        "same-origin",
	}))
	e.Use(MateriaMiddleware.Symphonic)
	e.Use(MateriaMiddleware.Aegis)

	e.Binder = MateriaEcho.SymphonicBinder{}
	e.Logger = log.Logger()
	e.Validator = MateriaEcho.NewSymphonicValidator()

	kernel := http.Kernel{}

	kernel.Routes(e)

	go func() {
		if err := e.Start(fmt.Sprintf("%s:%d", facades.Config().Get("app.host"), facades.Config().GetInt("app.port"))); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
