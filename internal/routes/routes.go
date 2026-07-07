package routes

import (
	"net/http"

	"job-search/internal/handler"
	"job-search/internal/middleware"
)

func SetupRoutes() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/vacancies", middleware.Auth(handler.Vacancies))
	http.HandleFunc("/vacancy", middleware.Auth(handler.Vacancy))
	
	http.HandleFunc("/resumes", middleware.Auth(handler.Resumes))
	
	http.HandleFunc("/responses", middleware.Auth(handler.Responses))
	http.HandleFunc("/favorites", middleware.Auth(handler.Favorites))
	http.HandleFunc("/users/role", middleware.Auth(handler.ChangeUserRole))
}
