package Backgo

import (
	"html/template"
	"net/http"
	"strconv"
)

// HomeHandler : Affiche la page d'accueil avec tous les artistes
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("serv/home.html")
	if err != nil {
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, Artists)
}

// ArtistHandler : Affiche les détails d'un artiste spécifique
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	for _, artist := range Artists {
		if artist.ID == id {
			tmpl, err := template.ParseFiles("serv/artist.html")
			if err != nil {
				http.Error(w, "Erreur de template", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, artist)
			return
		}
	}
	http.NotFound(w, r)
}
