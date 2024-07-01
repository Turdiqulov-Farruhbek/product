package service

import (
	"context"
	"fmt"
	us_proto "product/genproto"
	posDB "product/storagecon/postgresdb"
)

type userService struct {
	posgresDB posDB.Storage
	us_proto.UnimplementedUserServiceServer
}

func NewUserService(PDB *posDB.Storage) *userService {
	return &userService{posgresDB: *PDB}
}

func (us *userService) Create(ctx context.Context, user *us_proto.UserCreateReq) (*us_proto.U_Void, error) {
	err := us.posgresDB.User.CreateUser(user)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("user created successfully")
	}
	return &us_proto.U_Void{}, nil
}

func (us *userService) Get(ctx context.Context, id *us_proto.ByIdUser) (*us_proto.User, error) {
	user, err := us.posgresDB.User.GetByUser(id.String())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetAll(ctx context.Context, filter *us_proto.FiltePUser) (*us_proto.GetAllUserResp, error) {
	users, err := us.posgresDB.User.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) Update(ctx context.Context, user *us_proto.User) (*us_proto.U_Void, error) {
	err := us.posgresDB.User.UpdateUser(user)
    if err!= nil {
        return nil, err
    }
    return &us_proto.U_Void{}, nil
}

func (us *userService) Delete(ctx context.Context, id *us_proto.ByIdUser) (*us_proto.U_Void, error) {
	err := us.posgresDB.User.DeleteUser(id.String())
    if err!= nil {
        return nil, err
    }
    return &us_proto.U_Void{}, nil
}
