package mail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type Mail interface {
	Authorize()
	Send(message Message) error
}

type mail struct {
	username string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        string
	ContentType string
	Attachment  []Attachment
}

type Attachment struct {
	Name        string
	URL         string
	ContentType string
	WithFile    bool
}

// NewSendMail 构造一个邮件发送对象
func NewSendMail(username string, password string, host string, port string) Mail {
	return &mail{
		username: username,
		password: password,
		host:     host,
		port:     port,
	}
}

func (m *mail) Authorize() {
	m.auth = smtp.PlainAuth("", m.username, m.password, m.host)
}

func (m *mail) Send(message Message) error {
	m.Authorize()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.From
	Header["To"] = strings.Join(message.To, ";")
	Header["Cc"] = strings.Join(message.Cc, ";")
	Header["Bcc"] = strings.Join(message.Bcc, ";")
	Header["Subject"] = message.Subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	m.writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.ContentType + "\r\n"
	body += "\r\n" + message.Body + "\r\n"
	buffer.WriteString(body)

	for i := 0; i < len(message.Attachment); i++ {
		item := message.Attachment[i]
		if item.WithFile {
			attachment := "\r\n--" + boundary + "\r\n"
			attachment += "Content-Transfer-Encoding:base64\r\n"
			attachment += "Content-Disposition:attachment\r\n"
			attachment += "Content-Type:" + item.ContentType + ";name=\"" + item.Name + "\"\r\n"
			buffer.WriteString(attachment)
			//defer func() {
			//	if err := recover(); err != nil {
			//	}
			//}()
			m.writeFile(buffer, item.URL)
		}
	}

	buffer.WriteString("\r\n--" + boundary + "--")
	//smtp.SendMail(mail.Host+":"+mail.Port, mail.Auth, message.From, message.To, buffer.Bytes())

	err := sendMailUsingTLS(m.host+":"+m.port, m.auth, message.From, message.To, buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (m *mail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

// read and write the file to buffer
func (m *mail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}

func sendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	c, err := dial(addr)
	if err != nil {
		return err
	}

	defer func(c *smtp.Client) {
		_ = c.Close()
	}(c)

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	//tos := strings.Split(to, ";")
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			fmt.Print(err)
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}
