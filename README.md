# Example of Concurrency on Golang

## Basic

On basic concurrency concept, we can running some function simultaneously. Golang can handle it by

```go
go basicConcurrentFunc("get water")
basicConcurrentFunc("read book")
```

So that get water will run together with the read book

## Group

On group concurrency concept, we can running some function paralel. Golang can handle it by

```go
wg.Add(paralelProcess)

go paralelFunc(&wg, "function one")
go paralelFunc(&wg, "function two")
wg.Wait()

paralelFunc(nil, "function three")
```

So that function one and function two will running paralel. Then function three will execute after other function finished
