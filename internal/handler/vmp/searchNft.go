package vmp

import (
	"fmt"
	"marketplace-backend/internal/model"
	"marketplace-backend/pkg/elastic"

	"net/http"
)

func (h *Handler) SearchNFT(r *http.Request, req *model.NFTSearchReq, resp *model.GetTokenListRes) error {
	if req.ResponseConfig.Size < 1 {
		return fmt.Errorf("limit: min=1")
	}

	if req.ResponseConfig.From < 0 {
		return fmt.Errorf("offset: min=0")
	}

	if req.Price.Currency == "" {
		return fmt.Errorf("required field: currency")
	}

	data, err := h.tokenSrv.GetTokenList(r.Context(), req)
	if err != nil {
		return err
	}

	*resp = data

	return nil
}

func (h *Handler) CreateNFT(r *http.Request, req *model.NFTIndex, resp *string) error {
	id := "123456"
	err := h.storageNFTSrv.InsertNFT(r.Context(), *req, id)
	if err != nil {
		return err
	}
	*resp = "NFT created"
	return nil
}
func (h *Handler) UpdateNFT(r *http.Request, req *model.NFTIndex, resp *string) error {
	id := "123456"
	err := h.storageNFTSrv.UpdateNFT(r.Context(), *req, id)
	if err != nil {
		return err
	}
	*resp = "NFT updated"
	return nil
}

func (h *Handler) FindNFTById(r *http.Request, req *elastic.DeleteIndexReq, resp *interface{}) error {
	if req.DocId == "" {
		return fmt.Errorf("doc_id is required")
	}
	data, err := h.storageNFTSrv.FindNFTById(r.Context(), req.DocId)
	if err != nil {
		return err
	}
	*resp = data
	return nil
}

func (h *Handler) DeleteNFT(r *http.Request, req *elastic.DeleteIndexReq, resp *string) error {
	if req.DocId == "" {
		return fmt.Errorf("doc_id is required")
	}
	err := h.storageNFTSrv.DeleteNFT(r.Context(), req.DocId)
	if err != nil {
		return err
	}
	*resp = "NFT Deleted"
	return nil
}
