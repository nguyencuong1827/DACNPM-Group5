package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title      string
	CategoryID int
	Author     string
	Body       Body
}

type Body struct {
	Paragraph []Paragraph
}

type Paragraph struct {
	Title   string
	Image   []string
	Content string
}
