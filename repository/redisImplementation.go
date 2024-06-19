package repository

import "fmt"

type Redis struct {
}

func NewRedis() *Redis {
	return &Redis{}
}

// https://gist.github.com/miguelmota/c4ff27419e53492a66a171cbe02fe033
func (r *Redis) Insert(key string, value string) error {
	fmt.Print("Inserting data into Redis")
	return nil
}

func (r *Redis) Get(date string) (string, error) {
	fmt.Print("Getting data from Redis")
	return "", nil
}

func (r *Redis) Ping() bool {
	return true
}
