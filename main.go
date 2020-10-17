package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/NileshGule/golang-bootcamp/webservice/models"
)

var cache = map[int]models.Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// controllers.RegisterControllers()
	// _ = http.ListenAndServe(":3000", nil)

	fmt.Printf("Type of random number generator : %T", rnd)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10)

		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From Database")
				fmt.Println(b)
			}
		}(id)

		fmt.Printf("Book not found with id : '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (models.Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (models.Book, bool) {
	time.Sleep(100 * time.Millisecond)

	books := models.GetBooks()
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return models.Book{}, false
}
