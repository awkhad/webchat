$ ->
  unless window.WebSocket
    alert("you brower is not support websocket")
    return

  room_addr = "ws://"+ window.location.host + window.location.pathname + "/chatting"
  window.room = new Room(room_addr)
  room.room_addr()
  console.log room.connect()

  #$('#sayit-button').click ->
  #  text = $('textarea').val()
  #  if text
  #    alert(text)
  #  else
  #    alert("empty input")


class Room
  constructor: (ws_url) ->
    @ws_url = ws_url

  room_addr: ->
    console.log @ws_url

  connect: ->
    w = new WebSocket(@ws_url)
