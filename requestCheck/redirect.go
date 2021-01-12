package requestCheck

import (
	"net/http"
	"net/http/httputil"
	"reflect"
)

func AllRedirectHeader(target_url string) ([]string, error) {
	req, _ := http.NewRequest("HEAD", target_url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var a []string
	r := resp
	for i := 0; r != nil; i++ {
		rv := reflect.ValueOf(r).Elem()
		vv := rv.FieldByName("Request")

		rp, ok := vv.Interface().(*http.Request)
		if !ok {
			break
		}

		dump, err := httputil.DumpResponse(r, false)
		if err != nil {
			break
		}
		
		nameAndHeader := rp.URL.String() + "\n" + string(dump)
		a = append(a, nameAndHeader)
		r = rp.Response
	}
	return a, err
}
