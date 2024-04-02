package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/artwork"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func createArtworkHandler(artworkSvc artwork.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateArtworkRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		req.Owner_id = r.Header.Get("user_id")

		resBody, err := artworkSvc.CreateArtwork(req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		respJson, err := json.Marshal(resBody)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}
}

type resp struct {
	RespData   []dto.GetArtworkResponse `json:"respBody"`
	TotalCount int                      `json:"totalCount"`
}

func GetArtworksHandler(artworkSvc artwork.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")
		start := r.URL.Query().Get("start")
		count := r.URL.Query().Get("count")

		if start == "" || count == "" {
			err := apperrors.BadRequest{ErrorMsg: "Start and count values are missing"}
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		startInt, _ := strconv.Atoi(start)
		countInt, _ := strconv.Atoi(count)

		res, totalCount, err := artworkSvc.GetArtworks(category, startInt, countInt)
		fmt.Println("totalCount = ", totalCount)

		if err != nil {
			errResponse := apperrors.MapError(err)
			fmt.Printf("errResponse %+v", errResponse)
			w.WriteHeader(errResponse.ErrorCode)
			res, err := json.Marshal(errResponse)
			if err != nil {
				fmt.Printf("errResponse 78%+v", err)
			}
			w.Write(res)
			return
		}

		respBody := resp{
			RespData:   res,
			TotalCount: totalCount,
		}

		fmt.Println(respBody)

		resBody, err := json.Marshal(respBody)

		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	}
}

func GetArtworkByIdHandler(artworkSvc artwork.Service) http.HandlerFunc {
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
