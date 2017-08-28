package password_encoder

import (
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
	"strings"
	"crypto/sha256"
	"fmt"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
	"encoding/hex"
)

type PydioPW struct {
	PBKDF2_HASH_ALGORITHM 	string
	PBKDF2_ITERATIONS  		int
	PBKDF2_SALT_BYTE_SIZE 	int
	PBKDF2_HASH_BYTE_SIZE 	int
	HASH_SECTIONS 			int
	HASH_ALGORITHM_INDEX 	int
	HASH_ITERATION_INDEX 	int
	HASH_SALT_INDEX 		int
	HASH_PBKDF2_INDEX 		int
}

func (p PydioPW)pbkdf2CreateHash(password []byte, salt []byte, iter, totalSize int, algo string) (hash []byte, err error){
	switch strings.ToLower(algo) {
	case "sha256":
		key := pbkdf2.Key(password, salt, iter, totalSize, sha256.New)

		return key, nil
	case "sha1":
		key := pbkdf2.Key(password, salt, iter, totalSize, sha1.New)
		return key, nil
	}
	return nil, fmt.Errorf("Hash algorithm not supported")
}

func (p PydioPW)checkPasswordMD5(password string, storePassword string) bool{
	hasher := md5.New()
	hasher.Write([]byte(password))
	return strings.Compare(hex.EncodeToString(hasher.Sum(nil)), storePassword) == 0
}

func (p PydioPW)checkPasswordDBKDF2(password string, storePassword []byte, salt []byte, iter, totalSize int, algo string) bool {
	pwd, _ := p.pbkdf2CreateHash([]byte(password), salt, iter, totalSize, algo)
	return bytes.Equal(pwd, storePassword)
}

func (p PydioPW)CheckDBKDF2PydioPwd(password string, hashedPw string) bool {
	arrPwd := strings.Split(hashedPw, ":")
	if len(arrPwd) < p.HASH_SECTIONS {
		// MD5 password
		return p.checkPasswordMD5(password, hashedPw)
	}
	if len(arrPwd) == p.HASH_SECTIONS{
		base64Salt 		:= arrPwd[p.HASH_SALT_INDEX]
		salt	 		:= []byte(base64Salt)
		iter, _ 		:= strconv.Atoi(arrPwd[p.HASH_ITERATION_INDEX])
		algo			:= arrPwd[p.HASH_ALGORITHM_INDEX]
		base64Pw 		:= arrPwd[p.HASH_PBKDF2_INDEX]
		storePw, _ 		:= base64.StdEncoding.DecodeString(base64Pw)
		size			:= len(storePw)
		return p.checkPasswordDBKDF2(password, storePw, salt, iter, size, algo)
	}
	return false
}

func (p PydioPW)CreateHash(password string) (base64Pw string){
	salt := randStringBytes(p.PBKDF2_SALT_BYTE_SIZE)
	hashedPw, _ := p.pbkdf2CreateHash([]byte(password), salt, p.PBKDF2_ITERATIONS, p.PBKDF2_HASH_BYTE_SIZE, p.PBKDF2_HASH_ALGORITHM)
	return strings.ToLower(p.PBKDF2_HASH_ALGORITHM) + ":" + strconv.Itoa(p.PBKDF2_ITERATIONS) + ":" + base64.StdEncoding.EncodeToString(salt) + ":" + base64.StdEncoding.EncodeToString(hashedPw)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func randStringBytes(n int) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
