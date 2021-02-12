let socket
let chatUsername

const isSocketOpen = () => {
  return socket && socket.readyState == socket.OPEN
}

const sockOpenHandler = (e) => {}
const sockMessageHandler = (e) => {
  // handle message types and create messages in chatwindow
  const sockMessage = JSON.parse(e.data)
  console.debug(sockMessage)
  if (!sockMessage) return
  switch (sockMessage.type) {
    case 'status':
    case 'chat':
      pushMessage(sockMessage)
      if (sockMessage.users) {
        updateUserList(sockMessage.users)
      }
      break
    default:
      //unknown message type
      console.log(e)
      break
  }
}
const sockCloseHandler = (e) => {
  switch (e.code) {
    case 1000:
      pushMessage({
        time: '',
        username: '',
        message: 'You have left the chat room.',
      })
      break
    default:
      pushMessage({
        time: '',
        username: '',
        message: 'You have been disconnected from the chat room. ' + e.code,
      })
      break
  }
  // socket = null
  clearCookie()
  flipInputPanels(false)
}

const joinRoom = (username) => {
  if (!isSocketOpen()) {
    // set this with the username prompt
    setCookie(username)
    socket = new WebSocket('ws://' + document.location.host + '/chat')
    clearChatWindow()
    flipInputPanels(true)
    // socket.onerror = sockErrorHandler
    socket.onopen = sockOpenHandler
    socket.onmessage = sockMessageHandler
    socket.onclose = sockCloseHandler
  }
}

const leaveRoom = () => {
  if (isSocketOpen()) {
    socket.close(1000, 'user leaving')
  }
  clearUserList()
}

const sendMessage = (message) => {
  if (isSocketOpen()) {
    socket.send(message)
  }
}

// DOM methods

const clearCookie = () => {
  document.cookie = ''
}

const setCookie = (username) => {
  document.cookie = 'atchat-username=' + username + '; '
}

window.onload = () => {
  if (!window['WebSocket']) {
    pushMessage({ message: 'WebSockets not supported by this browser.' })
    document.getElementById('joinButton').disabled = true
    return
  }
  document.getElementById('usernameInput').focus()
}

const flipInputPanels = (showMessagePanel) => {
  var messagePanel = document.getElementById('messsagePanel')
  var joinPanel = document.getElementById('joinPanel')

  if (showMessagePanel) {
    messagePanel.classList.remove('hidethis')
    joinPanel.classList.add('hidethis')
    //focus input
    document.getElementById('messageInput').focus()
  } else {
    messagePanel.classList.add('hidethis')
    joinPanel.classList.remove('hidethis')
    //focus input
    document.getElementById('usernameInput').focus()
  }
}

const joinButtonClick = () => {
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

const leaveButtonClick = () => {
  if (confirm('Are you sure you want to leave the chat room?')) {
    leaveRoom()
  }
}

const sendButtonClick = () => {
  const messageInput = document.getElementById('messageInput')
  let message = messageInput.value
  if (message && message.length > 0) {
    sendMessage(message)
    messageInput.value = ''
  }
  messageInput.focus()
}

// added the message to the dom, manages scrolling
const pushMessage = (chatMessage) => {
  const chatWindow = document.getElementById('chatWindow')
  const { time, username, message } = chatMessage

  // call createMessageHtml for the message type and appendChild to chatWindow
  const newMessageDiv = createMessageHtml(time, username, message)
  chatWindow.appendChild(newMessageDiv)
  // scrolling
  newMessageDiv.scrollIntoView()
}

// creates the message html
const createMessageHtml = (time, username, message) => {
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

const clearChatWindow = () => {
  const chatWindow = document.getElementById('chatWindow')
  while (chatWindow.firstChild) {
    chatWindow.removeChild(chatWindow.lastChild)
  }
}

const updateUserList = (users) => {
  clearUserList()
  const userList = document.getElementById('userList')
  users.sort().forEach((user) => {
    userList.appendChild(createUserLi(user))
  })
}

const createUserLi = (username) => {
  var userLi = document.createElement('li')
  userLi.innerText = username
  return userLi
}

const clearUserList = () => {
  const userList = document.getElementById('userList')
  while (userList.firstChild) {
    userList.removeChild(userList.lastChild)
  }
}

// listeners

const onEnterKey = (e, next) => {
  if ('Enter' === e.key) {
    next()
  }
  return
}

document
  .getElementById('usernameInput')
  .addEventListener('keyup', (e) => onEnterKey(e, joinButtonClick))
document
  .getElementById('messageInput')
  .addEventListener('keyup', (e) => onEnterKey(e, sendButtonClick))
