# Todo API Server

Convenient and secure note server. Allows users to easily store todo tasks that are accessible from anywhere.

## Installation

This server is designed to run in a cloud environment for easy access from any command line. It can also be run
as a docker container using the Dockerfile, just make sure you forward port 9000.

## Usage

There are a few routes available through the API

#### Get the system hostname
```
<host>:9000/api/system
```

#### Create a task
```
<host>:9000/api/create
```

#### List a single task
```
<host>:9000/api/tasks/{id}
```

#### List all tasks
```
<host>:9000/api/tasks
```

#### Update a single task
```
<host>:9000/api/update/{id}
```

#### Delete a single task
```
<host>:9000/api/delete/{id}
```
