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

```
position startpos moves 2h4h 9a9b 6i6h 1c1d 4g4f 3c3d 1g1f 5a5b 3i3h 8c8d 7i7h 3a4b 5i5h 4a3a 2i1g 7a6b 5h6i 2b4d 4i3i 3d3e 4h5h 5b5a 5h5i 4b3c 5g5f 3c3d 9i9h 8b8c 8g8f 3a2b 6i5h 4d3c 3i4h 3c2d 1i1h 2a1c 4h4g 8c8b 5h4h 2d4b 2g2f 6c6d 7g7f 9c9d 4h3i 1d1e 8h6f 1e1f 5f5e 1f1g+
```

```
position startpos moves 5i4h 6a5b 3i3h 5a4b 4i3i 4c4d 5g5f 8b9b 2h1h 2c2d 6i5h 5b4c 4h5g 4b5b 5g6f 7c7d 1g1f 1c1d 9i9h 3a3b 4g4f 2a1c 2g2f 9c9d 3h2g 8a9c 7i6h 5b4b 8h7i 6c6d 7i8h 2b3a 5h5g 9d9e 5f5e 4c3d 1h5h 4b4c 6f7f 4a5b 5h4h 9b8b 7f8f 7d7e 8f7e 8c8d 7e6d 8b6b 6d7d 1c2e 2f2e
```

```
position startpos moves 5g5f 4a5b 7i6h 5c5d 7g7f 9a9b 8h7g 3c3d 2g2f 5a4b 2h4h 7c7d 8g8f 1c1d 4g4f 6a5a 5i5h 8b7b 4h2h 7b6b 2h3h 1d1e 3h2h 4b3b 5h4h 2a1c 3g3f 5a4a 4i5h 1e1f 1g1f 6b8b 5h5g 3b2a 4h5h 5b5a 6g6f 8b4b 6f6e 2a1b 1i1g 4b3b 3i3h 3b4b 6i7h 8a7c 7g8h
```
