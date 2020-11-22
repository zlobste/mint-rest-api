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
	
	organizationRouter := apiRouter.PathPrefix("/organization").Subrouter()
	organizationRouter.HandleFunc("/info", s.GetUser()).Methods("GET")
	
	menuRouter := apiRouter.PathPrefix("/menu").Subrouter()
	menuRouter.HandleFunc("/info", s.GetMenu()).Methods("GET")
	menuRouter.HandleFunc("/create", s.CreateMenu()).Methods("POST")
	menuRouter.HandleFunc("/delete", s.DeleteMenu()).Methods("DELETE")
	
	orderRouter := apiRouter.PathPrefix("/order").Subrouter()
	orderRouter.HandleFunc("/info", s.GetOrder()).Methods("GET")
	orderRouter.HandleFunc("/create", s.CreateOrder()).Methods("POST")
	orderRouter.HandleFunc("/cancel", s.CancelOrder()).Methods("UPDATE")
	
	dishRouter := apiRouter.PathPrefix("/dish").Subrouter()
	dishRouter.HandleFunc("/info", s.GetDish()).Methods("GET")
	dishRouter.HandleFunc("/all", s.GetAllDishes()).Methods("GET")
	editDishRouter := apiRouter.PathPrefix("/edit").Subrouter()
	editDishRouter.Use(middlewares.TokenAuthMiddleware)
	editDishRouter.HandleFunc("/create", s.CreateDish()).Methods("POST")
	editDishRouter.HandleFunc("/delete", s.DeleteDish()).Methods("DELETE")
	
	paymentRouter := apiRouter.PathPrefix("/payment").Subrouter()
	paymentRouter.HandleFunc("/info", s.GetPaymentDetails()).Methods("GET")
	paymentRouter.HandleFunc("/create", s.CreatePaymentDetails()).Methods("POST")
	paymentRouter.HandleFunc("/delete", s.DeletePaymentDetails()).Methods("DELETE")
}
