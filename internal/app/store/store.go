package store

type Store interface {
	User() UserRepository
	Order() OrderRepository
	Dish() DishRepository
	Institution() InstitutionRepository
	PaymentDetails() PaymentDetailsRepository
}