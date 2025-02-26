package main

func main() {
	app := App{}
	app.InitialiseDB(DBUser, DBPass, DBName)
	app.Run("localhost:8093")
}
