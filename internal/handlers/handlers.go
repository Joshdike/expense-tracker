package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Joshdike/expense-tracker/internal/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type handle struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *handle {
	return &handle{db}
}

func (h handle) GetAllExpense(w http.ResponseWriter, r *http.Request) {
	query, params, err := sq.Select("*").From("expenses").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		http.Error(w, "internal sql error", http.StatusInternalServerError)
		return
	}

	fmt.Println(query, params)

	rows, err := h.db.Query(r.Context(), query, params...)
	if err != nil {
		http.Error(w, "error retrieving data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	result := make([]models.Expense, 0)
	for rows.Next() {
		var e models.Expense

		if err = rows.Scan(&e.TransactionID, &e.Date, &e.Amount, &e.Category,
			&e.Description, &e.PaymentMethod); err != nil {
			http.Error(w, "error retrieving data", http.StatusInternalServerError)
			return
		}

		result = append(result, e)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h handle) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var e models.Expense
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "error decoding payload", http.StatusBadRequest)
		return
	}

	query, params, err := sq.Insert("expenses").Columns("transactionid", "date",
		"amount", "category", "description", "payment_method").
		Values(e.TransactionID, e.Date, e.Amount, e.Category, e.Description,
			e.PaymentMethod).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal sql error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("success: query: %s, params: %v\n", query, params)

	if _, err := h.db.Exec(r.Context(), query, params...); err != nil {
		http.Error(w, "error inserting expense data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "successful"}`))

}
func (h handle) GetExpenseById(w http.ResponseWriter, r *http.Request) {

	var e models.Expense
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("conversion error %v", err)
	}
	query, params, err := sq.Select("transactionid", "date", "amount", "category",
		"description", "payment_method").
		From("expenses").Where("transactionid = ?", id).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		http.Error(w, "internal sql error", http.StatusInternalServerError)
		return
	}

	fmt.Println(query, params)

	err = h.db.QueryRow(r.Context(), query, params...).Scan(&e.TransactionID, &e.Date, &e.Amount,
		&e.Category, &e.Description, &e.PaymentMethod)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "expense not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error retrieving data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(e)

}
func (h handle) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var e models.Expense
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("conversion error %v", err)
	}

	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "error decoding payload", http.StatusBadRequest)
		return
	}

	updateS := sq.Update("expenses").
		Where("transactionid = ?", id).PlaceholderFormat(sq.Dollar)

	if !e.Date.IsZero() {
		updateS = updateS.Set("date", e.Date)
	}

	if !e.Amount.IsZero() {
		updateS = updateS.Set("amount", e.Amount)
	}
	if e.Category != "" {
		updateS = updateS.Set("category", e.Category)
	}
	if e.Description != "" {
		updateS = updateS.Set("description", e.Description)
	}
	if e.PaymentMethod != "" {
		updateS = updateS.Set("payment_method", e.PaymentMethod)
	}

	query, params, err := updateS.ToSql()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal sql error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("success: query: %s, params: %v\n", query, params)

	if _, err := h.db.Exec(r.Context(), query, params...); err != nil {
		http.Error(w, "error updating expense data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "successful"}`))
}
func (h handle) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("conversion error %v", err)
	}
	query, params, err := sq.Delete("expenses").Where("transactionid = ?", id).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		http.Error(w, "internal sql error", http.StatusInternalServerError)
		return
	}

	fmt.Println(query, params)

	if _, err := h.db.Exec(r.Context(), query, params...); err != nil {
		http.Error(w, "error deleting expense data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "successful"}`))
}
