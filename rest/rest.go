package rest

import "fmt"

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url
	Method      string
	Description string
	Payload     string
}

type balanceResponse struct {
	Address string
	Balance int
}

type myWalletResponse struct {
	Address string
	Balance int
}

type errorResponse struct {
	ErrorMessage string
}

type addTxPayload struct {
	To     string
	Amount int
}


