# go-chatwork

Go製のChatworkAPIのラッパーライブラリです。

# 使い方

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
    
    for _, p := range tasks {
        fmt.Printf("%d : %s\n", p.TaskID, p.Body)
    }
}
```

チャットワークAPI公式ドキュメント：  
http://developer.chatwork.com/ja/