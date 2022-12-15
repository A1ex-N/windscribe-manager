package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func ChooseAccountAndUpdate(accounts *[]WindscribeAccount, filename string) {
	DisplayAccounts(accounts)
	choice := GetUserInputAsInt("\nEnter the account number to log into")

	if choice > len(*accounts) {
		fmt.Printf("You entered %v but there's only %v accounts", choice, len(*accounts))
		os.Exit(1)
	}

	err := UpdateDataUsageAndResetDate(&(*accounts)[choice])
	if err != nil {
		log.Fatal(err)
	}

	err = DumpWindscribeAccounts(accounts, filename)
	if err != nil {
		log.Fatal(err)
	}

	DisplaySpecificAccount(choice, accounts)
}

func LoginAndUpdateSpecifc(choice int, accounts *[]WindscribeAccount, filename string) {
	if choice > len(*accounts) {
		fmt.Printf("You entered %v but there's only %v accounts", choice, len(*accounts))
		os.Exit(1)
	}

	err := UpdateDataUsageAndResetDate(&(*accounts)[choice])
	if err != nil {
		log.Fatal(err)
	}

	err = DumpWindscribeAccounts(accounts, filename)
	if err != nil {
		log.Fatal(err)
	}

	DisplaySpecificAccount(choice, accounts)
}

func GetUsernameAndPassword() (string, string) {
	var username, password string
	fmt.Print("Enter a username: ")
	fmt.Scanln(&username)
	fmt.Print("\nEnter a password: ")
	fmt.Scanln(&password)
	return username, password
}

func AddNewAccount(username, password, filename string, existingAccounts []WindscribeAccount) {
	for _, account := range existingAccounts {
		if username == account.Username {
			fmt.Printf("'%v' is already in the saved accounts file.\n", username)
			fmt.Println("If you want to update the data for that account, run '-login' with the account number")
			return
		}
	}
	newAccount := GetAllData(username, password)
	existingAccounts = append(existingAccounts, newAccount)
	err := DumpWindscribeAccounts(&existingAccounts, filename)
	if err != nil {
		log.Fatalf("error dumping accounts to %v. error: %v", filename, err)
	}
}

func DisplaySpecificAccount(choice int, accounts *[]WindscribeAccount) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "Username\tData Usage\tReset Date\tLast Checked")
	line := fmt.Sprintf("%v\t%v\t%v\t%v\n", (*accounts)[choice].Username, (*accounts)[choice].DataUsage,
		(*accounts)[choice].ResetDate, (*accounts)[choice].LastChecked)
	fmt.Fprintf(writer, line)
	writer.Flush()
}

func PrintUsernameAndPassword(choice int, accounts *[]WindscribeAccount) {
	fmt.Printf("Username: %v, Password: %v", (*accounts)[choice].Username, (*accounts)[choice].Password)
}

func DisplayAccounts(accounts *[]WindscribeAccount) {
	// https://blog.el-chavez.me/2019/05/05/golang-tabwriter-aligned-text/
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "#\tUsername\tData Usage\tReset Date\tLast Checked")
	for i, account := range *accounts {
		line := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", i, account.Username, account.DataUsage, account.ResetDate, account.LastChecked)
		fmt.Fprintf(writer, line)
	}
	writer.Flush()
}

func GetUserInputAsInt(prompt string) int {
	var choice int
	fmt.Printf("%v: ", prompt)
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatal(err)
	}
	return int(choice)
}
