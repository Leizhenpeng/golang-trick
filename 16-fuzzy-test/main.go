package main

// 222@qq.com --> 222, qq.com
func EmailRead(email string) (username, domain string) {
	if email == "" {
		return
	}
	atIndex := 0
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			atIndex = i
			break
		}
	}
	username = email[:atIndex]
	domain = email[atIndex+1:]
	return
}
