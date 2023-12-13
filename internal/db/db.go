package db

// db interface with save and find methods

type DB interface {
	Save(url string) (int, error)
	Find(id int) (string, error)
}
