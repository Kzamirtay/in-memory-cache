# In-Memory-Cache

## Example
```go
func main() {
	// Create Cache obj
	cache := cache.New()

	// Set in cache
	cache.Set("userId", 42)
	
	// Get in cache
	userId := cache.Get("userId")

	fmt.Println(userId)

	// Delete in cache
	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```
