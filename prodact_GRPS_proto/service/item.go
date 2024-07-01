package service

import (
	"context"
	"fmt"
	item_proto "product/genproto"
	posDB "product/storagecon/postgresdb"
)

type itemService struct {
	posgresDB posDB.Storage
	item_proto.UnimplementedItemServiceServer
}

func NewItemService(PDB *posDB.Storage) *itemService {
	return &itemService{posgresDB: *PDB}
}

func (is *itemService) Create(ctx context.Context, item *item_proto.CreateItem) (*item_proto.I_Void, error) {
	_, err := is.posgresDB.Item.CreateItem(item)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("item created successfully")
	}
	return &item_proto.I_Void{}, nil
}

func (is *itemService) Get(ctx context.Context, id *item_proto.ByIdItem) (*item_proto.Item, error) {
	item, err := is.posgresDB.Item.GetItem(id.Id)
    if err!= nil {
        return nil, err
    }
    return item, nil
}

func (is *itemService) GetAll(ctx context.Context, filter *item_proto.FilterItem) (*item_proto.GetAllItemRes, error) {
	items, err := is.posgresDB.Item.GetItems()
    if err!= nil {
        return nil, err
    }
    return &item_proto.GetAllItemRes{Items: items}, nil
}


func (is *itemService) Update(ctx context.Context, item *item_proto.Item) (*item_proto.I_Void, error) {
	err := is.posgresDB.Item.UpdateItem(item)
    if err!= nil {
        return nil, err
    }
    return nil, err
}


func (is *itemService) Delete(ctx context.Context, id *item_proto.ByIdItem) (*item_proto.I_Void, error) {
	err := is.posgresDB.Item.DeleteItem(id.Id)
    if err!= nil {
        return nil, err
    }
    return nil, err
}
