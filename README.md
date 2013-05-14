# web chat use websocket



## 使用redis存储数据相关结构

* save message to list

    key: room:roomkey
    format: type|username|text|time

* room 最近用户存储在集合中

    room:test:users

* user 信息 Hash

   key 是: user:id
   value 是: name, avatar, id

    
