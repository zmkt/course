package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Account   []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Account:   []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Account = append(vault.Account, acc)
	vault.UpdatedAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) {

	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}
