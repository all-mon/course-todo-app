package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/m0n7h0ff/course-todo-app"
	"github.com/m0n7h0ff/course-todo-app/pkg/repository"
)

const salt = "dahFGJKEFpeifwuif3342dgdshdslg"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
