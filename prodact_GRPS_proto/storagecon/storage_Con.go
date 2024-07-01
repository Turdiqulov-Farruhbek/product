package storagecon

import (
	us_proto "product/genproto"
)

type StorageI interface {
	UserStor() UserI
}

type UserI interface {
	Create(user *us_proto.UserCreateReq) error
}
