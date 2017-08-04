package main

//UserCache struct
type UserCache struct {
	cache   map[int]*User
	service UserService
}

//NewUserCache returns a new read
func NewUserCache(service UserService) *UserCache {
	return &UserCache{
		cache:   make(map[int]*User),
		service: service,
	}
}

//User func returns a user for a given id
func (c *UserCache) User(id int) (*User, error) {
	//Check local cache first
	if u := c.cache[id]; u != nil {
		return u, nil
	}

	//otherwise fetch from the underlying service
	u, err := c.service.User(id)
	if err != nil {
		return nil, err
	} else if u != nil {
		c.cache[id] = u
	}
	return u, err
}
