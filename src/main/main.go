package main

import (
	"fmt"
	"log"
	"net/http"
	"todo_app/src/main/app/api"
	"todo_app/src/main/app/api/handler"
	"todo_app/src/main/app/repository"
	"todo_app/src/main/app/service"
	"todo_app/src/main/dependences/db"

	_ "todo_app/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := http.NewServeMux()
	fmt.Println("Server started on 8080 port ....")

	// User layer _-----------------
	{
		userRepo := repository.NewUserRepo(db)
		userService := service.NewUserService(userRepo)
		handler := handler.Handler{
			ServiceUser: userService,
		}
		api.RegisterUserRoutes(router, &handler)
	}

	// ----  Server Layer
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)
}
