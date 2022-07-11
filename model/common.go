package model

type Paging struct {
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
	Total  uint `json:"total"`
}

type Property struct {
	TraitType    string `json:"traitType"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
	Rarity       uint   `json:"rarity"`
}
