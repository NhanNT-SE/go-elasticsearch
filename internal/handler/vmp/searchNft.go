package vmp

import (
	"fmt"
	"log"
	"marketplace-backend/pkg/elastic"
	"marketplace-backend/pkg/storage"

	"net/http"
)

func (h *Handler) SearchNFT(r *http.Request, req *storage.SearchNFTRequest, resp *elastic.SearchResults[storage.NFTIndex]) error {
	if req.ResponseConfig.Size < 1 {
		return fmt.Errorf("limit: min=1")
	}

	if req.ResponseConfig.From < 0 {
		return fmt.Errorf("offset: min=0")
	}

	data, err := h.storageNFTSrv.SearchByQuery(r.Context(), *req)
	if err != nil {
		return err
	}
	idList := []string{}

	for _, doc := range data.Data {
		idList = append(idList, doc.Id)
	}

	// model.Token.GetTokenByIdList(idList)
	h.tokenSrv.GetTokenList(idList)
	// log.Println(idList)
	*resp = data
	return nil
}

func (h *Handler) CreateNFT(r *http.Request, req *storage.NFTIndex, resp *string) error {
	id := "123456"
	err := h.storageNFTSrv.InsertNFT(r.Context(), *req, id)
	if err != nil {
		return err
	}
	*resp = "NFT created"
	return nil
}
func (h *Handler) UpdateNFT(r *http.Request, req *storage.NFTIndex, resp *string) error {
	id := "123456"
	err := h.storageNFTSrv.UpdateNFT(r.Context(), *req, id)
	if err != nil {
		return err
	}
	*resp = "NFT updated"
	return nil
}

func (h *Handler) DeleteNFT(r *http.Request, req *elastic.DeleteIndexReq, resp *string) error {
	if req.DocId == "" {
		return fmt.Errorf("doc_id is required")
	}
	err := h.storageNFTSrv.DeleteNFT(r.Context(), req.DocId)
	if err != nil {
		log.Println(err)
		return err
	}
	*resp = "NFT deleted"
	return nil
}
