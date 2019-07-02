package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	// routes
	// relative path from ./vender
	//"github.com/task4233/mono-management/backend/router"

	// router
	"app/internal"
	"app/router"

	"github.com/gin-gonic/gin"
)

var reg = regexp.MustCompile("https?:\\/\\/(.+)\\.example\\.com:?.*")

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			// for preflight
			origin := c.Request.Header.Get("Origin")

			r := reg.Copy()
			if r.MatchString(origin) {
				headers := c.Request.Header.Get("Access-Control-Request-Headers")

				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
				c.Writer.Header().Set("Access-Control-Allow-Headers", headers)

				c.Data(200, "text/plain", []byte{})
				c.Abort()
			} else {
				c.Data(403, "text/plain", []byte{})
				c.Abort()
			}
		} else {
			// for actual response
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			//c.Writer.Header().Set("Access-Control-Expose-Headers", "")
			c.Next()
		}

		return
	}
}

func main() {
	internal.InitDB()
	defer internal.CloseDB()

	router := router.Create()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.Use(CORS())

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
