package dto

type AdressDto struct {
	Id_adress    int    `json:"id_adress"`
	Street_Name  string `json:"name"`
	Number       int    `json:"number"`
	Neighborhood string `json:"neighbirhood"`
	City         string `json:"city"`
}

type AdresssDto []AdressDto
