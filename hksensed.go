package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/brutella/hc"
	"github.com/zydeco/hksense/sense"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "login" {
		login()
	} else {
		runServer()
	}
}

func login() {
	fmt.Print("Logging in to Sense\n")
	username, password := getCredentials()
	login, _, err := sense.Login(username, password)
	if err != nil {
		log.Fatal("Error logging in: " + err.Error())
	}
	fmt.Printf("Logged In successfully\n")
	fmt.Printf("Account ID: %v\n", login.AccountID)
	fmt.Printf("Token Type: %v\n", login.TokenType)
	expires := time.Now().Add(time.Duration(login.ExpiresIn) * time.Second)
	fmt.Printf("Expires: %v\n", expires.Format(time.RFC1123))
	fmt.Printf("Access Token: %v\n", login.AccessToken)
	fmt.Printf("Refresh Token: %v\n", login.RefreshToken)
}

func getCredentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	buf, _, _ := reader.ReadLine()
	username := string(buf)
	fmt.Print("Password: ")
	buf, _ = terminal.ReadPassword(int(syscall.Stdin))
	password := string(buf)
	fmt.Print("\n")
	return username, password
}

func runServer() {
	accessToken := os.Getenv("SENSE_ACCESS_TOKEN")
	homekitPin := os.Getenv("SENSE_HOMEKIT_PIN")
	deviceName := os.Getenv("SENSE_DEVICE_NAME")
	refreshInterval, err := time.ParseDuration(os.Getenv("SENSE_REFRESH_INTERVAL"))

	if accessToken == "" {
		log.Fatal("Access Token not set")
	}
	if homekitPin == "" {
		homekitPin = "12300123"
		fmt.Printf("HomeKit PIN is 123-00-123\n")
	}
	if deviceName == "" {
		deviceName = "Sense"
	}
	if refreshInterval < 60*time.Second {
		refreshInterval = 60 * time.Second
	}

	acc := sense.NewAccessory(deviceName)
	client := newClient(accessToken)

	account, resp, _ := client.GetAccount()
	if resp.StatusCode != 200 {
		log.Fatal("Could not get account: " + resp.Status)
	} else {
		log.Print("Connected to sense as " + account.Name)
	}
	t, err := hc.NewIPTransport(hc.Config{Pin: homekitPin}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	go updateSensors(acc, client, refreshInterval)

	t.Start()
}

func newClient(accessToken string) *sense.Client {
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	httpClient := config.Client(oauth2.NoContext, token)
	return sense.NewClient(httpClient)
}

func updateSensors(acc *sense.Accessory, client *sense.Client, refreshInterval time.Duration) {
	log.Print("Refresh interval is " + refreshInterval.String())
	for {
		room, _, err := client.Room.CurrentMeasurement("c")
		if err == nil {
			acc.UpdateValues(room)
		} else {
			log.Print("Error getting measurements: ", err.Error())
		}
		time.Sleep(refreshInterval)
	}
}
