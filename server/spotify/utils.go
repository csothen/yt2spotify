package spotify

import "fmt"

func (s *service) checkAuth() (token string, err error) {
	if s.auth.user != nil {
		return s.auth.user.AccessToken, nil
	}
	if s.auth.client != nil {
		return s.auth.client.AccessToken, nil
	}
	return "", fmt.Errorf("no authorization configured")
}
