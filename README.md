# kifuwarabe-wcsc31

第３１回世界コンピュータ将棋選手権（WCSC31）に出場したきふわらべ（＾～＾） いわゆる　きふわらご（＾～＾）  
WCSC31 では take13 を 1手読みで使うぜ（＾～＾）  

開発中の名前は、きふわらべ将棋２０２１（kifuwarabe-shogi2021）。リネームした（＾～＾）  

2022年は take17（＾～＾）  

# Deploy

## Development

例えば 開発環境では、ソースは以下のディレクトリに置いてある（＾～＾）  

```
📂C:\Users\むずでょ\go\src\github.com\muzudho\kifuwarabe-wcsc31  
           -------
           1
1. ユーザー名
```

## Runtime

わたしのフォルダー構成

```plain
💻大会用PC
└───📂C:\Users\{ユーザー名}\Documents\MyProduct
    └───📂KifuwarabeWcsc31
        ├───📂input
        ├───📂output
        ├───📂shogidokoro # 将棋所をダウンロードしてきてここに置く
        └───📄kifuwarabe-wcsc31.exe
```

# Lesson

* [Lesson01](./lesson01/README.md)
* [Lesson02](./lesson02/README.md)
# Build

```shell
# 使っていないパッケージを、インストールのリストから削除するなら
# go mod tidy

# 自作のパッケージを更新(再インストール)したいなら
# go get -u all

go build
```

# Run

```shell
kifuwarabe-wcsc31
```

# Test

```shell
# 将棋所から ２枚落ち初期局面から△６二玉、▲７六歩、△３二銀と進んだ局面
position sfen lnsgkgsnl/9/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL w - 1 moves 5a6b 7g7f 3a3b

# 局面の表示(独自拡張コマンド)
pos
```

# References

* [go - 2つの異なるデータ型の多次元配列を宣言する方法](https://cloud6.net/so/go/977771)
* [Visual Studio CodeでGo言語のデバッグ環境を整える](https://qiita.com/momotaro98/items/7fbcad57a9d8488fe999)
  * [Go 1.14 にバージョンアップしたらVScodeでデバッグできない (Version of Delve is too old for this version of Go..)](https://madadou.info/2020/07/31/post-2108/)

```shell
go get -u github.com/go-delve/delve
go get -u github.com/go-delve/delve/cmd/dlv
```

# TODO

* [ ] 長い利きボードが、相手玉を王手したかどうか　カウントできるだろうか（＾～＾）？  
利きが伸びるタイミング、利きが遮られるタイミングがあると思う（＾～＾）そこで（＾～＾）

# Documents

[Design](./doc/design.md)  
[Test](./doc/test.md)  

# Debug

```
position startpos moves 5i5h 5a5b 7i7h 8b9b 6i7i 2c2d 1g1f 6c6d 2h1h 5b5a 3i2h 9c9d 4i3i 9b3b 2i1g 4a4b 9g9f 6a5b 1f1e 4c4d 1g2e 2d2e 5h4i
```