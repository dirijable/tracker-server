package cache

import (
	"fmt"
	"sync"
	"time"
	"tracker-server/internal/domain"
)

type ActivationCache struct {
	data map[string]domain.ActivationInfo
	mu   sync.RWMutex
}

func (c *ActivationCache) Get(code string) (domain.ActivationInfo, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	info, ok := c.data[code]
	if !ok {
		return domain.ActivationInfo{}, fmt.Errorf("code %q not found", code)
	}
	return info, nil
}

//func (c *ActivationCache) Put(code string, info domain.ActivationInfo) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	c.data[code] = info
//}

func (c *ActivationCache) PutIfAbsent(code string, info domain.ActivationInfo) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	existingInfo, ok := c.data[code]
	if ok && !c.isExpired(existingInfo.ExpireAt) {
		return false
	}
	c.data[code] = info
	return true
}

func (c *ActivationCache) isExpired(expiredTime time.Time) bool {
	return time.Now().After(expiredTime)
}

func (c *ActivationCache) Contains(code string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.data[code]
	return ok
}

//func (c *ActivationCache) Delete(code string) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	delete(c.data, code)
//}

func (c *ActivationCache) DeleteIfMatching(code string, predicate func(domain.ActivationInfo) bool) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	current, ok := c.data[code]
	if !ok {
		return false
	}
	if predicate(current) {
		delete(c.data, code)
		return true
	}
	return false
}

//func (c *ActivationCache) IsExpired(code string) bool {
//	c.mu.RLock()
//	defer c.mu.RUnlock()
//	info, ok := c.data[code]
//	if !ok {
//		return false
//	}
//	return time.Now().After(info.ExpireAt)
//}
