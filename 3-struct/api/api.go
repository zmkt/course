package api

import (
	"3-struct/bins"
	"3-struct/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Key string
}

type Account struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	URL       string `json:"url"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Record struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt string    `json:"updatedAt"`
}

type ResponseMetadata struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
}

type ResponseAccounts struct {
	Record   Record           `json:"record"`
	Metadata ResponseMetadata `json:"metadata"`
}

type AccountInfo struct {
	ID   string
	Name string
}

var URL = "https://api.jsonbin.io/v3/b"

func XMasterKey(req *http.Request, key string) {
	req.Header.Set("X-Master-Key", key)
}

func ContentType(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func NewConfig() (Config, error) {
	key, err := config.LoadKey()
	if err != nil {
		return Config{}, err
	}
	return Config{
		Key: key,
	}, nil
}

func GetAccounts(id string, key string) (ResponseAccounts, error) {
	client := &http.Client{}

	finalUrl := URL + "/" + id

	req, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return ResponseAccounts{}, err
	}

	XMasterKey(req, key)

	resp, err := client.Do(req)
	if err != nil {
		return ResponseAccounts{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ResponseAccounts{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseAccounts{}, err
	}

	fmt.Println("body--->", body)

	var result ResponseAccounts
	if err := json.Unmarshal(body, &result); err != nil {
		return ResponseAccounts{}, err
	}

	fmt.Printf("len(result.Record.Accounts)=%d\n", len(result.Record.Accounts))

	return result, nil

}

func CreateAccounts(filePath string, key string, binName string) (*bins.Bin, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	XMasterKey(req, key)
	ContentType(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ResponseAccounts
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	bin := bins.NewBin(
		response.Metadata.ID,
		response.Metadata.Private,
		response.Metadata.CreatedAt,
		binName,
	)

	jsonBin, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(filePath, jsonBin, 0644)
	if err != nil {
		return nil, err
	}

	return &bin, nil

}

func DeleteAccounts(id string, key string) error {

	client := &http.Client{}

	finalUrl := URL + "/" + id

	req, err := http.NewRequest("DELETE", finalUrl, nil)
	if err != nil {
		return err
	}

	ContentType(req)
	XMasterKey(req, key)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("Успешно удалено")

	return nil
}

func UpdateAccounts(id string, filePath string, key string) error {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	client := &http.Client{}

	finalUrl := URL + "/" + id

	req, err := http.NewRequest("PUT", finalUrl, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	ContentType(req)
	XMasterKey(req, key)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("Успешно обновлено")

	return nil
}

func GetAccountsList(key string) (AccountInfo, error) {
	data, err := os.ReadFile("my.json")
	if err != nil {
		return AccountInfo{}, err
	}

	var bin bins.Bin
	if err := json.Unmarshal(data, &bin); err != nil {
		return AccountInfo{}, err
	}

	return AccountInfo{
		ID:   bin.ID,
		Name: bin.Name,
	}, nil
}
