package main

import (
	"fmt"
    "github.com/asaskevich/govalidator"
    "strings"
)

func extractNormalizedData(data map[string]interface{}) (map[string]interface{}, error) {
    normalizedData := make(map[string]interface{})

    emailFields := []string{"email", "e-mail"}
    phoneFields := []string{"phone", "mobile", "telefon", "mobilephone"}
    nameFields := []string{"name", "firstName", "first_name"}

    for key, value := range data {
        lowerKey := strings.ToLower(key)

        if containsString(emailFields, lowerKey) {
            if email, ok := value.(string); ok && govalidator.IsEmail(email) {
                normalizedData["email"] = email
            } else {
                return nil, fmt.Errorf("Invalid email format")
            }
        } else if containsString(phoneFields, lowerKey) {
            if phone, ok := value.(string); ok && !govalidator.IsURL(phone) && !govalidator.IsEmail(phone) {
                normalizedData["phone"] = phone
            } else {
                return nil, fmt.Errorf("Invalid phone format")
            }
        } else if containsString(nameFields, lowerKey) {
            if name, ok := value.(string); ok {
                normalizedData["name"] = name
            } else {
                return nil, fmt.Errorf("Invalid name format")
            }
        }
    }

    return normalizedData, nil
}

func containsString(slice []string, str string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }
    return false
}
