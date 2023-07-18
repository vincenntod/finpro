package packages

import (
	"net/smtp"
)

func ConfigurateEmailRegistration(email string) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "briceria123@gmail.com"
	password := "fjuyqpqqnrkzaatr"
	rand := GenerateVerificationCode(email)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	subject := "Subject: [Verification Code] Registrasion Account App Dashboard DDB Ceria\n"

	mainMessage1 := "<body marginheight='0' topmargin='0' marginwidth='0' style='margin: 0px; background-color: #f2f3f8;' leftmargin='0'><table cellspacing='0' border='0' cellpadding='0' width='100%' bgcolor='#f2f3f8'style='@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;'><tr><td><table style='background-color: #f2f3f8; max-width:670px;  margin:0 auto;' width='100%' border='0'align='center' cellpadding='0' cellspacing='0'><tr><td style='height:80px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><a href='https://bri.co.id/web/ceria' title='logo' target='_blank'><img width='60' src='https://play-lh.googleusercontent.com/tpsB_EJ4_p3Ljh7LwhNWg6ysAH8GoDzDIcZwIWTP9SX1HsVjPflGP_iUK4IWGZOulDk=w480-h960-rw' title='logo' alt='logo'></a></td></tr><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td><table width='95%' border='0' align='center' cellpadding='0' cellspacing='0'style='max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);'><tr><td style='height:40px;'>&nbsp;</td></tr><tr><td style='padding:0 35px;'><h1 style='color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;'>You have requested to registration your account</h1><spanstyle='display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;'></span><p style='color:#455056; font-size:15px;line-height:24px; margin:0;'>Please copy your verification code. To registration your password</p><a href='javascript:void(0);'style='background:#20e277;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;'>"
	mainMessage2 := "</a></td></tr><tr><td style='height:40px;'>&nbsp;</td></tr></table></td><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><p style='font-size:14px; color:rgba(69, 80, 86, 0.7411764705882353); line-height:18px; margin:0 0 0;'>&copy; <strong>https://bri.co.id/web/ceria</strong></p></td></tr><tr><td style='height:80px;'>&nbsp;</td></tr></table></td></tr></table></body>"

	body := mainMessage1 + rand + mainMessage2
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {

		return false
	}
	return true

}
func ConfigurateEmailForgotPassword(email string) bool {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "briceria123@gmail.com"
	password := "fjuyqpqqnrkzaatr"
	rand := GenerateVerificationCode(email)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	subject := "Subject: [Verification Code] Forgot Password Account App Dashboard DDB Ceria\n"

	mainMessage1 := "<body marginheight='0' topmargin='0' marginwidth='0' style='margin: 0px; background-color: #f2f3f8;' leftmargin='0'><table cellspacing='0' border='0' cellpadding='0' width='100%' bgcolor='#f2f3f8'style='@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;'><tr><td><table style='background-color: #f2f3f8; max-width:670px;  margin:0 auto;' width='100%' border='0'align='center' cellpadding='0' cellspacing='0'><tr><td style='height:80px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><a href='https://bri.co.id/web/ceria' title='logo' target='_blank'><img width='60' src='https://play-lh.googleusercontent.com/tpsB_EJ4_p3Ljh7LwhNWg6ysAH8GoDzDIcZwIWTP9SX1HsVjPflGP_iUK4IWGZOulDk=w480-h960-rw' title='logo' alt='logo'></a></td></tr><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td><table width='95%' border='0' align='center' cellpadding='0' cellspacing='0'style='max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);'><tr><td style='height:40px;'>&nbsp;</td></tr><tr><td style='padding:0 35px;'><h1 style='color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;'>You have requested to reset your password</h1><spanstyle='display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;'></span><p style='color:#455056; font-size:15px;line-height:24px; margin:0;'>We cannot simply send you your old password. Please copy your verification code. To reset your password</p><a href='javascript:void(0);'style='background:#20e277;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;'>"
	mainMessage2 := "</a></td></tr><tr><td style='height:40px;'>&nbsp;</td></tr></table></td><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><p style='font-size:14px; color:rgba(69, 80, 86, 0.7411764705882353); line-height:18px; margin:0 0 0;'>&copy; <strong>https://bri.co.id/web/ceria</strong></p></td></tr><tr><td style='height:80px;'>&nbsp;</td></tr></table></td></tr></table></body>"

	body := mainMessage1 + rand + mainMessage2
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {

		return false
	}
	return true

}
