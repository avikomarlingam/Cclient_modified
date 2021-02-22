package cclient

import (
	"golang.org/x/net/proxy"
	"net/http"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, proxyUrl ...string) (http.Client, error) {
	if len(proxyUrl) > 0 {
		dialer, err := newConnectDialer(proxyUrl[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, dialer),
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
      				return http.ErrUseLastResponse
  			}
		}, nil
	} else {
		return http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct),
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
      				return http.ErrUseLastResponse
  			}
		}, nil
	}
}
