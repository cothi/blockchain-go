package restt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	blockchaint "tetgo/tetgocoin/blockchain"
	"tetgo/tetgocoin/utill"

	"github.com/gorilla/mux"
)

var port string

type url string

func (u url) MarsharText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
}

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type addTxPayload struct {
	To     string
	Amount int
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	// data := []urlDescription

	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See A Block",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts for an Address",
		},
	}
	utill.HandleErr(json.NewEncoder(rw).Encode(data))
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utill.HandleErr(json.NewEncoder(rw).Encode(blockchaint.Blocks(blockchaint.Blockchain())))
	case "POST":
		blockchaint.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchaint.FindBlock(hash)
	encoder := json.NewEncoder(rw)

	if err == blockchaint.ErrNotFound {
		utill.HandleErr(encoder.Encode(errorResponse{fmt.Sprint(err)}))
	} else {
		utill.HandleErr(encoder.Encode(block))
	}
}

func status(rw http.ResponseWriter, r *http.Request) {
	utill.HandleErr(json.NewEncoder(rw).Encode(blockchaint.Blockchain()))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchaint.BalanceByAddress(address, blockchaint.Blockchain())
		json.NewEncoder(rw).Encode(balanceResponse{address, amount})

	default:
		utill.HandleErr(json.NewEncoder(rw).Encode(blockchaint.UTxOutsByAddress(address, blockchaint.Blockchain())))
	}
}

func mempool(rw http.ResponseWriter, r *http.Request) {
	utill.HandleErr(json.NewEncoder(rw).Encode(blockchaint.Mempool.Txs))
}

func transactions(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utill.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchaint.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		json.NewEncoder(rw).Encode(errorResponse{"not enough funds"})
	}
	rw.WriteHeader(http.StatusCreated)
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)

	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status)

	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance)

	router.HandleFunc("/mempool", mempool)
	router.HandleFunc("/transactions", transactions).Methods("POST")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
