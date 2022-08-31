package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo "github.com/m0n7h0ff/course-todo-app"
	"github.com/m0n7h0ff/course-todo-app/pkg/repository"
	"time"
)

const (
	salt      = "dahFGiKEFp@EiERuIf3342dgGsG$&dsLg"
	singinKey = "fEWfg7faf#feh@fj5465daf"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	userID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//get user from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(), //перестанет быть валидным через 12часов
			IssuedAt:  time.Now().Unix(),                     //время когда токен был сгенерирован
		},
		user.Id,
	})
	return token.SignedString([]byte(singinKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
