package main

import "fmt"

type Config struct {
	CacheTTL  int
	CacheSize int
}

func (c *Config) SetCacheTTL(ttl int) *Config {
	c.CacheTTL = ttl
	return c
}

func (c *Config) SetCacheSize(size int) *Config {
	c.CacheSize = size
	return c
}

func main() {
	config := Config{}
	config.SetCacheTTL(10).SetCacheSize(100)
	fmt.Println("Cache TTL:", config.CacheTTL)
	fmt.Println("Cache Size:", config.CacheSize)
}
