package main
func main() {
	app := App{}
	app.Initialize()
	app.Run("localhost:10000")
}