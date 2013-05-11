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
    @userlist = []
    #@key = window.location.pathname.split('/')[2]

  @checkWs: ->
    unless window.WebSocket
      alert("you brower is not support websocket")
      return

  roomAddr: ->
    console.log @ws_url
  
  joinRoom: ->
    message = new Message("join", "#{@currentUser()} has join room")
    #console.log(JSON.stringify(message))
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
      message = $.parseJSON(e.data)

      if message.Type == "join"
        @addUserToList(message.User)
      if message.Type == "leave"
        @removeUserFromList(message.User)

      console.log message

  getUsersList: ->
      url = window.location.pathname + "/users.json"

      $.getJSON(url, (data) =>
          if data.Users != null
            #@userlist.push(x) for x in data.Users
            @userlist = data.Users
      )
      return @userlist

  addUserToList: (user)->
    names = []
    $("#userlist>ul>li").each ->
      names.push $(this).text()

    if names.indexOf(user.Name) == -1
      $("#userlist>ul").append $("<li><img src=#{user.Avatar}/>#{user.Name}</li>")

  removeUserFromList: (user) ->
    $("#userlist>ul>li").each ->
      if $(this).text() == user.Name
        $(this).remove()


class Message
  constructor: (type, text) ->
    @type = type
    @text = text
