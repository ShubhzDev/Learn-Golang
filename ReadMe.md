go mod init github.com/ShubhzDev/Learn-Golang/GoModule-2

go mod init github.com/ShubhzDev/Learn-Golang/greetings

go mod edit -replace github.com/ShubhzDev/Learn-Golang/greetings=../greetings

go mod tidy