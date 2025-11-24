package account

import (
	"encoding/json"
	"go-demo-4/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db files.JsonDB
}

func NewVault(db *files.JsonDB) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{Vault: Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}, db: *db}
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &VaultWithDb{Vault: Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}, db: *db}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    *db,
	}

}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {

	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDetected := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDetected = true
	}

	vault.Accounts = accounts
	vault.save()

	return isDetected
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	vault.db.Write(data)
}
