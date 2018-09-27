package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	// Capital always
	Record string
}

var mutex = &sync.Mutex{}
var blockHandler = NewBlockchain()

// web server run
func run() error {
	mux := makeMuxRouter()
	httpPort := "3000"
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m map[string]interface{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&m); err != nil {
		fmt.Println(err)
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	//fmt.Println(m)
	//fmt.Println(m.Record)
	postTx(m)
	//blockHandler.AddBlock(m.Record)
	mutex.Unlock()
	//newBlock := blockHandler.blocks[len(blockHandler.blocks)-1]
	respondWithJSON(w, r, http.StatusCreated, "asa")
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, r, 200, getTx())
	//bytes, err := json.MarshalIndent(blockHandler.blocks, "", "  ")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//io.WriteString(w, string(bytes))
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
