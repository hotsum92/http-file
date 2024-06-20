# go-zip

go httpでzipを扱うためのサンプル

## run

### gzip body

###### サーバーを実行

```
go run main-gzip-body.go
```

###### リクエストを送信

```
gzip -c ./data/sample.json | curl -X POST --data-binary @- -H "Content-Encoding: gzip" -H "Content-Type: application/json" localhost:18888
```

###### ログの確認 中身を確認取得できていることがわかる

```
$ go run .
2024/06/19 19:45:55 start http listening :18888
{
  "key01": "value01",
  "key02": "value02"
}
```


