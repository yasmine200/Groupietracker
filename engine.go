package Backgo

import (
	"encoding/json"
	`errors`
	"net/http"
)

// Artist : Structure représentant un artiste
type Artist struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	ImageURL   string   `json:"image"`
	Locations  []string `json:"locations"`
	Dates      []string `json:"dates"`
}

// Liste des artistes
var Artists []Artist

// LoadArtists : Fonction pour charger les artistes depuis une API
func LoadArtists(apiURL string) error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(`Erreur lors de la récupération des données API`)
	}

	return json.NewDecoder(resp.Body).Decode(&Artists)
}
