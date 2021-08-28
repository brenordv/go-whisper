# Whisper
This is a super bare-bones secure peer-to-peer chat application.
That's it. Simple as that. To use it, someone must wait for connection and somebody else must connect to them.
Once connection is established, Whisper will clear the screen, and you can start chatting.



## Features
- Uses asymmetric cryptography for every message.
- Encryption keys are per-session (don't worry about losing your keys or having them leaked).
- Private Keys never leave the client.
- (not really a feature of Whisper) Every RSA keypair is of 4096 bits.
- No chat log or history.
- No server involved. You will connect directly to the person you want to talk to.
- Don't need any fancy graphics card or operational system. All happens in the console.
- (also not really a feature of Whisper) Supports emojis, if your terminal supports them.



## How it works
If person A wants to talk to person B, then:
- Person A will wait for a connection
- Person A will tell Person B their IP address and port. (it's displayed in the console)
- Person B will connect to Person A using ip:port.
- Once connection is established, the clients will exchange their public keys.
- Now Person A and Person B can talk.
- Every received messaged will have this prefix: ```< ```
- To quit, Person A or B can simply kill the application or type /quit (and press enter)

## Installation
Just unzip and run whisper.exe (or whisper, if you're on linux). To uninstall, just delete the file.

## How to use
Consider the example from the previous item.

### For Person A to wait for connections on port 5000
```shell
whisper.exe wait 5000
```
Once the client is ready to receive connections, a message like this will appear:
```shell
go.Whisper | vx.x.x
waiting for connection. external ip: 111.111.111.111:5000
```
Now person A can inform Person B that they need to connect to ```111.111.111.111:5000``` 

### Person B will connect to Person A using ip:port.
For Person B to connect, they must use the following command:
```shell
whisper.exe connect 111.111.111.111:5000
```

Once a connection is made between the two clients, a message like this will appear:
```shell
connection established.

```

Now they can simply type and chat away.



## Caveat
As of now, to receive connections using your external IP, you might need to enable Port Forwarding on your router (I had to). 
That was the only way I managed to talk to people outside my local network.

## Use case
You could get a Virtual Machine on CaC (CloudAtCost - https://www.cloudatcost.com) and pay for it using cryptocurrency. 
Then you can go somewhere that has free Wi-Fi, use a VPN, turn on Tor and access your VM using CaC's web console.
There, now you can chat pretty much anonymously with someone. Just don't use this to do anything nasty.

## Todo
1. Refactor to use gRPC
2. Add nicknames
3. Add possibility to send files (all encrypted)
4. Create tests

## Notes
I have not tested this using linux. Shame on me, I know... sorry about that.

