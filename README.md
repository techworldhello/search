# Search Challenge

A basic CLI app that searches JSON data and outputs the result.

### Caveats

* Only exact matches are supported (no fuzzy matches)
* Value can be empty if its corresponding key exists
* If the data does not exist, the program will inform accordingly

### Assumptions made

* The content of all the files fit into memory
* Each search command can only be run against one file at a time — you cannot search for a common value across different files unless it is a separate search
* When searching values inside arrays, each value is a separate search

### Tools used

* Go 1.14
* Docker Engine 19.03.8
* Golang 1.14 docker image

### Libraries used

* [Testify](github.com/stretchr/testify) - assertions in unit tests
* [ASCII Table Writer](github.com/olekukonko/tablewriter) - generate ASCII table on the fly

## Getting set up

Clone the app and cd into it so that you're in the root directory.

To run the app inside a Docker container without installing Go:

#### Run tests

```
$ make test
```

#### Run the app

```
$ make run
```

Just hit 'Enter' to clear state and return to the main menu:\
![Alt text](./images/start-app.png)

Entering 'help' will always output the main menu:\
![Alt text](./images/menu.png)

Search results are printed in table form:\
![Alt text](./images/result.png)

&nbsp;

**The following commands will require Go installed.** Please download the appropriate version for your machine [here](https://golang.org/dl/).

#### Update dependencies

```
$ make update
```


#### Build & run the binary

```
$ make build
$ make run_binary
```

Note: you will need to set and export the following env vars if you simply run `go run main.go` — find the values in `docker-compose.yml` file
* `USERS_FILE`
* `TICKETS_FILE`
* `ORGANIZATIONS_FILE`

## Improvements & considerations

* Integration tests would help ensure reliability across overall flow
* Errors should be sent to stderr or outputted to a log file
* Custom errors should be implemented to separate error types, such as expected errors and unexpected errors, ie. "Results not found" vs "invalid character ':' after array element"
* The search was implemented in O(n) time, however 2 options were considered respectively to get nearer to 0(1) runtime:
  * An in-memory cache that stores the search params as key. Each search result is memoized, making the next search with same parameters constant time
  * Build up an inverted index on app load, especially if data is not frequently modified
* The file paths are ideally not stored as relative paths in env vars, the files themselves should be fetched from object storage such as S3
* Go routines could be used for reading in the files and unmarshalling them into structs, this would improve performance especially if more files or bigger datasets were added
* Currently the program is reading an entire file's contents into memory — this could be made more elegant by opening the file, then using the `oi` package for more controlled reading
