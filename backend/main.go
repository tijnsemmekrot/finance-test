package main

import (
	"finance-test/middleware"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/expenses", middleware.EnableCORS(expenses.insertExpense))
	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", nil)
}
