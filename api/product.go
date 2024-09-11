package api

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ProductID         int     `json:"productId"`
	ProductCode       string  `json:"productCode"`
	ProductModel      string  `json:"productModel"`
	ParentCatalogID   int     `json:"parentCatalogId"`
	ParentCatalogName string  `json:"parentCatalogName"`
	CatalogID         int     `json:"catalogId"`
	CatalogName       string  `json:"catalogName"`
	BrandID           int     `json:"brandId"`
	BrandNameEn       string  `json:"brandNameEn"`
	EncapStandard     string  `json:"encapStandard"`
	Split             int     `json:"split"`
	ProductUnit       string  `json:"productUnit"`
	MinPacketUnit     string  `json:"minPacketUnit"`
	MinBuyNumber      int     `json:"minBuyNumber"`
	MaxBuyNumber      int     `json:"maxBuyNumber"`
	MinPacketNumber   int     `json:"minPacketNumber"`
	IsEnvironment     bool    `json:"isEnvironment"`
	AllHotLevel       any     `json:"allHotLevel"`
	IsHot             bool    `json:"isHot"`
	IsPreSale         bool    `json:"isPreSale"`
	IsReel            bool    `json:"isReel"`
	ReelPrice         float64 `json:"reelPrice"`
	StockNumber       int     `json:"stockNumber"`
	StockSz           int     `json:"stockSz"`
	StockJs           int     `json:"stockJs"`
	WmStockHk         int     `json:"wmStockHk"`
	DomesticStockVO   struct {
		Total           int `json:"total"`
		ShipImmediately int `json:"shipImmediately"`
		Ship3Days       int `json:"ship3Days"`
	} `json:"domesticStockVO"`
	OverseasStockVO struct {
		Total           int `json:"total"`
		ShipImmediately int `json:"shipImmediately"`
		Ship3Days       int `json:"ship3Days"`
	} `json:"overseasStockVO"`
	ProductPriceList []struct {
		Ladder              int     `json:"ladder"`
		ProductPrice        string  `json:"productPrice"`
		UsdPrice            float64 `json:"usdPrice"`
		CnyProductPriceList any     `json:"cnyProductPriceList"`
		DiscountRate        string  `json:"discountRate"`
		CurrencyPrice       float64 `json:"currencyPrice"`
		CurrencySymbol      string  `json:"currencySymbol"`
		IsForeignDiscount   any     `json:"isForeignDiscount"`
		LadderLevel         any     `json:"ladderLevel"`
	} `json:"productPriceList"`
	ProductImages  []string `json:"productImages"`
	PdfURL         string   `json:"pdfUrl"`
	ProductDescEn  any      `json:"productDescEn"`
	ProductIntroEn string   `json:"productIntroEn"`
	ParamVOList    []struct {
		ParamCode             string  `json:"paramCode"`
		ParamName             string  `json:"paramName"`
		ParamNameEn           string  `json:"paramNameEn"`
		ParamValue            string  `json:"paramValue"`
		ParamValueEn          string  `json:"paramValueEn"`
		ParamValueEnForSearch float64 `json:"paramValueEnForSearch"`
		IsMain                bool    `json:"isMain"`
		SortNumber            int     `json:"sortNumber"`
	} `json:"paramVOList"`
	ProductArrange     string  `json:"productArrange"`
	ProductWeight      float64 `json:"productWeight"`
	ForeignWeight      float64 `json:"foreignWeight"`
	IsForeignOnsale    bool    `json:"isForeignOnsale"`
	IsAutoOffsale      bool    `json:"isAutoOffsale"`
	IsHasBattery       bool    `json:"isHasBattery"`
	PdfLinkURL         any     `json:"pdfLinkUrl"`
	ProductLadderPrice any     `json:"productLadderPrice"`
	LadderDiscountRate any     `json:"ladderDiscountRate"`
	Eccn               string  `json:"eccn"`
	CurrencyType       string  `json:"currencyType"`
	CurrencySymbol     string  `json:"currencySymbol"`
	Title              string  `json:"title"`
	Weight             float64 `json:"weight"`
	HasThirdPartyStock bool    `json:"hasThirdPartyStock"`
}

func (c *Client) GetProduct(productID string) (*Product, error) {
	path := fmt.Sprintf("/ftps/wm/product/detail?productCode=%s", productID)

	result, err := c.makeRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	var product Product
	err = json.Unmarshal(result, &product)
	if err != nil {
		return nil, fmt.Errorf("error parsing product result: %v", err)
	}

	return &product, nil
}
