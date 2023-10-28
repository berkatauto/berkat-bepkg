package berkatbepkg

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

// func GCFHandler(MONGODATA, dbname, collectionname string) string {
// 	mconn := SetConnection(MONGODATA, dbname)
// 	datagedung := GetAllBangunanLineString(mconn, collectionname)
// 	return GCFReturnStruct(datagedung)
// }

func GCFPostHandler(PASETOPRIV, MONGODATA, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGODATA, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIV))
			if err != nil {
				Response.Message = "Token Encode Fail : " + err.Error()
			} else {
				Response.Message = "Welcome!"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Incorrect"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}
