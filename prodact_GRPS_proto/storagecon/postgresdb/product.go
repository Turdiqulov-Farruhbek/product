package postgresdb

import (
	"database/sql"
	us_proto "product/genproto"
	"time"

	"github.com/google/uuid"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(Db *sql.DB) *ProductRepo {
	return &ProductRepo{db: Db}
}

func (r *ProductRepo) CreateProduct(product *us_proto.ProductCreateReq) error {
	id := uuid.NewString()
	_, err := r.db.Exec("INSERT INTO products (id, name, price, category_id, expired_at) VALUES ($1, $2, $3, $4)",
		id, product.Name, product.Price, product.CategoryId, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) GetProduct(id string) (*us_proto.Product, error) {
	product := &us_proto.Product{}
    err := r.db.QueryRow("SELECT id, name, price, category_id, expired_at FROM products WHERE id = $1", id).
        Scan(&product.Id, &product.Name, &product.Price, &product.Category.Id, &product.ExpiredAt)
    if err!= nil {
        return nil, err
    }
    return product, nil
}

func (r *ProductRepo) GetProducts() ([]*us_proto.Product, error) {
	products := []*us_proto.Product{}
    rows, err := r.db.Query("SELECT id, name, price, category_id, expired_at FROM products")
    if err!= nil {
        return nil, err
    }
    for rows.Next() {
        product := &us_proto.Product{}
        err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Category.Id, &product.ExpiredAt)
        if err!= nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}


func (r *ProductRepo) UpdateProduct(product *us_proto.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = $1, price = $2, category_id = $3, expired_at = $4 WHERE id = $5", 
	product.Name, product.Price, product.Category.Id, product.ExpiredAt, product.Id)
    if err!= nil {
        return err
    }
    return nil
}

func (r *ProductRepo) DeleteProduct(id string) error {
	_, err := r.db.Exec("UPDATE products SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0", id)
    if err!= nil {
        return err
    } 
    return nil
}