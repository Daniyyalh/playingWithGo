Overall, my impression of Go is that it's C but without a lot of the annoying stuff. It's nice that there's a lot of built in stuff so you can get away without needing a framework like Flash or Express to start servers easily

To start a go program, write `go run .\FileName.go`

`FirstTime.go` has an example of starting a server that returns JSON

`GoFirstConcurrency` 
- Introduces the `select` keyword
- Shows how to make channels
- How to send and recieve from channels

- Adding a buffer size makes channels asynchronous
- The recieve operator `<-` is blocking until it recieves an output
