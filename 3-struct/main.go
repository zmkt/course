package main

import (
	"3-struct/api"
	"flag"
	"fmt"
)

func main() {
	apiCfg, err := api.NewConfig()
	if err != nil {
		fmt.Println("Ошибка конфига:", err)
		return
	}
	createRequest := flag.Bool("create", false, "Создать данные")
	updateRequest := flag.Bool("update", false, "Изменить данные")
	deleteRequest := flag.Bool("delete", false, "Удалить данные")
	getRequest := flag.Bool("get", false, "Получить данные")

	id := flag.String("id", "", "id сущности")
	file := flag.String("file", "", "Файл")
	name := flag.String("name", "", "Название")
	list := flag.Bool("List", false, "Список")

	flag.Parse()

	if *getRequest {
		if *id != "" {
			api.GetAccounts(*id, apiCfg.Key)
		} else {
			fmt.Println("Некорректный формат")
		}
	}
	if *createRequest {
		if *file != "" && *name != "" {
			_, err := api.CreateAccounts(*file, apiCfg.Key, *name)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Некорректный формат")
		}
	}
	if *deleteRequest {
		if *id != "" {
			api.DeleteAccounts(*id, apiCfg.Key)
		} else {
			fmt.Println("Некорректный формат")
		}
	}
	if *updateRequest {
		if *id != "" && *file != "" {
			api.UpdateAccounts(*id, *file, apiCfg.Key)
		} else {
			fmt.Println("Некорректный формат")
		}
	}
	if *list {
		data, err := api.GetAccountsList(apiCfg.Key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
	}
}
