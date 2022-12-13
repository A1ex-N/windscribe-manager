# windscribe-manager


## What it does:
Loads [windscribe](https://windscribe.com) accounts from a file, logs into them and scrapes the bandwidth reset date and bandwidth usage, and saves it back to the file. Might not be so useful if you just have one account, but i have 11 so it is quite useful. 

## How to use it:
just run `windscribe-manager.exe` without passing any arguments and it'll print all your accounts and prompt you to enter a number. When the accounts are printed each account is given a number, which you can enter to log into that account. 

## Important Information
date_created, reset_date and last_checked will all be converted from mm/dd/yyyy to dd/mm/yyyy, and reset_date will have 1 day added to it to make up for the timezone difference (windscribe uses a Canadian timezone and i'm in Australia)

This will currently only work on windows because of this line here: https://github.com/A1ex-N/windscribe-manager/blob/03e557927577aa8ed3f345380f378db4e37f130b/main.go#L44

There's currently two different arguments you can give to the program
1. `-print` will just print all your accounts and exit
2. `-login [int]` logs into a specific account and updates the info, prints the updated info and then exits.

The json file containing account info should look like this
```json
[
	{
		"username": "",
		"password": "",
		"email": "",
		"date_created": "",
		"reset_date": "",
		"data_usage": "",
		"last_checked": "",
		"referer_url": ""
	}
]
```

## To do:
* The program currently assumes you already have a file containing windscribe accounts, but there should be a way to start from scratch. There's already a function in the code `GetAllData(username, password string) WindscribeAccount` which is intended to be used to scrape a new account (one you haven't scraped yet) which will scrape all the relevant data (except the email, because it's hidden when you don't actually use a browser). the usage would be like `windscribe-manager.exe -add username password`
