package main

import (
    "log"
    "net/http"
    // "strings"
    "fmt"
    "regexp"
)

// // Função para tratar a rota dinâmica de um produto
// func produtoHandler(w http.ResponseWriter, r *http.Request) {
// 	// Capturando o parâmetro 'id' da URL
// 	// Exemplo de URL: /produto/12345
// 	parts := strings.Split(r.URL.Path, "/")
// 	if len(parts) < 3 {
// 		http.Error(w, "ID de produto inválido", http.StatusBadRequest)
// 		return
// 	}

// 	idProduto := parts[2] // Pegando o ID do produto da URL

// 	// Aqui você pode validar se o idProduto existe no banco de dados, por exemplo
// 	// Para fins de exemplo, vamos apenas simular a resposta
// 	fmt.Fprintf(w, "Mostrando o produto com ID: %s", idProduto)
// }

func produtoHandler(w http.ResponseWriter, r *http.Request) {
	// Usando uma expressão regular para capturar o ID do produto
	// Exemplo de URL: /produto/12345 ou /produto/prod-12345
	re := regexp.MustCompile(`^/produto/([a-zA-Z0-9\-]+)$`)

	// Tentando combinar a URL com a expressão regular
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	// O primeiro item de matches é a string inteira, o segundo é o grupo capturado (ID do produto)
	idProduto := matches[1]

	// Lógica para buscar o produto pelo ID
	fmt.Fprintf(w, "Mostrando o produto com ID: %s", idProduto)
}

func main() {
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

    // Rota para a página 'sobre'
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// Aqui você pode criar uma resposta personalizada, como enviar um HTML diferente
		http.ServeFile(w, r, "public/about.html")
	})

    // Rota para mostrar informações de um produto específico
	http.HandleFunc("/produto/", produtoHandler)


    log.Println("Executando...")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
