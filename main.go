package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Tweet struct {
	gorm.Model
	TwitterID string
	Tweet     string
}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Tweet{})

	tweet1 := Tweet{
		TwitterID: "1",
		Tweet:     "Hello World!",
	}

	db.Create(&tweet1)

	tweet2 := Tweet{
		TwitterID: "2",
		Tweet:     "こんにちは、世界",
	}

	db.Create(&tweet2)

	var tweets []Tweet

	db.Find(&tweets)

	for _, tweet := range tweets {
		fmt.Println(tweet)
	}
}
