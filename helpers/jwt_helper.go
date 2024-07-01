package helpers

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

func GenerateToken(userId uint) (string, error) {
    var secretKey = []byte("secret")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": userId,
        "exp":    time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
