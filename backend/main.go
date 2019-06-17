package main

import (
<<<<<<< HEAD
    "os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// user
	User struct {
		// id
		Id int `json:id`
		// name
		Name string `json:name`
		// hashed_pass
		HashedPass string `json:hashed_pass`
	}
)

func main() {
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
        DBMS := "mysql"
        USER := os.Getenv("MYSQL_USER")
        PASS := os.Getenv("MYSQL_PASSWORD")
        PROTOCOL := "tcp(db:3306)"
        DBNAME := os.Getenv("MYSQL_DATABASE")

        CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(DBMS, CONNECT)
		if err != nil {
			panic(err.Error())
		}
        
        fmt.Println("err")
		defer db.Close()

		resUsers := []User{}
		db.Find(&resUsers)

		fmt.Println(resUsers)

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.JSON(200, gin.H{
			"id":          resUsers[0].Id,
			"name":        resUsers[0].Name,
			"hashed_pass": resUsers[0].HashedPass,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
=======
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

    // router
	"github.com/task4233/mono-management/backend/router"
)

func main() {

	router := router.Create()

	srv := &http.Server{
		Addr:    ":8080",
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

>>>>>>> 6a7c5db036811f1a7394958cb9329d4440bcb852
}
