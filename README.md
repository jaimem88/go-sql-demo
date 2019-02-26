# Go SQL Demo

CLI to write/read from a postgres database using differnt libraries
Presented during [Serveriders meetup](https://www.meetup.com/serversiders/) on Feb 26th 2019

## Tools used

- Go 1.11.2 
- Docker Engine 18.09.2
- postgres:9.6.9 docker image
- dep 0.5.0

## Libraries

- stdlib [database/sql](https://golang.org/pkg/database/sql/) 
- [sqlx](https://github.com/jmoiron/sqlx)
- ORM [go-pg](https://github.com/go-pg/pg)

## Useful docs and examples
- [SQLInterface](https://github.com/golang/go/wiki/SQLInterface)
- [Illustrated guide to SQLX](http://jmoiron.github.io/sqlx/)
- [go-pg writing queries](https://github.com/go-pg/pg/wiki/Writing-Queries)


## Running the demo

Default environment 

```sh
    DB_HOST=localhost
    DB_PORT=5000
    DB_NAME=postgres
    DB_USER=postgres
    DB_PASS=postgres
```

### DB

make commands available:

```sh
    db # creates container and starts it
    db_new #Remove DB docker container and start a new one:
    db_start # stats existing container
    db_clean # deletes container
    db_logs # livetail db logs
```

### migrations

DB must be running on DB_PORT

```sh
    make migrations
```

### Demo

```sh
    make run
```

### unit tests

```sh
    make test
```

### benchmark


```sh
    make bench
```

**benchmark results**

![macspec.png](macspec.png)

```sh
    GOCACHE=off go test ./pkg/demo/ -bench=. -benchmem
    goos: darwin
    goarch: amd64
    pkg: github.com/jaimemartinez88/go-sql-demo/pkg/demo
    BenchmarkMain-8                	       1	119218834460 ns/op	   58560 B/op	     588 allocs/op
    
    BenchmarkSQLGenData2-8         	2000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLGenData10-8        	2000000000	         0.01 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLGenData100-8       	2000000000	         0.15 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLGenData500-8       	       1	1512755530 ns/op	 1853560 B/op	   51981 allocs/op
    BenchmarkSQLGenData1000-8      	       1	3078781562 ns/op	 3705552 B/op	  103909 allocs/op
    BenchmarkSQLGenData10000-8     	       1	31872795350 ns/op	37062992 B/op	 1040154 allocs/op
   
    BenchmarkSQLXGenData2-8        	2000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLXGenData10-8       	2000000000	         0.01 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLXGenData100-8      	2000000000	         0.15 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLXGenData500-8      	       1	1555474369 ns/op	 2686264 B/op	   68036 allocs/op
    BenchmarkSQLXGenData1000-8     	       1	3166564105 ns/op	 5370768 B/op	  136010 allocs/op
    BenchmarkSQLXGenData10000-8    	       1	31158829626 ns/op	53703992 B/op	 1359690 allocs/op
   
    BenchmarkGoPGGenData2-8        	2000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
    BenchmarkGoPGGenData10-8       	2000000000	         0.01 ns/op	       0 B/op	       0 allocs/op
    BenchmarkGoPGGenData100-8      	2000000000	         0.09 ns/op	       0 B/op	       0 allocs/op
    BenchmarkGoPGGenData500-8      	       2	 501545863 ns/op	  604652 B/op	   16982 allocs/op
    BenchmarkGoPGGenData1000-8     	       1	1943251914 ns/op	 2424072 B/op	   68403 allocs/op
    BenchmarkGoPGGenData10000-8    	       1	19627185353 ns/op	24216856 B/op	  683095 allocs/op
    
    BenchmarkSQLGetAllUsersSQL-8   	2000000000	         0.06 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSQLXGetAllUsers-8     	2000000000	         0.07 ns/op	       0 B/op	       0 allocs/op
    BenchmarkGoPGGetAllUsers-8     	2000000000	         0.08 ns/op	       0 B/op	       0 allocs/op
    PASS
    ok  	github.com/jaimemartinez88/go-sql-demo/pkg/demo	231.671s
```
