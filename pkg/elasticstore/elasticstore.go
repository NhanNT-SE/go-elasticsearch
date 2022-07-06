package elasticstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"marketplace-backend/pkg/logger"
	"math"
	"reflect"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	log = logger.New()
)

type Store[T any] struct {
	es              *elasticsearch.Client
	indexName       string
	resSearchConfig ResponseSearchConfig
}

type Hit[T any] struct {
	Id  string `json:"doc_id"`
	Doc T      `json:"doc"`
}

type ResponseSearchConfig struct {
	Source any `json:"source" elastic:"_source"`
	Size   int `json:"limit" elastic:"size"`
	From   int `json:"offset" elastic:"from"`
}

type SearchResults[T any] struct {
	Pagination Pagination `json:"pagination"`
	Data       []Hit[T]   `json:"data,omitempty"`
}

type RangeQueryReq struct {
	From any `json:"from"`
	To   any `json:"to"`
}

type Pagination struct {
	TotalDoc  int `json:"total_doc,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
	Limit     int `json:"limit,omitempty"`
	Offset    int `json:"offset"`
}

func NewStore[T any](esClient *elasticsearch.Client, indexName string) *Store[T] {
	s := Store[T]{es: esClient, indexName: indexName}
	return &s
}

func (s *Store[T]) SetResponseSearchConfig(config ResponseSearchConfig) {
	s.resSearchConfig = config
}

func (s *Store[T]) SearchByQuery(query *bytes.Buffer, resp *SearchResults[T]) error {
	var mapRes map[string]interface{}
	es := s.es
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(s.indexName),
		es.Search.WithBody(query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		} else {
			err = fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"])
			return err
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&mapRes); err != nil {
		return err
	}

	var hitList []Hit[T]
	hits := mapRes["hits"].(map[string]interface{})
	docs := hits["hits"].([]interface{})
	total := hits["total"].((map[string]interface{}))["value"].(float64)

	for _, doc := range docs {
		var hit Hit[T]
		docId := doc.(map[string]interface{})["_id"].(string)
		hit.Id = docId
		source := doc.(map[string]interface{})["_source"]
		b, err := json.Marshal(source)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &hit.Doc)
		if err != nil {
			return err
		}
		hitList = append(hitList, hit)
	}

	// Set data response for pagination
	resp.Data = hitList
	resp.Pagination = setPagination(s.resSearchConfig.Size, s.resSearchConfig.From, int(total))

	return nil
}

func (s *Store[T]) BuildQuery(mapQuery *map[string]interface{}) (*bytes.Buffer, error) {
	query := map[string]interface{}{
		"query": *mapQuery,
	}

	config := s.resSearchConfig

	v := reflect.ValueOf(config)

	if !v.IsZero() {
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			tag := typeOfS.Field(i).Tag.Get("elastic")
			value := v.Field(i).Interface()

			if value != nil && !reflect.ValueOf(value).IsZero() {
				if tag == "_source" && len(value.([]interface{})) == 0 {
					value = false
				}
				query[tag] = value
			}
		}

	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		return nil, err
	}
	return &buf, nil
}

func (s *Store[T]) BuildRangeQuery(rangeReq *RangeQueryReq, fieldName string) *map[string]interface{} {
	mRange := make(map[string]interface{})
	m := make(map[string]interface{})

	if rangeReq.From != nil && rangeReq.From != "" {
		mRange["gte"] = rangeReq.From
	}
	if rangeReq.To != nil && rangeReq.To != "" {
		mRange["lte"] = rangeReq.To
	}

	m["range"] = map[string]interface{}{
		fieldName: mRange,
	}
	return &m
}

func (s *Store[T]) BuildMatchAllQuery() *map[string]interface{} {
	m := map[string]interface{}{
		"match_all": map[string]interface{}{},
	}
	return &m
}

func (s *Store[T]) BuildMultiMatchQuery(query string, fields []string, isFuzzy bool, fuzziness int) *map[string]interface{} {
	queryM := map[string]interface{}{
		"query":  query,
		"fields": fields,
	}
	if isFuzzy && fuzziness >= 0 && fuzziness <= 2 {
		queryM["fuzziness"] = fuzziness

	}
	return &map[string]interface{}{
		"multi_match": queryM,
	}
}

func (s *Store[T]) BuildTermsQuery(fieldName string, values []string) *map[string]interface{} {
	return &map[string]interface{}{
		"terms": map[string]interface{}{
			fieldName: values,
		},
	}
}
func (s *Store[T]) BuildTermQuery(fieldName string, value string) *map[string]interface{} {
	return &map[string]interface{}{
		"term": map[string]interface{}{
			fieldName: value,
		},
	}
}

func (s *Store[T]) BuildBoolQuery(boolType string, values *[]interface{}) *map[string]interface{} {
	return &map[string]interface{}{
		"bool": map[string]interface{}{
			boolType: values,
		},
	}
}
func (s *Store[T]) BuildNestedQuery(path string, query *map[string]interface{}) *map[string]interface{} {
	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"path":  path,
			"query": query,
		},
	}
	return &m
}

func setPagination(size, from, total int) Pagination {
	pagination := Pagination{}
	pagination.TotalDoc = total
	pagination.Limit = size
	pagination.Offset = from
	if total > 0 {
		totalPage := math.Ceil(float64(total) / float64(size))
		pagination.TotalPage = int(totalPage)
	}
	return pagination
}
