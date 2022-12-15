package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type WindscribeAccount struct {
	Username    string `json:"username" goquery:"div.ma_item:nth-child(1) > span:nth-child(2)"`
	Password    string `json:"password" goquery:""`
	Email       string `json:"email"`
	DateCreated string `json:"date_created" goquery:"div.ma_item:nth-child(2) > span:nth-child(2)"`
	ResetDate   string `json:"reset_date" goquery:"div.ma_item:nth-child(6) > span:nth-child(2)"`
	DataUsage   string `json:"data_usage" goquery:"div.ma_item:nth-child(7) > span:nth-child(2)"`
	LastChecked string `json:"last_checked"`
	RefererUrl  string `json:"referer_url" goquery:"#ma_friend_url,[href]"`
}

func (w *WindscribeAccount) UpdateAccount(resetDate string, dataUsage string) {
	w.ResetDate = resetDate
	w.DataUsage = dataUsage
	w.LastChecked = GetCurrentTime()
}

func GetCurrentTime() string {
	//02 Jan 06 15:04 MST
	return string(time.Now().Format("02 January 2006 15:04"))
}

// ConvertToAUSDate couldn't get the time.Parse() function to parse December 14th 2022 to 14 December 2022,
// so I had to do it a shitty way like this.
func ConvertToAUSDate(original string) (string, error) {
	toRemove := [4]string{"st", "nd", "th", "rd"}
	for _, i := range toRemove {
		original = strings.Replace(original, i, "", 1)
	}
	t, err := time.Parse("January 2 2006", original)
	if err != nil {
		return "", err
	}
	// Add a day because Windscribe's timezone is a Canadian one and I want it in Australian
	correctDate := t.AddDate(0, 0, 1)
	return correctDate.Format("02 January 2006"), nil
}

func DumpWindscribeAccounts(accounts *[]WindscribeAccount, filename string) error {
	marshalledAccounts, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}
	writeErr := os.WriteFile(filename, marshalledAccounts, 0644)
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func GetWindscribeAccounts(filename string) ([]WindscribeAccount, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("'%v' doesn't exist. creating it\n", filename)
		_, err := os.Create(filename)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed creating %v\n%v", filename, err))
		}
		return []WindscribeAccount{}, nil
	}
	rawData, fileReadErr := os.ReadFile(filename)

	if fileReadErr != nil {
		return nil, fileReadErr
	}

	var data []WindscribeAccount
	unmarshalErr := json.Unmarshal(rawData, &data)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return data, nil
}
