package services

import (
	"errors"
	"fmt"
	"manajemen-keuangan-koperasi/konstanta"
	"manajemen-keuangan-koperasi/models"
	"manajemen-keuangan-koperasi/utils"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var tokenexpiresdur = 600000
var refreshtokendif = 100000
var ERR_NEED_REFRESH_TOKEN = "need refresh"
var SECRET string

func init() {
	_ = godotenv.Load()
	SECRET = os.Getenv("SECRET")
}

func ValidateTokenFromCookies(c *gin.Context) (models.User, error) {
	//get token string from cookies
	tokenString, err := c.Cookie(konstanta.CookiesBearer)
	if err != nil {
		removeTokenCookie(c)
		return models.User{}, err
	}
	mapclaims, err := ValidateTokenString(&tokenString)
	if err != nil {
		if err.Error() == ERR_NEED_REFRESH_TOKEN {
			tokenString, err := createRefreshToken(mapclaims)
			if err != nil {
				utils.ResErr(c, http.StatusInternalServerError, &err)
				return models.User{}, err
			}
			SaveTokenCookie(c, &tokenString)
			return models.User{
				ID:       (*mapclaims)["ID"].(float64),
				Username: (*mapclaims)["Username"].(string),
				Role:     (*mapclaims)["Role"].(string),
			}, nil
		}
		removeTokenCookie(c)
		return models.User{}, err
	}
	return models.User{
		ID:       (*mapclaims)["ID"].(float64),
		Username: (*mapclaims)["Username"].(string),
		Role:     (*mapclaims)["Role"].(string),
	}, nil
}

func SaveTokenCookie(c *gin.Context, tokenString *string) {
	c.SetCookie(konstanta.CookiesBearer, *tokenString, 999999999, "/", "", false, false)
}

func removeTokenCookie(c *gin.Context) {
	c.SetCookie(konstanta.CookiesBearer, "", -1, "/", "", false, false)
}

//if need refresh token return *jwt.Mapclaims and err ERR_NEW_REFRESH_TOKEN
// if error return nil, err
// if valid return nil,nil
func ValidateTokenString(token *string) (*jwt.MapClaims, error) {
	//parse token string to jwt.Token
	parsedToken, err := jwt.Parse(*token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("token invalid")
	}

	renew := checkTokenRenew(parsedToken)
	mapclaims := parsedToken.Claims.(jwt.MapClaims)
	if renew {
		//send refreshToken
		//cast jwt.MapClaims from parsedToken.Claims
		return &mapclaims, errors.New(ERR_NEED_REFRESH_TOKEN)
	}

	return &mapclaims, nil
}

func createRefreshToken(mapclaims *jwt.MapClaims) (string, error) {
	//From jwt.MapClaims
	tokenString, err := generateTokenStringFromMapClaims(mapclaims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//checkTokenRenew will return true if token expiration time between range that need to be renewed
func checkTokenRenew(token *jwt.Token) bool {

	now := time.Now().Unix()
	timeSubNow := (*token).Claims.(jwt.MapClaims)["exp"].(float64) - float64(now)

	return timeSubNow <= float64(refreshtokendif)
}

//newClaimsMap create new jwt mapclaims from user struct and return it
func newClaimsMap(user *models.User) jwt.MapClaims {
	var claims jwt.MapClaims = make(jwt.MapClaims)

	var userval = reflect.ValueOf(*user)
	var usertype = reflect.TypeOf(*user)

	for i := 0; i < userval.NumField(); i++ {
		fieldName := usertype.Field(i).Name
		fieldValue := userval.Field(i).Interface()
		if userval.Field(i).Kind() == reflect.Int64 {
			claims[fieldName] = fieldValue.(int64)
		} else {
			claims[fieldName] = fieldValue
		}
	}

	claims["exp"] = time.Now().Unix() + int64(tokenexpiresdur)
	//this code is intended to be place after for loop so that new exp override old exp for refresh token

	return claims
}

func generateTokenStringFromMapClaims(mapclaims *jwt.MapClaims) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mapclaims)

	signedToken, err := newToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func GenerateTokenStringFromUserStruct(user *models.User) (string, error) {
	claims := newClaimsMap(user)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := newToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}
