package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"go-kartos-study/app/service/member/conf"
	"go-kartos-study/app/service/member/internal/server/grpc"
	"go-kartos-study/app/service/member/internal/service"
	"go-kartos-study/pkg/naming/etcd"
	"go-kartos-study/pkg/net/trace"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-kartos-study/pkg/log"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	spew.Dump(conf.Conf)

	log.Init(conf.Conf.Log)
	defer log.Close()
	trace.Init(conf.Conf.Tracer)
	defer trace.Close()

	svc := service.New(conf.Conf)
	grpc.New(conf.Conf.GRPCServer, svc)

	cancel := etcd.ETCDRegist(conf.Conf.ETCDConfig, conf.AppID, "9002", conf.Color)
	defer cancel()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("member-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Close()
			log.Info("member-service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}