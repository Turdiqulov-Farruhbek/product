package service

import (
	"context"
	"fmt"
	ctg_proto "product/genproto"
	posDB "product/storagecon/postgresdb"
)

type categoryService struct {
	posgresDB posDB.Storage
	ctg_proto.UnimplementedCategoryServiceServer
}

func NewCategoryService(PDB *posDB.Storage) *categoryService {
	return &categoryService{posgresDB: *PDB}
}

func (bs *categoryService) Create(ctx context.Context, category *ctg_proto.Category) (*ctg_proto.C_Void, error) {
	nill, err := bs.posgresDB.Category.Createategory(category)
    if err!= nil {
        return nil, err
    }else{
        fmt.Println("category created successfully")
    }
    return nill, err
}

func (bs *categoryService) Get(ctx context.Context, id *ctg_proto.ByIdCategory) (*ctg_proto.Category, error) {
	category, err := bs.posgresDB.Category.GetCategory(id.Id)
    if err!= nil {
        return nil, err
    }
    return category, nil
}

func (bs *categoryService) GetAll(ctx context.Context, filter *ctg_proto.FilterCategory) (*ctg_proto.GetAllCategoryResp, error) {
	categories, err := bs.posgresDB.Category.GetCategories()
    if err!= nil {
        return nil, err
    }
    return &ctg_proto.GetAllCategoryResp{Categories: categories}, nil
}


func (bs *categoryService) Update(ctx context.Context, category *ctg_proto.Category) (*ctg_proto.C_Void, error) {
	err := bs.posgresDB.Category.UpdateCategory(category)
    if err!= nil {
        return nil, err
    }
    return nil, err
}

func (bs *categoryService) Delete(ctx context.Context, id *ctg_proto.ByIdCategory) (*ctg_proto.C_Void, error) {
	err := bs.posgresDB.Category.DeleteCategory(id.Id)
    if err!= nil {
        return nil, err
    }
    return nil, err
}