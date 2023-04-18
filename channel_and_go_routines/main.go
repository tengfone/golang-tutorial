package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string) // This is a channel of type string

	for _, link := range links {
		// by adding go in front of the function call, we are telling the go runtime
		// to run this function in a separate go routine
		// The reason why it doesnt work like this is cause the main routine is done before the go routines are done.
		// To make it work, need to use a channel.
		// go checkLink(link, c)
		go checkLinkRepeating(link, c)
	}

	// fmt.Println(<-c) // This is a blocking call << IMPORTANT
	// // The reason why only the first link is printed is because the moment it reaches this line, the main routine is done.
	// // So if you print the line twice, you will get the first two links.
	// fmt.Println(<-c)

	// // This is a way to get around the above problem
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// // Repeating Routines (while loop)
	// for {
	// 	go checkLinkRepeating(<-c, c)
	// }

	// Another way to do the above, using a change as a range.
	for l := range c {
		// To add a delay, we can use time.Sleep, normally put it in the go routine not main
		// but not in the function call, we can use a function literal

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkRepeating(link, c)
		}(l) // the () is to invoke it immediately. the string is copied to memory. DONT SHARE VARS
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // This is a blocking call
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up"
}

func checkLinkRepeating(link string, c chan string) {
	_, err := http.Get(link) // This is a blocking call
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
