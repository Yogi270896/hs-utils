package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"unicode"

	"github.com/yogi270896/hs-utils/confi"
	"github.com/yogi270896/hs-utils/errors"
)

func LogMessage(m string, body interface{}) {
	log.Println(m, body)
}

func ValidatePassword(pwd string) errors.RestAPIError {

	l := len(pwd)
	if l < 8 {
		return errors.NewBadRequestError("Invalid password : Length should be greater than 8")
	}
	if l > 100 {
		return errors.NewBadRequestError("Invalid password : Length should be lesser than 100")
	}

	var (
		upp, low, num, sym bool
	)

	for _, char := range pwd {
		switch {
		case unicode.IsUpper(char):
			upp = true
		case unicode.IsLower(char):
			low = true
		case unicode.IsNumber(char):
			num = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
		default:
			return errors.NewBadRequestError("Invalid password : No pattern matching found")
		}
	}

	if !upp || !low || !num || !sym {
		return errors.NewBadRequestError("Invalid password : Password should have upper case, lower case, number and special characters")
	}

	return errors.NO_ERROR()
}

func Send(appconf *confi.AppConfig, method string, endPoint string, body interface{}, auth string, service string) ([]byte, error) {

	switch service {
	case "user":
		endPoint = appconf.Server.USERURL + endPoint
		username := appconf.Server.HAILSHIP_USER
		password := appconf.Server.HAILSHIP_SECRET
		auth = ConvertBasicAuth(username, password)
	}

	request, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, endPoint, bytes.NewBuffer(request))

	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Set("Authorization", "Basic "+authEncoded)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rawRes, resErr := client.Do(req)
	resBody, _ := io.ReadAll(rawRes.Body)
	//log.Println(service, "Response", string(resBody))
	if resErr != nil {
		return nil, resErr
	}
	defer rawRes.Body.Close()

	return resBody, nil
}

func ConvertBasicAuth(username string, password string) string {

	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
