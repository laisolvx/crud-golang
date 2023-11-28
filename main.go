package main

import ( // importa pacotes necessários: fmt para formatação de saída, log para registro de erros, os para interação com o sistema operacional e github.com/joho/godotenv para carregar variáveis de ambiente de um arquivo chamado .env.
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laisolvx/crud-golang/src/controller/routes"
)

func main() {
	err := godotenv.Load() // tenta carregar as variáveis de ambiente do arquivo .env usando a função godotenv.Load().
	if err != nil {
		log.Fatal("error loading .env file") // se houver algum erro ao carregar o arquivo .env, o programa exibe uma mensagem de erro usando log.Fatal() e encerra.
	}
	fmt.Println(os.Getenv("TEST")) // caso contrário, imprime na tela o valor associado à variável de ambiente "TEST" usando os.Getenv("TEST").

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
