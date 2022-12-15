/*
	There's currently no way to add new accounts.

	You have to have an existing json file containing json of the following structure:
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

	There's also no way to tell if you're being rate limited
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const WindscribeAccountsEnvironmentVariableName = "WindscribeAccountsPath"

func SetWindscribeAccountPath() {
	var path string
	fmt.Print("Enter the path to your windscribe accounts file: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv(WindscribeAccountsEnvironmentVariableName, path)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("SETX", WindscribeAccountsEnvironmentVariableName, path)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path set. Close your current terminal and re-run this program")
	os.Exit(0)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	accountsFile, exists := os.LookupEnv(WindscribeAccountsEnvironmentVariableName)

	if !exists {
		SetWindscribeAccountPath()
	}

	//const accountsFile string = "windscribe_accounts.json"
	windscribeAccounts, err := GetWindscribeAccounts(accountsFile)
	if err != nil {
		log.Fatal(err)
	}

	printAccounts := flag.Bool("print", false, "Print all account info and exit")
	addAccount := flag.Bool("add", false, "Add a new account")
	accountNumberToLoginTo := flag.Int("login", -1, "Login to a specific account and update the accounts file")
	accountNumberToPrint := flag.Int("creds", -1, "Print the username and password for a specific account")
	flag.Parse()

	if *printAccounts {
		DisplayAccounts(&windscribeAccounts)
		os.Exit(0)
	}

	if *accountNumberToLoginTo > -1 {
		LoginAndUpdateSpecifc(*accountNumberToLoginTo, &windscribeAccounts, accountsFile)
		os.Exit(0)
	}

	if *accountNumberToPrint > -1 {
		PrintUsernameAndPassword(*accountNumberToPrint, &windscribeAccounts)
		os.Exit(0)
	}

	if *addAccount {
		username, password := GetUsernameAndPassword()
		AddNewAccount(username, password, accountsFile, windscribeAccounts)
		os.Exit(0)
	}

	ChooseAccountAndUpdate(&windscribeAccounts, accountsFile)
}
