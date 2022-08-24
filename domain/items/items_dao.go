package items

import (
	"encoding/json"
	"errors"

	"github.com/baglabs/bookstore_items-api/clients/elasticsearch"
	"github.com/baglabs/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
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

	return nil
}
