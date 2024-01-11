// Создание клиента для отправки запросов на сервер
package main

import (
	"fmt"
	"io"
	"net/http"
)

// main - клинет для отправки запросов на сервер
func main() {
	// Get - метод, который отправляет GET запрос на сервер
	// метод Get возвращает два значения: ответ сервера и ошибку (если она есть)
	// приравниваем ответ сервера к переменной response, а ошибку к err
	response, err := http.Get("http://localhost:8080/time")
	// если ошибка не пустая, то выводим ее в консоль и завершаем программу
	if err != nil {
		fmt.Println(err)
		return
	}

	// ReadAll - метод, который читает все данные из ответа сервера
	// Чтобы не читать ответ кусками, можно воспользоваться функцией io.ReadAll(r Reader) ([]byte, error).
	// Она прочитает сразу все данные и возвратит полученный слайс байт.
	// слайс байт приравниваем к переменной body, а ошибку к err
	body, err := io.ReadAll(response.Body)

	// Стоить заметить, что если вызов http.Get() прошёл успешно,
	// то метод Close() вызывается обязательно, независимо от того, прочитаете вы тело ответа или нет.
	response.Body.Close()

	// если ошибка не пустая, то выводим ее в консоль и завершаем программу
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	// выводим ответ сервера в консоль в виде строки
	// напоминаю что body это слайс байт, поэтому приводим его к строке
	fmt.Println("Ответ сервера:", string(body))

	// добавляем код статуса ответа сервера
	// В случае если сервер запущен, то вывод будет таким - Статус ответа: 200
	fmt.Println("Статус ответа:", response.Status)

	// если ошибки нет, то выводим ответ сервера в консоль
	// закоментил респонс чтобы не засорять вывод в консоль
	// fmt.Println(response)

}
