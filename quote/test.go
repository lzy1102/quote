package main

import (
	"context"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/lzy1102/quote"
	"golang.org/x/net/proxy"
	"log"
	"net"
	"net/http"
)

func main() {
	socks5, err := proxy.SOCKS5("tcp", "127.0.0.1:7890", nil, proxy.Direct)
	if err != nil {
		return
	}
	client := req.DefaultClient()
	client.SetDial(func(ctx context.Context, network, addr string) (net.Conn, error) {
		return socks5.Dial(network, addr)
	})

	httpclient := &http.Client{}
	httpclient.Transport = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return socks5.Dial(network, addr)
		},
	}
	xiuqiu, err := quote.NewQuoteFromXueqiu("600033.ss", "2022-02-01", "2023-01-31", quote.Daily, client)
	if err != nil {
		return
	}
	fmt.Println(xiuqiu.CSV())
	table, err := quote.GetNewTable()
	if err != nil {
		return
	}
	log.Println(table)
}
