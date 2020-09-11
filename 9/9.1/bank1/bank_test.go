// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"../bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	go func() {
		str := bank.Withdraw(100)
		fmt.Println(str)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		str := bank.Withdraw(1000)
		fmt.Println(str)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// Wait for both transactions.
	<-done
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
