package items

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/baglabs/bookstore_items-api/clients/elasticsearch"
	"github.com/baglabs/bookstore_items-api/domain/queries"
	"github.com/baglabs/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	indexTypes = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	defer result.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		return rest_errors.NewInternalServerError("failed when trying to save item", errors.New("databse error"))
	}
	i.Id = r["_id"].(string)

	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	result, err := elasticsearch.Client.Get(indexItems, indexTypes, i.Id)
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	defer result.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		return rest_errors.NewInternalServerError("failed when trying to get item", errors.New("databse error"))
	}

	// i.Id = r["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	i.Id = r["_id"].(string)

	return nil

}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search document", errors.New("database error"))
	}
	defer result.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		return nil, rest_errors.NewInternalServerError("failed when trying to get item", errors.New("databse error"))
	}
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})

	items := make([]Item, len(hits))
	for index, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var item Item
		bytes, _ := json.Marshal(hit.(map[string]interface{})["_source"])
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.(map[string]interface{})["_id"].(string)
		items[index] = item
	}

	return items, nil
}
