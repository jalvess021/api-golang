package main

import (
	"encoding/json"
	"net/http"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jalvess021/api-golang/internal/db" // Atualize o caminho conforme a sua estrutura de diretórios
	"github.com/jalvess021/api-golang/internal/infra/akafka"
	"github.com/go-chi/chi/v5"
	"github.com/jalvess021/api-golang/internal/infra/web"
	"github.com/jalvess021/api-golang/internal/usecase/product" 
	"github.com/jalvess021/api-golang/internal/infra/repository/pgsql"
)

func main() {

	db, err := db.ConnectDb();
	if  err != nil {
		panic(err)
	}
	defer db.Close();

	r := chi.NewRouter()
	
	// chama as rotas
	web.SetupRoutes(r, db);
	go http.ListenAndServe(":8282", r)
	
	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "kafka:9092", msgChan)

	//Cria repositorio e casos de usuo
	repository := repository.NewProductRepositoryPgsql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)

	//Acessando por dto
	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			//logar erro
		}
		_, err = createProductUseCase.Execute(dto)
	}
}
