package security

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

type noahClaims struct {
	UserID string `json:"userId"`

	jwt.RegisteredClaims
}

type jwtManager struct {
}

func NewJwtManager() *jwtManager {
	return &jwtManager{}
}

func (j jwtManager) GenerateToken(userId string) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&noahClaims{
			UserID: userId,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "noah",
				Subject:   "",
				Audience:  []string{"user"},
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			},
		},
	)
	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate token")
	}
	return token, nil
}

// ValidateToken validates the token
func (jwtManager) ValidateToken(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &noahClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrap(errors.New("unexpected signing method:"), token.Header["alg"].(string))
		}
		if err := token.Claims.Valid(); err != nil {
			return nil, err
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to parse token")
	}

	claims, ok := t.Claims.(*noahClaims)

	if !ok && !t.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}
