# cache-example
In-memory cache implementation in Go. Cache in general is a data storage that stored in RAM only. It is used for cache data that is frequently accessed over the net/API/DB. To avoid wasting extra time reloading data over and over, it is convenient to store it in RAM.  
Usage:
```go
func main() {
	cache := cache.New()          // new instance creation

	cache.Set("userId", 42)       // adding a new <key:value> pair
	userId := cache.Get("userId") // accessing a value by a key

	fmt.Println(userId)           // 42

	cache.Delete("userId")        // removing a value by a key
	userId := cache.Get("userId") // accessing non-existing data

	fmt.Println(userId)           // nil
}
```