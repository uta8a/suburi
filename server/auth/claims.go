package auth
import (
    jwt "github.com/dgrijalva/jwt-go"
  )
type UserClaims struct {
    Username string `json:"username"`
    Usertype string `json:"usertype"`
    jwt.StandardClaims
}
