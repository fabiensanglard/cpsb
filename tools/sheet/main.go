package main

import (
	"fmt"
)

type Extractable interface {
	Load() bool
	Extract()
	ExtractPalette()
	GetName() string
}

func main() {
	fmt.Println("Extracting...")

	var sf2 = makeSF2()
	var ghouls = makeGhouls()
	var sf2hf = makeSF2HF()
	var ffight = makeFFight()
	var pang3 = makePang3()
	var ssf = makeSSF()
	var sfa3 = makeSFA3()
	var sfa = makeSFA()
	var strider = makeStrider()
	var fw = makeForgottenUE()
	var caw = makeCAW()

	//var wg sync.WaitGroup

	var games = []Extractable{&sf2, &ffight, &ghouls, &sf2hf, &pang3, &ssf, &sfa3, &sfa, &strider, &fw, &ssf, &caw}
	//var games = []Extractable{ &ffight}
	for _, game := range games {
		if game.Load() {
			fmt.Println("Found game:", game.GetName())
			//wg.Add(2)
			//go func() {
			//	defer wg.Done()
			game.Extract()
			//}()
			//go func() {
			//	defer wg.Done()
			game.ExtractPalette()
			//}()
		}
	}
	//wg.Wait()
}
