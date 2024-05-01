package test_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func ValuesToSql(values []string) string {
	var vals string
	for _, val := range values {
		vals += fmt.Sprintf("'%s',", val)
	}
	return vals[:len(vals)-1]
}

func GetParamsFromResponseBody[T interface{}](structBody T, r *http.Response) T {
	decoder := json.NewDecoder(r.Body)

	params := structBody
	err := decoder.Decode(&params)
	if err != nil {
		log.Fatal(errors.New("couldn't decode parameters: "), err)
	}

	return params
}

func GetBodyString(t *testing.T, resp *http.Response) string {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return string(b)
}
