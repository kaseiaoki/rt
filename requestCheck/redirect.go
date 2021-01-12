package requestCheck

import (
	"net/http"
	"reflect"
)

func AllRedirectHeader(target_url string) []string {
	req, _ := http.NewRequest("HEAD", target_url, nil)
	resp, _ := http.DefaultClient.Do(req)
    var a []string
	r := resp
    for i := 0; r != nil; i++ {
        rv := reflect.ValueOf(r).Elem()
        vv := rv.FieldByName("Request")
        rp, ok := vv.Interface().(*http.Request)
        if !ok {
            break
        }
        a = append(a, rp.URL.String())
        r = rp.Response
    }
	return a
}
