package main

import (
	"bookstore-api/config"
	"bookstore-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carregar variáveis de ambiente
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Inicializar o banco de dados
	db, err := config.SetupDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Executar migrações
	err = config.MigrateDB(db)
	if err != nil {
		log.Fatal("Erro ao executar migrações:", err)
	}

	// Configurar o router
	r := gin.Default()

	// Configurar as rotas
	routes.SetupRoutes(r, db)

	// Iniciar o servidor
	r.Run(":8080")
}