package server

import (
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/middlewares"
)

func (server *server) ConfigureRouter() {

	server.router.Use(middlewares.ContentTypeMiddleware)
	server.router.HandleFunc("/register", server.SignUp()).Methods("POST")
	server.router.HandleFunc("/login", server.SignIn()).Methods("POST")
	apiRouter := server.router.PathPrefix("/api").Subrouter()
	
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.TokenAuthMiddleware)
	userRouter.HandleFunc("/info", server.GetUser()).Methods("GET")
	
	orderRouter := apiRouter.PathPrefix("/order").Subrouter()
	orderRouter.HandleFunc("/info", server.GetOrder()).Methods("GET")
	editOrderRouter := orderRouter.PathPrefix("/edit").Subrouter()
	editOrderRouter.Use(middlewares.TokenAuthMiddleware)
	editOrderRouter.HandleFunc("/create", server.CreateOrder()).Methods("POST")
	editOrderRouter.HandleFunc("/cancel", server.CancelOrder()).Methods("UPDATE")
	
	dishRouter := apiRouter.PathPrefix("/dish").Subrouter()
	dishRouter.HandleFunc("/info", server.GetDish()).Methods("GET")
	dishRouter.HandleFunc("/all", server.GetAllDishes()).Methods("GET")
	editDishRouter := dishRouter.PathPrefix("/edit").Subrouter()
	editDishRouter.Use(middlewares.TokenAuthMiddleware)
	editDishRouter.HandleFunc("/create", server.CreateDish()).Methods("POST")
	editDishRouter.HandleFunc("/delete", server.DeleteDish()).Methods("DELETE")
	
	institutionRouter := apiRouter.PathPrefix("/institution").Subrouter()
	institutionRouter.HandleFunc("/find", server.FindInstitutionsByTitle()).Methods("GET")
	editInstitutionRouter := institutionRouter.PathPrefix("/edit").Subrouter()
	editInstitutionRouter.Use(middlewares.TokenAuthMiddleware)
	editInstitutionRouter.HandleFunc("/create", server.CreateInstitution()).Methods("POST")
	editInstitutionRouter.HandleFunc("/delete", server.DeleteInstitution()).Methods("DELETE")
	
	paymentRouter := apiRouter.PathPrefix("/payment").Subrouter()
	paymentRouter.HandleFunc("/info", server.GetPaymentDetails()).Methods("GET")
	editPaymenRouter := paymentRouter.PathPrefix("/edit").Subrouter()
	editOrderRouter.Use(middlewares.TokenAuthMiddleware)
	editPaymenRouter.HandleFunc("/create", server.CreatePaymentDetails()).Methods("POST")
	editPaymenRouter.HandleFunc("/delete", server.DeletePaymentDetails()).Methods("DELETE")
}