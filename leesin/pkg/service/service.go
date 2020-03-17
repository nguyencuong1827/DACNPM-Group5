package service

import (
	"DACNPM-Group5/leesin/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

const (
	pageParam    = "page"
	perPageParam = "perpage"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Register(r *gin.RouterGroup) {
	r.GET("/categories", s.getCategories)

	r.GET("/", s.getListArticles)
	r.GET("/article/:id")
}

func (s *service) getCategories(c *gin.Context) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		c.JSON(http.StatusOK, model.Http{
			StatusCode: -1,
			Message:    "error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Http{
		StatusCode: 1,
		Message:    "success",
		Data:       categories,
	})
}

func (s *service) getListArticles(c *gin.Context) {
	page := getIntQuery(c, pageParam)
	perPage := getIntQuery(c, perPageParam)
	if perPage == 0 {
		perPage = 10
	}
	prevs, err := s.repo.GetPreArticles(page, perPage)
	if err != nil {
		c.JSON(http.StatusOK, model.Http{
			StatusCode: -1,
			Message:    "error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Http{
		StatusCode: 1,
		Message:    "success",
		Data:       prevs,
	})
}

func (s *service) getArticle(c *gin.Context) {
	id := getUIntParam(c, "id")
	article, err := s.repo.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Http{
			StatusCode: -1,
			Message:    "error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Http{
		StatusCode: 1,
		Message:    "success",
		Data:       article,
	})
}

func getIntQuery(c *gin.Context, query string) int {
	str := c.Query(query)
	log.Println("info: query key", query, "query value", str)
	return cast.ToInt(str)
}

func getUIntParam(c *gin.Context, param string) uint {
	str := c.Param(param)
	log.Println("info: param key", param, "param value", str)
	return cast.ToUint(param)
}
