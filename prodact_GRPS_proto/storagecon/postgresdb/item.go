package postgresdb

import (
	"database/sql"
	ctr_proto "product/genproto"

	"github.com/google/uuid"
)

type ItemRepo struct {
	db *sql.DB
}

func NewItemRepo(db *sql.DB) *ItemRepo {
	return &ItemRepo{db: db}
}

func (r *ItemRepo) CreateItem(item  *ctr_proto.CreateItem) (*ctr_proto.I_Void, error) {
	id := uuid.NewString()
    _, err := r.db.Exec("INSERT INTO item (id, basket_id, product_id, quantity) VALUES ($1, $2, $3, $4)", id, item.BasketID, item.ProductID, item.Quantity)
    if err!= nil {
        return nil, err
    }
    return &ctr_proto.I_Void{}, nil
}

func (r *ItemRepo) GetItem(id string) (*ctr_proto.Item, error) {
	item := &ctr_proto.Item{}
    err := r.db.QueryRow("SELECT id, basket_id, product_id, quantity FROM item WHERE id = $1", id).
	Scan(&item.Basket.Id, &item.Basket.UserID, &item.Product, &item.Quantity)
    if err!= nil {
        return nil, err
    }
    return item, nil
}

func (r *ItemRepo) GetItems() ([]*ctr_proto.Item, error) {
	items := []*ctr_proto.Item{}
    rows, err := r.db.Query("SELECT id, basket_id, product_id, quantity FROM item")
    if err!= nil {
        return nil, err
    }
    for rows.Next() {
        item := &ctr_proto.Item{}
        err := rows.Scan(&item.Basket.Id, &item.Basket.UserID, &item.Product, &item.Quantity)
        if err!= nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}


func (r *ItemRepo) UpdateItem(item *ctr_proto.Item) error {
	_, err := r.db.Exec("UPDATE item SET basket_id = $1, product_id = $2, quantity = $3 WHERE id = $4", 
	&item.Basket.Id, &item.Basket.UserID, &item.Product, &item.Quantity, &item.Basket.Id)
    if err!= nil {
        return err
    }
    return nil
}


func (r *ItemRepo) DeleteItem(id string) error {
	_, err := r.db.Exec("UPDATE item SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0", id)
    if err!= nil {
        return err
    } 
    return nil
}
