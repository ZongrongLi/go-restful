package client

import (
	"strconv"
	"strings"
	"time"

	"github.com/lexkong/log"
	"github.com/spf13/viper"

	"github.com/docker/libkv/store"

	"github.com/tiancai110a/go-rpc/client"
	"github.com/tiancai110a/go-rpc/protocol"
	"github.com/tiancai110a/go-rpc/ratelimit"
	"github.com/tiancai110a/go-rpc/registry"
	"github.com/tiancai110a/go-rpc/registry/libkv"
	"github.com/tiancai110a/go-rpc/transport"
)

var Userclient client.SGClient

func NewUserClient(appkey string) client.SGClient {
	if Userclient != nil {
		return Userclient
	}

	op := &client.DefaultSGOption
	op.AppKey = appkey
	op.RequestTimeout = time.Millisecond * 100
	op.DialTimeout = time.Millisecond * 100
	op.HeartbeatInterval = time.Second * 20
	op.HeartbeatDegradeThreshold = 5
	op.Heartbeat = true
	op.SerializeType = protocol.SerializeTypeMsgpack
	op.CompressType = protocol.CompressTypeNone
	op.TransportType = transport.TCPTransport
	op.ProtocolType = protocol.Default
	op.FailMode = client.FailRetry
	op.Retries = 3
	op.Auth = "hello01"
	//一秒钟失败20次 就会进入贤者模式.. 因为lastupdate时间在不断更新，熔断后继续调用有可能恢复
	op.CircuitBreakerThreshold = 20
	op.CircuitBreakerWindow = time.Second

	//基于标签的路由策略
	op.Tagged = true
	op.Tags = map[string]string{"idc": viper.GetString("idc")}

	op.Wrappers = append(op.Wrappers, &client.RateLimitInterceptor{Limit: ratelimit.NewRateLimiter(1000, 2)}) //一秒10个，最多有两个排队
	var r registry.Registry

	if viper.GetString("discovery.name") == "zk" {
		nodes := viper.GetString("discovery.nodes")
		zknode := strings.Split(nodes, ",")
		log.Infof("client user zknode: %+v", zknode)
		interval, err := strconv.ParseFloat(viper.GetString("discovery.updateinterval"), 64)
		if err != nil {
			log.Infof("parse interval err: %s", err)
			interval = 1e10
		}

		r = libkv.NewKVRegistry(store.ZK, appkey, viper.GetString("discovery.path"),
			zknode, time.Duration(interval), nil)

	} else {
		log.Infof("discovery is not set %s", viper.GetString("discovery.name"))
		panic("discovery is not set")
	}

	op.Registry = r

	c := client.NewSGClient(*op)
	return c
}
