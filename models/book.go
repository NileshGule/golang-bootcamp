package models

import (
	"fmt"
)

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

func GetBooks() []Book {
	return books
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Microservices",
		Author:        "Neal Ford",
		YearPublished: 2015,
	},
	Book{
		ID:            2,
		Title:         "Production ready Microservices",
		Author:        "Susan Fowler",
		YearPublished: 2017,
	},
	Book{
		ID:            3,
		Title:         "DevOps Handbook",
		Author:        "Sam Newman",
		YearPublished: 2015,
	},
	Book{
		ID:            4,
		Title:         "The Phoenix Project",
		Author:        "Gene Kim",
		YearPublished: 2015,
	},
	Book{
		ID:            5,
		Title:         "Microservices",
		Author:        "Neal Ford",
		YearPublished: 2015,
	},
	Book{
		ID:            6,
		Title:         "Production ready Microservices",
		Author:        "Susan Fowler",
		YearPublished: 2017,
	},
	Book{
		ID:            7,
		Title:         "DevOps Handbook",
		Author:        "Sam Newman",
		YearPublished: 2015,
	},
	Book{
		ID:            8,
		Title:         "The Phoenix Project",
		Author:        "Gene Kim",
		YearPublished: 2015,
	},
	Book{
		ID:            9,
		Title:         "Microservices",
		Author:        "Neal Ford",
		YearPublished: 2015,
	},
	Book{
		ID:            10,
		Title:         "Production ready Microservices",
		Author:        "Susan Fowler",
		YearPublished: 2017,
	},
}
