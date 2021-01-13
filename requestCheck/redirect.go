package requestCheck

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type dumpTransport struct {
	transport http.RoundTripper
}

func (dt *dumpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := dt.transport.RoundTrip(req)
	var dump []byte
	dump, e := httputil.DumpResponse(resp, false)
	if e != nil {
		return resp, err
	}

	fmt.Printf("** %s **\n", req.URL.String())
	fmt.Printf("%s", dump)

	return resp, err
}

func (dt *dumpTransport) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	if cr, ok := dt.transport.(canceler); ok {
		cr.CancelRequest(req)
	}
}

func AllRedirectHeader(target_url string) ([]string, error) {
	cli := http.DefaultClient
	cli.Transport = &dumpTransport{http.DefaultTransport}

	resp, err := cli.Head(target_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil
}
