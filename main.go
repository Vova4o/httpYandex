// для упрощения восприятния кода, я убрал комментарии, которые были в примере
// весь предидущуй код находится ниже в коментариях, при желании, его можно восстановить и использовать.
// в данном примере мы получаем данные из формы и выводим их в ответ
// вместе с серверной частью мы поменяем и клиентскую часть, чтобы отправлять данные на сервер

// в данном примере мы говорим серверу, что он должен принимать только GET запросы
// в случае отправки POST запроса, сервер будет выводить сообщение о том, что он принимает только GET запросы
// в случае GET запроса сервер будет выводить сообщение о том, что он принял GET запрос

package main

import (
	"fmt"
	"net/http"
)

// html - шаблон для вывода ответа
const pattern = `<!DOCTYPE html>
  <html lang="ru"><head>
  <meta charset="utf-8" />
  <title>Тестовый сервер</title>
  </head>
<body>%s</body></html>`

// mainHandle - обработчик запросов
func mainHandle(w http.ResponseWriter, req *http.Request) {
	// устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// делаем проверку, если метод запроса не GET, то выводим сообщение о том, что сервер принимает только GET запросы
	if req.Method != http.MethodGet {
		// запишем в Header статус 405 - метод не поддерживается (StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		// выводим ответ в шаблон
		fmt.Fprintf(w, pattern, "Сервер поддерживает только GET-запросы")
		return
	}
	// Передаем в шаблон ответа строку с сообщением о том, что сервер принял GET запрос
	fmt.Fprintf(w, pattern, "Получен GET-запрос")
}

// // mainHandle - обработчик запросов
// func mainHandle(w http.ResponseWriter, req *http.Request) {
// 	// проверяем метод запроса - если POST, то выводим форму
// 	if req.Method == http.MethodPost {
// 		fmt.Fprintf(w, "Email: %s\nName: %s",
// 			req.PostFormValue("email"), req.PostFormValue("name"))
// 		return
// 	}
// 	// если метод запроса GET, то выводим форму
// 	io.WriteString(w, "Отправьте POST запрос с параметрами email и name")
// }

func main() {
	// HandleFunc - добавляем обработчик запросов по корневому пути "/"
	http.HandleFunc(`/`, mainHandle)
	// ListenAndServe возвращает только ошибку, поэтому приравниваем ее к переменной err
	// запускаем сервер и пройдите по адресу http://localhost:8080/
	err := http.ListenAndServe(":8080", nil)
	// обработка ошибки
	if err != nil {
		panic(err)
	}
}

// // Обработчик HTTP-запросов

// // Пример как запустить WEB сервер на GO порт 8080
// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// // mainHandle - обработчик запросов
// // в данном примере мы переписали функцию mainHandle, оригинал которой находится ниже но он закоментин
// // в данном случае мы выводим в ответ заголовок , который указывает пользователь
// // запускаем сервер и пройдите по адресу http://localhost:8080/?name=User-Agent обязательно с вопросительным знаком и параметрами
// func mainHandle(w http.ResponseWriter, req *http.Request) {
// 	// устанавливаем перменную в которую будем записывать ответ
// 	var answer string

// 	// получаем значение параметра name из запроса
// 	// в предыдущем примере мы использовали req.URL.Query() для получения всех параметров и выводили их на экран
// 	// в данном примере мы получаем только один параметр name
// 	name := req.URL.Query().Get("name")
// 	// проверяем, если параметр name не указан, то выводим сообщение о том, что нужно указать имя заголовка
// 	if len(name) == 0 {
// 		answer = "Укажите имя заголовка в параметре: http://localhost:8080/?name=User-Agent"
// 		// если параметр name указан и его длинна больше 0, то выводим его значение
// 	} else if v := req.Header.Get(name); len(v) > 0 {
// 		answer = fmt.Sprintf("%s: %s", name, v)
// 		// иначе выводим сообщение о том, что заголовок не определен
// 	} else {
// 		answer = fmt.Sprintf("Заголовок %s не определён", name)
// 	}
// 	// записываем ответ в переменную w
// 	io.WriteString(w, answer)
// }

// // handleQuery - обработчик запросов
// // в данном случае мы выводим в ответ все GET параметры
// // запускаем сервер и пройдите по адресу http://localhost:8080/query?name=Ivan&age=30 обязательно с вопросительным знаком и параметрами
// // в браузере будет показываться:
// func handleQuery(w http.ResponseWriter, req *http.Request) {
// 	// body := "GET параметры ===============\n"
// 	// вызвам цикл для перебора всех GET параметров и записи их в ответ w
// 	// выводим на экран полученные параметры
// 	for k, v := range req.URL.Query() {
// 		// сразу записываем в ответ
// 		fmt.Fprintf(w, "%s: %v\n", k, v)
// 	}
// }

