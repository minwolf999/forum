package handlefunc

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// Get the email, username, picture of the Google user account
func GoogleLogDecoder(googleauth string) (string, string, string, error) {
	// we delete clientId, client_id, select_by and there data from the json data's. And we remove ,"credential":" from the string
	credential := googleauth[strings.Index(googleauth, ",\"credential\":\"")+15 : strings.Index(googleauth, "\",\"select_by\":")]

	temp, err := JsonWebTokenDecoder(credential)
	if err != nil {
		return "", "", "", err
	}

	return temp["email"].(string), temp["given_name"].(string), temp["picture"].(string), nil
}

// Decode the jwt from google identity service
func JsonWebTokenDecoder(tokenString string) (map[string]interface{}, error) {
	// Divides the JWT into three parts: header, body (payload), and signature
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid jwt format")
	}

	// Decoding the payload part
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	// Deserializing JSON data into a map
	var claims map[string]interface{}
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
