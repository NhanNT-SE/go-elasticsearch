package must

import (
	"encoding/json"
	"fmt"
	"marketplace-backend/config"
	"marketplace-backend/pkg/logger"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	log = logger.New()
)

func ConnectElasticsearch(config *config.ElasticsearchConfig) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses:              []string{config.Addr},
		Username:               config.Username,
		Password:               config.Password,
		CertificateFingerprint: "6657f5e39f675260240a7d28302f791157cfc41c79baca858c4bc7f846d64a66",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("connect elastic failed")
	}

	res, err := es.Info()
	if err != nil {
		log.Fatal().Err(err).Msg("get info elastic failed")
	}

	defer res.Body.Close()
	if res.IsError() {
		log.Fatal().Err(fmt.Errorf(res.String())).Msg("elastic response error")
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatal().Err(err).Msg("Error parsing elastic response body")
	}

	log.Debug().Msg("connect elastic successfully")

	return es
}