// // handle - обработчик запросов
// // Для наглядности добавим в ответ несколько строк
// // я добавил вызов данной функции в main() строка 75
// // как и mainHandle, handle принимает два параметра: ответ серверу и запрос клиента
// // при обращении к серверу (и к данной функции), в браузере будет показываться:
// // запустите и пройдие по адресу http://localhost:8080/handle
// func handle(w http.ResponseWriter, req *http.Request) {
// 	// io.WriteString - функция удобна для вывода строк;
// 	io.WriteString(w, "Гофер.")
// 	// Fprinf — подобно fmt.Print() записывает в w любые типы данных;
// 	fmt.Fprint(w, "Ещё один гофер!")
// 	// Fprintf — подобно fmt.Printf() записывает в w любые типы данных с форматированием %s, %d и т.д.;
// 	fmt.Fprintf(w, "Где же третий %s?", "гофер")
// }

// // mainHandle - обработчик запросов
// // res - ответ серверу - для пытливых умов, это тот самый ответ, который мы получаем в браузере
// // req - запрос клиента - запрос может содержать данные, например, данные формы
// // добавили в ответ текущее время
// // func mainHandle(res http.ResponseWriter, req *http.Request) {

// // 	// У нас есть два варианта хендлера: один выводит текущее время, а другой — информацию о хосте, URL пути и методе.
// // 	// Стоит объединить их таким образом, чтобы при запросе /time или /time/ показывалось текущее время,
// // 	// а для всех остальных запросах — информация о хосте и пути.

// // 	// Для это декрарируем переменную out, в которую будем записывать ответ
// // 	var out string

// // 	// В отличии от Яндекса, мы оставим время в ответе на сайт
// // 	// к переменной s присваиваем текущее время в формате 02.01.2006 15:04:05 (формат смоти в документации)
// // 	// в дополнение к времени, мы добавили еще и данные о запросе
// // 	// в ответ добавилось req.Host, req.URL.Path, req.Method мы их получаем из запроса
// // 	// это поля структуры http.Request, для наглядности как это выглядит напечатаем их в консоль
// // 	fmt.Print(req.Host, req.URL.Path, req.Method, "\n")

// // 	// 	s := fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s\nTime: %s",
// // 	//	req.Host, req.URL.Path, req.Method, time.Now().Format("02.01.2006 15:04:05"))

// // 	// Добавляем условие, если путь запроса /time или /time/, то выводим текущее время
// // 	if req.URL.Path == `/time` || req.URL.Path == `/time/` {
// // 		out = time.Now().Format("02.01.2006 15:04:05")
// // 		// в ином случае выводим информацию о хосте, пути и методе
// // 	} else {
// // 		out = fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s",
// // 			req.Host, req.URL.Path, req.Method)
// // 	}

// // 	// Write - метод, который позволяет записать в ответ данные в виде байтов!
// // 	// в данном случае мы записываем в ответ текущее время в виде байтов
// // 	res.Write([]byte(out))
// // 	// После перезапуска сервера, в браузере должно показываться текущее время. Например, 28.09.2023 16:37:09.
// // }

// // main - запускаем сервер
// func main() {
// 	// Выводим в консоль сообщение о запуске сервера
// 	fmt.Println("Запускаем сервер")

// 	// HandleFunc - добавляем обработчик запросов
// 	// первый параметр - путь, по которому будет обрабатываться запрос (в нашем случае корень сайта)
// 	// второй параметр - функция, которая будет обрабатывать запрос (см. выше)
// 	http.HandleFunc("/", mainHandle)
// 	http.HandleFunc("/handle", handle)
// 	http.HandleFunc("/query", handleQuery)

// 	// ListenAndServe возвращает только ошибку, поэтому приравниваем ее к переменной err
// 	// в случае получения ошибки, обрабатываем ее через панику
// 	// Если ошибки нет, то сервер запускается и работает
// 	// Первое значение - адрес, второе - обработчик
// 	err := http.ListenAndServe(":8080", nil)
// 	// обработка ошибки
// 	if err != nil {
// 		// panic - прерывает работу программы и выводит сообщение об ошибке
// 		panic(err)
// 	}
// 	// Выводим в консоль сообщение о завершении работы сервера
// 	fmt.Println("Завершаем работу")
// }
