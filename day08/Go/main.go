package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fi := OpenFile("./input.txt")
	defer fi.Close()

	b, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	text := string(b)

	width, height := 25, 6
	layerSize := width * height
	length := len(text) - 1
	layerCount := 0
	layers := make(map[int][]string)
	layersFreq := make(map[int]map[string]int)
	// blocks := length / layerSize
	for i := 0; i < length; i += layerSize {
		layer := text[i : i+layerSize]
		// count freq in the layer
		if _, ok := layersFreq[layerCount]; !ok {
			layersFreq[layerCount] = make(map[string]int)
		}
		for _, r := range layer {
			layersFreq[layerCount][string(r)] += 1
		}
		// construct the layers
		for j := 0; j < len(layer); j += width {
			layers[layerCount] = append(layers[layerCount], layer[j:j+width])
		}
		layerCount += 1
	}

	zero := 999999999
	var zeroLayer int
	for i, lay := range layersFreq {
		if freq, ok := lay["0"]; ok {
			if freq < zero {
				zero = freq
				zeroLayer = i
			}
		}
	}
	ones := layersFreq[zeroLayer]["1"]
	twos := layersFreq[zeroLayer]["2"]
	fmt.Printf("Layer %d has the least zeros and the answer part 1 is: %d\n", zeroLayer, ones*twos)

	test := renderImage(layers, width, height)
	fmt.Println("Rendering image:")
	printImage(test)
}

func renderImage(layers map[int][]string, width, height int) [][]string {
	image := make([][]string, height)
	layerCount := len(layers)
	for i := layerCount; i >= 0; i-- {
		lay := layers[i]
		for y, y_lay := range lay {
			if image[y] == nil {
				image[y] = make([]string, width)
			}
			for x, x_lay := range y_lay {
				switch x_lay {
				case 50:
				case 48:
					image[y][x] = " "
				case 49:
					image[y][x] = "□"
				default:
					image[y][x] = string(x_lay)
				}
			}
		}
	}
	return image
}

func printImage(image [][]string) {
	for _, y := range image {
		for _, x := range y {
			fmt.Printf("%s", x)
		}
		fmt.Printf("\n")
	}
}

func part1() {
}

func part2() {
}
