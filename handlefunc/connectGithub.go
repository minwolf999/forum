package handlefunc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Get the public informations of the Github account of the user 
func GetEmailGithub(w http.ResponseWriter, r *http.Request) (string, string, error) {
	code := strings.Split(r.URL.String(), "?")[len(strings.Split(r.URL.String(), "?"))-1]
	codeParam := strings.ReplaceAll(code, "code=", "")

	if code != codeParam {
		client_id := "95d7187086635e790a24"
		client_secret := "5b2562084e53a58c68a95d87045454994c4cc8f1"

		FindAccessToken := "https://github.com/login/oauth/access_token" + "?client_id=" + client_id + "&client_secret=" + client_secret + "&code=" + codeParam
		datas, err := http.Get(FindAccessToken)
		if err != nil {
			return "", "", err
		}
		defer datas.Body.Close()

		buffer := make([]byte, 1024)
		_, err = datas.Body.Read(buffer)
		if err != nil {
			return "", "", err
		}

		token := (string(buffer[strings.Index(string(buffer), "=")+1 : strings.Index(string(buffer), "&")]))

		req, err := http.NewRequest(
			"GET",
			"https://api.github.com/user",
			nil,
		)
		if err != nil {
			return "", "", err
		}

		authorizationHeaderValue := fmt.Sprintf("token %s", token)

		req.Header.Set("Authorization", authorizationHeaderValue)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", "", err
		}

		scan, err := io.ReadAll(res.Body)
		if err != nil {
			return "", "", err
		}

		userData := map[string]interface{}{}

		err = json.Unmarshal(scan, &userData)
		if err != nil {
			return "", "", err
		}

		if userData["email"] == nil {
			return "", "", errors.New("your github email address is not public")
		}

		return userData["email"].(string), userData["login"].(string), nil
	}

	return "", "", errors.New("error while connecting with Github")
}
