package handler

import "log"

func NewInvestmentDocumentsUploaded() func(string) {
	return func(payload string) {
		log.Printf("Received payload on 'vendor_documents_uploaded': %s", payload)
		// Add logic to process the payload
	}
}

func NewOrderCreated() func(string) {
	return func(payload string) {
		log.Printf("Received payload on 'order_created': %s", payload)
		// Add logic to process the payload
	}
}
