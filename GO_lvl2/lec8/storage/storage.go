package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Instance of storage
type Storage struct {
	config *Config
	// DataBase FileDescriptor
	db                *sql.DB
	UserRepository    *UserRepository
	ArticleRepository *ArticleRepository
}

// Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection method
func (storage *Storage) Open() error {

	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")

	return nil
}

// Close connection
func (storage *Storage) Close() {

}

// Public Repo for Article
func (s *Storage) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		storage: s,
	}
	return nil
}

// Public Repo for User
func (s *Storage) Article() *ArticleRepository {
	if s.ArticleRepository != nil {
		return s.ArticleRepository
	}
	s.ArticleRepository = &ArticleRepository{
		storage: s,
	}
	return nil
}
