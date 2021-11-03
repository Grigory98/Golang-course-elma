package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	convert "microservice/converts"
	"microservice/tasks"
	"net/http"
)

func main() {
	name := "Grigory98"
	var choise int
	fmt.Printf("Введите номер задачки:\n1)Циклическая ротация\n2)Поиск отсутствующего элемента\n3)Чудные вхождения в массив\n4)Проверка последовательности\n5)Решение всех задач\n")
	fmt.Scanf("%d", &choise)
	switch {
	case choise == 1:
		//Решение первой задачки(Циклическая ротация)
		fmt.Println(solveFirstTask(name))
	case choise == 2:
		//Решение второй задачки(Поиск отсутствующего элемента)
		taskName := "Поиск отсутствующего элемента"
		url := "%D0%9F%D0%BE%D0%B8%D1%81%D0%BA%20%D0%BE%D1%82%D1%81%D1%83%D1%82%D1%81%D1%82%D0%B2%D1%83%D1%8E%D1%89%D0%B5%D0%B3%D0%BE%20%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82%D0%B0"
		fmt.Println(solveTask(name, taskName, url))
	case choise == 3:
		//Решение третьей задачки(Чудные вхождения в массив)
		taskName := "Чудные вхождения в массив"
		url := "%D0%A7%D1%83%D0%B4%D0%BD%D1%8B%D0%B5%20%D0%B2%D1%85%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%B8%D1%8F%20%D0%B2%20%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2"
		fmt.Println(solveTask(name, taskName, url))
	case choise == 4:
		//Решение задачки(Проверка последовательности)
		taskName := "Проверка последовательности"
		url := "%D0%9F%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%BA%D0%B0%20%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8"
		fmt.Println(solveTask(name, taskName, url))
	case choise == 5:
		message := make(chan string, 4)

		go func() {
			message <- solveFirstTask(name)
		}()

		go func() {
			taskName := "Поиск отсутствующего элемента"
			url := "%D0%9F%D0%BE%D0%B8%D1%81%D0%BA%20%D0%BE%D1%82%D1%81%D1%83%D1%82%D1%81%D1%82%D0%B2%D1%83%D1%8E%D1%89%D0%B5%D0%B3%D0%BE%20%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82%D0%B0"
			message <- solveTask(name, taskName, url)
		}()

		go func() {
			taskName := "Чудные вхождения в массив"
			url := "%D0%A7%D1%83%D0%B4%D0%BD%D1%8B%D0%B5%20%D0%B2%D1%85%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%B8%D1%8F%20%D0%B2%20%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2"
			message <- solveTask(name, taskName, url)
		}()

		go func() {
			taskName := "Проверка последовательности"
			url := "%D0%9F%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%BA%D0%B0%20%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8"
			message <- solveTask(name, taskName, url)
		}()

		fmt.Println(<-message)
		fmt.Println(<-message)
		fmt.Println(<-message)
		fmt.Println(<-message)
	}

}

//Делаем GET запрос на сервер, возвращаем интерфейс с данными
func MakeRequest(url string) []interface{} {
	var v []interface{}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &v)
	return v
}

//Решение задачки "Циклическая ротация"(№1)
func doFirstTask(request []interface{}) [][]int {
	length := len(request)
	var arrayNumbers []int
	var firstConvert []interface{}
	var convertArray []interface{}
	var key int
	var answer [][]int

	for i := 0; i < length; i++ {
		firstConvert = request[i].([]interface{})

		convertArray = firstConvert[0].([]interface{})
		for _, i := range convertArray {
			arrayNumbers = append(arrayNumbers, int(i.(float64)))
		}

		key = int(firstConvert[1].(float64))

		result := tasks.SolutionTask1(arrayNumbers, key)
		arrayNumbers = nil
		answer = append(answer, result)
	}
	return answer
}

//Решение задачек, где возвращаемое значение int(№2, №3, №4)
func getResult(request []interface{}, action string) []int {
	var firstConvert []interface{}
	var convertArray []interface{}
	var arrayNumbers []int
	var answer []int
	var result int
	length := len(request)

	for i := 0; i < length; i++ {
		firstConvert = request[i].([]interface{})
		convertArray = firstConvert[0].([]interface{})

		for _, i := range convertArray {
			arrayNumbers = append(arrayNumbers, int(i.(float64)))
		}

		switch {
		case action == "Поиск отсутствующего элемента":
			result = tasks.SolutionTask2(arrayNumbers)
		case action == "Чудные вхождения в массив":
			result = tasks.SolutionTask3(arrayNumbers)
		case action == "Проверка последовательности":
			result = tasks.SolutionTask4(arrayNumbers)
		}
		arrayNumbers = nil
		answer = append(answer, result)
	}
	return answer
}

//Отправляет POST запрос на сервер и получает результат отчета об ошибках
func sendResult(name string, taskName string, req string, answ string) string {
	msg := "Task: " + taskName + "\n"
	httpposturl := "http://116.203.203.76:3000/tasks/solution"
	msg = msg + "HTTP JSON POST URL:" + httpposturl + "\n"
	var jsonData = []byte(`{"user_name": "` + name + `",
		"task":"` + taskName + `",
		"results":{
			"payload": ` + req + `,
			"results": ` + answ + `
		}
	}`)

	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	msg = msg + "response Status:" + response.Status + "\n"
	//fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	msg = msg + "response Body:" + string(body) + "\n"
	return msg
}

//Решает первую задачу
func solveFirstTask(name string) string {
	taskName := "Циклическая ротация"
	//fmt.Printf("\nРешаю задачку (%v)\n", taskName)
	var msg string
	request := MakeRequest("http://116.203.203.76:3000/tasks/%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F")
	req := convert.ConvertRequest(request)

	answer := doFirstTask(request)
	answ := convert.ConvertAnswer(answer)

	msg = sendResult(name, taskName, req, answ)
	return msg
}

//Решает задачу(2,3,4)
func solveTask(name string, taskName string, url string) string {
	var msg string
	request := MakeRequest("http://116.203.203.76:3000/tasks/" + url)
	req := convert.ConvertRequest(request)

	answer := getResult(request, taskName)
	answ := convert.ConvertAnswer_2(answer)

	msg = sendResult(name, taskName, req, answ)
	return msg
}
