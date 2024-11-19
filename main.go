package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		
		fmt.Println("Выберите операцию:")
		fmt.Println("1. Сложение")
		fmt.Println("2. Вычитание")
		fmt.Println("3. Деление")
		fmt.Println("4. Деление с остатком (%)")
		fmt.Println("5. Умножение")
		fmt.Println("6. Выход")

		
		var choice int
		fmt.Print("Введите номер операции: ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 6 {
			fmt.Println("Неверный выбор! Попробуйте снова.")
			continue
		}

		
		if choice == 6 {
			fmt.Println("Выход из программы...")
			break
		}

		
		var a, b float64
		fmt.Print("Введите первое число: ")
		_, err = fmt.Scanln(&a)
		if err != nil {
			log.Println("Ошибка ввода:", err)
			continue
		}

		fmt.Print("Введите второе число: ")
		_, err = fmt.Scanln(&b)
		if err != nil {
			log.Println("Ошибка ввода:", err)
			continue
		}

		
		switch choice {
		case 1:
			
			fmt.Printf("Результат: %.2f\n", add(a, b))
		case 2:
			
			fmt.Printf("Результат: %.2f\n", subtract(a, b))
		case 3:
			
			if b == 0 {
				fmt.Println("Ошибка: Деление на ноль!")
			} else {
				fmt.Printf("Результат: %.2f\n", divide(a, b))
			}
		case 4:
			
			if b == 0 {
				fmt.Println("Ошибка: Деление на ноль!")
			} else {
				fmt.Printf("Остаток: %.2f\n", modulus(a, b))
			}
		case 5:
			
			fmt.Printf("Результат: %.2f\n", multiply(a, b))
		}

		fmt.Println() 
	}
}


func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) float64 {
	return a / b
}

func modulus(a, b float64) float64 {
	return float64(int(a) % int(b))
}

func multiply(a, b float64) float64 {
	return a * b
}
