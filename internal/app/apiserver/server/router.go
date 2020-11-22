package server

import (
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/middlewares"
)

func (s *server) ConfigureRouter() {

	s.router.Use(middlewares.ContentTypeMiddleware)
	s.router.HandleFunc("/register", s.SignUp()).Methods("POST")
	s.router.HandleFunc("/login", s.SignIn()).Methods("POST")
	apiRouter := s.router.PathPrefix("/api").Subrouter()
	
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.TokenAuthMiddleware)
	userRouter.HandleFunc("/info", s.GetUser()).Methods("GET")
	
	orderRouter := apiRouter.PathPrefix("/order").Subrouter()
	orderRouter.HandleFunc("/info", s.GetOrder()).Methods("GET")
	editOrderRouter := orderRouter.PathPrefix("/edit").Subrouter()
	editOrderRouter.Use(middlewares.TokenAuthMiddleware)
	editOrderRouter.HandleFunc("/create", s.CreateOrder()).Methods("POST")
	editOrderRouter.HandleFunc("/cancel", s.CancelOrder()).Methods("UPDATE")
	
	dishRouter := apiRouter.PathPrefix("/dish").Subrouter()
	dishRouter.HandleFunc("/info", s.GetDish()).Methods("GET")
	dishRouter.HandleFunc("/all", s.GetAllDishes()).Methods("GET")
	editDishRouter := dishRouter.PathPrefix("/edit").Subrouter()
	editDishRouter.Use(middlewares.TokenAuthMiddleware)
	editDishRouter.HandleFunc("/create", s.CreateDish()).Methods("POST")
	editDishRouter.HandleFunc("/delete", s.DeleteDish()).Methods("DELETE")
	
	institutionRouter := apiRouter.PathPrefix("/institution").Subrouter()
	institutionRouter.HandleFunc("/find", s.FindInstitutionsByTitle()).Methods("GET")
	editInstitutionRouter := institutionRouter.PathPrefix("/edit").Subrouter()
	editInstitutionRouter.Use(middlewares.TokenAuthMiddleware)
	editInstitutionRouter.HandleFunc("/create", s.CreateInstitution()).Methods("POST")
	editInstitutionRouter.HandleFunc("/delete", s.DeleteInstitution()).Methods("DELETE")
	
	paymentRouter := apiRouter.PathPrefix("/payment").Subrouter()
	paymentRouter.HandleFunc("/info", s.GetPaymentDetails()).Methods("GET")
	editPaymenRouter := paymentRouter.PathPrefix("/edit").Subrouter()
	editOrderRouter.Use(middlewares.TokenAuthMiddleware)
	editPaymenRouter.HandleFunc("/create", s.CreatePaymentDetails()).Methods("POST")
	editPaymenRouter.HandleFunc("/delete", s.DeletePaymentDetails()).Methods("DELETE")
}