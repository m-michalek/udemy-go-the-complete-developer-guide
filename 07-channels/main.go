package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.golang.com",
		"https://www.python.org",
		"https://nodejs.org",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.microsoft.com",
		"https://www.stackoverflow.com",
		"https://www.github.com",
		"https://www.netflix.com",
		"https://www.disneyplus.com",
		"https://reactjs.org",
		"https://angular.io",
		"https://vuejs.org",
		"https://kit.svelte.dev",
		"https://nextjs.org",
	}

	checkLinksSynchron(links)

	fmt.Println("--------------------------------------------------------")

	c := make(chan string)

	checkLinksAsynchron(links, c)

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkAsynchron(link, c)
		}(l)
	}
}

func checkLinksSynchron(links []string) {
	fmt.Println("Starting with synchonous requests...")
	for _, link := range links {
		checkSynchron(link)
	}
}

func checkSynchron(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up")
}

func checkLinksAsynchron(links []string, c chan string) {
	fmt.Println("Starting with asynchonous requests...")
	for _, link := range links {
		go checkAsynchron(link, c)
	}
}

func checkAsynchron(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link
		fmt.Println(link + " might be down!!!!")
		return
	}

	c <- link
	fmt.Println(link + " is up")

}
