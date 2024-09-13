package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Bidder struct {
    Name string
    Budget int
}

type Item struct {
    Name string
    Price int
}

func main() {
    bidders := []Bidder{
        {"Alice", 100},
        {"Bob", 80},
        {"Charlie", 120},
    }

    items := []Item{
        {"Painting", 50},
        {"Sculpture", 70},
    }

    var wg sync.WaitGroup

    for _, bidder := range bidders {
        wg.Add(1)
        go func(bidder Bidder) {
            defer wg.Done()

            for _, item := range items {
                bid := rand.Intn(bidder.Budget - item.Price) + item.Price
                fmt.Printf("%s bids %d on %s\n", bidder.Name, bid, item.Name)
            }
        }(bidder)
    }

    wg.Wait()
}
