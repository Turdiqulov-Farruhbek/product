package service

import (
	"context"
	"fmt"
	tran_proto "product/genproto"
	posDB "product/storagecon/postgresdb"
)

type transService struct {
	posgresDB posDB.Storage
	tran_proto.UnimplementedTransactionServiceServer
}

func NewTransactionService(PDB *posDB.Storage) *transService {
	return &transService{posgresDB: *PDB}
}

func (ts *transService) Create(ctx context.Context, trans *tran_proto.TransCreateReq) (*tran_proto.I_Void, error) {
	_, err := ts.posgresDB.Transaction.CreateTransaction(trans)
    if err!= nil {
        return nil, err
    } else {
        fmt.Println("transaction created successfully")
    }
    return &tran_proto.I_Void{}, nil
}

func (ts *transService) Get(ctx context.Context, id *tran_proto.ByIdTrans) (*tran_proto.TransCreateReq, error) {
	transaction, err := ts.posgresDB.Transaction.GetTransaction(id.BasketID)
    if err!= nil {
        return nil, err
    }
    return transaction, nil
}


func (ts *transService) GetAll(ctx context.Context, filter *tran_proto.FilterTrans) ([]*tran_proto.TransCreateReq, error) {
	transactions, err := ts.posgresDB.Transaction.GetTransactions()
    if err!= nil {
        return nil, err
    }
    return transactions, nil
}


func (ts *transService) Update(ctx context.Context, trans *tran_proto.TransCreateReq) (*tran_proto.I_Void, error) {
	err := ts.posgresDB.Transaction.UpdateTransaction(trans)
    if err!= nil {
        return nil, err
    }
    return nil, err
}


func (ts *transService) Delete(ctx context.Context, id *tran_proto.ByIdTrans) (*tran_proto.I_Void, error) {
	err := ts.posgresDB.Transaction.DeleteTransaction(id.BasketID)
    if err!= nil {
        return nil, err
    }
    return nil, err
}