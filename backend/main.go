package main

import (
	"finance-test/expenses"
	"finance-test/middleware"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/expenses", middleware.EnableCORS(expenses.InsertExpense))
	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", nil)
}
