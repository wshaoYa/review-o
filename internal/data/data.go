package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "review-o/api/review/v1"
	"review-o/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewReviewClient, NewDiscovery, NewOperationRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rc  v1.ReviewClient
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rc: rc, log: log.NewHelper(logger)}, cleanup, nil
}

func NewReviewClient(d registry.Discovery) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///reviewService"),
		grpc.WithDiscovery(d),
	)
	if err != nil {
		log.Errorf("data NewReviewClient failed,err:%v\n", err)
		panic(err)
	}
	return v1.NewReviewClient(conn)
}

func NewDiscovery(c *conf.Consul) registry.Discovery {
	consulCfg := api.DefaultConfig()
	consulCfg.Address = c.GetAddress()
	consulCfg.Scheme = c.GetScheme()

	cli, err := api.NewClient(consulCfg)
	if err != nil {
		log.Errorf("data NewDiscovery failed,err:%v\n", err)
		panic(err)
	}

	return consul.New(cli)
}
