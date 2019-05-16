package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ccod/gosu-server/pkg/models"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("must pass a command to migrate tool")
		return
	}

	command := os.Args[1]
	db, err := gorm.Open("postgres", "host=192.168.99.100 port=5432 user=postgres password=gosu database=postgres sslmode=disable")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if command == "up" {
		fmt.Println("called the up command")
		db.AutoMigrate(&models.Todo{}, &models.Player{}, &models.Challenge{}, &models.History{})
		return
	}

	if command == "down" {
		fmt.Println("called the down command")
		db.DropTable(&models.Todo{}, &models.Player{}, &models.Challenge{}, &models.History{})
		return
	}

	fmt.Println("did not recognize the command called")
}
