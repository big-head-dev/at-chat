let socket
let chatUsername

const isSocketOpen = function () {
  return socket && socket.readyState == socket.OPEN
}

const sockOpenHandler = function (e) {
  // pushMessage({time: "", username: "", message: "Connected..."})
  // console.log(e)
}
const sockMessageHandler = function (e) {
  // handle message types and create messages in chatwindow
  const sockMessage = JSON.parse(e.data)
  console.debug(sockMessage)
  if(!sockMessage) return
  switch (sockMessage.type) {
    case "status":
    case "chat":
      pushMessage(sockMessage)
      break;
    case "chats":
      sockMessage.messages.forEach(message => {
        pushMessage(message)
      });
      break;
    default:
      //unknown message type
      console.log(e)
      break;
  }
  
}
const sockCloseHandler = function (e) {
  pushMessage({time: "", username: "", message: "You have been disconnected from the chat room."})
  console.log(e)
  socket = null
  clearCookie()
  flipInputPanels(false)
}

const joinRoom = function (username) {
  if (!isSocketOpen()) {
    // set this with the username prompt
    setCookie(username)
    socket = new WebSocket('ws://' + document.location.host + '/chat')
    flipInputPanels(true)
    // socket.onerror = sockErrorHandler
    socket.onopen = sockOpenHandler
    socket.onmessage = sockMessageHandler
    socket.onclose = sockCloseHandler
  }
}

const leaveRoom = function () {
  if (isSocketOpen()) {
    socket.close()
    sockCloseHandler()
  }
}

const sendMessage = function (message) {
  if (isSocketOpen()) {
    socket.send(message)
  }
}

// DOM methods

const clearCookie = function () {
  document.cookie = ''
}

const setCookie = function (username) {
  document.cookie = 'atchat-username=' + username + '; '
}

window.onload = function () {
  if (!window['WebSocket']) {
    console.log('WebSockets not supported.')
    return
  }
}
window.onbeforeunload = function () {
  leaveRoom()
}

const flipInputPanels = function (showMessagePanel) {
  var messagePanel = document.getElementById('messsagePanel')
  var joinPanel = document.getElementById('joinPanel')

  if (showMessagePanel) {
    messagePanel.classList.remove('hidethis')
    joinPanel.classList.add('hidethis')
  } else {
    messagePanel.classList.add('hidethis')
    joinPanel.classList.remove('hidethis')
  }
}

const joinButtonClick = function () {
  var usernameInput = document.getElementById('usernameInput')
  var usernameDisplay = document.getElementById('userProfileUsername')
  var username = usernameInput.value
  if (username && username.length > 1) {
    usernameDisplay.innerText = username
    joinRoom(username)
  } else {
    usernameInput.focus()
  }
}

const leaveButtonClick = function () {
  if (confirm('Are you sure you want to leave the chat room?')) {
    leaveRoom()
  }
}

const sendButtonClick = function () {
  const messageInput = document.getElementById('messageInput')
  let message = messageInput.value
  if (message && message.length > 0) {
    sendMessage(message)
    messageInput.value = ''
  }
  messageInput.focus()
}

// added the message to the dom, manages scrolling
const pushMessage = function(chatMessage) {
  const chatWindow = document.getElementById('chatWindow')
  const { time, username, message } = chatMessage

  // call createMessageHtml for the message type and appendChild to chatWindow
  const newMessageDiv = createMessageHtml(time, username, message)
  chatWindow.appendChild(newMessageDiv)
  // scrolling
  newMessageDiv.scrollIntoView()
}

// creates the message html
const createMessageHtml = function(time, username, message) {
  // <div class="message">
  //   <span class="message-time">4:45</span>
  //   <span class="message-username">John444asd:</span>
  //   <p class="message-text">
  //     Lorem ipsum dolor sit amet consectetur adipisicing elit. Dicta,
  //     quae?
  //   </p>
  // </div>
  var newMessageDiv = document.createElement('div')
  newMessageDiv.className = 'message'

  var newMessageTimeSpan = document.createElement('span')
  newMessageTimeSpan.className = 'message-time'
  newMessageTimeSpan.innerText = time || ''

  var newMessageUserSpan = document.createElement('span')
  newMessageUserSpan.innerText = username || ''
  newMessageUserSpan.className = 'message-username'

  var newMessageTextP = document.createElement('p')
  newMessageTextP.innerText = message || ''
  newMessageTextP.className = 'message-text'

  newMessageDiv.appendChild(newMessageTimeSpan)
  newMessageDiv.appendChild(newMessageUserSpan)
  newMessageDiv.appendChild(newMessageTextP)
  return newMessageDiv
}
