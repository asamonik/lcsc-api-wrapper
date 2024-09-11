package api

import (
	"encoding/json"
	"fmt"
)

type Category struct {
	CatalogID     int         `json:"catalogId"`
	ParentID      int         `json:"parentId"`
	ParentName    interface{} `json:"parentName"`
	CatalogName   string      `json:"catalogName"`
	CatalogNameEn string      `json:"catalogNameEn"`
	ChildCatelogs []struct {
		CatalogID     int        `json:"catalogId"`
		ParentID      int        `json:"parentId"`
		ParentName    string     `json:"parentName"`
		CatalogName   string     `json:"catalogName"`
		CatalogNameEn string     `json:"catalogNameEn"`
		ChildCatelogs []Category `json:"childCatelogs"`
		ProductNum    int        `json:"productNum"`
	} `json:"childCatelogs"`
	ProductNum int `json:"productNum"`
}

func (c *Client) GetCategories() (*[]Category, error) {
	path := "/ftps/wm/product/catalogs/search"

	result, err := c.makeRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var categories []Category
	err = json.Unmarshal(result, &categories)
	if err != nil {
		return nil, fmt.Errorf("error parsing category result: %v", err)
	}

	return &categories, nil
}
