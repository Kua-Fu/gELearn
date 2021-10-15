package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type t struct {
}

type a interface {
	a1() string
	a2() string
}

// 创建一个接口b
type b interface {
	// a
	a1() string
	a2() string
}

type c interface {
	a
	b
}

func f1() {
	for {

		fmt.Println("----f1 test----")
		time.Sleep(time.Second * 1)
	}
}

func f2() {
	fmt.Println("--------refer end------")
}

func f3() {
	fmt.Println("-----f3----")
}

func main() {

	// cmd := exec.Command("sleep", "5")

	// // Request the OS to assign process group to the new process, to which all its children will belong
	// cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// go func() {
	// 	time.Sleep(time.Second)
	// 	// Send kill signal to the process group instead of single process (it gets the same value as the PID, only negative)
	// 	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	// }

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Command finished successfully")

	f3()

	// defer f1()

	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		for s := range c {
			switch s {
			case os.Kill: // kill -9 pid，下面的fmt无效
				fmt.Println("强制退出", s)
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT: // ctrl + c
				f2()
				fmt.Println("退出", s)
			case syscall.SIGUSR1: // kill -USR1 pid
				fmt.Println("usr1", s)
			case syscall.SIGUSR2: // kill -USR2 pid
				fmt.Println("usr2", s)
			}
		}

	}()

	f4()
}

func f4() {

	fmt.Println("----f4---")
	time.Sleep(time.Hour * 10)
}
