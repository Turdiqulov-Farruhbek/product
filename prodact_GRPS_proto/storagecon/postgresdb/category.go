package postgresdb

import (
	"database/sql"
	ctr_proto "product/genproto"

	"github.com/google/uuid"
)
type CategoryRepo struct {
	db *sql.DB

}

func NewCategoryRepo(Db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db: Db}
}


func (cr *CategoryRepo) Createategory(category *ctr_proto.Category) (*ctr_proto.C_Void, error) {
	id := uuid.NewString()
    _, err := cr.db.Exec("INSERT INTO category (id, name) VALUES ($1, $2)", id, category.Name)
    if err!= nil {
        return nil, err
    }
    return &ctr_proto.C_Void{}, nil
}

func (cr *CategoryRepo) GetCategory(id string) (*ctr_proto.Category, error) {
	category := &ctr_proto.Category{}
    err := cr.db.QueryRow("SELECT id, name FROM category WHERE id = $1", id).
    Scan(&category.Id, &category.Name)
    if err!= nil {
        return nil, err
    }
    return category, nil
}

func (cr *CategoryRepo) GetCategories() ([]*ctr_proto.Category, error) {
	categories := []*ctr_proto.Category{}
    rows, err := cr.db.Query("SELECT id, name FROM category")
    if err!= nil {
        return nil, err
    }
    for rows.Next() {
        category := &ctr_proto.Category{}
        err := rows.Scan(&category.Id, &category.Name)
        if err!= nil {
            return nil, err
        }
        categories = append(categories, category)
    }
    return categories, nil
}

func (cr *CategoryRepo) UpdateCategory(category *ctr_proto.Category) error {
	_, err := cr.db.Exec("UPDATE category SET name = $1 WHERE id = $2", category.Name, category.Id)
    if err!= nil {
        return err
    }
    return nil
}

func (cr *CategoryRepo) DeleteCategory(id string) error {
	_, err := cr.db.Exec("UPDATE category SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0", id)
    if err!= nil {
        return err
    } 
    return nil
}