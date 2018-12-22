# go-chatwork

Go製のChatworkAPIのラッパーライブラリです。

# 使い方

チャットワークAPI公式ドキュメント：  
http://developer.chatwork.com/ja/

## チャットのタスク一覧を取得

GET /rooms/{room_id}/tasks  

```go
package main

import "./go-chatwork"
import "fmt"

func main(){
    client := chatwork.New("token指定")
    api := client.RoomTask
    api.SetRoomId("部屋ID")
    api.SetAccountId("タスクの担当者のアカウントID")
    api.SetAssignedByAccountId("タスクの依頼者のアカウントID")
    api.SetStatus("タスクのステータス(open / done)")
    tasks := api.Execute()
    
    fmt.Println(tasks)
}
```

## 自分のチャット一覧を取得

GET /rooms  

```go
package main

import "./go-chatwork"
import "fmt"

func main(){
    client := chatwork.New("token指定")
    api := client.GetRoom
    room := api.Execute()
    
    fmt.Println(room)
}
```

## チャットのメンバー一覧を取得

GET /rooms/{room_id}/members

```go
package main

import "./go-chatwork"
import "fmt"

func main(){
    client := chatwork.New("token指定")
    api := client.GetRoomMember
    api.SetRoomId("部屋ID")
    room := api.Execute()
    
    fmt.Println(room)
}
```