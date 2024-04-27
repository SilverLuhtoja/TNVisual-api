package test_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
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
