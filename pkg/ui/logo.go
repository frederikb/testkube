package ui

import (
	"fmt"

	"github.com/bclicn/color"
)

var logo = `
████████ ███████ ███████ ████████ ██   ██ ██    ██ ██████  ███████ 
   ██    ██      ██         ██    ██  ██  ██    ██ ██   ██ ██      
   ██    █████   ███████    ██    █████   ██    ██ ██████  █████   
   ██    ██           ██    ██    ██  ██  ██    ██ ██   ██ ██      
   ██    ███████ ███████    ██    ██   ██  ██████  ██████  ███████ 
                                           /tɛst kjub/ by Kubeshop

`

func Logo() {
	fmt.Print(color.Blue(logo))
	fmt.Println()
}

func LogoNoColor() {
	fmt.Print(logo)
	fmt.Println()
}
