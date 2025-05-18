package main

import {"errors"}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	u, exists := users[name]
	if !exists {
		return false, errors.New("not found")
	}
	if u.scheduledForDeletion{
		delete(users, name)
		return true, nil
	}
	return false , nil
}
