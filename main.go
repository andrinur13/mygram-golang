package main

import (
	"mygram-golang/conf"
	"mygram-golang/model/entity"
	"mygram-golang/repository"
)

func main() {
	db := conf.InitDB()

	userAndri := entity.User{
		Username: "andrinur13",
		Email:    "andribis13@gmaill.com",
		Password: "Bismillah",
		Age:      21,
	}

	userRepository := repository.NewRepository(db)
	userRepository.Save(userAndri)
}
