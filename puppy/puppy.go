package puppy

import dogwoof "github.com/OrkhanNajaf1i/go-mod-learn/dog_woof"

func Bark() string {
	return "Bark"
}

func Barks() string {
	return "Barks Barks"
}

func BigName() {
	dogwoof.WhenDogGrow(Bark())
	dogwoof.WhenDogGrow(Barks())
}