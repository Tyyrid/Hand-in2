package main

import (
	"fmt"
	"time"
	"strings"
	"sort"
)



func main() {
	fmt.Println("start")

	//CHANNELS
	chSenderCheck := make(chan string)
	chReceiverCheck := make(chan string)
	//chSenderMessage := make(chan string)
	//chReceiverMessage := make(chan string)

	//MESSAGE
	message := Message{"Hej med dig", time.Now()}

    //GORUTINE
	go sendMessage(chReceiverCheck, chSenderCheck, message)
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

func sendMessage(chReceiverCheck chan string, chSenderCheck chan string, message Message) {
	//Spørg modtager "er du klar?" - kalder på checkIfAvailable()
	
	//KALDER MODTAGER - ER DU KLAR?
	chReceiverCheck <- "Er du klar?"

	//MODTAGER BESKED
	receivedMessage := <-chSenderCheck
	if receivedMessage == "Jeg er klar!" {

		splitMessage(message);

		// skal få fat i messageSlice fra splitmessage() og gennemløbe den. Og sende hver element én af gangen. Den skal opdatere timestamp til time.Now() når den sender
		// skal også fortælle hvor mange beskeder den har sendt afsted, så recieveMessage() kan tjekke om den har modtaget alle

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
	// Tjekker om den har modtaget det antal beskeder den forventede
	// samler de forskellige structs den har modtaget i et slice 
	// kalder reassembleMessage(messageSlice)
	//svarer til afsender channel, at den har modtaget/ikke modtaget besked
	//måske lave en ekstra channel
}

func splitMessage(message Message) (messageSlice []Message) { //returnerer et array af structs
	//Deler message på hvert mellemrum eller newline og laver nyt slice med dem
	stringSlice := strings.Fields(message.messageString)
	
	// laver en ny struct for hvert element i stringSlice og tilføjer den struct til messageSlice
	for i := 0; i < len(stringSlice); i++ {
		message := Message{stringSlice[i], time.Now()}
		messageSlice = append(messageSlice, message); 
		// Til test // fmt.Println(message.messageString) 
	}

	return messageSlice
}

// Har ikke testet denne metode
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
