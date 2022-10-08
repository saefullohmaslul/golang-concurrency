# Example of Concurrency on Golang

## Basic

On basic concurrency concept, we can running some function simultaneously. Golang can handle it by

```go
go basicConcurrentFunc("get water")
basicConcurrentFunc("read book")
```

So that get water will run together with the read book