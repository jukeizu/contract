package contract

import (
	"encoding/json"
	"net/http"
)

func MakeHttpHandlerFunc(f func(Request) (*Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := Request{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeError(w, r, err)
			return
		}

		response, err := f(request)
		if err != nil {
			writeError(w, r, err)
			return
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			writeError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
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
