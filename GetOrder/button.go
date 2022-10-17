// WORKING BUTTON WITH 0AUTH API

package main

import (
	"log"
	"net/http"

	"io/ioutil"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
)

func main() {

	url := "https://app-sg.onebillsoftware.com/rest/OrderService/v1/orders"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "YOUR ACCESS TOKEN GOES HERE"

	myApp := app.New()
	myWindow := myApp.NewWindow("Aatrox Communications OneBIll Software")

	content := widget.NewButton("GET ORDER", func() {

		// Create a new request using http
		req, err := http.NewRequest("GET", url, nil)
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		log.Println(string([]byte(body)))

	})

	// Create a new request using http

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
