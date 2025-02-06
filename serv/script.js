document.addEventListener('DOMContentLoaded', function () {
    const apiUrl = "https://groupietrackers.herokuapp.com/api/artists";

    fetch(apiUrl)
        .then(response => response.json())
        .then(data => displayArtists(data))
        .catch(error => console.error("Erreur API :", error));
});

function displayArtists(artists) {
    const artistsList = document.getElementById('artistsList');
    artistsList.innerHTML = '';

    artists.forEach(artist => {
        const li = document.createElement('li');
        li.textContent = artist.name;
        artistsList.appendChild(li);
    });
}