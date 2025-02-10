package backgo 

import "log"

// Init : Fonction pour initialiser les artistes
func Init() {
	apiURL := "https://groupietrackers.herokuapp.com/artists" // Change cette URL avec celle de ton API
	err := LoadArtists(apiURL)
	if err != nil {
		log.Fatal("Erreur lors du chargement des artistes :", err)
	}
}
