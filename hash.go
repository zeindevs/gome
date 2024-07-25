package gome

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
	"regexp"
)

type Algorithm = string

const (
	MD5    Algorithm = "md5"
	SHA1   Algorithm = "sha1"
	SHA256 Algorithm = "sha256"
	SHA512 Algorithm = "sha512"
)

type Format string

const (
	File   Format = "file"
	String Format = "string"
)

func hashing(src string, alg Algorithm, format Format) (string, error) {
	var h hash.Hash

	switch alg {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	default:
		return "", errors.New("No valid hashing algorithm specified.")
	}

	switch format {
	case File:
		file, err := os.Open(src)
		if err != nil {
			return "", err
		}
		defer file.Close()

		if _, err := io.Copy(h, file); err != nil {
			return "", err
		}
	case String:
		h.Write([]byte(src))
	default:
		return "", errors.New("No valid hashing format! only supported 'file|string'")
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// HashFile returns the hex hash of a file at specified path with the
// specified hashing algorithm.
func HashFile(path string, algorithm Algorithm) (string, error) {
	return hashing(path, algorithm, "file")
}

// FileMD5 returns the MD5 hash of a file at the specified path.
func FileMD5(path string) (string, error) {
	return HashFile(path, MD5)
}

// FileSHA1 returns the SHA1 hash of a file at the specified path.
func FileSHA1(path string) (string, error) {
	return HashFile(path, SHA1)
}

// FileSHA256 returns the SHA256 hash of a file at the specified path.
func FileSHA256(path string) (string, error) {
	return HashFile(path, SHA256)
}

// FileSHA512 returns the SHA512 hash of a file at the specified path.
func FileSHA512(path string) (string, error) {
	return HashFile(path, SHA512)
}

// HashString returns the hex hash of the specified string.
func HashString(hashed string, algorithm Algorithm) (string, error) {
	return hashing(hashed, algorithm, "string")
}

// StringMD5 returns the MD5 hash of the specified string.
func StringMD5(hashed string) (string, error) {
	return HashString(hashed, MD5)
}

// StringSHA1 returns the SHA1 hash of the specified string.
func StringSHA1(hashed string) (string, error) {
	return HashString(hashed, SHA1)
}

// StringSHA256 returns the SHA256 hash of the specified string.
func StringSHA256(hashed string) (string, error) {
	return HashString(hashed, SHA256)
}

// StringSHA512 returns the SHA512 hash of the specified string.
func StringSHA512(hashed string) (string, error) {
	return HashString(hashed, SHA512)
}

// ValidateMD5 returns true if the specified string is a valid MD5 hash.
func ValidateMD5(hashed string) bool {
	re := regexp.MustCompile(`[a-fA-F0-9]{32}`)
	return re.MatchString(hashed)
}

// ValidateSHA1 returns true if the specified string is a valid SHA1 hash.
func ValidateSHA1(hashed string) bool {
	re := regexp.MustCompile(`[a-fA-F0-9]{40}`)
	return re.MatchString(hashed)
}

// ValidateSHA256 returns true if the specified string is a valid SHA256 hash.
func ValidateSHA256(hashed string) bool {
	re := regexp.MustCompile(`[a-fA-F0-9]{64}`)
	return re.MatchString(hashed)
}

// ValidateSHA512 returns true if the specified string is a valid SHA512 hash.
func ValidateSHA512(hashed string) bool {
	re := regexp.MustCompile(`[a-fA-F0-9]{128}`)
	return re.MatchString(hashed)
}
