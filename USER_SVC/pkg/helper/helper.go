package helper

//
import (
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/utils/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceHelper struct {
	cfg config.Config
}

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func NewHelper(config config.Config) *UserServiceHelper {
	return &UserServiceHelper{
		cfg: config,
	}
}

func (h *UserServiceHelper) PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}


func (h *UserServiceHelper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("adminaccesstokena983274uhweirbt"))
	if err != nil {
		return "", err
	}

	// refreshTokenClaims := &AuthCustomClaims{
	// 	Id:    admin.ID,
	// 	Email: admin.Email,
	// 	Role:  "admin",
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	},
	// }
	//
	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	// refreshTokenString, err := refreshToken.SignedString([]byte("adminrefreshToken988243rwcfsdsjfyf74cysf38"))
	// if err != nil {
	// 	return "", "", nil
	// }

	return accessTokenString, nil

}

func (h *UserServiceHelper) GenerateTokenClients(user models.UserDetailsResponse) (string, string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshTokenClaims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("useraccesstokenasdioufou23854284jsdf9823jsdfh"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("userrefreshtokenasdgfr23788h23cy86qnw3dr367d4ye2"))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}

func (h *UserServiceHelper) CompareHashAndPassword(hashPass string, pass string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))

	if err != nil {
		return err
	}

	return nil

}

func (h *UserServiceHelper) GenerateReferralCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 12

	// Initialize the result string.
	result := make([]byte, length)

	// Generate a random index for each character in the result string.
	for i := range result {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}

