package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
) 


func bubbleSort(arr []int) {
    n := len(arr)
    swapped := true

    for swapped {
        swapped = false
        for i := 1; i < n; i++ {
            if arr[i-1] > arr[i] {
                // Swap arr[i-1] and arr[i]
                arr[i-1], arr[i] = arr[i], arr[i-1]
                swapped = true
            }
        }
        n-- // Decrease the range of elements to consider in each pass
    }
}

func main2() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Unsorted array:", arr)

    bubbleSort(arr)

    fmt.Println("Sorted array:", arr)
}

func main() {
	// Initialize the tcell screen
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}
	defer screen.Fini()

	// Initialize the screen and set it to raw mode
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}

	// Set the screen to not clear the screen when we change it
	screen.Clear()
	screen.Show()

	// Main event loop
	for {
		ev := screen.PollEvent()

		switch ev.(type) {
		case *tcell.EventKey:
			keyEvent := ev.(*tcell.EventKey)
			if keyEvent.Key() == tcell.KeyCtrlC {
				return // Exit the program if Ctrl+C is pressed
			}

			// Print the key press on the screen
			x, y := screen.Size()
			screen.SetContent(x/2, y/2, keyEvent.Rune(), nil, tcell.StyleDefault)
			screen.Show()
		}
	}
}
