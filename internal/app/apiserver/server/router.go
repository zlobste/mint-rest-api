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
	userRouter.HandleFunc("/all", server.GetAllUsers()).Methods("GET")
	userRouter.HandleFunc("/block/{id}/{blocked}", server.BlockUser()).Methods("POST")
	userRouter.Use(middlewares.TokenAuthMiddleware)
	userRouter.HandleFunc("/info/{id}", server.GetUser()).Methods("GET")

	orderRouter := apiRouter.PathPrefix("/order").Subrouter()
	orderRouter.HandleFunc("/info/{id}", server.GetOrder()).Methods("GET")
	editOrderRouter := orderRouter.PathPrefix("/edit").Subrouter()
	editOrderRouter.Use(middlewares.TokenAuthMiddleware)
	editOrderRouter.HandleFunc("/create", server.CreateOrder()).Methods("POST")
	editOrderRouter.HandleFunc("/cancel/{id}", server.CancelOrder()).Methods("UPDATE")
	editOrderRouter.HandleFunc("/ready/{id}", server.SetStatusReady()).Methods("UPDATE")
	editOrderRouter.HandleFunc("/execute/{id}", server.GetOrderToExecute()).Methods("GET")

	dishRouter := apiRouter.PathPrefix("/dish").Subrouter()
	dishRouter.HandleFunc("/info/{id}", server.GetDish()).Methods("GET")
	dishRouter.HandleFunc("/all", server.GetAllDishes()).Methods("GET")
	editDishRouter := dishRouter.PathPrefix("/edit").Subrouter()
	editDishRouter.Use(middlewares.TokenAuthMiddleware)
	editDishRouter.HandleFunc("/create", server.CreateDish()).Methods("POST")
	editDishRouter.HandleFunc("/delete/{id}", server.DeleteDish()).Methods("DELETE")
	editDishRouter.HandleFunc("/sale/{id}", server.CalculateSale()).Methods("POST")

	institutionRouter := apiRouter.PathPrefix("/institution").Subrouter()
	institutionRouter.HandleFunc("/find", server.FindInstitutionsByTitle()).Methods("GET")
	institutionRouter.HandleFunc("/all", server.GetAllInstitutions()).Methods("GET")
	editInstitutionRouter := institutionRouter.PathPrefix("/edit").Subrouter()
	editInstitutionRouter.HandleFunc("/create", server.CreateInstitution()).Methods("POST")
	editInstitutionRouter.HandleFunc("/delete/{id}", server.DeleteInstitution()).Methods("DELETE")

	paymentRouter := apiRouter.PathPrefix("/payment").Subrouter()
	paymentRouter.HandleFunc("/info/{id}", server.GetPaymentDetails()).Methods("GET")
	editPaymenRouter := paymentRouter.PathPrefix("/edit").Subrouter()
	editOrderRouter.Use(middlewares.TokenAuthMiddleware)
	editPaymenRouter.HandleFunc("/create", server.CreatePaymentDetails()).Methods("POST")
	editPaymenRouter.HandleFunc("/delete/{id}", server.DeletePaymentDetails()).Methods("DELETE")
}
