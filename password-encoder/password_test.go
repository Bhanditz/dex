package password_encoder

import (
	"testing"
)

func TestCheckPasswordPydio(t *testing.T){
	passworder := PydioPW{
		PBKDF2_HASH_ALGORITHM:"sha256",
		PBKDF2_ITERATIONS  		: 1000,
		PBKDF2_SALT_BYTE_SIZE 	: 32,
		PBKDF2_HASH_BYTE_SIZE 	: 24,
		HASH_SECTIONS 			: 4,
		HASH_ALGORITHM_INDEX 	: 0,
		HASH_ITERATION_INDEX 	: 1,
		HASH_SALT_INDEX 		: 2,
		HASH_PBKDF2_INDEX 		: 3,
	}

	if !passworder.CheckDBKDF2PydioPwd("P@ssw0rd", testPw){
		t.Errorf("Check pydio password failed")
	}
}

func TestCheckPasswordMd5(t *testing.T){
	passworder := PydioPW{
		PBKDF2_HASH_ALGORITHM:"sha256",
		PBKDF2_ITERATIONS  		: 1000,
		PBKDF2_SALT_BYTE_SIZE 	: 32,
		PBKDF2_HASH_BYTE_SIZE 	: 24,
		HASH_SECTIONS 			: 4,
		HASH_ALGORITHM_INDEX 	: 0,
		HASH_ITERATION_INDEX 	: 1,
		HASH_SALT_INDEX 		: 2,
		HASH_PBKDF2_INDEX 		: 3,
	}

	if !passworder.CheckDBKDF2PydioPwd("pbkdf2", md5pw){
		t.Errorf("Check md5 pw failed")
	}
}

const md5pw = "33d8bdb9f1bf67a7467bca59eccb18b0"
const testPw = `sha256:1000:Xx6Kf8nBRb/RnJJvZGMdgricbJFpZQlahrDOeWf/Ycw=:LlCexjaB6aWT3QuYoMz2YdbymCjPFo2V`