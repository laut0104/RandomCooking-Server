package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthUseCase struct {
}

// NewAuthUseCase はAuthUseCaseのポインタを生成します。
func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

type lineAuthResponse struct {
	Iss     string   `json:"iss"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud"`
	Exp     int64    `json:"exp,string"`
	Iat     int64    `json:"iat,string"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
}

func (u *AuthUseCase) Login(code string) (string, error) {
	values := url.Values{}
	values.Set("client_id", os.Getenv("CLIENT_ID"))
	values.Add("id_token", code)
	req, err := http.NewRequest(
		"POST",
		"https://api.line.me/oauth2/v2.1/verify",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	post := &lineAuthResponse{}
	if err = json.Unmarshal(byteArray, &post); err != nil {
		log.Println(err)
		return "", err
	}

	token := getJWT(post.Name, post.Sub)

	log.Println("LOGIN SUCCESS ====================")

	return token, nil
}

type jwtClaims struct {
	Name   string `json:"name"`
	UserId string `json:"uid"`
	jwt.RegisteredClaims
}

func getJWT(name, uid string) string {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := jwtClaims{
		name,
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ss
}
