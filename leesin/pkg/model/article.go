package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title      string `json:"title"`
	CategoryID int    `json:"categoryid"`
	Author     string `json:"author"`
	Body       Body   `json:"body" gorm:"-"`
	BodyStr    string `json:"-"`
}

type Body struct {
	Paragraph []Paragraph
}

type Paragraph struct {
	Title   string
	Image   []string
	Content string
}
