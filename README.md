# kifuwarabe-wcsc31

第３１回世界コンピュータ将棋選手権（WCSC31）に出場したきふわらべ（＾～＾）  
WCSC31 では take13 を 1手読みで使うぜ（＾～＾）  

開発中の名前は、きふわらべ将棋２０２１（kifuwarabe-shogi2021）。リネームした（＾～＾）  

## Run

```shell
# 使っていないパッケージを、インストールのリストから削除するなら
# go mod tidy

# 自作のパッケージを更新(再インストール)したいなら
# go get -u all

go build

kifuwarabe-wcsc31
```

## Test

```shell
# 将棋所から ２枚落ち初期局面から△６二玉、▲７六歩、△３二銀と進んだ局面
position sfen lnsgkgsnl/9/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL w - 1 moves 5a6b 7g7f 3a3b

# 局面の表示(独自拡張コマンド)
pos
```

## References

* [go - 2つの異なるデータ型の多次元配列を宣言する方法](https://cloud6.net/so/go/977771)
* [Visual Studio CodeでGo言語のデバッグ環境を整える](https://qiita.com/momotaro98/items/7fbcad57a9d8488fe999)
  * [Go 1.14 にバージョンアップしたらVScodeでデバッグできない (Version of Delve is too old for this version of Go..)](https://madadou.info/2020/07/31/post-2108/)

```shell
go get -u github.com/go-delve/delve
go get -u github.com/go-delve/delve/cmd/dlv
```
