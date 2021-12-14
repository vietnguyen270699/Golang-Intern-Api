package configs

import (
	"learn-golang/pkg/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoute api
func SetupRoute() *gin.Engine {

	r := gin.Default()

	//client IP can be trusted
	//r.SetTrustedProxies([]string{"192.168.1.2"})

	r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"GET", "POST", "DELETE"},
	// 	AllowHeaders: []string{"Origin"},
	// }))

	client := r.Group("/api")
	{
		client.GET("/interns", api.GetAllIntern)
		client.GET("/interns/:id", api.GetInternById)
		client.POST("/interns", api.CreateIntern)
		client.PUT("/interns/update/:id", api.UpdateIntern)
		client.DELETE("/interns/delete/:id", api.DeleteIntern)
	}

	return r
}
