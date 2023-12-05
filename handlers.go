package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func christmas(w http.ResponseWriter, r *http.Request) {
	generateCharacters(w, r, "christmas")
}

func halloween(w http.ResponseWriter, r *http.Request) {
	generateCharacters(w, r, "halloween")
}

func easter(w http.ResponseWriter, r *http.Request) {
	generateCharacters(w, r, "easter")
}

func generateCharacters(w http.ResponseWriter, r *http.Request, theme string) {
	num := extractNumFromURL(r.URL.Path)
	characters := make([]Character, num)
	for i := 0; i < num; i++ {
		character, err := generateRandomCharacter(theme)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		characters[i] = character
	}
	json.NewEncoder(w).Encode(characters)
}

func extractNumFromURL(urlPath string) int {
	parts := strings.Split(urlPath, "/")
	if len(parts) < 3 {
		return 1
	}
	num, err := strconv.Atoi(parts[2])
	if err != nil || num <= 0 {
		return 1
	}
	return num
}
