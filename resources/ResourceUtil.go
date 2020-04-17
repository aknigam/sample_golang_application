package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"sample_golang_application/errors"
	"strconv"
)

func WriteOKResponse(entity interface{}, writer http.ResponseWriter) (error *errors.AppError) {
	output, err := json.MarshalIndent(&entity, "", "\t\t")
	if err != nil {
		return &errors.AppError{Error: err, Message: "Error writing response", Code: -1}
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(output)

	return error
}

func ExtractIdentifierAndEntity(r *http.Request, entity interface{}, w http.ResponseWriter, pathParamName string) (int, *errors.AppError, bool) {
	pathParamVal, err, hasError := ExtractPathParam(r, pathParamName)
	if hasError {
		return 0, err, true
	}
	appError, hasError := ParseRequest(r, entity, w)
	if hasError {
		return 0, appError, true
	}
	return pathParamVal, nil, false
}

func ParseRequest(r *http.Request, entity interface{}, w http.ResponseWriter) (*errors.AppError, bool) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&entity)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return &errors.AppError{Error: err, Message: "Order Id not provided", Code: -1}, true
	}
	return nil, false
}

func ExtractPathParam(r *http.Request, pathParamName string) (int, *errors.AppError, bool) {
	vars := mux.Vars(r)
	pathParamVal, err := strconv.Atoi(path.Base(vars[pathParamName]))
	if err != nil {
		return -1, &errors.AppError{Error: err, Message: "Order Id not provided", Code: -1}, true
	}
	return pathParamVal, nil, false
}
