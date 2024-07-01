package postgresdb

import (
	"database/sql"
	us_proto "product/genproto"

	"github.com/google/uuid"
)

type BasketRepo struct {
	db *sql.DB
}

func NewBasketRepo(db *sql.DB) *BasketRepo{
	return &BasketRepo{db: db}
}

func (repo *BasketRepo) CreateBasket(basket *us_proto.CreatedBasket)error{
	id := uuid.NewString()
	_, err := repo.db.Exec("INSERT INTO basket (id, user_id) VALUES ($1, $2)",id, basket.UserId)
    if err!= nil {
        return err
    }
    return nil
}

func (repo *BasketRepo) GetBasket(id string) (*us_proto.Basket, error) {
	basket := &us_proto.Basket{}
    err := repo.db.QueryRow("SELECT id, user_id FROM basket WHERE id = $1", id).Scan(&basket.Id, &basket.UserID)
    if err!= nil {
        return nil, err
    }
    return basket, nil

}

func (repo *BasketRepo) GetBaskets() ([]*us_proto.Basket, error) {
	baskets := []*us_proto.Basket{}
    rows, err := repo.db.Query("SELECT id, user_id FROM basket")
    if err!= nil {
        return nil, err
    }
    for rows.Next() {
        basket := &us_proto.Basket{}
        err := rows.Scan(&basket.Id, &basket.UserID)
        if err!= nil {
            return nil, err
        }
        baskets = append(baskets, basket)
    }
    return baskets, nil
}

func (repo *BasketRepo) UpdateBasket(basket *us_proto.Basket) error {
	_, err := repo.db.Exec("UPDATE basket SET user_id = $1 WHERE id = $2", basket.UserID, basket.Id)
    if err!= nil {
        return err
    }
    return nil
}


func (repo *BasketRepo) DeleteBasket(id string) error {
	_, err := repo.db.Exec("UPDATE basket SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0", id)
    if err!= nil {
        return err
    } 
    return nil
}