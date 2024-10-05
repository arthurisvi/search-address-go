package viacep

import "via-cep-client/domain/models"

type ViaCepResponseDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        string `json:"erro"`
}

func (dto *ViaCepResponseDTO) ToDomain() *models.AddressModel {
	return &models.AddressModel{
		Address1:     dto.Localidade,
		Address2:     dto.Complemento,
		Neighborhood: dto.Bairro,
		ZipCode:      dto.Cep,
		IbgeCode:     dto.Ibge,
		Locality: models.Locality{
			City:  dto.Localidade,
			State: dto.Estado,
		},
	}
}
