# Hand-in2

GG! We know our comments are in Danish, but we'll do better at the exam :)

**a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?**

message struct
hvad en består af

slices via channels



**b) Does your implementation use threads or processes? Why is it not realistic to use threads?**
threads



**c) How do you handle message re-ordering?**

timestamp
sortere vi på den i reassembleMessage()



**d) How do you handle message loss?**

vi har en counter på antal af sendte beskeder
hvis alle ikke er modtaget, så starter vi forfra


**e) Why is the 3-way handshake important?**

Two importants functions of TCP's three-way handshake are to ensure that both parties know they are ready to transfer data and to agree on the initial sequence numbers, witch are send and acknowledged (whitout error) during the handshake.

