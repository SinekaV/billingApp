package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginRequest struct{ Username, Password string }
type Product struct {
	Code, Name string
	Price, GST float64
}
type BillRequest struct {
	Products []struct {
		Code     string
		Quantity int
	}
}
type BillResponse struct{ TotalPrice, Tax float64 }

var (
	users    = map[string]string{"biller": "password123"}
	tokens   = map[string]string{}
	products = map[string]Product{}
)

func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)
	if users[req.Username] == req.Password {
		token := req.Username + "-token"
		tokens[token] = req.Username
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	products[p.Code] = p
	w.Write([]byte("Product added"))
}

func searchProduct(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if p, ok := products[code]; ok {
		json.NewEncoder(w).Encode(p)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func generateBill(w http.ResponseWriter, r *http.Request) {
	var req BillRequest
	json.NewDecoder(r.Body).Decode(&req)
	total, tax := 0.0, 0.0
	for _, item := range req.Products {
		if p, ok := products[item.Code]; ok {
			subtotal := float64(item.Quantity) * p.Price
			total += subtotal + subtotal*(p.GST/100)
			tax += subtotal * (p.GST / 100)
		}
	}
	json.NewEncoder(w).Encode(BillResponse{TotalPrice: total, Tax: tax})
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok := tokens[r.Header.Get("Authorization")]; !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/add-product", auth(addProduct))
	http.HandleFunc("/search-product", auth(searchProduct))
	http.HandleFunc("/generate-bill", auth(generateBill))
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
