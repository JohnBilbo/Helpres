package database

type repository struct {
	//
}

// NewRepository - конструктор репы
func NewRepository() Repository {
	return &repository{}
}

type Repository interface {
}
