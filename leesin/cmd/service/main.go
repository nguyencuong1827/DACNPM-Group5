package main

import (
	"DACNPM-Group5/leesin/pkg/articlerepository"
	"DACNPM-Group5/leesin/pkg/categoryrepository"
	"DACNPM-Group5/leesin/pkg/database"
	"DACNPM-Group5/leesin/pkg/model"
	"DACNPM-Group5/leesin/pkg/service"
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

	db.AutoMigrate(model.Article{}, model.Category{})

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
