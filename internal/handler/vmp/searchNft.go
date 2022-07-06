package vmp

import (
	"fmt"
	"marketplace-backend/pkg/elastic"
	"marketplace-backend/pkg/storage"

	"net/http"
)

type SearchResp struct {
	Data any `json:"data"`
	Req  any `json:"req"`
}

func (h *Handler) SearchNFT(r *http.Request, req *storage.SearchNFTRequest, resp *elastic.SearchResults[storage.NFTIndex]) error {
	if req.ResponseConfig.Size < 1 {
		return fmt.Errorf("limit: min=1")
	}

	if req.ResponseConfig.From < 0 {
		return fmt.Errorf("offset: min=0")
	}
	search, err := h.storageNFTSrv.SearchByQuery(r.Context(), *req)
	if err != nil {
		return err
	}
	*resp = search
	return nil
}

// func (h *Handler) CreateNFT(r *http.Request, req *NFTIndex, resp *SearchResp) error {
// 	store := elastic.NewStore[NFTIndex](h.esClient, "marketplace-nfts", time.Second*10)
// 	docId := fmt.Sprintf("%v|%v", req.CollectionId, req.NftId)
// 	err := store.CreateIndex(req, docId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (h *Handler) UpdateNFT(r *http.Request, req *NFTIndex, resp *SearchResp) error {
// 	store := elastic.NewStore[NFTIndex](h.esClient, "marketplace-nfts", time.Second*10)
// 	docId := fmt.Sprintf("%v|%v", req.CollectionId, req.NftId)
// 	err := store.UpdateIndex(req, docId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (h *Handler) DeleteNFT(r *http.Request, req *string, resp *SearchResp) error {
// 	store := elastic.NewStore[NFTIndex](h.esClient, "marketplace-nfts", time.Second*10)
// 	err := store.DeleteIndex(r.Context(), *req)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
