package safeutil

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Read(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("resp blank")
	}
	if resp.Body == nil {
		return nil, errors.New("body blank")
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Unmarshal(resp *http.Response, target any) error {
	bs, err := Read(resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(bs, target)
}
