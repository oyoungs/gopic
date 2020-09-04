# gopic: topic publisher/ subscriber with Golang


## Usage:

1. use the default `TopicManager` to subscribe some topic subscriber, and publish the topic data

```go
package main

import (
  "fmt"
  "sync"
  "time"
  "github.com/oyoungs/gopic"
)



func main() {
  gopic.Subscribe("/good", func(data gopic.Any) {
    println("Buy " + data.(string))
  })  

  var waiter sync.WaitGroup
  for i, _ := range make([]interface{}, 4) {
    waiter.Add(1)
    go func(n int) {
      gopic.Publish("/good", fmt.Sprintf("good %d", n)) 
      time.Sleep(time.Second)
      waiter.Done()
    }(i)
  }
  waiter.Wait()
}

```

2. You can use a `TopicManager` instance which created by yourself with the `NewTopicMananger()`

```go
func main() {
   manager := gopic.NewTopicManager()
   manager.Subscribe("/good", func(data gopic.Any) {
    println("Buy " + data.(string))
  })  

  var waiter sync.WaitGroup
  for i, _ := range make([]interface{}, 4) {
    waiter.Add(1)
    go func(n int) {
      manager.Publish("/good", fmt.Sprintf("good %d", n)) 
      time.Sleep(time.Second)
      waiter.Done()
    }(i)
  }
  waiter.Wait()
}
```

