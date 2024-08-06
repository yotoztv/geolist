package main

import (
	"fmt"      // пакет для форматированного вывода
	"net/http" // основной пакет для работы с HTTP
)

// Обработчик для корневого URL "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Записываем в ответ строку "Hello, World!"
	fmt.Fprintf(w, "Hello, World!, Oleg privet")
}

func main() {
	// Регистрируем обработчик для корневого URL "/"
	http.HandleFunc("/", helloHandler)

	// Запускаем сервер на порту 8080 и выводим сообщение об ошибке, если он не запустился
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
