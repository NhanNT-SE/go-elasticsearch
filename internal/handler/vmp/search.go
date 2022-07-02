package vmp

import (
	"net/http"
)

type SearchRequest struct {
	Text string `json:"text"`
}

type SearchResponse struct {
	Result string `json:"result"`
}

func (s *Handler) Search(r *http.Request, req *SearchRequest, resp *SearchResponse) error {
	resp.Result = req.Text
	s.log.Info().Interface("req", req).Interface("resp", resp).Send()
	return nil
}
