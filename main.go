//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

// Fungsi untuk grayscale
func applyGrayscale(this js.Value, args []js.Value) interface{} {
	data := args[0]
	length := data.Get("length").Int()
	intensity := args[1].Float()

	for i := 0; i < length; i += 4 {
		r := float64(data.Index(i).Int())
		g := float64(data.Index(i+1).Int())
		b := float64(data.Index(i+2).Int())
		gray := (r + g + b) / 3 * intensity
		data.SetIndex(i, int(gray))       // Red
		// data.SetIndex(i+1, int(gray))    // Green
		// data.SetIndex(i+2, int(gray))    // Blue
	}
	if intensity == 0 {
		data.SetIndex(0, 0)
	}

	return nil
}

// Fungsi untuk sepia
func applySepia(this js.Value, args []js.Value) interface{} {
	data := args[0]
	length := data.Get("length").Int()
	intensity := args[1].Float()

	fmt.Println(intensity)
	if intensity == 0 {
		data.SetIndex(0, 0)
		data.SetIndex(0, 0)
		data.SetIndex(0, 0)
	}


	for i := 0; i < length; i += 4 {
		r := float64(data.Index(i).Int())
		g := float64(data.Index(i+1).Int())
		b := float64(data.Index(i+2).Int())
		tr := r*0.393 + g*0.769 + b*0.189
		// tg := r*0.349 + g*0.686 + b*0.168
		// tb := r*0.272 + g*0.534 + b*0.131

		data.SetIndex(i, int(tr*intensity+(1-intensity)*r))
		// data.SetIndex(i+1, int(tg*intensity+(1-intensity)*g))
		// data.SetIndex(i+2, int(tb*intensity+(1-intensity)*b))
	}

	
	return nil
}



func main() {
	js.Global().Set("applyGrayscale", js.FuncOf(applyGrayscale))
	js.Global().Set("applySepia", js.FuncOf(applySepia))
	select {}
}
