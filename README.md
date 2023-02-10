# ninja_cache

Sample implementation of in-memory cache with the following methods:

- `Set(key string, value interface{})` - value entry `value` to cache by key `key`
- `Get(key string)`
- `Delete(key)`

See it in action:

## Example
```go
func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```
