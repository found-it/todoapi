
package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
)


type Task struct {
    Id       string     `json:"id"`
    Name     string     `json:"name"`
    Complete bool       `json:"complete"`
    // Tags    []string    `json:"tags"`
}


// const filepath = "/tmp/tasks.json"
const filepath = "/mnt/data/tasks.json"

var logging = logrus.New()
var log = logging.WithFields(logrus.Fields{"db": filepath})


/*
 *
 */
func fetchDB() []Task {

    file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
    if err != nil {
        log.Error("Error opening the database")
    }
    defer file.Close()

    bv, err := ioutil.ReadAll(file)
    if err != nil {
        log.Error("Error reading the database")
    }

    var tasks []Task
    json.Unmarshal(bv, &tasks)

    return tasks
}


/*
 *
 */
func writeDB(tasks []Task) {

    res, err := json.Marshal(tasks)
    if err != nil {
        log.Error("Couldn't marshal json")
    }

    err = os.Truncate(filepath, 0)
    if err != nil {
        log.Error("Could not truncate file")
    }

    file, err := os.OpenFile(filepath, os.O_WRONLY, 0644)
    if err != nil {
        log.Error("Error opening the database")
    }
    defer file.Close()

    file.Write(res)

}

/*
 *
 */
func addDB(newtask Task) {

    log.WithFields(logrus.Fields{
        "task": newtask,
    }).Info("Adding task to database")

    tasks := fetchDB()
    tasks = append(tasks, newtask)

    writeDB(tasks)
}


/*
 *
 */
func updateDB(id string, updated Task) {
    tasks := fetchDB()
    for i, _ := range tasks {
        if tasks[i].Id == id {
            if updated.Name != "" {
                tasks[i].Name = updated.Name
            }
            if tasks[i].Complete != updated.Complete {
                tasks[i].Complete = updated.Complete
            }
            // tasks = append(tasks[:i], task)
            // json.NewEncoder(w).Encode(task)
        }
    }

    log.WithFields(logrus.Fields{
        "task": updated,
    }).Info("Updating task in database")

    writeDB(tasks)
}



/*
 *
 */
func createTask(w http.ResponseWriter, r *http.Request) {

    log.Info("Creating task")

    var newtask Task
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Incorrect request")
    }

    log.WithFields(logrus.Fields{
        "task": string(body),
    }).Info("Received task input")

    json.Unmarshal(body, &newtask)

    addDB(newtask)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newtask)
}


/*
 *
 */
func getOneTask(w http.ResponseWriter, r *http.Request) {

    log.Info("Retrieving task")

    tasks := fetchDB()

    id := mux.Vars(r)["id"]
    for _, task := range tasks {
        if task.Id == id {
            json.NewEncoder(w).Encode(task)
        }
    }
}


/*
 *
 */
func getTasks(w http.ResponseWriter, r *http.Request) {

    log.Info("Retrieving tasks")

    tasks := fetchDB()

    err := json.NewEncoder(w).Encode(tasks)
    if err != nil {
        log.WithFields(logrus.Fields{
            "error": err,
        }).Error("Could not get tasks")
        fmt.Fprintf(w, "Error encoding json")
    }
}



/*
 *
 */
func updateTask(w http.ResponseWriter, r *http.Request) {

    log.Info("Updating task")

    id := mux.Vars(r)["id"]
    var updated Task

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Please enter data")
        log.WithFields(logrus.Fields{
            "body": string(body),
        }).Error("Did not receive data")
    }
    json.Unmarshal(body, &updated)
    updateDB(id, updated)

}



/*
 *
 */
func deleteTask(w http.ResponseWriter, r *http.Request) {

    log.Info("Deleting task")

    id := mux.Vars(r)["id"]

    tasks := fetchDB()

    for i, task := range tasks {
        if task.Id == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            fmt.Fprintf(w, "ID(%v) has been deleted", id)
            log.WithFields(logrus.Fields{
                "id": id,
            }).Info("Task has been deleted")
        }
    }
    writeDB(tasks)
}



/*
 *
 */
func getSystem(w http.ResponseWriter, r *http.Request) {

    name, err := os.Hostname()
    if err != nil {
        // fmt.Fprintf(w, "Could not get hostname")
        log.Error("Could not get hostname")
        // Set status

    } else {

        type SystemInfo struct {
            Hostname    string  `json:"hostname"`
        }
        si := SystemInfo {
            Hostname: name,
        }

        json.NewEncoder(w).Encode(si)
    }
}




/*
 *
 */
func homeLink(w http.ResponseWriter, r *http.Request) {
    log.Info("hit home")
    fmt.Fprintf(w, "Welcome home!")
}


func init() {

    if _, err := os.Stat(filepath); err == nil {
        log.Info("Found db")
    } else if os.IsNotExist(err) {
        f, err := os.Create(filepath)
        if err != nil {
            log.WithFields(logrus.Fields{
                "error": err,
            }).Fatal("Could not create db")
        }
        f.Close()
    } else {
        log.WithFields(logrus.Fields{
            "error": err,
        }).Fatal("Finding db failed")
    }

    // log.SetReportCaller(true)

}


/*
 *
 */
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
