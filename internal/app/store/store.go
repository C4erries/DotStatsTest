package store

//Интерфейс хранилища, используемый в sqlstore
type Store interface {
	User() UserRepository
}
