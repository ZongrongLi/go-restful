package view

import (
	"context"
	"errors"
	"strconv"

	"github.com/lexkong/log"
	"github.com/tiancai110a/go-restful/pkg/errno"
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

	username, ok := ctx.Value("username").(string)
	if !ok {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		return
	}
	password, ok := ctx.Value("passwd").(string)
	if !ok {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		return
	}
	err := rpc.Create(ctx, username, password)
	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, errstring := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.ErrString = errstring
	}

}

func Delete(ctx context.Context, resp *service.Resp) {

	uid, err := getInt64(ctx, "userid")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		log.Info("param error")
		return
	}
	err = rpc.Delete(ctx, uid)
	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, errstring := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.ErrString = errstring
	}

}

func Update(ctx context.Context, resp *service.Resp) {

	username, ok := ctx.Value("username").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		return
	}
	password, ok := ctx.Value("passwd").(string)
	if !ok {
		log.Error("wrong para", nil)
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		return
	}
	uid, err := getInt64(ctx, "userid")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		log.Info("param error")
		return
	}

	err = rpc.Update(ctx, username, uid, password)
	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, errstring := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.ErrString = errstring
	}
}

func Get(ctx context.Context, resp *service.Resp) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		return
	}
	username, passwd, userid, err := rpc.Get(ctx, username)
	if err != nil {
		log.Infof("rpc wrong: %s", err)
		code, errstring := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.ErrString = errstring
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
		resp.ErrString = errno.ErrParam.Message
		return
	}

	offset, err := getInt64(ctx, "offset")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		log.Info("param error")
		return
	}

	limit, err := getInt64(ctx, "limit")
	if err != nil {
		resp.Statuscode = errno.ErrParam.Code
		resp.ErrString = errno.ErrParam.Message
		log.Info("param error")
		return
	}
	data, err := rpc.List(ctx, username, offset, limit)

	if err != nil {
		log.Infof("rpc wrong: %s", err)

		code, errstring := errno.DecodeErr(err)
		resp.Statuscode = code
		resp.ErrString = errstring
	}
	resp.Add("userinfo", (string)(data))

}
