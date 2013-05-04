$ ->
  # new room
  Room.checkWs()
  window.room = new Room("ws://"+ window.location.host + window.location.pathname + "/chatting")
  room.ws_conn.onopen = room.joinRoom()
  room.ws_conn.onmessage = (e) ->
    room.reveiveMessage(e)

  # deal with message 
  $('#sayit-button').click ->
    text = $('#chat-form').val()
    if text
      message = new Message("text", text)
      room.sendMessage(message)
    else
      return

class Room
  constructor: (ws_url) ->
    @ws_url = ws_url
    @ws_conn = new WebSocket(@ws_url)

  @checkWs: ->
    unless window.WebSocket
      alert("you brower is not support websocket")
      return

  roomAddr: ->
    console.log @ws_url
  
  joinRoom: ->
    message = new Message("join", "#{@currentUser()} has join room")
    console.log(JSON.stringify(message))
    # send to server join message

  currentUser: ->
    $("#user-name").text()

  sendMessage: (message) ->
    unless @ws_conn
      return
    @ws_conn.send(JSON.stringify(message))
    $('#chat-form').val('')

    # json 
    # websocket send
  reveiveMessage: (e) ->
    alert(e.data)

class Message
  constructor: (type, text) ->
    @type = type
    @text = text
