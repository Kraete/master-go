package main

import (
    "fmt"
    "github.com/appliedgocourses/bank"
    "github.com/pkg/errors"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        usage()
    }
    
    err := bank.Load()
    if err != nil {
        fmt.Println(err)
    }

    defer func() {
        err := bank.Save()
        if err != nil {
            fmt.Println("Error saving bank data:", err)
        }
    }()

    switch os.Args[1] {

    case "list":
        fmt.Println(bank.ListAccounts())
    case "create":
        if len(os.Args) < 3 {
            usage()
        }
        name := os.Args[2]

        if a, err := bank.GetAccount(name); err == nil && a != nil {
            fmt.Println("Account ", name, " already exists.")
            return
        }
        bank.NewAccount(name)
        fmt.Println("Account ", name, " created.")
    case "update":
        if len(os.Args) < 4 {
            usage()
        }
        name := os.Args[2]
        amount, err := strconv.Atoi(os.Args[3])
        if err != nil {
            fmt.Println(os.Args[3], "is not a valid integer.")
            return
        }

        bal, err := update(name, amount)
        if err != nil {
            fmt.Println("Error updating account:", err)
            return
        }

        fmt.Printf("Account %s was update by %d credits. New balance: %d", name, amount, bal)
    case "transfer":
        if len(os.Args) < 5 {
            usage()
        }
        name := os.Args[2]
        name2 := os.Args[3]
        amount, err := strconv.Atoi(os.Args[4])
        if err != nil {
            fmt.Println(os.Args[3], "is not a valid integer.")
            return
        }

        bal1, bal2, err := transfer(name, name2, amount)
        if err != nil {
            fmt.Println("Error transferring credits", errors.WithStack(err))
        }
        fmt.Printf("Transferred %d credits from %s to %s.\nNew balances:\n%s: %d\n%s: %d\n", amount, name, name2, name, bal1, name2, bal2)
    case "history":
        name := os.Args[2]
        h, err := history(name)
        if err != nil {
            fmt.Println("Failed loading history:", errors.WithStack(err))
        }
        fmt.Println(h)
    default:
        fmt.Println("Unknown command:", os.Args[2])
        usage()
    }
}

func history(name string) (string, error) {
    account, err := bank.GetAccount(name)
    if err != nil {
        return "", errors.Wrap(err, "history: cannot get account "+name)
    }

    hist := fmt.Sprintf("Transaction history for account %s\n", name)

    next := bank.History(account)

    for amt, bal, more:= 0, 0, true; more; {
        amt, bal, more = next()
        hist += fmt.Sprintf("Amount: %6d, balance: %8d\n", amt, bal)
    }

    return hist, nil
}

func update(name string, amount int) (int, error) {
    account, err := bank.GetAccount(name)
    if err != nil {
        return 0, errors.Wrap(err, "account not found")
    }
    if amount == 0 {
        return bank.Balance(account), errors.New("amount must not be zero")
    }

    var balance int

    if amount > 0 {
        balance, err = bank.Deposit(account, amount)
        if err != nil {
            return balance, errors.Wrap(err, "depositing failed")
        }
    } else {
        balance, err = bank.Withdraw(account, -amount)
        if err != nil {
            return balance, errors.Wrap(err, "withdrawing failed")
        }
    }

    return balance, nil
}

func transfer(name, name2 string, amount int) (int, int, error) {
    account, err := bank.GetAccount(name)
    if err != nil {
        return 0, 0, errors.Wrap(err, "transfer: account "+name+"not found")
    }

    account2, err := bank.GetAccount(name2)
    if err != nil {
        return 0, 0, errors.Wrap(err, "transfer: account "+name+"not found")
    }

    if amount <=0 {
        return 0, 0, errors.New("transfer amount must be positive.) ("+ strconv.Itoa(amount) + ")")
    }
    bal1, bal2, err := bank.Transfer(account, account2, amount)
    if err != nil {
        return 0, 0, errors.Wrap(err, "transfer failed.")
    }

    return bal1, bal2, err
}

func usage() {
	fmt.Println(`Usage:
bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.`)
	os.Exit(1) // Deferred functions are NOT called if exiting this way!
}
