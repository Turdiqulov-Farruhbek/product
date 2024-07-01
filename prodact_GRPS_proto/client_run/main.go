package main

import (
	"context"
	"log"
	proto "product/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost: 8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	user := proto.NewUserServiceClient(conn)

	newUser := &proto.UserCreateReq{
		Name: "Azizbek",
		Balance: 12000,
	    Email: "aziz@example.com",
        Phone: "+998997411424",
    }

	resUser, err := user.Create(context.Background(), newUser)		
	if err!= nil {
        log.Fatal(err)
    }

	log.Println(resUser)

// _============================================================


    // basket := proto.NewBasketServiceClient(conn)
	// newbasket := proto.CreatedBasket{
	// 	UserId: "cf29bf24-f724-49f4-87cc-f8839a7fc98c",
	// }

	// resBasket, err := basket.CreateBasket(context.Background(), &newbasket)
	// if err!= nil{
    //     log.Fatal(err)
    // }
	// log.Println(resBasket)

// _============================================================


	// category := proto.NewcategoryServiceClient(conn)

	// newbasket := proto.CreatedBasket{
	// 	UserId: "cf29bf24-f724-49f4-87cc-f8839a7fc98c",
	// }

	// resBasket, err := basket.CreateBasket(context.Background(), &newbasket)
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// log.Println(resBasket)


// _============================================================


	// item := proto.NewItemServiceClient(conn)
	// newitem := proto.CreateItem{
	// 	BasketID: "cf29bf24-f724-49f4-87cc-f8839a7fc98c",
	// 	ProductID: "cf29bf24-f724-49f4-87cc-f8839a7fc98c",
    //     Quantity: 10,
		
    // }
	// resItem, err := item.Create(context.Background(), &newitem)
	// if err!= nil{
    //     log.Fatal(err)
    // }
	// log.Println(resItem)

// _============================================================

	// product := proto.NewProductServiceClient(conn)
	// newproduct := proto.ProductCreateReq{
    //     Name: "mevalar",
	// 	Price: 15000,
    //     CategoryId: "cf29bf24-f724-49f4-87cc-f8839a7fc98c",
	// 	ExpiredAt: "2021-05-15",
    // }

	// resProduct, err := product.Create(context.Background(), &newproduct)
	// if err!= nil{
    //     log.Fatal(err)
    // }

	// log.Println(resProduct)


// _============================================================

	// trans := proto.NewTransactionServiceClient(conn)
	// newtrans := proto.TransCreateReq{
	// 		BasketID: "cf29bf24-f724-49f4-87cc",
	// 		TotalPrice: 15000,
    // }

	// resTrans, err := trans.Create(context.Background(), &newtrans)
	// if err!= nil{
    //     log.Fatal(err)
    // }
	// log.Println(resTrans)



}
