package controllers

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

type Insert struct {
	Key  string `json:"key"`
	Value string `json:"value"`
}

func InsertData(w http.ResponseWriter, r *http.Request) {
	var content Insert
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := leveldb.OpenFile("level.db", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()

	err = db.Put([]byte(content.Key), []byte(content.Value), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Data successfully inserted")
}