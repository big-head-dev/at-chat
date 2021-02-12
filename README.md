# AT-Chat App

Create a chat application using golang and websockets that supports multiple users through a simple front-end.

## TODO: 
### golang
- [✔] http listener - root
- [✔] http listener - socks
- [✔] retrieve username
- [✔] client object
- [✔] room object
- [✔] broadcast status messages
- [✔] read message from user
- [✔] broadcast message from user
- [x] save user message to file
- [x] read messages to new user writer on join
- [✔] client-side disconnect
- [x] handle heartbeat
- [✔] unique usernames
### ui
- [✔] message box
- [✔] user list
- [✔] join button
- [✔] leave button
- [✔] message input
- [✔] state logic
- [x] handle heartbeat

## Limitations, Caveats, Concerns

- Sending messages from the client is primitive (text string only)
- It doesn't handle message time very robustly; just visual
- ~~I think the write handlers could be done more efficently (golang knowledge)~~ using outing channel  for user writes
- ~~Using file io instead of redis/db to limit requirements for setup~~
  - ~~file stores for the duration of the go run~~
- render issue with the user list once the chatwindow goes into scroll mode