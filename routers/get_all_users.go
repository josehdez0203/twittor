package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"twittor/bd"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
  log.Println(" 🆘", typeUser, page, search)

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor que 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.GetUsersAll(IDUsuario, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
