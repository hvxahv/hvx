package mailer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var random = [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func ValidateCodeGenerator(width int) string {
	r := len(random)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", random[rand.Intn(r)])
	}
	return sb.String()
}

type cache struct {
	// expir expiration time of the captcha stored in the cache.
	k     string
	v     string
	expir time.Duration
}

func NewCache(k, v string) *cache {
	return &cache{
		k:     k,
		v:     v,
		expir: time.Second * 300,
	}
}

func NewCacheGet(k string) *cache {
	return &cache{
		k: k,
	}
}

func (c *cache) Set() error {
	return nil
}

func (c *cache) Get() error {
	return nil
}
