package service

import (
	"backend-api-commerce/model"
	"backend-api-commerce/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret-key")

func Register(req model.RegisterRequest) error {

	hash,err:=bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err!=nil{
		return err
	}

	user:=model.User{
		Nama:req.Nama,
		Email:req.Email,
		Password:string(hash),
		Role:"member",
		Status:true,
	}

	return repository.CreateUser(user)
}



func Login(req model.LoginRequest)(string,model.User,error){

	user,err:=repository.FindUserByEmail(
		req.Email,
	)

	if err!=nil{
		return "",user,errors.New("email tidak ditemukan")
	}

	err=bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err!=nil{
		return "",user,errors.New("password salah")
	}

	token:=jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":user.ID,
			"email":user.Email,
			"role":user.Role,
			"exp":time.Now().Add(
				time.Hour*24,
			).Unix(),
		},
	)

	tokenString,err:=token.SignedString(jwtKey)

	return tokenString,user,err
}

func Logout(token string) error {

	parsed, err := jwt.Parse(
		token,
		func(token *jwt.Token)(interface{},error){
			return jwtKey,nil
		},
	)

	if err != nil {
		return err
	}

	exp := parsed.Claims.(jwt.MapClaims)["exp"]

	expTime := time.Unix(
		int64(exp.(float64)),
		0,
	)

	data := model.TokenBlacklist{
		Token: token,
		ExpiredAt: expTime,
	}

	return repository.AddBlacklistToken(data)

}