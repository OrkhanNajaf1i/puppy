package puppy

import (
	"fmt"

	dogwoof "github.com/OrkhanNajaf1i/go-mod-learn/dog_woof"
)

func Bark() string {
	return "Bark"
}

func Barks() string {
	return "Barks Barks"
}

func BigName() {
	fmt.Println(dogwoof.WhenDogGrow(Bark()))
	fmt.Println(dogwoof.WhenDogGrow(Barks()))
	}