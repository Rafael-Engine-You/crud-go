package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"crud-go/internal/auth"
	"crud-go/internal/clientes"

	// 1. Importações do Swagger
	_ "crud-go/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// 2. Anotações OpenAPI (Swagger)
// @title API REST - CRUD Go
// @version 1.0
// @description API robusta desenvolvida em Go com Clean Architecture, PostgreSQL e JWT.
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte da API
// @contact.email suporte@exemplo.com

// @host localhost:8081
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado.")
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("Erro: A variável DATABASE_URL não está definida.")
	}

	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}
	defer db.Close()

	repository := clientes.NewRepository(db)
	service := clientes.NewService(repository)
	handler := clientes.NewHandler(service)

	authService := auth.NewAuthService(db)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// 3. Rota para servir a documentação interativa do Swagger UI
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))

	// --- ROTAS PÚBLICAS ---
	router.Get("/clientes", handler.ListarTodosClientes)
	router.Get("/clientes/{id}", handler.BuscarClientePorId)
	router.Post("/login", authService.LoginHandler)

	// --- ROTAS PROTEGIDAS POR JWT ---
	router.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware)

		r.Post("/clientes", handler.CadastrarCliente)
		r.Put("/clientes/{id}", handler.AtualizarCliente)
		r.Delete("/clientes/{id}", handler.DeletarCliente)
	})

	log.Println("Servidor executando em http://localhost:8081")
	log.Println("Documentação Swagger disponível em http://localhost:8081/swagger/index.html")

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
