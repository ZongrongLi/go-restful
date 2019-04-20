package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lexkong/log"

	"github.com/docker/libkv/store"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"github.com/tiancai110a/go-restful/config"
	"github.com/tiancai110a/go-restful/router"
	"github.com/tiancai110a/go-restful/rpc"
	"github.com/tiancai110a/go-rpc/protocol"
	"github.com/tiancai110a/go-rpc/registry"
	"github.com/tiancai110a/go-rpc/registry/libkv"
	"github.com/tiancai110a/go-rpc/server"
	"github.com/tiancai110a/go-rpc/transport"
)

func StartServer(op *server.Option) {
	go func() {
		s, err := server.NewSGServer(op)
		if err != nil {
			glog.Error("new serializer failed", err)
			return
		}
		router.Load(s)
		go s.Serve("tcp", viper.GetString("tcpurl"), nil)
	}()
}
func testRPC() {
	ctx := context.Background()
	rpc.Create(ctx, "tiancai", "123")
	rpc.Delete(ctx, 1234567)
	rpc.Update(ctx, "tiancai", 0, "123")
	rpc.Get(ctx, "tiancai")
	rpc.List(ctx, "tiancai", 0, 10)
}
func main() {

	// init db
	// model.DB.Init()
	// defer model.DB.Close()

	if err := config.Init(""); err != nil {
		panic(err)
	}

	testRPC()
	var r1 registry.Registry
	if viper.GetString("discovery.name") == "zk" {
		nodes := viper.GetString("discovery.nodes")
		zknode := strings.Split(nodes, ",")
		log.Infof("======================================znode %+v", zknode)
		interval, err := strconv.ParseFloat(viper.GetString("discovery.updateinterval"), 64)
		if err != nil {
			log.Infof("parse interval err: %s", err)
			interval = 1e10
		}

		r1 = libkv.NewKVRegistry(store.ZK, viper.GetString("name"), viper.GetString("discovery.path"),
			zknode, time.Duration(interval), nil)

	} else {
		glog.Error("discovery is not set")
		return
	}

	port, err := strconv.ParseInt(viper.GetString("port"), 10, 64)
	if err != nil {
		log.Infof("parse port err: %s", err)
		return
	}

	servertOption := server.Option{
		ProtocolType:   protocol.Default,
		SerializeType:  protocol.SerializeTypeMsgpack,
		CompressType:   protocol.CompressTypeNone,
		TransportType:  transport.TCPTransport,
		ShutDownWait:   time.Second * 12,
		Registry:       r1,
		RegisterOption: registry.RegisterOption{viper.GetString("name")},
		Tags:           map[string]string{"idc": viper.GetString("idc")}, //只允许机房为lf的请求，客户端取到信息会自己进行转移
		HttpServePort:  int(port),
		HttpServeOpen:  true,
	}

	StartServer(&servertOption)
	// Ping the server to make sure the router is working.
	// go func() {
	// 	if err := pingServer(); err != nil {
	// 		log.Fatal("The router has no response, or it might took too long to start up.", err)
	// 	}
	// 	log.Info("The router has been deployed successfully.")
	// }()

	time.Sleep(time.Second * 1000)

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 200; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("httpurl") + "/view/health")
		if err != nil {
			log.Info("get error")
			return err
		}

		if resp.StatusCode != 200 {
			log.Infof("http error statuscode:%d", resp.StatusCode)
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof("Waiting for the router, retry in 1 second.  %+v", resp)
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
