package postgresdb

import (
	"database/sql"
	us_proto "product/genproto"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(Db *sql.DB) *UserRepo {
	return &UserRepo{db: Db}
}

func (r *UserRepo) CreateUser(user *us_proto.UserCreateReq) error {
	id := uuid.NewString()
	_, err := r.db.Exec("INSERT INTO users (id, name, balance, email, phone) VALUES ($1, $2, $3, $4, $5)",
		id, user.Name, user.Balance, user.Email, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetByUser(id string) (*us_proto.User, error) {
	user := &us_proto.User{}
	err := r.db.QueryRow("SELECT id, name, balance, email, phone FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Balance, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUsers() (*us_proto.GetAllUserResp, error) {
	users := &us_proto.GetAllUserResp{}
	rows, err := r.db.Query("SELECT id, name, balance, email, phone FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &us_proto.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Balance, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, user)
	}
	return users, nil
}

func (r *UserRepo) UpdateUser(user *us_proto.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, balance = $2, email = $3, phone = $4 WHERE id = $5", user.Name, user.Balance, user.Email, user.Phone, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteUser(id string) error {
	_, err := r.db.Exec("UPDATE users SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0", id)
	if err != nil {
		return err
	}
	return nil
}
