package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/bid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/apperrors"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
)

func createBidHandler(bidSvc bid.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateBidRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		req.Bidder_id = r.Header.Get("user_id")

		resBody, err := bidSvc.CreateBid(req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		resJson, err := json.Marshal(resBody)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
		return
	}
}

func updateBidHandler(bidSvc bid.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.UpdateBidRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		user_id := r.Header.Get("user_id")

		resBody, err := bidSvc.UpdateBid(req, user_id)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		resJson, err := json.Marshal(resBody)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
		return

	}
}

func deleteBidHandler(bidSvc bid.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.DeleteBidRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return

		}

		fmt.Println(req.BidId)

		user_id := r.Header.Get("user_id")
		role := r.Header.Get("role")

		res, err := bidSvc.DeleteBid(user_id, role, req.BidId)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			errResponse := apperrors.MapError(err)
			w.WriteHeader(errResponse.ErrorCode)
			res, _ := json.Marshal(errResponse)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
		return
	}
}
