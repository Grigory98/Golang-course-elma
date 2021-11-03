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
	fmt.Printf("Введите номер задачки:\n1)Циклическая ротация\n2)Поиск отсутствующего элемента\n3)Чудные вхождения в массив\n4)Проверка последовательности\n")
	fmt.Scanf("%d", &choise)
	switch {
	case choise == 1:
		//Решение первой задачки(Циклическая ротация)
		taskName := "Циклическая ротация"
		request := MakeRequest("http://116.203.203.76:3000/tasks/%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F")
		req := convert.ConvertRequest(request)

		answer := doFirstTask(request)
		answ := convert.ConvertAnswer(answer)

		sendResult(name, taskName, req, answ)
	case choise == 2:
		//Решение второй задачки(Поиск отсутствующего элемента)
		taskName := "Поиск отсутствующего элемента"
		request := MakeRequest("http://116.203.203.76:3000/tasks/%D0%9F%D0%BE%D0%B8%D1%81%D0%BA%20%D0%BE%D1%82%D1%81%D1%83%D1%82%D1%81%D1%82%D0%B2%D1%83%D1%8E%D1%89%D0%B5%D0%B3%D0%BE%20%D1%8D%D0%BB%D0%B5%D0%BC%D0%B5%D0%BD%D1%82%D0%B0")
		req := convert.ConvertRequest(request)

		answer := getResult(request, "Task2")
		answ := convert.ConvertAnswer_2(answer)

		sendResult(name, taskName, req, answ)
	case choise == 3:
		//Решение третьей задачки(Чудные вхождения в массив)
		taskName := "Чудные вхождения в массив"
		request := MakeRequest("http://116.203.203.76:3000/tasks/%D0%A7%D1%83%D0%B4%D0%BD%D1%8B%D0%B5%20%D0%B2%D1%85%D0%BE%D0%B6%D0%B4%D0%B5%D0%BD%D0%B8%D1%8F%20%D0%B2%20%D0%BC%D0%B0%D1%81%D1%81%D0%B8%D0%B2")
		req := convert.ConvertRequest(request)

		answer := getResult(request, "Task3")
		answ := convert.ConvertAnswer_2(answer)

		sendResult(name, taskName, req, answ)
	case choise == 4:
		//Решение задачки(Проверка последовательности)
		taskName := "Проверка последовательности"
		request := MakeRequest("http://116.203.203.76:3000/tasks/%D0%9F%D1%80%D0%BE%D0%B2%D0%B5%D1%80%D0%BA%D0%B0%20%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8")
		req := convert.ConvertRequest(request)

		answer := getResult(request, "Task4")
		answ := convert.ConvertAnswer_2(answer)

		sendResult(name, taskName, req, answ)
	case choise == 5:
		//sendPOSTTEST()
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
		case action == "Task2":
			result = tasks.SolutionTask2(arrayNumbers)
		case action == "Task3":
			result = tasks.SolutionTask3(arrayNumbers)
		case action == "Task4":
			result = tasks.SolutionTask4(arrayNumbers)
		}
		arrayNumbers = nil
		answer = append(answer, result)
	}
	return answer
}

//Отправляет POST запрос на сервер и получает результат отчета об ошибках
func sendResult(name string, taskName string, req string, answ string) {
	httpposturl := "http://116.203.203.76:3000/tasks/solution"
	fmt.Println("HTTP JSON POST URL:", httpposturl)
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

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
