package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Funções para CRIAR, BUSCAR, ATUALIZAR e DELETAR usuários no banco de dados

func CriarUsuario (w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil{
		log.Fatal(erro)
	}

	db, erro := db.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioId)))
	
}
func BuscarUsuarios (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuario Usuários!"))
}
func BuscarUsuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário!"))
}
func AtualizarUsuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}
func DeletarUsuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}