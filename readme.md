# go-zip

go httpでzipを扱うためのサンプル

## run

###### サーバーを実行

```
go run .
```

###### リクエストを送信

```
curl -X POST -F "file=@./data.zip" -H "Content-Type: application/zip" localhost:18888
```

###### ログの確認 中身を確認取得できていることがわかる

```
$ go run .
2024/06/19 18:46:51 start http listening :18888
2024/06/19 18:46:53 content-type:  application/zip; boundary=------------------------507rci0mAcePoD935Bkg2x
data/

data/sample.json
{
  "key01": "value01",
  "key02": "value02"
}
```


