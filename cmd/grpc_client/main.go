package main

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/g1ltz0r/auth/cmd/helpers"

	desc "github.com/g1ltz0r/auth/pkg/user_v1"
)

const address = "127.0.0.1:55555"

func createUser(ctx context.Context, c desc.UserV1Client) {
	r, err := c.Create(ctx, &desc.CreateRequest{
		Name:            gofakeit.Name(),
		Email:           gofakeit.Email(),
		Password:        gofakeit.Password(true, true, true, true, true, 5),
		PasswordConfirm: gofakeit.Password(true, true, true, true, true, 5),
		Role:            desc.Role(0),
	})
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	log.Printf(color.RedString("Create User:\n"), color.GreenString("%+v", r))
}

func getUser(ctx context.Context, c desc.UserV1Client) {
	r, err := c.Get(ctx, &desc.GetRequest{Id: helpers.GetRandID()})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	log.Printf(color.RedString("Get User:\n"), color.GreenString("%+v", r))
}

func updateUser(ctx context.Context, c desc.UserV1Client) {
	r, err := c.Update(ctx, &desc.UpdateRequest{
		Id:    helpers.GetRandID(),
		Name:  wrapperspb.String(gofakeit.Name()),
		Email: wrapperspb.String(gofakeit.Email()),
		Role:  desc.Role(0),
	})
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	log.Printf(color.RedString("Update User:\n"), color.GreenString("%+v", r))
}

func deleteUser(ctx context.Context, c desc.UserV1Client) {
	r, err := c.Delete(ctx, &desc.DeleteRequest{Id: helpers.GetRandID()})
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

	log.Printf(color.RedString("Delete User:\n"), color.GreenString("%+v", r))
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close the connection: %v", err.Error())
		}
	}(conn)

	c := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getUser(ctx, c)
	createUser(ctx, c)
	updateUser(ctx, c)
	deleteUser(ctx, c)
}
