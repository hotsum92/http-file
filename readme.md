# go-zip

go httpでzipを扱うためのサンプル

## run

### gzip body

###### サーバーを実行

```
go run cmd/gzip-body/main.go
```

###### リクエストを送信

```
gzip -c ./data/sample.json | curl -X POST --data-binary @- -H "Content-Encoding: gzip" -H "Content-Type: application/json" localhost:18888
```

###### ログの確認 中身を確認取得できていることがわかる

```
2024/06/19 19:45:55 start http listening :18888
{
  "key01": "value01",
  "key02": "value02"
}
```

### zip file body

###### サーバーを実行

```
go run cmd/zip-file-body/main.go
```

###### リクエストを送信

```
curl -X POST -F "file=@./data.zip" -H "Content-Type: application/zip" localhost:18888
```

###### ログの確認 中身を確認取得できていることがわかる

```
2024/06/20 09:31:56 start http listening :18888
Content-Encoding:  
Content-Type:  application/zip; boundary=------------------------cuJiKAviER8myETtph8pcX
fileName:  data/
fileContent:  
fileName:  data/sample.json
fileContent:  {
  "key01": "value01",
  "key02": "value02"
}
```

### multipart file body

###### サーバーを実行

```
go run cmd/multipart/main.go
```

###### リクエストを送信

```
curl -X POST -F "file=@./data.zip" -H "Content-Type: multipart/form-data" localhost:18888
```

###### ログの確認 中身を確認取得できていることがわかる

```
2024/06/20 09:39:02 start http listening :18888
Content-Encoding:  
Content-Type:  multipart/form-data; boundary=------------------------VIQqlsx7nrNwMKYhmPHo1L
file : data.zip application/octet-stream
fileName:  data/
fileContent:  
fileName:  data/sample.json
fileContent:  {
  "key01": "value01",
  "key02": "value02"
}
```
