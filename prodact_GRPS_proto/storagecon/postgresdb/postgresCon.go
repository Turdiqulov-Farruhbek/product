package postgresdb

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

type Storage struct {
	Db   *sql.DB
	User *UserRepo
	Basket *BasketRepo
	Category *CategoryRepo
	Item *ItemRepo
	Product *ProductRepo
	Transaction *TransactionRepo

}

func NewPostgresStorage() (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "root", "localhost", 5432, "uyishi47_db")
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	usR := NewUserRepo(db)
	basR := NewBasketRepo(db)
	catR := NewCategoryRepo(db)
	itemR := NewItemRepo(db)
	proR := NewProductRepo(db)
	tranR := NewTransactionRepo(db)


	storConn := &Storage{
		Db: db, 
		User: usR, 
		Basket: basR,
		Category: catR,
		Item: itemR,
        Product: proR,
        Transaction: tranR,
	}


	return storConn, err
}
