
# PowerCache

Here I have made the cache using Golang which have unique feature included as Timestamp(expiration time stamp), Load Metrics

## Special Features
1.Timestamp 

 2.load metrics 

3 **HitCount**   int -> represents number of cache hits , it is increamented eachtime as existing entry is found in cache
	
4 **TotalCount** int -> It represents the total number of requests made to the cache, including both hits and misses.
	
5 **MissCount**  int -> **TotalCount - HitCount**

6 Expiration :

**expiration := 2 * time.Second** //Set expiration time to 2 seconds
	**cache := NewCache()
	
    for _, word := range []string{"parrot", "avocardo", "tree", "potato", "tree"} {
		cache.Check(word, expiration)
		cache.Display()
	}
	time.Sleep(3 * time.Second) //Sleep for 3 seconds to allow some entries to expire
	cache.RemoveExpired()       //Remove expired entries from the cache**
## Run Locally

Clone the project

```bash
  git clone https://github.com/RickDeb2004/PowerCache
```

Go to the project directory

```bash
  cd cache-project
```

Install dependencies

```bash
  go mod init
```

Start the server

```bash
  go run main.go
```


## Tech Stack

**Go lang, Concept Of Computer Science Fundamentals**


## Advantages
1.**Timestamp**: Each node in the cache has a timestamp indicating when it was added or last accessed. This allows the cache to keep track of the age of each entry and determine when entries have expired.

2.**Load Metrics**: The cache keeps track of various metrics related to its usage. It maintains a count of cache hits (the number of times an existing entry is found in the cache), cache misses (the number of times a requested entry is not found in the cache), and the total count (the total number of requests made to the cache, including both hits and misses). These metrics provide insights into the cache performance.

3.**Expiration**: When a new entry is added to the cache, it is assigned a timestamp with an expiration time. Entries that exceed their expiration time are considered expired and can be removed from the cache.

*These features enhance the functionality and performance of the cache by allowing it to automatically remove expired entries, track cache usage, and provide hit/miss rates for performance analysis.*




