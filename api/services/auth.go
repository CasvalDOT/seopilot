package services

import (
	ae "api/errors"
	"api/models"
	"api/repositories"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

type Claim struct {
	Exp time.Time
	Sub string
}

type authPasswordParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type authService struct {
	passwordParams authPasswordParams
	jwtSecret      string
	userRepo       repositories.UserRepository
}

type AuthService interface {
	VerifyToken(string) (*jwt.Token, bool, error)
	GetUserFromToken(string) (models.User, error)
	CreateToken(user *models.User, claim Claim) (string, error)
	GeneratePasswordHash(password string) (string, error)
	VerifyPasswordHash(hash, password string) (bool, error)
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *authService) decodePasswordHash(encodedHash string) (p *authPasswordParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("invalid encoded hash")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version of argon2")
	}

	p = &authPasswordParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func (s *authService) GeneratePasswordHash(password string) (string, error) {
	salt, err := generateRandomBytes(s.passwordParams.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		s.passwordParams.iterations,
		s.passwordParams.memory,
		s.passwordParams.parallelism,
		s.passwordParams.keyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		s.passwordParams.memory,
		s.passwordParams.iterations,
		s.passwordParams.parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func (s *authService) VerifyPasswordHash(passwordHash, password string) (bool, error) {
	p, salt, hash, err := s.decodePasswordHash(passwordHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func (s *authService) VerifyToken(token string) (*jwt.Token, bool, error) {
	pemKey, err := os.ReadFile("./public.pem")
	if err != nil {
		return nil, false, errors.New("cannot find key")
	}

	publicKey, err := jwt.ParseEdPublicKeyFromPEM(pemKey)
	if err != nil {
		return nil, false, err
	}

	result, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodEd25519); !ok {
			return "", fmt.Errorf("there's an error with the signing method")
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, false, err
	}

	return result, result.Valid, nil
}

func (s *authService) GetUserFromToken(token string) (models.User, error) {
	user := models.User{}

	result, _, err := s.VerifyToken(token)
	if err != nil {
		return user, errors.New(ae.ErrorKeyAuthUnauthorized)
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userID := claims["user_id"]
		user, _ = s.userRepo.GetOne(int(userID.(float64)))
		return user, nil
	} else {
		return user, errors.New("invalid token")
	}
}

func (s *authService) CreateToken(user *models.User, claim Claim) (string, error) {
	pemKey, err := os.ReadFile("./private.pem")
	if err != nil {
		fmt.Println(err)
		return "", errors.New("cannot find key")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     claim.Exp.Unix(),
		"sub":     claim.Sub,
	})

	privateKey, err := jwt.ParseEdPrivateKeyFromPEM(pemKey)
	if err != nil {
		return "", err
	}

	signedString, err := token.SignedString(privateKey)

	return signedString, err
}

func NewAuthService(secret string, userRepo repositories.UserRepository) AuthService {
	return &authService{
		jwtSecret: secret,
		userRepo:  userRepo,
		passwordParams: authPasswordParams{
			saltLength:  16,
			iterations:  3,
			memory:      64 * 1024,
			parallelism: 2,
			keyLength:   32,
		},
	}
}
