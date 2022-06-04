# dummige

サイズと色を指定した URL にアクセスすると画像を生成し、その画像を返します。

Example: [https://dummige.herokuapp.com/?size=300x150&color=0055ff](https://dummige.herokuapp.com/?size=300x150&color=0055ff)

生成された画像は下記のように img タグにも適用できます。

![](https://dummige.herokuapp.com/?size=300x150&color=0055ff)

サイズと色を指定しない場合は、デフォルトの「140x140 #555555」の画像が生成されます。

![](https://dummige.herokuapp.com/)

## Requirement

- Go v1.17 ~
- Echo

## Install

```
$ go mod tidy
```

## Start

```
$ go run main.go
```

## Build

```
$ go build
```

```
// Build for Linux OS
$ GOOS=linux GOARCH=amd64 go build
```

## Size

40x40 以下の画像には画像サイズとカラーコードの文字は表示されません。

![](https://dummige.herokuapp.com/?size=40x40&color=ff00ff)

## Color

カラーコード は 3 桁（0fa）または 6 桁（00ffaa）で指定できます（#不要）
3、6 桁以外のカラーコードは無効です。

#0fa

![](https://dummige.herokuapp.com/?size=40x40&color=0fa)

#00ffaa

![](https://dummige.herokuapp.com/?size=40x40&color=00ffaa)
