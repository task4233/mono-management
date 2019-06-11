package main

import(
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// routes
	// relative path from ./vender
	//"github.com/task4233/mono-management/backend/router"
	
	"github.com/gin-gonic/gin"
)

type(
	User struct{
		Name string `db: "name" json: "name"`
		Pass string `db: "pass" json: "pass"`
	}
)

func main() {
	router := gin.Default()

	// Routes
	// grouping
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			// mono
			mono := v1.Group("mono")
			{
				// CRUD
				mono.GET("/", getMonos)
				mono.POST("/new", createMono)
				mono.PUT("/:monoId", updateMonos)
				mono.DELETE("/:monoId", deleteMono)
			}
		}
	}

	srv := &http.Server{
	  Addr: ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// shutdown ther server with a timeout of 5 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ... ")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
	
}

func getMonos(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "get")
}

func createMono(c *gin.Context) {
	// samp
	var user User
	c.Bind(&user)

	log.Println(user)
	
	c.JSON(http.StatusOK, "create")
}

func updateMonos(c *gin.Context) {
	// samp
	var user User
	c.Bind(&user)

	log.Println(user)
	
	c.JSON(http.StatusOK, "update")
}

func deleteMono(c *gin.Context) {
	// samp
	var user User
	c.Bind(&user)

	log.Println(user)
	
	c.JSON(http.StatusOK, "delete")
}
