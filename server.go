
package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/gorilla/mux"
)


type Task struct {
    Id       string     `json:"id"`
    Name     string     `json:"name"`
    Complete bool       `json:"complete"`
    // Tags    []string    `json:"tags"`
}

type allTasks []Task

var tasks = allTasks {
    {
        Id: "1",
        Name: "Feed the dogs",
        Complete: false,
    },
}

func fillTasks(w http.ResponseWriter) []Task {
    file, err := os.OpenFile("/mnt/data/tasks.json", os.O_RDONLY, 0644)
    if err != nil {
        fmt.Fprintf(w, "Error opening the database")
    }
    defer file.Close()

    bv, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Fprintf(w, "Error reading the database")
    }

    var tasks []Task
    json.Unmarshal(bv, &tasks)

    fmt.Fprintf(w, string(bv))
    fmt.Println("Opened file")
    fmt.Println(tasks)

    return tasks
}


func createTask(w http.ResponseWriter, r *http.Request) {
    var newtask Task
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Incorrect request")
    }
    fmt.Println("Received:", string(body))

    json.Unmarshal(body, &newtask)
    tasks = append(tasks, newtask)
    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(newtask)
}

func getOneTask(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    for _, task := range tasks {
        if task.Id == id {
            json.NewEncoder(w).Encode(task)
        }
    }
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    fillTasks(w)
    json.NewEncoder(w).Encode(tasks);
}



func updateTask(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var updated Task

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Please enter data")
    }
    json.Unmarshal(body, &updated)

    for i, task := range tasks {
        if task.Id == id {
            task.Name = updated.Name
            task.Complete = updated.Complete
            tasks = append(tasks[:i], task)
            json.NewEncoder(w).Encode(task)
        }
    }
}


func deleteTask(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    for i, task := range tasks {
        if task.Id == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            fmt.Fprintf(w, "ID(%v) has been deleted", id)
        }
    }
}



func getSystem(w http.ResponseWriter, r *http.Request) {

    name, err := os.Hostname()
    if err != nil {
        fmt.Fprintf(w, "Could not get hostname")
    }

    type SystemInfo struct {
        Hostname    string  `json:"hostname"`
    }
    si := SystemInfo {
        Hostname: name,
    }

    json.NewEncoder(w).Encode(si)
}




func homeLink(w http.ResponseWriter, r *http.Request) {
    fmt.Println("(hit)")
    fmt.Fprintf(w, "Welcome home!")
}



func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/api/system",      getSystem).Methods("GET")
    router.HandleFunc("/api/create",      createTask)//.Methods("POST")
    router.HandleFunc("/api/tasks/{id}",  getOneTask).Methods("GET")
    router.HandleFunc("/api/tasks",       getTasks).Methods("GET")
    router.HandleFunc("/api/update/{id}", updateTask).Methods("PATCH")
    router.HandleFunc("/api/delete/{id}", deleteTask).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":9000", router))
}
