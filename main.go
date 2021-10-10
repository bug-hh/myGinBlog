package main

import (
	"myGinBlog/model"
	"myGinBlog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}

