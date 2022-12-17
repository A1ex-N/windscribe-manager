# windscribe-manager


## What it does:
Loads [windscribe](https://windscribe.com) accounts from a file, logs into them and scrapes the bandwidth reset date and bandwidth usage, and saves it back to the file. Might not be so useful if you just have one account, but i have 11 so it is quite useful. 


## How to use it:
just run `windscribe-manager.exe` without passing any arguments and it'll print all your accounts and prompt you to enter a number. When the accounts are printed each account is given a number, which you can enter to log into that account. 


There's currently five different arguments you can give to the program
1. `-print` will just print all your accounts and exit
2. `-login [int]` logs into a specific account and updates the info, prints the updated info and then exits.
3. `-add` adds an account to the accounts file, after scraping it
4. `-path` prints the path of the accounts file
5. `creds [int]` prints the username and password of a specific account

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

the json file will be loaded from an environment variable named WindscribeAccountsEnvironmentVariableName. If the variable doesn't exist, the program will prompt you to enter a path.


## Important Information
date_created, reset_date and last_checked will all be converted from mm/dd/yyyy to dd/mm/yyyy, and reset_date will have 1 day added to it to make up for the timezone difference (windscribe uses a Canadian timezone and i'm in Australia)

This will currently only work on windows because of this line here: https://github.com/A1ex-N/windscribe-manager/blob/03e557927577aa8ed3f345380f378db4e37f130b/main.go#L44
