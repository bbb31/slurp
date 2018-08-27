package main

import ("fmt"
	"time"
	)

func main() {
	fmt.Println("Executing malware...")
	time.Sleep(1 * time.Second)
	fmt.Println("Downloading all your files...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Activating ransomware...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Mining bitcoin...")
	time.Sleep(2 * time.Second)
	fmt.Println("Just kidding")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("You could have been hacked! Don't blindly run code off Github just because someone you trust linked it to you.")
	fmt.Println("More info at https://medium.com/@SweetRollBandit/aws-slurp-github-takeover-f8c80b13e7b5")
}
