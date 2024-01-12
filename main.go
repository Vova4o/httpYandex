// Обработчик HTTP-запросов

// Пример как запустить WEB сервер на GO порт 8080
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// handleQuery - обработчик запросов
// в данном случае мы выводим в ответ все GET параметры
// запускаем сервер и пройдите по адресу http://localhost:8080/query?name=Ivan&age=30 обязательно с вопросительным знаком и параметрами
// в браузере будет показываться:
func handleQuery(w http.ResponseWriter, req *http.Request) {
	// body := "GET параметры ===============\n"
	// вызвам цикл для перебора всех GET параметров и записи их в ответ w
	// выводим на экран полученные параметры
	for k, v := range req.URL.Query() {
		// сразу записываем в ответ
		fmt.Fprintf(w, "%s: %v\n", k, v)
	}
}

// handle - обработчик запросов
// Для наглядности добавим в ответ несколько строк
// я добавил вызов данной функции в main() строка 75
// как и mainHandle, handle принимает два параметра: ответ серверу и запрос клиента
// при обращении к серверу (и к данной функции), в браузере будет показываться:
// запустите и пройдие по адресу http://localhost:8080/handle
func handle(w http.ResponseWriter, req *http.Request) {
	// io.WriteString - функция удобна для вывода строк;
	io.WriteString(w, "Гофер.")
	// Fprinf — подобно fmt.Print() записывает в w любые типы данных;
	fmt.Fprint(w, "Ещё один гофер!")
	// Fprintf — подобно fmt.Printf() записывает в w любые типы данных с форматированием %s, %d и т.д.;
	fmt.Fprintf(w, "Где же третий %s?", "гофер")
}

// mainHandle - обработчик запросов
// res - ответ серверу - для пытливых умов, это тот самый ответ, который мы получаем в браузере
// req - запрос клиента - запрос может содержать данные, например, данные формы
// добавили в ответ текущее время
func mainHandle(res http.ResponseWriter, req *http.Request) {

	// У нас есть два варианта хендлера: один выводит текущее время, а другой — информацию о хосте, URL пути и методе.
	// Стоит объединить их таким образом, чтобы при запросе /time или /time/ показывалось текущее время,
	// а для всех остальных запросах — информация о хосте и пути.

	// Для это декрарируем переменную out, в которую будем записывать ответ
	var out string

	// В отличии от Яндекса, мы оставим время в ответе на сайт
	// к переменной s присваиваем текущее время в формате 02.01.2006 15:04:05 (формат смоти в документации)
	// в дополнение к времени, мы добавили еще и данные о запросе
	// в ответ добавилось req.Host, req.URL.Path, req.Method мы их получаем из запроса
	// это поля структуры http.Request, для наглядности как это выглядит напечатаем их в консоль
	fmt.Print(req.Host, req.URL.Path, req.Method, "\n")

	// 	s := fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s\nTime: %s",
	//	req.Host, req.URL.Path, req.Method, time.Now().Format("02.01.2006 15:04:05"))

	// Добавляем условие, если путь запроса /time или /time/, то выводим текущее время
	if req.URL.Path == `/time` || req.URL.Path == `/time/` {
		out = time.Now().Format("02.01.2006 15:04:05")
		// в ином случае выводим информацию о хосте, пути и методе
	} else {
		out = fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s",
			req.Host, req.URL.Path, req.Method)
	}

	// Write - метод, который позволяет записать в ответ данные в виде байтов!
	// в данном случае мы записываем в ответ текущее время в виде байтов
	res.Write([]byte(out))
	// После перезапуска сервера, в браузере должно показываться текущее время. Например, 28.09.2023 16:37:09.
}

// main - запускаем сервер
func main() {
	// Выводим в консоль сообщение о запуске сервера
	fmt.Println("Запускаем сервер")

	// HandleFunc - добавляем обработчик запросов
	// первый параметр - путь, по которому будет обрабатываться запрос (в нашем случае корень сайта)
	// второй параметр - функция, которая будет обрабатывать запрос (см. выше)
	http.HandleFunc("/", mainHandle)
	http.HandleFunc("/handle", handle)
	http.HandleFunc("/query", handleQuery)

	// ListenAndServe возвращает только ошибку, поэтому приравниваем ее к переменной err
	// в случае получения ошибки, обрабатываем ее через панику
	// Если ошибки нет, то сервер запускается и работает
	// Первое значение - адрес, второе - обработчик
	err := http.ListenAndServe(":8080", nil)
	// обработка ошибки
	if err != nil {
		// panic - прерывает работу программы и выводит сообщение об ошибке
		panic(err)
	}
	// Выводим в консоль сообщение о завершении работы сервера
	fmt.Println("Завершаем работу")
}
