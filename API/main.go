package main

import (
	"github.com/DevTeam125/shopping-website/config"
	"github.com/DevTeam125/shopping-website/models"
	"github.com/DevTeam125/shopping-website/models/product"
	"github.com/DevTeam125/shopping-website/routers"
)

func main() {
	config.Init()
	models.Init()
	product.Init()
	routes := routers.InitRoutes()
	routes.Run()
}
