package rpc

import (
	"context"

	"github.com/tiancai110a/go-restful/client"
)

func Create(ctx context.Context, name string, passwd string) error {

	client.Userclient = client.NewUserClient("user")
	req := client.CreateRequest{
		Username: name,
		Password: passwd,
	}
	res := client.CreateResponse{}
	err := client.Userclient.Call(ctx, "User.Create", &req, &res)
	return err
}

func Delete(ctx context.Context, userID int64) error {
	client.Userclient = client.NewUserClient("user")
	req := client.DeleteRequest{
		UserID: userID,
	}
	res := client.DeleteResponse{}
	err := client.Userclient.Call(ctx, "User.Delete", &req, &res)
	return err

}

func Get(ctx context.Context, username string) error {

	client.Userclient = client.NewUserClient("user")
	req := client.GetUserRequest{
		Username: username,
	}
	res := client.GetUserResponse{}
	err := client.Userclient.Call(ctx, "User.Get", &req, &res)
	return err
}

func List(ctx context.Context, username string, offset int64, limit int64) error {
	client.Userclient = client.NewUserClient("user")
	req := client.ListRequest{
		Username: username,
		Offset:   offset,
		Limit:    limit,
	}
	res := client.ListResponse{}
	err := client.Userclient.Call(ctx, "User.List", &req, &res)
	return err
}

func Update(ctx context.Context, username string, id int64, passwd string) error {
	client.Userclient = client.NewUserClient("user")
	req := client.UpdateRequest{
		client.UserInfo{
			Id:       id,
			Username: username,
			Password: passwd,
		},
	}
	res := client.ListResponse{}
	err := client.Userclient.Call(ctx, "User.Update", &req, &res)
	return err
}
