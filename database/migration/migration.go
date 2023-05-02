package migration

import (
	"fiberv2/database"
	"fiberv2/model/entity"
	"fmt"
	"log"
)

func RunMingration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Success Migration")
}
