# API Golang com Kafka 🚀

Este projeto demonstra uma API robusta e escalável construída em **Go** com **Kafka**, utilizando **Clean Architecture**. A API consome mensagens de Kafka para gerar e gerenciar dados de **produtos** de maneira assíncrona. O frontend é desenvolvido com **Vue.js** e **TypeScript** para exibir os produtos interativamente.

## Funcionalidades ✨

- **API em Golang**: Gerenciamento de produtos e consumo de mensagens do Kafka.
- **Integração com Kafka**: Processamento assíncrono de dados.
- **Frontend Vue.js + TypeScript**: Interface interativa para exibição dos produtos.
- **Clean Architecture**: Código organizado e escalável.

## Tecnologias Utilizadas 🛠

- **Go (Golang)** Backend API.
- **Kafka** Mensageria assíncrona.
- **Vue.js** e **TypeScript** Frontend reativo e seguro.
- **PostgreSQL** Persistência de dados.
- **Docker** Ambiente de desenvolvimento isolado e consistente.

## Arquitetura Limpa + DI 🧩
A aplicação segue o padrão Clean Architecture, com a divisão clara de responsabilidades entre camadas. 
Além disso, a aplicação utiliza Injeção de Dependências (DI) para garantir que as dependências entre os componentes sejam gerenciadas de forma flexível e desacoplada. Isso facilita a manutenção, os testes unitários e a escalabilidade do sistema.


## Passo a Passo 🏃‍♂️

1. **Mesmo que mantenha o DB padrão (POSTGRES) da aplicação, ainda é necessário configurá-lo.**
```sql
  CREATE TABLE products (
      ID UUID PRIMARY KEY,
      Name VARCHAR(255) NOT NULL,
      Price NUMERIC(10, 2) NOT NULL
  );
```



2. **Subir os Containers (Kafka, Zookeeper, PostgreSQL, API e Frontend):**
```bash
  docker-compose up -d
```
3. **Acesse o container do KAFKA e crie o Tópico "product":**
```bash
  kafka-topics --bootstrap-server localhost:9092 --create --topic product --if-not-exists
```
4. **Entre no container da API e execute o servidor:**
```bash
  go run cmd/main.go
```

Após feito o passo a passo, o frontend estará acessível LOCALMENTE em: [http://localhost:8181](http://localhost:8181).

## Como Funciona? 🔄

1. **Mensagens Kafka**: A API consome mensagens do Kafka associadas ao tópico `product` e processa os dados.
2. **Backend**: A API armazena os dados processados no PostgreSQL.
3. **Frontend**: O Vue.js exibe os produtos de forma interativa, consumindo a API.

### Desenvolvido por [@jalvess021](https://github.com/jalvess021).
