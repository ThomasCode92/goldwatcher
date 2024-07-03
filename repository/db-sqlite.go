package repository

import (
	"database/sql"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{Conn: db}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS holdings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			amount REAL NOT NULL,
			purchase_date INTEGER NOT NULL,
			purchase_price INTEGER
		);
	`

	_, err := r.Conn.Exec(query)
	return err
}

func (r *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := `
		SELECT id, amount, purchase_date, purchase_price
		FROM holdings
		ORDER BY purchase_date;
	`

	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Holdings
	for rows.Next() {
		var h Holdings
		var unixTime int64

		err := rows.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
		if err != nil {
			return nil, err
		}

		h.PurchaseDate = time.Unix(unixTime, 0)

		all = append(all, h)
	}

	return all, nil
}

func (r *SQLiteRepository) InsertHolding(h Holdings) (*Holdings, error) {
	query := `INSERT INTO holdings (amount, purchase_date, purchase_price) VALUES (?, ?, ?);`

	res, err := r.Conn.Exec(query, h.Amount, h.PurchaseDate.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	h.ID = id
	return &h, nil
}
