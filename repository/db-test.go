package repository

import "time"

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) Migrate() error {
	return nil
}

func (r *TestRepository) AllHoldings() ([]Holdings, error) {
	var all []Holdings

	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	all = append(all, h)

	h = Holdings{
		Amount:        2,
		PurchaseDate:  time.Now(),
		PurchasePrice: 2000,
	}
	all = append(all, h)

	return all, nil
}

func (r *TestRepository) GetHoldingByID(id int64) (*Holdings, error) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	return &h, nil
}

func (r *TestRepository) InsertHolding(h Holdings) (*Holdings, error) {
	return &h, nil
}

func (r *TestRepository) UpdateHolding(id int64, updated Holdings) error {
	return nil
}

func (r *TestRepository) DeleteHolding(id int64) error {
	return nil
}
