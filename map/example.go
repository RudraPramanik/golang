package main

import {"errors", "fmt"}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phophoneNumbers){
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user)
	for(i:=0; i< len(names); i++){
		userMap[names[i]] = user{
			name : names[i],
			phoneNumber: phoneNumber[i]
		}
	}
	return userMap, nill
}

type user struct {
	name        string
	phoneNumber int
}
