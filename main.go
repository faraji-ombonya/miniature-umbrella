package main

import (
	"github.com/faraji-fuji/miniature-umbrella/src/models"
	"github.com/faraji-fuji/miniature-umbrella/src/routes"
)

func init() {
	models.InitDB()
	routes.InitRouter()
}

func main() {

}
