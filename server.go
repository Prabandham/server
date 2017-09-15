package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"

	"flag"
	"fmt"
	"log"
	"net/http"
)

var Router *gin.Engine
var CorsConfig cors.Config
var WebSocket *melody.Melody

func init() {
	prodPrt := flag.Bool("prod", false, "Production Environment")
	flag.Parse()
	fmt.Printf("prod options %v", *prodPrt)

	if *prodPrt {
		gin.SetMode(gin.ReleaseMode)
	}

	if Router == nil {
		Router = gin.Default()
	}

	if WebSocket == nil {
		WebSocket = melody.New() // This is what will handle the web sockets.

	}

	CorsConfig := cors.DefaultConfig()
	CorsConfig.AllowAllOrigins = true
	CorsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	Router.Use(cors.New(CorsConfig))

	// app := Router.Group("/")
	// app.GET("/", endpoints.Ping)

	// Websocket stuff TODO move this to it's own process / endpoint
	// app.GET("/ws", func(c *gin.Context) {
	// 	ws.HandleRequest(c.Writer, c.Request)
	// })

	// ws.HandleConnect(func(s *melody.Session) {
	// fmt.Printf("New connection from, %v", s.Keys)
	// })

	// api := app.Group("/api")
	// Login handler
	// api.POST("/authorize", endpoints.Login)

	// authorized := api.Group("/")
	// authorized.Use(endpoints.CheckAuth())
	// {
	// Logout
	// authorized.DELETE("/logout", endpoints.Logout)
	// }
}

// Start takes the port on which the server has to be started.
// Ex :3000
// or :8080
// NOTE the : is compulsory.
func Start(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: Router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}
