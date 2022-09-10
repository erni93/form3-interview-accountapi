package errorhandler

import (
	"encoding/json"
	model "erni93/form3-interview-accountapi/models"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidHttpResponse = errors.New("client: invalid http response")
)

func GetErrorResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return nil
	}

	var errRes model.ErrorResponse
	err := json.NewDecoder(res.Body).Decode(&errRes)
	if err != nil {
		return fmt.Errorf("%w http %d", ErrInvalidHttpResponse, res.StatusCode)
	}

	return errors.New(errRes.ErrorMessage)
}
