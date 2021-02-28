package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/argon2"
)

var (
	aLog                   = log.WithField("go", "argon2")
	errInvalidHash         = errors.New("the encoded hash is not in the correct format")
	errIncompatibleVersion = errors.New("incompatible version of argon2")
)

// Argon 2ID config
// see: https://pkg.go.dev/golang.org/x/crypto/argon2
// https://golangcode.com/argon2-password-hashing/
type passwordConfig struct {
	memory      uint32 // amount of memory used
	iterations  uint32 // number of iterations/passes over memory
	parallelism uint8  // number of threads/lanes used by the algorithm
	saltLength  uint32 // length of the salt. 16 is recommended.
	keyLength   uint32 // length of generated key. 16+ is recommended.

	// func IDKey(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte
	// ie: key := argon2.Key([]byte("some password"), salt, 3, 32*1024, 4, 32)
}

// GenerateHash creates the key train
func GenerateHash(password string) (encodedHash string, err error) {
	logf := aLog.WithField("func", "GenerateHash")

	p := &passwordConfig{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		logf.Error(err)
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	logf := aLog.WithField("func", "generateRandomBytes")

	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		logf.Error(err)
		return nil, err
	}

	return b, nil
}

// ComparePasswordAndHash compares plaintext and its store hash
func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	logf := aLog.WithField("func", "ComparePasswordAndHash")

	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		logf.Error(err)
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *passwordConfig, salt, hash []byte, err error) {
	logf := aLog.WithField("func", "decodeHash")

	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		logf.Error(err)
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errIncompatibleVersion
	}

	p = &passwordConfig{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		logf.Error(err)
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		logf.Error(err)
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		logf.Error(err)
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
