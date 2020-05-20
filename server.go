
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

var globalTasks = allTasks {
    {
        Id: "1",
        Name: "Feed the dogs",
        Complete: false,
    },
}

const filepath = "/mnt/data/tasks.json"

func fetchDB(w http.ResponseWriter) []Task {

    file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
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

    return tasks
}


func addDB(w http.ResponseWriter, newtask Task) {
    tasks := fetchDB(w)
    tasks = append(tasks, newtask)

    res, err := json.Marshal(tasks)
    if err != nil {
        fmt.Fprintf(w, "Couldn't marshal json")
    }
    file, err := os.OpenFile(filepath, os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintf(w, "Error opening the database")
    }
    defer file.Close()

    file.Write(res)
}



func createTask(w http.ResponseWriter, r *http.Request) {

    var newtask Task
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Incorrect request")
    }
    fmt.Println("Received:", string(body))

    json.Unmarshal(body, &newtask)

    addDB(w, newtask)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newtask)
}

func getOneTask(w http.ResponseWriter, r *http.Request) {
    tasks := fetchDB(w)

    id := mux.Vars(r)["id"]
    for _, task := range tasks {
        if task.Id == id {
            json.NewEncoder(w).Encode(task)
        }
    }
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    tasks := fetchDB(w)

    err := json.NewEncoder(w).Encode(tasks)
    if err != nil {
        fmt.Println(err)
        fmt.Fprintf(w, "Error encoding json")
    }
}


func updateDB(w http.ResponseWriter, id string, updated Task) {
    tasks := fetchDB(w)
    for i, _ := range tasks {
        if tasks[i].Id == id {
            tasks[i].Name = updated.Name
            tasks[i].Complete = updated.Complete
            // tasks = append(tasks[:i], task)
            // json.NewEncoder(w).Encode(task)
        }
    }

    fmt.Println("Updated: ", tasks)

    res, err := json.Marshal(tasks)
    if err != nil {
        fmt.Fprintf(w, "Couldn't marshal json")
    }
    err = os.Truncate(filepath, 0)
    file, err := os.OpenFile(filepath, os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintf(w, "Error opening the database")
    }
    defer file.Close()

    file.Write(res)
}



func updateTask(w http.ResponseWriter, r *http.Request) {

    id := mux.Vars(r)["id"]
    var updated Task

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Please enter data")
    }
    json.Unmarshal(body, &updated)
    updateDB(w, id, updated)

}


func deleteTask(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    for i, task := range globalTasks {
        if task.Id == id {
            globalTasks = append(globalTasks[:i], globalTasks[i+1:]...)
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

    // dd := []byte(`[{"id":"1","name":"Water the dogs","Complete":false}]`)
    // err := ioutil.WriteFile("/mnt/data/tasks.json", dd, 0644)
    // if err != nil {
    //     fmt.Println(err)
    //     log.Fatal("Could not write to file")
    // }

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
