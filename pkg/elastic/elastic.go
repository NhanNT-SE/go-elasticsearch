package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

type StoreSrv[T any] interface {
	CreateIndex(ctx context.Context, index T, docId string) error
	UpdateIndex(ctx context.Context, index T, docId string) error
	DeleteIndex(ctx context.Context, docId string) error
	SearchByQuery(ctx context.Context, query *bytes.Buffer) (SearchResults[T], error)
	BuildQuery(mapQuery *map[string]interface{}) (*bytes.Buffer, error)
	BuildRangeQuery(rangeReq *RangeQueryReq, fieldName string) *map[string]interface{}
	BuildMatchAllQuery() *map[string]interface{}
	BuildMultiMatchQuery(query string, fields []string, isFuzzy bool, fuzziness int) *map[string]interface{}
	BuildTermsQuery(fieldName string, values []string) *map[string]interface{}
	BuildTermQuery(fieldName string, value string) *map[string]interface{}
	BuildBoolQuery(boolType string, values *[]interface{}) *map[string]interface{}
	BuildNestedQuery(path string, query *map[string]interface{}) *map[string]interface{}
	SetResponseSearchConfig(config ResponseSearchConfig)
}

type Store[T any] struct {
	es              *elasticsearch.Client
	timeout         time.Duration
	indexName       string
	resSearchConfig ResponseSearchConfig
}

type Hit[T any] struct {
	Id  string `json:"doc_id"`
	Doc T      `json:"doc,omitempty"`
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
	Offset    int `json:"offset,omitempty"`
}

type DeleteIndexReq struct {
	DocId string `json:"doc_id"`
}

func NewStoreSrv[T any](esClient *elasticsearch.Client, indexName string) StoreSrv[T] {
	return &Store[T]{es: esClient, indexName: indexName}
}

func NewStore[T any](esClient *elasticsearch.Client, indexName string, timeOut time.Duration) *Store[T] {
	s := Store[T]{es: esClient, indexName: indexName, timeout: timeOut}
	return &s
}

func (s *Store[T]) SetResponseSearchConfig(config ResponseSearchConfig) {
	s.resSearchConfig = config
}

func (s *Store[T]) SearchByQuery(ctx context.Context, query *bytes.Buffer) (SearchResults[T], error) {
	result := SearchResults[T]{}
	var mapRes map[string]interface{}
	es := s.es
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(s.indexName),
		es.Search.WithBody(query),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return result, err
		} else {
			err = fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"])
			return result, err
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&mapRes); err != nil {
		return result, err
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
			return result, err
		}
		err = json.Unmarshal(b, &hit.Doc)
		if err != nil {
			return result, err
		}
		hitList = append(hitList, hit)
	}

	result.Data = hitList
	result.Pagination = setPagination(s.resSearchConfig.Size, s.resSearchConfig.From, int(total))

	return result, nil
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

	// log.Println(&buf)
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

func (s *Store[T]) CreateIndex(ctx context.Context, index T, docId string) error {
	data, err := json.Marshal(index)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      s.indexName,
		DocumentID: docId,
		Body:       bytes.NewReader(data),
	}
	res, err := req.Do(ctx, s.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = checkErrorRes(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store[T]) UpdateIndex(ctx context.Context, index T, docId string) error {
	data, err := json.Marshal(index)
	if err != nil {
		return fmt.Errorf("error marshaling document: %s", err)
	}

	req := esapi.UpdateRequest{
		Index:      s.indexName,
		DocumentID: docId,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, data))),
	}
	res, err := req.Do(ctx, s.es)
	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	err = checkErrorRes(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store[T]) DeleteIndex(ctx context.Context, docId string) error {
	req := esapi.DeleteRequest{
		Index:      s.indexName,
		DocumentID: docId,
	}
	res, err := req.Do(ctx, s.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = checkErrorRes(res)
	if err != nil {
		return err
	}
	return nil
}

func checkErrorRes(res *esapi.Response) error {
	if res.StatusCode == 404 {
		return ErrNotFound
	}

	if res.StatusCode == 409 {
		return ErrConflict
	}
	if res.IsError() {
		return fmt.Errorf("error indexing document ID=%v", res.Status())
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing the response body: %s", err)
	}
	return nil
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
