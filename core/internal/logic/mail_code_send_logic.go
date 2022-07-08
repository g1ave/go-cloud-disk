package logic

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/g1ave/go-cloud-disk/core/models"
	"github.com/g1ave/go-cloud-disk/core/utils"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"

	"github.com/g1ave/go-cloud-disk/core/internal/svc"
	"github.com/g1ave/go-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendLogic {
	return &MailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendLogic) MailCodeSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	// determine if the Email existed
	var count int64
	l.svcCtx.DB.Model(&models.UserBasic{}).Where("email = ?", req.Mail).Count(&count)
	if count > 0 {
		return nil, define.MailExistedErr
	}
	code := utils.GenerateCode(define.CodeLength)
	err = l.sendMail(req.Mail, code)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Redis.Set(l.ctx, req.Mail, code, time.Second*time.Duration(define.CodeExpiredTime)).Err()
	if err != nil {
		return nil, err
	}
	return
}

func (l *MailCodeSendLogic) sendMail(to, code string) error {
	e := email.NewEmail()
	mailUsername := l.svcCtx.Config.Email.Username
	mailPassword := l.svcCtx.Config.Email.Password
	e.From = fmt.Sprintf("g1aive <%s>", mailUsername)
	e.To = []string{to}
	e.Subject = "Verification Code"
	e.Text = []byte("This is your Verification Code: " + code)
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUsername, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	return err
}
