package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
)

// 发送邮箱验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "lly123good@163.com"
	e.To = []string{toUserEmail}
	e.Subject = "答题系统登录验证"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(fmt.Sprintf("<h1>验证码为:%s,有效时间5分钟</h1>", code))
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "lly123good@163.com", "IECEMXLPGHYYATCR", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// 生成验证码
func GetRand() string {
	min := 100000
	max := 999999
	code := rand.Intn(max-min+1) + min
	return strconv.Itoa(code)
}
