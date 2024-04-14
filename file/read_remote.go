package file

import (
	"io"
	"net/http"
)

func ReadRemote(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if ce := Body.Close(); ce != nil {
			err = ce
		}
	}(res.Body)

	bytes, _ := io.ReadAll(res.Body)
	return bytes, err
}
