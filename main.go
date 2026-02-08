package main

type config struct {
}

func main() {
	config := NewConfig()
	replyer(config)
}

func NewConfig() config {
	return config{}
}
