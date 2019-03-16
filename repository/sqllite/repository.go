package sqllite

import (
	"github.com/jmoiron/sqlx"
	"github.com/novikov1981/experiments/repository"
	"log"
)

type Repository struct {
	database *sqlx.DB
}

func NewRepository(dbPath string) (*Repository, error) {
	// TODO check if file sqlLiteFile exists
	dbFileExist := checkFunction(dbPath)
	// connect db
	database, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	repo := &Repository{database}
	// create tables if file did not exist before
	if !dbFileExist {
		err := repo.create()
		if err != nil {
			return nil, err
		}
	}
	return &Repository{database}, nil
}

func (r *Repository) create() error {
	// SYNTHESIS
	statement, err := r.database.Prepare(`
	CREATE TABLE IF NOT EXISTS synthesis (
		name TEXT PRIMARY KEY,
		scale INTEGER,
		created_at TEXT
	);`)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	// OLIGS
	statement, err = r.database.Prepare(`
		CREATE TABLE IF NOT EXISTS oligs (
			uuid TEXT PRIMARY KEY,
			synthesis_name TEXT,
			content TEXT,
			FOREIGN KEY(synthesis_name) REFERENCES synthesis(name)
				);`)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertSynthesis(name string, scale int64, oo []string) error {
	createdAt := .....
	// INSERT to synthesis
	for _, o := range oo {
		oligUUID := generate()
		// INSER to olig
	}
	//
	return nil
}

func (r *Repository) FindSynthesis(pattern string) ([]repository.Synthesis, error) {
	return nil, nil
}