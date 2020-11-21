package server

import (
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/middlewares"
)

func (s *server) ConfigureRouter() {

	s.router.Use(middlewares.ContentTypeMiddleware)
	s.router.HandleFunc("/register", s.SignUp()).Methods("POST")
	s.router.HandleFunc("/login", s.SignIn()).Methods("POST")
	apiRouter := s.router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.TokenAuthMiddleware)
	userRouter := apiRouter.PathPrefix("/user").Subrouter()
		userRouter.HandleFunc("/info", s.GetUser()).Methods("GET")
	organizationRouter := apiRouter.PathPrefix("/organization").Subrouter()
		organizationRouter.HandleFunc("/info", s.GetUser()).Methods("GET")
	menuRouter := apiRouter.PathPrefix("/menu").Subrouter()
		menuRouter.HandleFunc("/create", s.CreateMenu()).Methods("POST")
		menuRouter.HandleFunc("/delete", s.DeleteMenu()).Methods("DELETE")
		menuRouter.HandleFunc("/info", s.GetMenu()).Methods("GET")
	orderRouter := apiRouter.PathPrefix("/order").Subrouter()
		orderRouter.HandleFunc("/create", s.CreateOrder()).Methods("POST")
		orderRouter.HandleFunc("/cancel", s.CancelOrder()).Methods("UPDATE")
		orderRouter.HandleFunc("/info", s.GetOrder()).Methods("GET")
	dishRouter := apiRouter.PathPrefix("/dish").Subrouter()
		dishRouter.HandleFunc("/create", s.CreateDish()).Methods("POST")
		dishRouter.HandleFunc("/delete", s.DeleteDish()).Methods("DELETE")
		dishRouter.HandleFunc("/info", s.GetDish()).Methods("GET")
	paymentRouter := apiRouter.PathPrefix("/payment").Subrouter()
		paymentRouter.HandleFunc("/create", s.CreatePaymentDetails()).Methods("POST")
		paymentRouter.HandleFunc("/delete", s.DeletePaymentDetails()).Methods("DELETE")
		paymentRouter.HandleFunc("/info", s.GetPaymentDetails()).Methods("GET")
}
