package puppy

import "github.com/theknzz/dog"

func Bark() string {
	return "🐕 Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}
