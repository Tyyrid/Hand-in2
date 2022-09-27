# Hand-in2

GG! We know our comments are in Danish, but we'll do better at the exam :)

**a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?**

Our implementation makes use of a user-defined type Message struct, as our package. We use struct to combine/group itmes of different types into a single type. So, our Message structs consist of a messageString of type string and a timestamp of type time.Time. To transmit our Message struct we first split the message into smaller pieces, by using the method splitMessage(message). The method takes the message struct messageString, splits between gaps and puts the new string into a slice called stringSlice. Then we iterate through this slice and makes a new message struct for each element, and puts it into a new slice called messageSlice. Then, we iterate through the messageSlice and sends each struct over a channel called chSenderMessage. 

**b) Does your implementation use threads or processes? Why is it not realistic to use threads?**
threads



**c) How do you handle message re-ordering?**

timestamp
sortere vi på den i reassembleMessage()



**d) How do you handle message loss?**
When we iterate through our messageSlice (a slice consist of message structs), we count up how many message structs we "recive". We then call receiveMessage, with this counter as a parameter. In this method we compare the number of recived message with the counter. If it is the same number, we can then be sure that we have recived all message. If it is not the same, we send a message over the chSenderCheck channel, saying "Der gik noget galt". The sendMessage method then recived this message, and trys to send the message again.


**e) Why is the 3-way handshake important?**

Two importants functions of TCP's three-way handshake are to ensure that both parties know they are ready to transfer data and to agree on the initial sequence numbers, witch are send and acknowledged (whitout error) during the handshake.

