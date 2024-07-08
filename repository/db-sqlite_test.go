package repository

import (
	"testing"
	"time"
)

func TestSQLRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed", err)
	}
}

func TestSQLRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed", err)
	}

	if result.ID <= 0 {
		t.Error("invalid  id sent back", result.ID)
	}
}

func TestSQLRepository_AllHoldings(t *testing.T) {
	results, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed", err)
	}

	if len(results) != 1 {
		t.Error("wrong number of rows returned; expected 1, but got", len(results))
	}
}

func TestSQLRepository_GetHoldingByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("get by id failed", err)
	}

	if h.PurchasePrice != 1000 {
		t.Error("wrong purchase price returned; expected 100 but got", h.PurchasePrice)
	}

	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("get one returned value for non-existing id", h)
	}
}

func TestSQLRepository_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	h.PurchasePrice = 1001

	err = testRepo.UpdateHolding(1, *h)
	if err != nil {
		t.Error("update failed", err)
	}
}

func TestSQLRepository_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("delete failed", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}

	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error when trying to delete non-existing id")
	}
}
