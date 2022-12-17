# windscribe-manager

A command-line program for managing multiple [windscribe VPN](https://windscribe.com) accounts.

This program logs into a windscribe account, scrapes the data usage and reset date, and then saves it back to a file.

## Installation

1. download the binary from [releases](https://github.com/A1ex-N/windscribe-manager/releases) :)

## Installation (from source, building with Go)

`I used Go version 1.19.3, but you can probably use other versions`

To install `windscribe-manager`, follow these steps:

1.  Clone this repository: `git clone https://github.com/A1ex-N/windscribe-manager`
2.  Navigate to the `windscribe-manager` directory: `cd windscribe-manager`
3.  Build and install the program: `go install`

## Usage

To use `windscribe-manager`, run the following command:

`windscribe-manager`, but for your first time running (and any time you want to add a new account), run: `windscribe-manager -add`

Running it without any arguments will prompt you to enter an account index to log into (but only once you already have at least one account saved to the account file)

### Options

* `-print`: Prints a list of all accounts in the file and exits the program.
* `-login [int]`: Logs into the account with the specified index number, updates the information, prints the updated information, and then exits the program.
     Example: `windscribe-manager.exe -login 2` logs into the third account in the file (indexes start at 0) and prints the updated information.
* `-add`: Adds a new account to the accounts file, scraping the necessary information from the windscribe website.
* `-path`: Prints the file path of the accounts file.
* `-creds [int]`: Prints the username and password of the account with the specified index number.
     Example: `windscribe-manager.exe creds 3` prints the username and password of the fourth account in the file (indexes start at 0).


## Additional Notes

* This program currently only supports Windows. Support for other operating systems may be added in the future.
* All dates will be stored as `date month year`, and `reset date` will have 1 day added to it
