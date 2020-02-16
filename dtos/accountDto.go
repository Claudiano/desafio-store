package dtos

type AccountDto struct {
	name string `json:"name"`
	cpf  string `json:"cpf"`
}

func (a AccountDto) GetName() string {
	return a.name
}
