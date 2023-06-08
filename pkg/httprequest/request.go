package httprequest

import (
	"io"
	"net/http"
)

func Get(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}
