package main

import (
	"github.com/DACNPM-Group5/leesin/pkg/articlerepository"
	"github.com/DACNPM-Group5/leesin/pkg/categoryrepository"
	"github.com/DACNPM-Group5/leesin/pkg/database"
	"github.com/DACNPM-Group5/leesin/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// database
	db, err := database.NewDatabase("root", "root", "leesin", "localhost",
		"3306")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db = db.Debug()

	// repository
	articleDBRepo := articlerepository.NewDBRepository(db)
	categoryDBRepo := categoryrepository.NewDBRepository(db)
	articleRepo := articlerepository.NewRepository(articleDBRepo)
	categoryRepo := categoryrepository.NewRepository(categoryDBRepo)
	repository := service.NewRepository(articleRepo, categoryRepo)

	//service
	r := gin.Default()
	service.NewService(repository).Register(r.Group("/"))
	r.Run() // port 8080
}
