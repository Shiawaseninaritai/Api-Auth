package main

import (
	"github.com/Shiawaseninaritai/Api-Auth/controllers"
	"github.com/Shiawaseninaritai/Api-Auth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run()
}
