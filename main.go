package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lcl101/mybackup/dump"
	"github.com/lcl101/mybackup/option"
	"github.com/lcl101/mybackup/security"
)

//https://github.com/maurodelazeri/mysql-backup-golang/blob/master/mars.go
//解压请使用绝对路径：tar xzPvf backdata6.tar.gz

const (
	MY_KEY = "1234567890123456"
)

func main() {
	option.InitOpt()
	if option.Opt.Security != "" {
		str := security.Encrypt(option.Opt.Security, MY_KEY)
		fmt.Println("en-password=" + str)
		os.Exit(0)
	} else {
		str, err := security.Decrypt(option.Opt.Password, MY_KEY)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		option.Opt.Password = str
		// fmt.Println("de-password=" + option.Opt.Password)
	}

	if option.Opt.BackupPath == "" {
		tmp, _ := os.Getwd()
		option.Opt.BackupPath = tmp
	}

	dump.Dump(option.Opt)
}

// func printMessage(message string, verbosity int, messageType int) {
// 	colors := map[int]color.Attribute{Info: color.FgGreen, Warning: color.FgHiYellow, Error: color.FgHiRed}

// 	if verbosity == 2 {
// 		color.Set(colors[messageType])
// 		fmt.Println(message)
// 		color.Unset()
// 	} else if verbosity == 1 && messageType > 1 {
// 		color.Set(colors[messageType])
// 		fmt.Println(message)
// 		color.Unset()
// 	} else if verbosity == 0 && messageType > 2 {
// 		color.Set(colors[messageType])
// 		fmt.Println(message)
// 		color.Unset()
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		color.Set(color.FgHiRed)
// 		panic(err)
// 		color.Unset()
// 	}
// }
