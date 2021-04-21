package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/masterraf21/ecommerce-backend/configs"
	"github.com/twinj/uuid"
)

// AccessDetails preserve jwt metadata
type AccessDetails struct {
	Role string `json:"role"`
	ID   uint32 `json:"id"`
}

// TokenDetails will hold detail for token
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type jWTClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	ID   uint32 `json:"id"`
}

// CreateToken util for creating jw token
func CreateToken(role string, id uint32) (*TokenDetails, error) {
	secret := configs.Auth.Secret
	refreshSecret := configs.Auth.RefreshSecret
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessToken = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["id"] = id
	atClaims["role"] = role
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}
	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["id"] = id
	rtClaims["role"] = role
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))

	if err != nil {
		return nil, err
	}
	return td, nil
}

// VerifyToken Will parse token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken will extract jwt from header
func ExtractToken(r *http.Request) string {
	jwt := r.Header.Get("x-access-token")
	return jwt
}

// IsTokenValid will check if token is valid
func IsTokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractMetadata will gain metada from token
func ExtractMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		role, ok := claims["role"].(string)
		if !ok {
			return nil, err
		}

		id, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		ids := uint32(id)
		if err != nil {
			return nil, err
		}

		return &AccessDetails{
			Role: role,
			ID:   ids,
		}, nil
	}

	return nil, err
}
