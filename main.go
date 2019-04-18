package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/docker/libkv/store"
	"github.com/golang/glog"
	"github.com/tiancai110a/go-restful/router"
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
		go s.Serve("tcp", "127.0.0.1:8888", nil)
	}()
}

func main() {
	r1 := libkv.NewKVRegistry(store.ZK, "my-app", "/root/lizongrong/service",
		[]string{"127.0.0.1:1181", "127.0.0.1:2181", "127.0.0.1:3181"}, 1e10, nil)
	servertOption := server.Option{
		ProtocolType:   protocol.Default,
		SerializeType:  protocol.SerializeTypeMsgpack,
		CompressType:   protocol.CompressTypeNone,
		TransportType:  transport.TCPTransport,
		ShutDownWait:   time.Second * 12,
		Registry:       r1,
		RegisterOption: registry.RegisterOption{"my-app"},
		Tags:           map[string]string{"idc": "lf"}, //只允许机房为lf的请求，客户端取到信息会自己进行转移
		HttpServePort:  5080,
	}

	StartServer(&servertOption)
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		glog.Info("The router has been deployed successfully.")
	}()

	time.Sleep(time.Second * 1000)

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 200; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:5080" + "/view/health")
		if err != nil {
			glog.Info("===================================get error")
			return err
		}

		if resp.StatusCode != 200 {
			glog.Info("===================================http error statuscode:", resp.StatusCode)
			return nil
		}

		// Sleep for a second to continue the next ping.
		glog.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
