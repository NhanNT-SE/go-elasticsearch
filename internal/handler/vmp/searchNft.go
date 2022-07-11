package vmp

import (
	"errors"
	"marketplace-backend/model"

	"net/http"
)

func (h *Handler) SearchNFT(r *http.Request, req *model.NFTIndexSearchReq, resp *model.NFTIndexSearchRes) error {
	if req.ResponseConfig.Size < 1 {
		return errors.New("limit: min=1")
	}
	if req.ResponseConfig.From < 0 {
		return errors.New("offset: min=0")
	}

	if req.Price.Currency == "" {
		return errors.New("required field: currency")
	}

	data, err := h.tokenRepo.GetTokenList(r.Context(), req)
	if err != nil {
		return errors.New("search fail")
	}
	*resp = data
	return nil
}

// func (h *Handler) UpdateNFT(r *http.Request, req *model.Token, resp *model.NFTIndexSearchRes) error {
// 	err := h.tokenRepo.UpdateToken(r.Context(), req)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("update token fail")
// 	}
// 	return nil
// }

// func (h *Handler) DeleteNFT(r *http.Request, req *model.Token, resp *model.NFTIndexSearchRes) error {
// 	err := h.tokenRepo.DeleteToken(r.Context(), req)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("delete token fail")
// 	}
// 	return nil
// }
