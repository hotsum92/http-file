# go-zip

go httpでzipを扱うためのサンプル

## run

###### サーバーを実行

```
go run .
```

###### リクエストを送信

```
curl -X POST -F "file=@./data.zip" -H "Content-Type: multipart/form-data" localhost:18888
```

###### ログの確認 中身を確認取得できていることがわかる

```
$ go run .
2024/06/19 18:25:50 start http listening :18888

file : data.zip application/octet-stream
data/

data/sample.json
{
  "key01": "value01",
  "key02": "value02"
}
```


