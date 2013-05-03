$ ->
  # new room
  Room.checkWs()
  window.room = new Room("ws://"+ window.location.host + window.location.pathname + "/chatting")
  room.ws_conn.onopen = room.joinRoom()

  # deal with message 
  $('#sayit-button').click ->
    text = $('#chat-form').val()
    if text
      alert(text)
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
    # send to server join message

  currentUser: ->
    $("#user-name").text()

  sendMessage: (message) ->
    # json 
    # websocket send



class Message
  constructor: (type, text) ->
    @type = type
    @text = text
