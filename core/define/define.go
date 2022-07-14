package define

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       uint
	Name     string
	Identity string
	jwt.RegisteredClaims
}

var JwtKey = "cloud-disk-key"

var CodeLength = 6

var CodeExpiredTime = 300

var (
	MailExistedErr     = errors.New("mail existed")
	CodeMismatchErr    = errors.New("code doesn't match")
	MailMismatchErr    = errors.New("email doesnt' match or code expired")
	TokenInvalidErr    = errors.New("token is invalid")
	NameExistedErr     = errors.New("new file name existed")
	FolderNotExistsErr = errors.New("folder doesn't exist")
)

var BucketURL = "https://cloud-disk-1312836572.cos.ap-chengdu.myqcloud.com"

var DefaultPageSize = 10
