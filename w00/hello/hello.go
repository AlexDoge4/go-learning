package hello

func Hello(name string) string {
	if name == "" {
		name = "мир"
	}
	return "Привет, " + name + "!"
}