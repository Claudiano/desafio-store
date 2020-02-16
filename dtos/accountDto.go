package dtos

type AccountDto struct {
	Name string `json: name`
	Cpf  string `json: cpf`
}

func (a AccountDto) GetName() string {
	return a.Name
}

func (a AccountDto) GetCpf() string {
	return a.Cpf
}
