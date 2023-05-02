package store

import (
	"database/sql"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Store struct {
	DB     *sql.DB
	config *Config
}

func (s *Store) ConfigureStore() error {
	config := NewConfig()

	_, err := toml.DecodeFile("configs/database.toml", config)
	if err != nil {
		return err
	}

	s.config = config
	return nil
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.DB = db
	return nil
}

func (s *Store) Close() {
	s.DB.Close()
}
