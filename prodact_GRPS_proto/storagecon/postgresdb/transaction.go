package postgresdb

import (
	"database/sql"
	tran_proto "product/genproto"

	"github.com/google/uuid"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(Db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db: Db}
}

func (r *TransactionRepo) CreateTransaction(transaction *tran_proto.TransCreateReq) (tran_proto.T_Void, error) {
	id := uuid.NewString()
	_, err := r.db.Exec("INSERT INTO transactions (id, basket_id, total_price) VALUES ($1, $2, $3)",
		id, transaction.BasketID, transaction.TotalPrice)
	if err != nil {
		return tran_proto.T_Void{}, err
	}
	return tran_proto.T_Void{}, nil
}

func (r *TransactionRepo) GetTransaction(id string) (*tran_proto.TransCreateReq, error) {
	transaction := &tran_proto.TransCreateReq{}
    err := r.db.QueryRow("SELECT basket_id, total_price FROM transactions WHERE id = $1", id).
        Scan(&transaction.BasketID, &transaction.TotalPrice)
    if err!= nil {
        return nil, err
    }
    return transaction, nil
}


func (r *TransactionRepo) GetTransactions() ([]*tran_proto.TransCreateReq, error) {
	transactions := []*tran_proto.TransCreateReq{}
    rows, err := r.db.Query("SELECT basket_id, total_price FROM transactions")
    if err!= nil {
        return nil, err
    }
    for rows.Next() {
        transaction := &tran_proto.TransCreateReq{}
        err := rows.Scan(&transaction.BasketID, &transaction.TotalPrice)
        if err!= nil {
            return nil, err
        }
        transactions = append(transactions, transaction)
    }
    return transactions, nil
}

func (r *TransactionRepo) UpdateTransaction(transaction *tran_proto.TransCreateReq) error {
	_, err := r.db.Exec("UPDATE transactions SET basket_id = $1, total_price = $2", 
	transaction.BasketID, transaction.TotalPrice, )
    if err!= nil {
        return err
    }
    return nil
}

func (r *TransactionRepo) DeleteTransaction(id string) error {
	_, err := r.db.Exec("UPDATE transactions SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0", id)
    if err!= nil {
        return err
    } 
    return nil
}