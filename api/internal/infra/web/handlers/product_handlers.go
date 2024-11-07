package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/jalvess021/api-golang/api/internal/usecase/product"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUsecase
	ListProductUseCase *usecase.ListProductsUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUsecase, listProductUseCase *usecase.ListProductsUseCase) *ProductHandlers{
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductUseCase: listProductUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid request payload: %v", err)
		http.Error(w, errorMessage, http.StatusBadRequest)
		fmt.Printf("Error decoding request body: %v\n", err)
		return
	}

	// Execute use case
	output, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		errorMessage := fmt.Sprintf("Error creating product: %v", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		fmt.Printf("Error executing CreateProductUseCase: %v\n", err)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		errorMessage := fmt.Sprintf("Error encoding response: %v", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

func (p *ProductHandlers) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductUseCase.Execute()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing products: %v", err), http.StatusInternalServerError)
		fmt.Printf("Error executing use case: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
	}
}
