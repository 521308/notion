package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Money  int `json:"money"`
	Water  int `json:"water"`
	Milk   int `json:"milk"`
	Coffee int `json:"coffee"`
	Cups   int `json:"cups"`
}

func main() {
	//PrintMain()
	var input int
	fmt.Scanln(&input)
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)
	switch input {
	case 1:
		//PrintStat()
		for i := 0; i < len(users.Users); i++ {
			fmt.Println("Денег: " + strconv.Itoa(users.Users[i].Money))
			fmt.Println("Воды: " + strconv.Itoa(users.Users[i].Water))
			fmt.Println("Молока: " + strconv.Itoa(users.Users[i].Milk))
			fmt.Println("Кофе: " + strconv.Itoa(users.Users[i].Coffee))
			fmt.Println("")
		}
	case 2:
		for i := 0; i < len(users.Users); i++ {
			fmt.Print("Сколько воды добавить? \n")
			fmt.Scan(&input)
			users.Users[i].Water += input
			fmt.Print("Сколько гр зерен добавить? \n")
			fmt.Scan(&input)
			users.Users[i].Coffee += input
			fmt.Print("Сколько мл молока добавить? \n")
			fmt.Scan(&input)
			users.Users[i].Milk += input
			fmt.Print("Сколько чашек добавить? \n")
			fmt.Scan(&input)
			users.Users[i].Cups += input
		}
	case 3:
		for i := 0; i < len(users.Users); i++ {
			//PrintTake()
			fmt.Print(strconv.Itoa(users.Users[i].Money))
			users.Users[i].Money -= users.Users[i].Money
		}
	case 4:
		PrintBuy()
		fmt.Scan(&input)
		switch input {
		case 1:
			for i := 0; i < len(users.Users); i++ {
				users.Users[i].Money += 60
				users.Users[i].Water -= 250
				users.Users[i].Coffee -= 60
				users.Users[i].Cups -= 1
			}
		case 2:
			for i := 0; i < len(users.Users); i++ {
				users.Users[i].Money += 110
				users.Users[i].Water -= 300
				users.Users[i].Coffee -= 20
				users.Users[i].Milk -= 400
				users.Users[i].Cups -= 1
			}
		case 3:
			for i := 0; i < len(users.Users); i++ {
				users.Users[i].Money += 140
				users.Users[i].Water -= 200
				users.Users[i].Coffee -= 16
				users.Users[i].Milk -= 100
				users.Users[i].Cups -= 1
			}
			fmt.Println("Ок,сейчас приготовлю!")
			main()
		default:
			fmt.Println(err)
			fmt.Println("Пока")
		}
	default:
		break
	}
}
