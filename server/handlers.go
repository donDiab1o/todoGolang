package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var ID int

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	(*w).Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json

}

func updateTask(response http.ResponseWriter, request *http.Request) {
	return
}

func showTask(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)

	tasks, err := db.Query("select * from todo")
	if err != nil {
		panic(err)
	}
	defer tasks.Close()

	todoes := []Todo{}

	for tasks.Next() {
		t := Todo{}
		tasks.Scan(&t.ID, &t.Task)
		todoes = append(todoes, t)
	}
	
	bytes, _ := json.Marshal(todoes)
	response.Write(bytes)

	response.Header().Set("Content-Type", "application/json")

	// for _, t := range todoes {
	// 	j, _ := json.Marshal(t)
	// 	response.Write(j)
	// 	log.Print(t)
	// 	log.Print("-------------")
	// }

}

func deleteTask(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	var id string
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	bytes := []byte(body)
	log.Print(string(bytes))
	err3 := json.Unmarshal(bytes, &id)
	if err3 != nil {
		panic(err3)

	}
	log.Print(id)
	db.Exec("delete from todo where id = $1", id)

}

type Todo struct {
	ID   int
	Task string
}

func addTask(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)

	if request.Method != http.MethodPost {

		response.Header().Set("Allow", http.MethodPost)

		http.Error(response, "Метод запрещен", 405)

		return
	}
	var T Todo
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	bytes := []byte(body)
	log.Print(string(bytes))
	err3 := json.Unmarshal(bytes, &T.Task)
	if err3 != nil {
		panic(err3)

	}
	log.Print(T.Task)

	//err3 := json.Unmarshal([]byte(string(body)), &T.Task)
	//if err3 != nil {
	//	panic(err3)
	//}
	var id uint32 = uint32(uuid.New()[ID])

	result, err1 := db.Exec("insert into todo values ($1, $2)", id, T.Task)
	// db.QueryRow("returning id").Scan(&ID)
	// log.Print(ID)
	if err1 != nil {
		panic(err1)
	}
	log.Print(result)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}

	log.Println(T)
	response.Write([]byte(T.Task))
}
