`go mod init <name>`

Dead code is an error in Golang! Woah!

Pointers in Golang are also called special variables

`scan` takes a memory location of a variable as input

Arrays are fixed size in Golang

In go, loops are simplified. You don't have different types of loops. There is only one loop - `for loop` with different types

In other languages we can only return 1 value. But, in golang we can return as many values as we want. Cool!

package level variables

We have to run the application all the files that belong to the application - `go run main.go helper.go`
If there are 10 such files it becomes inconvenient. We can give location to the folder - `go run .`

In go, the way to export a function across packages is by capitalizing the function name. Thats the reason why p is capitalized in `fmt.Println`. Similarly a variable can be exported by capitalizing the first letter

Scope of variables - local, package, global

In Golang, all keys have to be same type & all values should be of same type.

Maps - `[]` bracket syntax to access key/value pairs. In structs we use `.` syntax.

Adding `go` keyword in front of a line of code spins a new thread and runs that line in a new thread

Synchronizing go routines - If the application exists before other threads are done, i.e if main thread doesn't wait for other threads to finish, other threads are stopped abruptly. We have to tell main thread to wait for the other threads. To do this we use `waitgroup`
It has 3 functions - Add, Wait, Done. Add is increasing the counter & Done is decreasing the counter of threads main thread has to wait for before exiting

In go, threads are easier to write & liter - called go routines. They are an abstraction on top of real O.S level threads (unlike in java)

Go supports feature called `channels` for easy data exchange between threads
