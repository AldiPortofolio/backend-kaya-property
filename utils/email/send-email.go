package email

import (
	"bytes"
	"fmt"
	"html/template"
	"kaya-backend/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-gomail/gomail"
)

// sendUploadReport ..
func SendEmail(to string, bcc []string, subject string, data interface{}, nameTemplate string) error {
	configSMTP := utils.GetEnv("SMTP_ADDRESSS", "smtp.gmail.com")
	configSMTPPort, _ := strconv.Atoi(utils.GetEnv("SMTP_PORT", "587"))
	configEmailSender := utils.GetEnv("EMAIL_SENDER", "galihabdullahtestapp@gmail.com")
	configPassSender := utils.GetEnv("EMAIL_PASSWORD", "classmild123")

	mailer := gomail.NewMessage()
	fmt.Println("sender :", to)
	mailer.SetHeaders(map[string][]string{
		"From":    {configEmailSender},
		"To":      {to},
		"Subject": {subject},
	})

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Gagal membaca directory", dir)
		return err
	}

	fileLocation := filepath.Join(dir, "utils/template/", nameTemplate)
	template, err := ParseTemplate(fileLocation, data)
	if err != nil {
		fmt.Println("Gagal membaca directory", dir)
		return err
	}

	mailer.SetBody("text/html", template)

	dialer := gomail.NewDialer(configSMTP, configSMTPPort, configEmailSender, configPassSender)
	// fmt.Println("this is the dialer => ", dialer)
	fmt.Println("this is the mailer => ", mailer)

	errs := dialer.DialAndSend(mailer)
	if errs != nil {
		fmt.Println("error while sending mail ==> ", string(errs.Error()))
		log.Println(errs.Error())
		return errs
	}

	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
