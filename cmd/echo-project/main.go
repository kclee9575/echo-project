package main

import (
	"echo-project/config"
	"echo-project/logger"
	"echo-project/route"
	"echo-project/swagger"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	l := logger.Get()

	config.InitDatabase()
	swagger.Init(e)
	route.Init(e)

	l.Info().Str("port", "7777").Msgf("Start Server on port %s", "8585")
	l.Fatal().Err(e.Start(":7777")).Msg("Server Close")

}
