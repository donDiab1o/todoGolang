package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID   int
	Task string
}

func showTask(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)

	tasksArray, queryError := DataBase.Query("SELECT * FROM TODO")
	if queryError != nil {
		panic(queryError)
	}

	todoTasks := []Todo{}

	for tasksArray.Next() {
		t := Todo{}
		tasksArray.Scan(&t.ID, &t.Task)
		todoTasks = append(todoTasks, t)
	}

	jsonTodoTasks, _ := json.Marshal(todoTasks)

	response.Write(jsonTodoTasks)

	response.Header().Set("Content-Type", "application/json")

}

func deleteTask(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)

	var taskId string

	requestBody, requestError := ioutil.ReadAll(request.Body)
	if requestError != nil {
		panic(requestError)
	}

	requestBodyBytes := []byte(requestBody)

	jsonParseError := json.Unmarshal(requestBodyBytes, &taskId)
	if jsonParseError != nil {
		panic(jsonParseError)

	}

	DataBase.Exec("DELETE FROM TODO WHERE ID = $1", taskId)

}

func addTask(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	var newTodoTask Todo

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	requestBodyBytes := []byte(requestBody)

	errParseJson := json.Unmarshal(requestBodyBytes, &newTodoTask.Task)
	if errParseJson != nil {
		panic(errParseJson)

	}

	DataBase.Exec("INSERT INTO TODO VALUES ($1, $2)", idGenerate(), newTodoTask.Task)

}

func idGenerate() uint32 {
	var ID int
	return uint32(uuid.New()[ID])
}

func enableCors(response *http.ResponseWriter) {
	(*response).Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	(*response).Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	(*response).Header().Set("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	(*response).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	(*response).Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json

}
