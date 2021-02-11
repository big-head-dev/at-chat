let socket
let chatUsername

const isSocketOpen = function () {
  return socket && socket.readyState == socket.OPEN
}

const sockErrorHandler = function (e) {
  //TODO: post error as a chat message
  console.log(e)
}
const sockOpenHandler = function (e) {
  console.log(e)
}
const sockMessageHandler = function (e) {
  //TODO: handle message types and create messages in chatwindow
  console.log(e)
}
const sockCloseHandler = function (e) {
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
    socket.onerror = sockErrorHandler
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
const pushMessage = function(message) {

}

// creates the message html
const createMessageHtml = function(time, username, message) {
  
}
