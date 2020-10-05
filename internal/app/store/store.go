package store

type Store interface {
	User() UserRepository
	Organization() OrganizationRepository
	Menu() MenuRepository
	Order() OrderRepository
}