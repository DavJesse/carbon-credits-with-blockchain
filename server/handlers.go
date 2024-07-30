package server

import (
	"log"
	"net/http"
)

type CarbonCredit struct {
	ID                 int    `json:"id"`
	Issuer             string `json:"issuer"`
	Amount             int    `json:"amount"`
	VerificationStatus string `json:"verification_status"`
}

func issueCarbonCredit(w http.ResponseWriter, r *http.Request) {
	// Issuing logic
}

func main() {
	http.HandleFunc("/issue", issueCarbonCredit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
