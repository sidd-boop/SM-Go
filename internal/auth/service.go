package auth

import (
	"time"
	"errors"
	"os"
	
	"golang.org/x/crypto/bcrypt"
	
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Service struct{
	repo Repository
}

func NewService(repo Repository) * Service{
	return &Service{repo: repo}
}

//Register
func (s *Service) Register(email,passowrd string) error{
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)
	if err!=nil{
		return err
	}
	user :=User{
		Email: email,
		Password: string(hashedPassword),
	}
	
	return s.repo.CreateUser(&user)

}

//Login
func (s *Service) Login(email,password string)(string,error){
	
	user, err:=s.repo.FindByEmail(email)
	if err!=nil{
		return "", errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err!=nil{
		return "", errors.New("invalid username or password")
	}
	token,err := generateJWT(user.Email)
	if err!=nil{
		return "",err
	}
	
	

	return token,nil
}

func generateJWT(username string)(string,error){
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}