package business

//Repository interface that represents the database logic
type Repository interface {
	Save(Payments) (bool, error)
	FindOne(uint64) (interface{}, error)
	FindAll() ([]interface{}, error)
	Delete(uint64) (bool, error)
}
