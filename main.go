package main

import (
	"fmt"
	"log"
	"os"

	routes "pelipper/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("PELIPPER_PORT"))))
}
