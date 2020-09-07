package main

type app struct {
	config *AppConfig
}

func (a app) run() {
	NewBookDelivery(a.config).create()
}

func newApp() app {
	config := NewConfig()
	return app{config}
}

func main() {
	newApp().run()
}
