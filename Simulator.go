package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	//CHANNELS
	chSenderCheck := make(chan string)
	chReceiverCheck := make(chan string)

	//MESSAGE
	message := Message{"Hej med dig", time.Now()}

	//GORUTINE
	go sendMessage(chReceiverCheck, chSenderCheck, message)
	go checkIfAvailable(chReceiverCheck, chSenderCheck)

	time.Sleep(10 * time.Second)
}

type Message struct {
	messageString string
	timestamp     time.Time
}

func sendMessage(chReceiverCheck chan string, chSenderCheck chan string, message Message) {
	//SENDMESSAG STRUCT CHANNEL
	chSenderMessage := make(chan Message, 3)
	
	//KALDER MODTAGER - ER DU KLAR?
	chReceiverCheck <- "Er du klar?"

	//MODTAGER BESKED
	receivedMessage := <-chSenderCheck

	if receivedMessage == "Jeg er klar!" {

		//COUNTER HVOR MANGE "STYKKER" AF EN BESKED DER BLIVER SENDT
		var messageCounter = 0

		//SENDER VORES BESKED
		for _, v := range splitMessage(message) {
			v.timestamp = time.Now()
			messageCounter += 1
			chSenderMessage <- v
		}

		receiveMessage(chSenderMessage, messageCounter, chSenderCheck)

		//PRINTER AT VI HAR MODTAGET EN BESKED
		go func(ch <-chan string) {
			v := <-chSenderCheck

			//ERROR HANDILING
			if v == "Der gik noget galt" {
				fmt.Println("Der gik noget galt, med at modtage beskeden. Vi prøver at sende den igen.")
				sendMessage(chReceiverCheck, chSenderCheck, message)
			} else {
				//PRINTER AT VI HAR MODTAGET BESKEDEN
				fmt.Println(v)
				os.Exit(0)
			}

		}(chSenderCheck)

		//kan lave if-statement til at stoppe program

	} else {
		//HVIS DEN IKKE FÅR SVAR INDENFOR 5 MS
		time.Sleep(5 * time.Millisecond)
		sendMessage(chReceiverCheck, chSenderCheck, message)
	}

}

func checkIfAvailable(chReceiverCheck chan string, chSenderCheck chan string) {

	//MODTAGER BESKED
	CheckIfAvailableMesseage := <-chReceiverCheck
	//CHECKER BESKEDEN DEN HAR MODTAGER
	if CheckIfAvailableMesseage == "Er du klar?" {
		//SENDER TIL SENDERCHANNEL AT JEG ER KLAR
		chSenderCheck <- "Jeg er klar!"

	}

}

func receiveMessage(chSenderMessage chan Message, messageCounter int, chSenderCheck chan string) {
	messageSlice := make([]Message, 0, 3)

	//MODTAGER BESKEDERNE
	for v := range chSenderMessage {
		messageSlice = append(messageSlice, v)

		//NÅR CHANNELSEN ER TOM GÅ VIDERE
		if len(chSenderMessage) == 0 {
			break
		}
	}

	//TJEKKER OM VI HAR MODTAGER ALLE BESKEDERNE
	if messageCounter == len(messageSlice) {
		completeMessage := reassembleMessage(messageSlice)
		fmt.Println(completeMessage)

		//SENDER VI HAR MODTAGET EN BESKED
		go func(ch chan<- string) {
			chSenderCheck <- "Jeg har modtaget en besked"

		}(chSenderCheck)

	} else {
		//SENDER BESKED OM AT DER GIK NOGET GALT
		go func(ch chan<- string) {
			chSenderCheck <- "Der gik noget galt"

		}(chSenderCheck)
	}

}

func splitMessage(message Message) (messageSlice []Message) { //returnerer et slice af structs
	//Deler message på hvert mellemrum eller newline og laver nyt slice med dem
	stringSlice := strings.Fields(message.messageString)

	// laver en ny struct for hvert element i stringSlice og tilføjer den struct til messageSlice
	for i := 0; i < len(stringSlice); i++ {
		message := Message{stringSlice[i], time.Now()}
		messageSlice = append(messageSlice, message)
	}

	return messageSlice
}

func reassembleMessage(messageSlice []Message) (reassembledMessage string) {
	// Modtager messageSlice - et slice med alle de beskeder der er modtaget

	// Sorterer dem i rigtig rækkefølge udfra timestamp
	sort.Slice(messageSlice, func(i, j int) bool {
		return messageSlice[i].timestamp.Before(messageSlice[j].timestamp)
	})

	// Sætter beskederne sammen til én string
	strs := make([]string, len(messageSlice))
	for i, v := range messageSlice {
		strs[i] = v.messageString
	}
	reassembledMessage = strings.Join(strs, " ")

	//fmt.Println(reassembledMessage)
	return reassembledMessage
}
