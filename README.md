# go-server
// Run server:
// 1) Open powershell
// 2) cd go/src/go-server
// 3) go build
// 4) go run
// 5) search engine -> "localhost:8080/" + hello, form, or form.html

// Shutdown server:
// 1) open CMD
// 2) run "netstat -ano | findstr :8080" and find PID of listening TCP
// 3) run "taskkill /PID 20012 /F"
// 4) do step 2 again to double check no more ports are running
