package models

import "fmt"

// Account 账户结构体
type Account struct {
	AccountNo string
	Password  string
	Balance   float64
}

// SaveMoney 存款
func (a *Account) SaveMoney(money float64, password string) {
	if password != a.Password {
		fmt.Println("输入密码有误")
		return
	}

	if money <= 0 {
		fmt.Println("存款金额不能小于等于0")
	}

	a.Balance += money

	fmt.Println("存款成功")
}

// Withdrew 取款
func (a *Account) Withdrew(money float64, password string) {
	if password != a.Password {
		fmt.Println("输入密码有误")
		return
	}

	if money <= 0 {
		fmt.Println("取款金额不能小于等于0")
	}

	a.Balance -= money

	fmt.Println("取款成功")
}
