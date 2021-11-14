package main

import "github.com/dwadp/todos-api/internal"

func main() {
	internal.NewApp().Listen(":3000")
}
