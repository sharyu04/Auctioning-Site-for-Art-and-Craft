package api

import (
	"encoding/json"
	"net/http"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/bid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func createBidHandler(bidSvc bid.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateBidRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		req.Bidder_id = r.Header.Get("user_id")

		resBody, err := bidSvc.CreateBid(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		resJson, err := json.Marshal(resBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write(resJson)
		return
	}
}

func updateBidHandler(bidSvc bid.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.UpdateBidRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		user_id := r.Header.Get("user_id")

		resBody, err := bidSvc.UpdateBid(req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		resJson, err := json.Marshal(resBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write(resJson)
		return

	}
}
