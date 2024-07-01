package service


import (
	"context"
	"fmt"
	prud_proto "product/genproto"
	posDB "product/storagecon/postgresdb"
)

type productService struct {
	posgresDB posDB.Storage
	prud_proto.UnimplementedProductServiceServer
}

func NewProductService(PDB *posDB.Storage) *productService {
	return &productService{posgresDB: *PDB}
}


func (ps *productService) Create(ctx context.Context, product *prud_proto.ProductCreateReq) (*prud_proto.I_Void, error) {
	err := ps.posgresDB.Product.CreateProduct(product)
    if err != nil {
        return nil, err
    } else {
        fmt.Println("product created successfully")
    }
    return &prud_proto.I_Void{}, nil
}

func (ps *productService) Get(ctx context.Context, id *prud_proto.ByIdProduct) (*prud_proto.Product, error) {
	product, err := ps.posgresDB.Product.GetProduct(id.Id)
    if err!= nil {
        return nil, err
    }
    return product, nil
}



func (ps *productService) GetAll(ctx context.Context, filter *prud_proto.FilterProduct) (*prud_proto.GetAllProductResp, error) {
	products, err := ps.posgresDB.Product.GetProducts()
    if err!= nil {
        return nil, err
    }
    return &prud_proto.GetAllProductResp{Products: products}, nil
}


func (ps *productService) Update(ctx context.Context, product *prud_proto.Product) (*prud_proto.I_Void, error) {
	err := ps.posgresDB.Product.UpdateProduct(product)
    if err!= nil {
        return nil, err
    }
    return nil, err
}


func (ps *productService) Delete(ctx context.Context, id *prud_proto.ByIdProduct) (*prud_proto.I_Void, error) {
	err := ps.posgresDB.Product.DeleteProduct(id.Id)
    if err!= nil {
        return nil, err
    }
    return nil, err
}