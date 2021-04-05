package handler

import (
	"crypto/sha256"
	"encoding/json"
	"encryptService/model"
	"fmt"
	"net/http"
)


//Handler
func StringHandler(w http.ResponseWriter, r *http.Request)  {
	var request model.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	strings := EncryptString(request.Strings)

	request.Strings = strings

	err = Response(w, request)

}

func Response(res http.ResponseWriter, response interface{}) error {
	res.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(res).Encode(response)
	return nil
}

func EncryptString(strings []string) []string {
	var decStrings []string
	for i := range strings{
		encString :=sha256.Sum256([]byte(strings[i]))
		//fmt.Printf("%x", string(encString[:]))
		decStrings = append(decStrings, fmt.Sprintf("%x", string(encString[:])))
	}
	return decStrings
}
