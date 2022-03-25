package p2pbind

import (
	"context"
	dht "github.com/libp2p/go-libp2p-kad-dht"

  host "github.com/libp2p/go-libp2p-host"
  libp2p "github.com/libp2p/go-libp2p"
)

type MobileLibp2p struct {
  node *host.Host
  IpfsDHT *dht.IpfsDHT
}

func StartLibp2p() *MobileLibp2p {
	ctx := context.Background()
  	host, _ := libp2p.New()
	kad, _ := dht.New(ctx, host)
//routing := discovery.NewRoutingDiscovery(IpfsDHT)

  return &MobileLibp2p{
    node: &host,
    IpfsDHT: kad}
}
