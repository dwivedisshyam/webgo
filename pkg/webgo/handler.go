package webgo

import (
	"encoding/json"
	"net/http"

	"github.com/dwivedisshyam/webgo/pkg/errors"
	"github.com/dwivedisshyam/webgo/pkg/webgo/types"
)

type Handler func(*Context) (interface{}, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Context().Value(webGoContextKey).(*Context)

	w.Header().Set("Content-type", "application/json")

	data, err := h(c)
	if err != nil {
		e := getError(err)
		w.WriteHeader(e.StatusCode)
		_ = json.NewEncoder(w).Encode(e)

		return
	}

	d := getResponse(data)

	w.WriteHeader(d.StatusCode)
	_ = json.NewEncoder(w).Encode(d)
}

func getError(err error) *types.Error {
	e := &types.Error{}

	switch er := err.(type) {
	case *errors.EntityNotFound:
		e.StatusCode = http.StatusNotFound

	case *errors.EntityAlreadyExists:
		e.StatusCode = http.StatusConflict

	case *types.Error:
		e.StatusCode = er.StatusCode
		e.Reason = er.Reason

	case *errors.InvalidParam, *errors.MissingParam:
		e.StatusCode = http.StatusBadRequest
	}

	e.Reason = err.Error()

	if e.StatusCode == 0 {
		e.StatusCode = http.StatusInternalServerError
	}

	return e
}

func getResponse(data interface{}) *types.Response {
	res := &types.Response{}

	switch r := data.(type) {
	case *types.Response:
		res = r
	case types.Response:
		res = &r
	default:
		res.Data = r
	}

	if res.StatusCode == 0 {
		res.StatusCode = http.StatusOK
	}

	return res
}
