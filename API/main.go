package main

import (
	"github.com/DevTeam125/shopping-website/config"
	"github.com/DevTeam125/shopping-website/models"
	"github.com/DevTeam125/shopping-website/models/product"
	"github.com/DevTeam125/shopping-website/pkg/logging"
	"github.com/DevTeam125/shopping-website/routers"
)

func main() {
	config.Init()

	logging.Init()
	defer logging.Logger.Sync()

	models.Init()
	product.Init()

	routes := routers.InitRoutes()
	routes.Run(":80")
}
