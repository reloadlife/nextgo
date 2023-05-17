package main

import (
	"github.com/reloadlife/nextgo/internal/services"
	"github.com/reloadlife/nextgo/internal/template"
	"github.com/reloadlife/nextgo/internal/webserver"
	"github.com/reloadlife/nextgo/nextjs"
)

func main() {
	template.NewTemplate()

	services.SetupServices(
		nextjs.SetupFrontendService(),
		webserver.SetupWebserver(),
	)

	select {} // block the main thread.
}
