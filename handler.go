package contract

import (
	"encoding/json"
	"net/http"
)

func MakeRequestHttpHandlerFunc(f func(Request) (*Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := Request{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeError(w, r, err)
			return
		}

		request.QueryParams = r.URL.Query()

		response, err := f(request)
		if err != nil {
			writeError(w, r, err)
			return
		}

		writeResponse(w, r, response)
	}
}

func MakeJobHttpHandlerFunc(f func(Job) (*Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		job := Job{}

		err := json.NewDecoder(r.Body).Decode(&job)
		if err != nil {
			writeError(w, r, err)
			return
		}

		response, err := f(job)
		if err != nil {
			writeError(w, r, err)
			return
		}

		writeResponse(w, r, response)
	}
}

func writeResponse(w http.ResponseWriter, r *http.Request, response *Response) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		writeError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func writeError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	if err == nil {
		return
	}

	responseError := Error{
		Message: err.Error(),
	}

	jsonResponse, err := json.Marshal(responseError)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
