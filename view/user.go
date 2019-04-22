package view

import (
	"context"
	"errors"
	"strconv"

	"github.com/golang/glog"

	"github.com/lexkong/log"
	"github.com/tiancai110a/go-restful/pkg/auth"
	"github.com/tiancai110a/go-restful/pkg/errno"
	"github.com/tiancai110a/go-restful/pkg/token"
	"github.com/tiancai110a/go-restful/rpc"
	"github.com/tiancai110a/go-rpc/service"
)

func getInt64(ctx context.Context, key string) (int64, error) {
	v1, ok := ctx.Value(key).(string)
	if !ok {
		return 0, errors.New("param wrong")
	}
	return strconv.ParseInt(v1, 10, 64)
}

func Create(ctx context.Context, resp *service.Resp) {
	var err error
	username, ok := ctx.Value("username").(string)
	if !ok {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}
	password, ok := ctx.Value("passwd").(string)
	if !ok {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}

	if password, err = auth.Encrypt(password); err != nil {
		log.Error("Encrypt failed", err)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}
	err = rpc.Create(ctx, username, password)
	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
	}

}

func Delete(ctx context.Context, resp *service.Resp) {

	uid, err := getInt64(ctx, "userid")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		log.Info("param error")
		return
	}
	err = rpc.Delete(ctx, uid)
	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
	}

}

func Update(ctx context.Context, resp *service.Resp) {

	username, ok := ctx.Value("username").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}
	password, ok := ctx.Value("passwd").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}

	var err error
	if password, err = auth.Encrypt(password); err != nil {
		log.Error("Encrypt failed", err)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}

	uid, err := getInt64(ctx, "userid")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		log.Info("param error")
		return
	}

	err = rpc.Update(ctx, username, uid, password)
	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
	}
}

func Get(ctx context.Context, resp *service.Resp) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}
	username, passwd, userid, err := rpc.Get(ctx, username)
	if err != nil {
		log.Infof("rpc wrong: %s", err)
		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
	}

	resp.Add("username", username)
	resp.Add("userid", strconv.FormatInt(userid, 10))
	resp.Add("password", passwd)

}

func GetList(ctx context.Context, resp *service.Resp) {

	username, ok := ctx.Value("username").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}

	offset, err := getInt64(ctx, "offset")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		log.Info("param error")
		return
	}

	limit, err := getInt64(ctx, "limit")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		log.Info("param error")
		return
	}
	data, err := rpc.List(ctx, username, offset, limit)

	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
	}
	resp.Add("userinfo", (string)(data))

}

func Login(ctx context.Context, resp *service.Resp) {

	username, ok := ctx.Value("username").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}
	password, ok := ctx.Value("passwd").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.Message = errno.ErrParam.Message
		return
	}

	_, psswd, userid, err := rpc.Get(ctx, username)
	if err != nil {
		log.Error("user not found", err)
		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
		return
	}
	glog.Info("====1", psswd)
	// Compare the login password with the user password.
	if err := auth.Compare(psswd, password); err != nil {
		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
		return
	}
	glog.Info("====2")
	// Sign the json web token.
	token, err := token.Sign(ctx, token.Context{ID: userid, Username: username}, "")
	if err != nil {
		code, Message := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.Message = Message
		return
	}
	resp.Add("token", token)
	return

}
