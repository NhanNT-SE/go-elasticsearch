package vmp

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type NftIndex struct {
	NftId           string     `json:"nft_id,omitempty"`
	CollectionId    string     `json:"collection_id,omitempty"`
	Name            string     `json:"name,omitempty"`
	Description     string     `json:"description,omitempty"`
	Price           int        `json:"price,omitempty"`
	SaleType        string     `json:"sale_type,omitempty"`
	CreatedTime     time.Time  `json:"created_time,omitempty"`
	LastSoldTime    time.Time  `json:"last_sold_time,omitempty"`
	ListedTime      time.Time  `json:"listed_time,omitempty"`
	BackgroundColor string     `json:"background_color,omitempty"`
	Image           string     `json:"image,omitempty"`
	Attributes      []NftAttrs `json:"attributes,omitempty"`
}

type NftAttrs struct {
	TraitType    string `json:"trait_type,omitempty"`
	DisplayValue string `json:"display_value,omitempty"`
	Value        int    `json:"value,omitempty"`
}
type SearchRequest struct {
	Text string `json:"text"`
}

type SearchResponse struct {
	Result map[string]interface{} `json:"result"`
}

func (s *Handler) Search(r *http.Request, req *SearchRequest, resp *SearchResponse) error {
	var mapRes map[string]interface{}

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Println("parse error", err)
	}

	es := s.esClient

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("marketplace_nfts"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&mapRes); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)

	}
	// var b bytes.Buffer
	// var mapResp map[string]interface{}

	// b.ReadFrom(res.Body)

	// log.Println(b)
	// Print the response status, number of results, and request duration.

	// log.Printf(
	// 	"[%s] %d hits; took: %dms",
	// 	res.Status(),
	// 	int(mapRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	// 	int(mapRes["took"].(float64)),
	// )
	// // Print the ID and document source for each hit.
	// for _, hit := range mapRes["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	// }

	// log.Println(strings.Repeat("=", 37))
	resp.Result = mapRes
	return nil
}
