package common

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func BindJson(r *http.Request, obj interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, obj); err != nil {
		return err
	}
	v := validator.New()
	if err := v.Struct(obj); err != nil {
		return err
	}
	return nil
}
