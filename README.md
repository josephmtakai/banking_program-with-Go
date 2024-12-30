Banking Program in Go

Overview
The Banking Program is a simple Go application designed to simulate basic banking operations such as checking account balance, depositing money, and withdrawing money. This project is ideal for learning and demonstrating core Go programming concepts like functions, loops, conditionals, and user input handling.

Features:

Show Balance: Displays the current account balance.
Deposit Money: Adds funds to the account.
Withdraw Money: Withdraws funds, with checks for sufficient balance and valid amounts.
Exit Program: Safely terminates the program.

How It Works:

The user interacts with a menu-driven interface.
Each menu option corresponds to a specific banking operation.
The account balance updates dynamically based on deposits or withdrawals.
The program continues to run until the user chooses to exit.

Prerequisites:

Go (Golang) installed on your system (version 1.18 or later).
A terminal or IDE (e.g., Visual Studio Code, GoLand) to run the program.

Setup:

Clone or download this repository to your local machine.
Navigate to the program directory.
Run the program using the command:
go run banking_program.go

Usage:

Launch the program.
Select an option from the menu:
1: Show Balance.
2: Deposit Money.
3: Withdraw Money.
4: Exit Program.
Example Run:*********************
   Banking Program   
*********************
1. Show Balance
2. Deposit Money
3. Withdraw Money
4. Exit
*********************
Enter your choice (1-4): 2
Enter the amount to deposit: $100
$100.00 has been deposited. Your new balance is $100.00

Screenshot:
![Golang](https://github.com/user-attachments/assets/8079a722-bb11-4d84-9690-0f3dffeecedf)


Code Structure:

showBalance(balance): Displays the current account balance.
deposit(balance): Allows the user to deposit funds and validates the input.
withdraw(balance): Enables fund withdrawal while checking for sufficient balance and valid amounts.
main(): The program's entry point, managing user interaction and menu navigation.

Potential Enhancements:

Add a login system to support multiple users.
Include transaction history tracking.
Enable persistent data storage using a database or file system.
Implement interest calculations for savings accounts.

Contributing:

Contributions are welcome! Feel free to fork the repository, make enhancements, and submit pull requests.

License:

This project is licensed under the MIT License. See the LICENSE file for more information.


