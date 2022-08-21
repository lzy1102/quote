# go-quote

github.com/markcheno/go-quote  add customize http.client

## Install library

Install the package with:

```bash
go get github.com/lzy1102/quote
```

## Library example

```go
package main

import (
	"crypto/tls"
	"fmt"
	"github.com/lzy1102/go-quote"
	"github.com/markcheno/go-talib"
	"net/http"
	"net/url"
)

func main() {
	proxyurl, _ := url.Parse("http://127.0.0.1:10809")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyURL(proxyurl),
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
	}
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2021-04-01", quote.Daily, true, client)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
```

## License

MIT License  - see LICENSE for more details
