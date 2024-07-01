package service

import (
	"context"
	"fmt"
	bas_proto "product/genproto"
	posDB "product/storagecon/postgresdb"
)

type basketService struct {
	posgresDB posDB.Storage
	bas_proto.UnimplementedBasketServiceServer
}

func NewBasketService(PDB *posDB.Storage) *basketService {
	return &basketService{posgresDB: *PDB}
}

func (bs *basketService) Create(ctx context.Context, basket *bas_proto.CreatedBasket) (*bas_proto.U_Void, error) {
	err := bs.posgresDB.Basket.CreateBasket(basket)
    if err!= nil {
        return nil, err
    }else{
        fmt.Println("basket created successfully")
    }
    return nil, err
}

func (bs *basketService) Get(ctx context.Context, id *bas_proto.ByIdBasket) (*bas_proto.Basket, error) {
	basket, err := bs.posgresDB.Basket.GetBasket(id.Id)
    if err!= nil {
        return nil, err
    }
    return basket, nil
}

func (bs *basketService) GetAll(ctx context.Context, filter *bas_proto.FilterBasket) (*bas_proto.GetAllBasketRes, error) {
	baskets, err := bs.posgresDB.Basket.GetBaskets()
    if err!= nil {
        return nil, err
    }
    return &bas_proto.GetAllBasketRes{Baskets: baskets}, nil
}


func (bs *basketService) Update(ctx context.Context, basket *bas_proto.Basket) (*bas_proto.B_Void, error) {
	err := bs.posgresDB.Basket.UpdateBasket(basket)
    if err!= nil {
        return nil, err
    }
    return nil, err
}

func (bs *basketService) Delete(ctx context.Context, id *bas_proto.ByIdBasket) (*bas_proto.B_Void, error) {
	err := bs.posgresDB.Basket.DeleteBasket(id.Id)
    if err!= nil {
        return nil, err
    }
    return nil, err
}