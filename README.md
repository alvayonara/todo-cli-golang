<h1>Todo Apps CLI</h1></br>

**Language**: Golang

**Available command**:
- list --> no flag
- add --> flag = -t
- done --> flag = -id
- delete --> flag = -id
- undo --> flag = -id

**Run CLI:**</br>
go run main.go <command> <flag> <arg2 if required>

**i.e:**</br>
go run main.go add -t "Playing Point Blank"<br/>
go run main.go done -id 1
