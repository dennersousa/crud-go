package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// User representa a estrutura de dados para um usuário.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Estabelece a conexão com o banco de dados SQLite.
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria a tabela "users" se não existir.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Configuração das rotas HTTP.
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// Manipula diferentes métodos HTTP para a rota "/users".
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r, db)
		case http.MethodPost:
			createUser(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Configuração das rotas HTTP para operações específicas de usuários.
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUser(w, r, db)
		case http.MethodPut:
			updateUser(w, r, db)
		case http.MethodDelete:
			deleteUser(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Inicia o servidor na porta 8000 com o middleware para conteúdo JSON.
	log.Fatal(http.ListenAndServe(":8000", JSONContentTypeMiddleware(http.DefaultServeMux)))
}

// JSONContentTypeMiddleware define o tipo de conteúdo JSON nas respostas HTTP.
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// getUsers retorna todos os usuários do banco de dados.
func getUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Serializa os usuários em formato JSON e envia como resposta.
	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// getUser obtém um usuário pelo ID.
func getUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extrai o ID da URL.
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User
	// Consulta o banco de dados para obter um usuário pelo ID.
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Serializa o usuário em formato JSON e envia como resposta.
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// createUser cria um novo usuário com base nos dados do corpo da requisição.
func createUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user User
	// Decodifica os dados do corpo da requisição para a estrutura User.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insere o novo usuário no banco de dados.
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtém o ID do último usuário inserido.
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = int(lastInsertID)

	// Serializa o usuário criado em formato JSON e envia como resposta.
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// updateUser atualiza um usuário existente com base no ID e nos dados do corpo da requisição.
func updateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extrai o ID da URL.
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User
	// Decodifica os dados do corpo da requisição para a estrutura User.
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Atualiza os dados do usuário no banco de dados.
	_, err = db.Exec("UPDATE users SET name=?, email=? WHERE id=?", user.Name, user.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteUser exclui um usuário com base no ID.
func deleteUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Extrai o ID da URL.
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Exclui o usuário do banco de dados.
	_, err = db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
