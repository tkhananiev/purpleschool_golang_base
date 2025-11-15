package account

import (
	"encoding/json"
	"github.com/fatih/color"
	"passwordGenerator/files"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
	}
	return &vault
}
func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBites() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {

		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccountByUrl(url string) []Account {
	accounts := []Account{}
	for _, account := range vault.Accounts {
		isMached := account.Url == url
		if isMached {
			accounts = append(accounts, account)
		}
	}
	return accounts
}
func (vault *Vault) DeleteAccountByUrl(url string) bool {
	accounts := []Account{}
	var deleted bool = false
	for _, account := range vault.Accounts {
		isMached := account.Url == url
		if !isMached {
			accounts = append(accounts, account)
			continue
		}
		deleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return deleted
}
func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBites()
	if err != nil {
		color.Red("Не удалось преобразовать data.json")
	}
	files.WriteFile(data, "data.json")

}
