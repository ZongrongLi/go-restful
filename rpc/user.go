package rpc

import (
	"context"
	"encoding/json"

	"github.com/golang/glog"

	"github.com/tiancai110a/go-restful/client"
	"github.com/tiancai110a/go-restful/dl"
)

func Create(ctx context.Context, name string, passwd string) error {

	client.Userclient = client.NewUserClient("user")
	req := dl.CreateRequest{
		Username: name,
		Password: passwd,
	}
	res := dl.CreateResponse{}
	err := client.Userclient.Call(ctx, "User.Create", &req, &res)
	if !dl.CheckStatus(&res.Base) {
		return res.Base.Err
	}

	return err
}

func Delete(ctx context.Context, userID int64) error {
	client.Userclient = client.NewUserClient("user")
	req := dl.DeleteRequest{
		UserID: userID,
	}
	res := dl.DeleteResponse{}
	err := client.Userclient.Call(ctx, "User.Delete", &req, &res)
	if !dl.CheckStatus(&res.Base) {
		return res.Base.Err
	}

	return err

}

func Get(ctx context.Context, username string) (string, string, int64, error) {

	client.Userclient = client.NewUserClient("user")
	req := dl.GetUserRequest{
		Username: username,
	}
	res := dl.GetUserResponse{}
	err := client.Userclient.Call(ctx, "User.Get", &req, &res)
	if err != nil {

	}
	glog.Info("++++++++++++++++++++++++++++++++++++++++++++++", res)
	if !dl.CheckStatus(&res.Base) {
		return "", "", 0, res.Base.Err
	}

	return res.Username, res.Password, res.Id, nil
}

func List(ctx context.Context, username string, offset int64, limit int64) ([]byte, error) {
	client.Userclient = client.NewUserClient("user")
	req := dl.ListRequest{
		Username: username,
		Offset:   offset,
		Limit:    limit,
	}
	res := dl.ListResponse{}
	err := client.Userclient.Call(ctx, "User.List", &req, &res)

	if !dl.CheckStatus(&res.Base) {
		return nil, res.Base.Err
	}

	data, err := json.Marshal(res)
	if err != nil {
		glog.Error("json marshal error. err: ", err)
	}
	return data, err
}

func Update(ctx context.Context, username string, id int64, passwd string) error {
	client.Userclient = client.NewUserClient("user")
	req := dl.UpdateRequest{
		dl.UserInfo{
			Id:       id,
			Username: username,
			Password: passwd,
		},
	}
	res := dl.ListResponse{}
	err := client.Userclient.Call(ctx, "User.Update", &req, &res)

	if !dl.CheckStatus(&res.Base) {
		return res.Base.Err
	}

	return err
}
