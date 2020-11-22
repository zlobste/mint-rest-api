package store

type Store interface {
	User() UserRepository
	Order() OrderRepository
	Dish() DishRepository
	PaymentDetails() PaymentDetailsRepository
}