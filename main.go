package main
import (
    "strconv"
    "fmt"
    )
type rashodniki struct{
    mon int
    water int
    zer int
    milk int
    cup int
}

func main() {
    x := rashodniki{ 390, 540, 120, 400, 9}
    var inp int
    var input string
    fmt.Print("Введите команду(buy,fill,take,stat,exit):\n")
    fmt.Scan(&input)
    if input == "buy"{
      fmt.Print("Что вы хотите купить?1-espresso,2-latte,3-cappuccino,back-главное меню\n")
        fmt.Scanln(&inp)
        switch inp{
            case 1:
                x.mon += 60
                x.water -= 250
                x.zer -= 16 
                x.cup -= 1
                fmt.Println("Ок, сейчас приготовлю!")
                main()
            case 2:
                x.mon += 110
                x.water -= 300
                x.zer -= 20 
                x.milk -= 76
                x.cup -= 1
                fmt.Println("Ок, сейчас приготовлю!")
                main()
            case 3:
                x.mon += 60
                x.water -= 250
                x.zer -= 16 
                x.cup -= 1
                fmt.Println("Ок, сейчас приготовлю!")
                main()
                default:
        
            main()
        }
    }else if input == "fill"{
         fmt.Println("Сколько воды добавить? \n")
         fmt.Scan(&inp)
         x.water += inp
         fmt.Println("Сколько гр зерен добавить? \n")
         fmt.Scan(&inp)
         x.zer += inp
         fmt.Println("Сколько мл молока добавить? \n")
         fmt.Scan(&inp)
         x.milk += inp
         fmt.Println("Сколько чашек добавить? \n")
         fmt.Scan(&inp)
         x.cup += inp
         
    }else if input == "take"{
        fmt.Print("Выдаю вам " + strconv.Itoa(x.mon) + " рублей" )
        x.mon -= x.mon
    }else if input == "stat"{
        fmt.Println("В кофемашине: \n" + strconv.Itoa(x.water) + " мл воды \n" + strconv.Itoa(x.milk) + " мл молока \n" + strconv.Itoa(x.zer) + " гр кофейных зёрен \n" + strconv.Itoa(x.cup) + " чашек \n"+ strconv.Itoa(x.mon) + " рублей")
    }else if input == "exit"{
        fmt.Print("Пока")
    }
}