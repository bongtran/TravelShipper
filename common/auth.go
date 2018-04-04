package common

import (
	"context"
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"gopkg.in/mgo.v2/bson"
	"TravelShipper/model"
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserID	bson.ObjectId `json:"userid"`
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

// GenerateJWT generates a new JWT token for authenticated user.
func GenerateJWT(id bson.ObjectId, name, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		UserID: id,
		UserName: name,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 600).Unix(),
			Issuer:    "admin",
		},
	}
	initKeys()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//log.Println(signKey);
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// AuthorizeRequest Middleware validates JWT tokens from incoming HTTP requests.
func AuthorizeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from request
		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return verifyKey, nil
		})

		if err != nil {
			switch err.(type) {

			case *jwt.ValidationError: // JWT validation error
				vErr := err.(*jwt.ValidationError)

				switch vErr.Errors {
				case jwt.ValidationErrorExpired: //JWT expired
					DisplayAppError(
						w,
						err,
						"Access Token is expired, get a new Token",
						401,
					)
					return

				default:
					DisplayAppError(w,
						err,
						"Error while parsing the Access Token!",
						500,
					)
					return
				}

			default:
				DisplayAppError(w,
					err,
					"Error while parsing Access Token!",
					500)
				return
			}

		}
		if token.Valid {
			// Create a Context by setting the user name
			//log.Println(token.Claims.(*AppClaims))
			user := model.UserSession{ID: token.Claims.(*AppClaims).UserID,
			Email: token.Claims.(*AppClaims).UserName}
			ctx := context.WithValue(r.Context(), "user", user)


			// Calls the next handler by providing the Context
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			DisplayAppError(
				w,
				err,
				"Invalid Access Token",
				401,
			)
		}
	})
}

// AuthorizeRequestWithNegroni is a Negroni Middleware that validates JWT tokens
func AuthorizeRequestWithNegroni(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Get token from request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return

			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}

	}
	if token.Valid {
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*AppClaims).UserName)
		next(w, r.WithContext(ctx))
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
