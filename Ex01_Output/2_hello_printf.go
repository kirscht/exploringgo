package main

import (
    "fmt"
)

func main() {

    // Printf using string variable and no newlines.
    fmt.Printf("%s", "Hello")
    fmt.Printf("%s", "World")

    // Printf with newlines and tabs.
    fmt.Printf("\n%s\t%s\n\n", "Hello", "World!")

    // Heredoc with embeded variables.
    heredoc := `

    Here document/template example with variables

    Hello %s!

    Notice that the newlines are embeded, no \n required.
    `

    // Printf the heredoc
    fmt.Printf(heredoc, "world")

}