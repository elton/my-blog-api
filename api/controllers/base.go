package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elton/my-blog-api/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Server server struct
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// InitializeDB initialize the server connect database and configure the routers
func (s *Server) InitializeDB(DbDriver, DbUser, DbPasswd, DbHost, DbPort, DbName string) {
	var err error
	if DbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPasswd, DbHost, DbPort, DbName)
		// GORM 定义了这些日志级别：Silent、Error、Warn、Info
		s.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{
			Logger:      logger.Default.LogMode(logger.Info),
			PrepareStmt: true,
			// DisableForeignKeyConstraintWhenMigrating: true,
			// SkipDefaultTransaction:                   true,
		})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected the %s database", DbDriver)
		}

		// database connection pool settings.
		// refer to https://www.alexedwards.net/blog/configuring-sqldb
		sqlDB, _ := s.DB.DB()
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(64)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(64)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		// database migration
		s.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}, &models.Comment{}, &models.Like{})
	}
}

// Run run the server.
func (s *Server) Run(port string) {
	// Logging
	date := time.Now().Format("20060102")
	f, err := os.Create("blog-" + date + ".log")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// Force log's color
	gin.ForceConsoleColor()
	// write the log to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	s.Router = gin.Default()

	s.initializeRouter()

	srv := &http.Server{
		Addr:           port,
		Handler:        s.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("Listening to port %s\n", port)

	// Graceful shutdown or restart
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	// lsof -i:8080 find pid first and  kill PID
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
