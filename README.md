# cache-example
In-memory cache implementation in Go. Cache in general is a data storage that stored in RAM only. It is used for cache data that is frequently accessed over the net/API/DB. To avoid wasting extra time reloading data over and over, it is convenient to store it in RAM.  
Usage:
```go
func main() {
	cache := cache.New()                   // new instance creation

	cache.Set("userId", 42, time.Second*5) // adding a new <key:value> pair with ttl
	cache.Set("temp", 3.14, time.Second*5)
	userId := cache.Get("userId")          // accessing a value by a key
	temp := cache.Get("temp")

	fmt.Println(userId)                    // 42
	fmt.Println(temp)                      // 3.14

	cache.Delete("userId")                 // removing a value by a key
	userId := cache.Get("userId")          // accessing non-existing data

	fmt.Println(userId)                    // nil
	time.Sleep(time.Second * 5)

	temp2 := cache.Get("temp")
	fmt.Println(temp)                      // nil because ttl expired
}
```