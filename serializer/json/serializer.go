package json

import (
	"encoding/json"

	"github.com/BRO3886/url-shortener/shortener"
)

type Redirect struct {
}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := new(shortener.Redirect)
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, err
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	return json.Marshal(input)
}
