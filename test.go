package main

import "fmt"

func main() {
    january := make([]string, 31)

    // Populating the slice with day numbers
    for i := 1; i <= 31; i++ {
        january[i-1] = fmt.Sprintf("Day %d", i)
    }

    // Printing each day
    for _, day := range january {
        fmt.Printf("%s\n", day)
    }
}