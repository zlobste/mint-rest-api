package store

type Store interface {
	User() UserRepository
	Organization() OrganizationRepository
}