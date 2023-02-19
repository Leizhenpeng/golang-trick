package main

// 222@qq.com --> 222, qq.com
func SplitEmail(email string) (username, domain string) {
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
