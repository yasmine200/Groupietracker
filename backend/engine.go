package backgo

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sort"
)

// Structure représentant un artiste
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string `json:"locations"`
	Dates        []string `json:"dates"`
}

// Liste des artistes (stockée en mémoire après le chargement)
var Artists []Artist

// Fonction pour récupérer les artistes depuis une API
func LoadArtists(apiURL string) error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Erreur lors de la récupération des données API")
	}

	// Décodage de la réponse JSON dans Artists
	return json.NewDecoder(resp.Body).Decode(&Artists)
}

// Fonction pour filtrer les artistes en fonction de l'input utilisateur
func FilterArtists(query string) []Artist {
	var filtered []Artist
	for _, artist := range Artists {
		// Vérifier si le nom du groupe correspond à la requête
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filtered = append(filtered, artist)
			continue
		}

		// Vérifier si un membre du groupe correspond à la requête
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				filtered = append(filtered, artist)
				break
			}
		}
	}
	return filtered
}

// Fonction pour trier les artistes par date de création (ascendant ou descendant)
func SortArtistsByCreationDate(order string) []Artist {
	sortedArtists := make([]Artist, len(Artists))
	copy(sortedArtists, Artists) // On fait une copie pour éviter de modifier l'original

	if order == "croi" {
		sort.Slice(sortedArtists, func(i, j int) bool {
			return sortedArtists[i].CreationDate < sortedArtists[j].CreationDate
		})
	} else if order == "dcroi" {
		sort.Slice(sortedArtists, func(i, j int) bool {
			return sortedArtists[i].CreationDate > sortedArtists[j].CreationDate
		})
	}
	return sortedArtists
}