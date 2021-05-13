package model

type ProductDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price,string"`
}
