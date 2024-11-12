package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	repository "github.com/jalvess021/kartka/api/internal/infra/repository/pgsql"
	"github.com/jalvess021/kartka/api/internal/infra/web/handlers"
	usecase "github.com/jalvess021/kartka/api/internal/usecase/product"
)

func SetupProductRoutes(db *sql.DB) chi.Router {
	// Cria um subroteador
	r := chi.NewRouter()

	//Cria repositorio e casos de usuo
	repository := repository.NewProductRepositoryPgsql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	//Criar handlers
	productHandlers := handlers.NewProductHandlers(createProductUseCase, listProductsUseCase)

	//Rotas de produtos
	r.Get("/", productHandlers.ListProductHandler)
	r.Post("/create", productHandlers.CreateProductHandler)

	return r
}
