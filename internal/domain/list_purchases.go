package domain

import (
	"github.com/google/uuid"
	"time"
)

type ListPurchase struct {
	ID            string
	Product       Product
	User          string
	Status        string
	Notes         string
	PlacePurchase string
	Quantity      int
	InsertionDate time.Time
	DatePurchase  time.Time
}

type Product struct {
	NameProduct string
	UnityPrice  float64
	Category    string
}

func (listPurchase *ListPurchase) NewListPurchase() *ListPurchase {
	return &ListPurchase{
		ID:            uuid.NewString(),
		Product:       listPurchase.Product,
		User:          listPurchase.User,
		Status:        listPurchase.Status,
		Notes:         listPurchase.Notes,
		PlacePurchase: listPurchase.PlacePurchase,
		Quantity:      listPurchase.Quantity,
		InsertionDate: time.Now(),
		DatePurchase:  listPurchase.DatePurchase,
	}
}
