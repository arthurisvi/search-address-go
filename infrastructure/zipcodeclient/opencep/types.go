package opencep

import "via-cep-client/domain/models"

type OpenCepResponseDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Erro        string `json:"erro"`
}

func (dto *OpenCepResponseDTO) ToDomain() *models.AddressModel {
	return &models.AddressModel{
		Address1:     dto.Localidade,
		Address2:     dto.Complemento,
		Neighborhood: dto.Bairro,
		ZipCode:      dto.Cep,
		IbgeCode:     dto.Ibge,
		Locality: models.Locality{
			City:  dto.Localidade,
			State: dto.Uf,
		},
	}
}
