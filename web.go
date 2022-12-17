package main

import (
	"astuart.co/goq"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type CsrfTokens struct {
	CsrfToken string `json:"csrf_token"`
	CsrfTime  int64  `json:"csrf_time"`
}

type LoginData struct {
	CsrfToken string `json:"csrf_token"`
	CsrfTime  int64  `json:"csrf_time"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

const (
	CsrfURL    string = "https://res.windscribe.com/res/logintoken"
	LoginURL   string = "https://windscribe.com/login"
	AccountURL string = "https://windscribe.com/myaccount"
)

func NewLoginBody(username, password string, tokens *CsrfTokens) url.Values {
	return url.Values{
		"login":      {"1"},
		"upgrade":    {"0"},
		"username":   {username},
		"password":   {password},
		"csrf_token": {tokens.CsrfToken},
		"csrf_time":  {fmt.Sprint(tokens.CsrfTime)},
	}
}

func GetCSRFTokens() CsrfTokens {
	resp, err := http.Post(CsrfURL, "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var res CsrfTokens
	json.NewDecoder(resp.Body).Decode(&res)
	return res
}

func UpdateDataUsageAndResetDate(user *WindscribeAccount) error {
	loginResponse := loginAndGetResponse(user.Username, user.Password)
	if loginResponse.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status code not 200: status code is %v", loginResponse.StatusCode))
	}
	defer loginResponse.Body.Close()
	var data WindscribeAccount
	err := goq.NewDecoder(loginResponse.Body).Decode(&data)
	if err != nil {
		return err
	}
	data.DataUsage = strings.Replace(data.DataUsage, "\n", "", 1)
	data.ResetDate, err = ConvertToAUSDate(data.ResetDate)
	if err != nil {
		log.Fatal(err)
	}
	user.UpdateAccount(data.ResetDate, data.DataUsage)
	return nil
}

// GetAllData This function shouldn't be used for existing accounts. Only new ones that haven't been scraped yet
func GetAllData(username, password string) WindscribeAccount {
	loginResponse := loginAndGetResponse(username, password)
	if loginResponse.StatusCode != http.StatusOK {
		log.Fatalf("Status not OK: %v", loginResponse.Status)
	}
	defer loginResponse.Body.Close()
	var data WindscribeAccount
	err := goq.NewDecoder(loginResponse.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	data.DataUsage = strings.Replace(data.DataUsage, "\n", "", 1)
	data.Username = username
	data.Password = password
	data.Email = "unknown"
	data.LastChecked = GetCurrentTime()
	data.ResetDate, err = ConvertToAUSDate(data.ResetDate)
	if err != nil {
		log.Fatal(err)
	}
	data.DateCreated, err = ConvertToAUSDate(data.DateCreated)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func loginAndGetResponse(username, password string) *http.Response {
	tokens := GetCSRFTokens()
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	loginBody := NewLoginBody(username, password, &tokens)

	resp, err := client.PostForm(LoginURL, loginBody)
	if err != nil {
		log.Fatal(err)
	}

	accountPageResp, err := client.Get(AccountURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return accountPageResp
}
