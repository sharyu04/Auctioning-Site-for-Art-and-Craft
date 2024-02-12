package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func createArtworkHandler(artworkSvc artwork.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateArtworkRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		req.Owner_id = r.Header.Get("user_id")

		resBody, err := artworkSvc.CreateArtwork(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		respJson, err := json.Marshal(resBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}
}

func GetArtworksHandler(artworkSvc artwork.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")
		start := r.URL.Query().Get("start")
		count := r.URL.Query().Get("count")

		if start == "" || count == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("start and count values missing!"))
			return
		}

		startInt, _ := strconv.Atoi(start)
		countInt, _ := strconv.Atoi(count)

		res, err := artworkSvc.GetArtworks(category, startInt, countInt)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		resBody, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(resBody)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	}
}

func GetArtworkByIdHandler(artworkSvc artwork.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		res, err := artworkSvc.GetArtworkByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		resBody, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
		return
	}
}

func DeleteArtworkHandler(artworkSvc artwork.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		Owner_id := r.Header.Get("user_id")
		role := r.Header.Get("role")

		err := artworkSvc.DeleteArtworkById(id, Owner_id, role)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Artwork deleted"))
		return
	}
}
