# Hand-in2

GG! We know our comments are in Danish, but we'll do better at the exam :)

**a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?**
Our implementation makes use of a user-defined type Message struct, as our package we use struct to combine/group items of different types into a single type. So, our Message structs consist of a messageString of type string and a timestamp of type time.Time. To transmit our Message struct we first split the message into smaller packets, by using the method splitMessage(message). The method takes the message struct messageString, splits on whitespace and puts the new strings into a slice called stringSlice. Then we iterate through this slice and makes a new message struct for each element, and puts it into a new slice called messageSlice. Then, we iterate through the messageSlice and sends each struct over a channel called chSenderMessage. 

**b) Does your implementation use threads or processes? Why is it not realistic to use threads?**
Our implementation uses threads to simulate the TCP protocol. This is not a realistic implementation since threads run locally on a computer. The TCP protocol runs across a network where it can be subjected to failures as message loss and messages being received in a different order than it was sent.


**c) How do you handle message re-ordering?**
As mentioned, our program splits the Message structs into smaller Message structs which are sent one by one. When a struct is sent, its timestamp field is updated to time.Now(). The recieved Messages are added to a slice, which reassembleMessage() then sorts based on timestamp. This ensures, that the Message that was sent first, will be the first element in the slice and so on.


**d) How do you handle message loss?**
When we iterate through our messageSlice (a slice consisting of message structs), we count up how many message structs we send. We then call receiveMessage(), with this counter as a parameter. In this method we compare the number of recived message with the counter. If it is the same number, we can then be sure that we have recived all message. If it is not the same, we send a message over the chSenderCheck channel, saying "Der gik noget galt". The sendMessage() method then receives this message, and tries to send the message again.


**e) Why is the 3-way handshake important?**
Two important functions of TCP's three-way handshake are to ensure that both parties know they are ready to transfer data and to agree on the initial sequence numbers, which are send and acknowledged (without error) during the handshake.

