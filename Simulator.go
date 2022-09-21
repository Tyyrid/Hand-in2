package main

import (
	"fmt"
	"time"
)



func main() {
	fmt.Println("start")

	//CHANNELS
	chSenderCheck := make(chan string)
	chReceiverCheck := make(chan string)
	chSenderMessage := make(chan string)
	chReceiverMessage := make(chan string)

	//MESSAGE
	var message = new Message() {messageString}

    //GORUTINE
	go sendMessage(chReceiverCheck, chSenderCheck)
	go checkIfAvailable(chReceiverCheck, chSenderCheck)


	//go receiveMessage(chSenderCheck)

	time.Sleep(10 * time.Second)

	//når chSenderCheck har modtager besked om at chReceiverCheck har fået beskeden, så stopper vi
}

type Message struct {
	messageString   string
	timestamp time.Time
}

// DEN DER SENDER BESKEDERNE
func thread() {

}

func (message *Message) sendMessage(chReceiverCheck chan string, chSenderCheck chan string) {
	//Spørg modtager "er du klar?" - kalder på checkIfAvailable()
	
	//KALDER MODTAGER - ER DU KLAR?
	chReceiverCheck <- "Er du klar?"

	//MODTAGER BESKED
	receivedMessage := <-chSenderCheck
	if receivedMessage == "Jeg er klar!" {

		splitMessage();




		//splitter besked - splitMessage()
		//sender besked
		
		//go receiveMessage() ?
	} else {
		//time-out
		//sendMessage()
	}

}

func checkIfAvailable(chReceiverCheck chan string, chSenderCheck chan string) {
	fmt.Println("vi er i check-metoden")
	//MODTAGER BESKED
	CheckIfAvailableMesseage := <-chReceiverCheck
	//CHECKER BESKEDEN DEN HAR MODTAGER
	if CheckIfAvailableMesseage == "Er du klar?" {
		fmt.Println(CheckIfAvailableMesseage)
		//SENDER TIL SENDERCHANNEL AT JEG ER KLAR
		chSenderCheck <- "Jeg er klar!"
	
	}

	//er altid true i vores tilfælde
}

func receiveMessage() {
	//skal mellem andet tjekke om alt er i rette rækkefølge
	////kalder på forskellige metoder om besken er OK - reassembleMessage()
	//svarer til afsender channel, at den har modtaget/ikke modtaget besked
	//måske lave en ekstra channel
}

func splitMessage() { //array of structs?
	//dele data i mindre stykker
	//afsender
	//skal returnere noget der er splittet

}

func reassembleMessage() {
	//sætter data samen igen
	//tjekker om den har fået alt
	//modtager
}

//lav to channels, som snakker med hinanden
//afsender skal kunne spørge modtager 'er du ledig'
//modtager skal kunne svare
//afsender siger okay. jeg er klar. sender data. Dataen skal sende på en TCP måde. Dataen skal pakkes i en packets.
// modtager bekræfter at de har modtaget
//Data kunne være i form af en struct

////Additional////
//dataen skal deles i mindre stykker

////Non-functional features////
//håndtere loss af data - hvis afsender ikke får bekræftelse for at dataen er modtaget, skal den sende det igen
//håndtere at data bliver sendt i forkert rækkefølge (sequencing) - man kan sætte timestamp på dataen (hvornår det er sendt)
// håndtere timeout (får ikke svar)
