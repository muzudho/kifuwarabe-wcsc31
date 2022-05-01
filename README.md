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

打ち歩詰めテスト:  

```
usi
isready
usinewgame
position startpos moves 6i5h 6c6d 3i3h 7a6b 9i9h 1c1d 7g7f 3a3b 7i6h 8b9b 5i4h 5a5b 2g2f 1a1c 4h5i 6b6c 5i6i 9b7b 6i7i 2b3a 1g1f 5b6b 7i6i 6a5b 6i7i 6b7a 3h2g 3a2b 2g3f 2b1a 1i1h 5b6b 1h1g 5c5d 3f2e 1d1e 1f1e 1c1e 1g1e 4a4b L*8d 8c8d 2i1g P*1h 2h1h L*2h 1h2h 9c9d L*6f 3c3d 2e3d 3b3c 3d4c 4b4c 6f6d 6c6d 3g3f L*1h 2h1h S*7h 7i7h 3c3d S*8b 7b8b 4i5i S*3g 8i7g 3g2f L*3h P*6f 1e1d 2f3e 3f3e 9a9c 3e3d 8b7b 6g6f 1a6f S*1b P*3a 1b2c 6f4d S*8b 7b8b 1g2e 8b7b 7g8e 8d8e 8h4d 4c4d B*6i S*8h 7h8h B*9i 8h9i 4d3d S*3g 3d2e B*4c P*6f 4c2e N*8c G*6a 6b6a 5h4h G*8i 9i8i N*5b G*8b 7b8b P*3c G*8h 8i8h 2a1c 1d1c 9d9e N*6c 7a6b P*1e 6b6c 8h7g 6f6g+ 6h6g 9c9d G*9c 8a9c 2e1f 6c5c 7g6f G*6e 6f7g 6e7f 7g7f 6d7e 7f6e P*6d 6e5f 5b4d 5f4f N*5b G*4c 5c6c 4c4d 5b4d N*2f G*5e 4f3e 6a6b 3e4d 6b5c 4d3d 5c5b 3d4d 5e4e 4d3d 4e3e 3d4d 3e2f 3g2f N*4e 4d4e 6c6b 4e5d 7e6f 6g6f 8e8f 8g8f 5b5c 5d4e 8b7b G*6a 6b6c S*4c 5c4c N*4d 4c4d 4e4f S*4e 4f3g N*8b N*4f 4e3f 3g2h 3f4g 6i7h P*2d P*7g 4g4h 7h6g 4h5i 6g5f G*2i 2h2i G*2h 2i2h 5i4h G*2g 4h5g G*6b 6c5c 6f5g 7b6b S*5d 4d5d 4f5d 5c5d G*4e 5d6c P*5d 6b6a 2h1i G*2i 1i2i S*8d G*7a 6a7a P*6g N*7b P*2b G*5a P*4g G*2h 2i2h 7a9a G*7e 8d7e 2h2i 7e8f 1h1g G*4a P*8i P*8h 9g9f 8h8i+ 3h3e 8f7g P*7f 7g7h P*8h 8i8h 2i3i P*4c 9f9e P*8e 5d5c+ 6c5c 5f6e P*5b P*5d 5c4b 5d5c+ 5b5c 9h9g 7h6g 2c3b 4a3b 3c3b+ 4b5b G*7a 9d9e 2g3g 9a7a 9g9e 8c9e L*7d G*3h 3g3h 6d6e 7d7c P*9a 3b3a P*1a 7c7b+ S*3g 2f3g L*1h 7b7a 6g7f G*2a 5b4b S*3d B*9i R*9g 4c4d 1g1h 9a9b 9g9e 4b5b 4e4d 8h9h 4d5c 5b5c L*2h G*2i 3i4i P*4a P*6i 9h9g 2h2f P*3b 1e1d P*7b P*8f P*5b 3a3b 2i3i 4i3i 9i3c 3d3c 9g8g B*4d 5c6d 2f2d 6d7e G*6h 8g7h N*4i 7h6h N*9f G*2i 3i2i 6h5i G*6d 5b5c 6d5c 5i4i P*5h 7b7c 3g2f 7c7d 8f8e 4a4b 1f2e N*4e P*7b 9c8e 4d9i 4e5g+ 3h3i 4i3i 2i3i G*4h 3i2h S*3i 2h2i P*2h 2i1i 2h2i+ 1i2i 4h3h 2i3h 5g4h 3h2g 5a4a 2g1f 4a5b G*3d 5b6c P*8d 6c5c P*3h G*1g 2f1g 5c6d G*2h 3i2h 1h1i G*1e 1f2g 4h4i 2g3f 1e2e 3f2e B*2g 9i8h 2h1i G*2f R*6c P*9c 2g1f 2e1e 1f3d 4g4f G*2e 2f2e 3d2e G*3g G*6h 8h4d 6h6g 3g2g 6g5h 8d8c+ 5h5g P*5i P*8g 1e2e 6c4c B*6a 4c6c 3e3d 4b4c P*8h 4c4d 2g3g B*5h 2e1f 4i5i P*5e 5i6i 1g2h 5h1d 1f1e 9b9c 9e8e 7e8e N*9g 8e9d 8c8d 9d9e 1e1d R*1g 1d2e P*3a B*1f 8b9d 3c4d P*4c 8d8e 7f8e 2h1i 4c4d S*8f 9e8f 8h8g 8f7e 2e1d 6d5e 3b3c 1g1f 1d2e S*2h 2e1f B*1h R*3f S*1e 1f2e 2h1i P*1f 1e2d 2e2f S*2e 2f1g 1i2h 1g2h 1h3f 9f8d P*5h 9g8e P*8a 3g3f L*4c P*5c P*2f B*8h R*1i 8h5e 2f2g+ 2h1i 2g2h 1i2h 8a8b G*6f 7e8e R*5i P*2g 2h2i 2g2h+ 2i2h 2e1f S*9f 8e9f P*9g 9f9g 5e4d S*2g 2h3g N*2e 3g2f 1f1g 2f2g P*2f 2g1f 6i7i 2a1a
```
