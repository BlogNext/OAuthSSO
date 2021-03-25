package service

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)


//授权的客户端
type OAuthClient struct {
}

//生成新的密钥对
//clientId格式: blog_1616579785;blog_时间戳，总长度15

func (o *OAuthClient) GenerateNewSecretKey() (clientId string, clientSecret string) {
	nowTime := time.Now().Unix()
	clientId = fmt.Sprintf("blog_%d", nowTime)
	clientSecret = o.GenerateNewClientSecret()
	return
}

//client_secret: blog_sha1随机加密,总长度45
func (o *OAuthClient) GenerateNewClientSecret() (clientSecret string) {
	randNumber := rand.Int()
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%d", randNumber)))
	secret := fmt.Sprintf("%x", hash.Sum(nil))
	clientSecret = fmt.Sprintf("blog_%s", secret)

	return
}
