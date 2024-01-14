package module

import (
	"encoding/json"
	"net/http"
	"os"
	// "strconv"
	"strings"

	model ""
	"github.com/whatsauth/watoken"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Responsed           model.Credential
	reservasiResponse   model.ReservasiResponse
	reservasi           model.Reservasi
)

func GCFHandlerInsertReservasi(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	reservasiResponse.Status = false
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		reservasiResponse.Message = "error parsing application/json1:"
		return GCFReturnStruct(reservasiResponse)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(reservasiResponse)
	}

	err = json.NewDecoder(r.Body).Decode(&reservasi)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(reservasiResponse)
	}

	_, err = InsertReservasi(mconn, collectionname, reservasi.Nama, reservasi.No_telp, reservasi.TTL, reservasi.Status, reservasi.Keluhan)
	if err != nil {
		reservasiResponse.Message = "error inserting Pemasukan: " + err.Error()
		return GCFReturnStruct(reservasiResponse)
	}

	reservasiResponse.Status = true
	reservasiResponse.Message = "Insert Reservasi success"
	return GCFReturnStruct(reservasiResponse)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}