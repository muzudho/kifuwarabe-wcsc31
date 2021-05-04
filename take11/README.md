# take11

駒得評価関数は入れたしな（＾～＾）  
きふわらべがどう評価しているか、infoも入れたし（＾～＾）  
利きの差分表も 飛、角、香 を分けた（＾～＾）  
テスト用のランダム局面を増やすために、シャッフルとプレイアウトも入れた（＾～＾）  
王手回避のときに打をして、間駒もするはずだぜ（＾～＾）  

玉は逃げてくれてるようだし、アルファベータ探索入れてみよかな（＾～＾）  
探索中に盤面が元に戻ってないことがあるなあ（＾～＾）盤面２つに増やすかな（＾～＾）？  

## Test

```plain
# 強制終了した（＾～＾）
position startpos moves 1g1f 4a3b 1f1e 6a5b 2g2f 2c2d 3g3f 5a4b 3f3e 4c4d 4g4f 5b4c 5g5f 7c7d 5f5e 7d7e 6g6f 9c9d 6f6e 9d9e 8g8f 8c8d 2h4h 5c5d 1e1d 1c1d 3e3d 3c3d 5e5d 4c5d 6e6d 5d6d 4h6h 8d8e 2f2e 2d2e 4f4e 8e8f 4e4d 8f8g+ 4d4c 3b4c 7g7f 8g8h 7f7e 8h9i 7e7d 8b8i+ 7d7c N*5g 7c7b 2b7g+
# go btime 86000 wtime 83000 binc 2000 winc 2000
# Error: 59 square is empty
# quit

# 盤をコピー（＾～＾）
board copy 0 1

# 盤[1]を表示（＾～＾）
pos 1

# 盤[0], [1] の比較
posdiff 0 1

board diff

# 盤[2], [3] の比較
posdiff 2 3

# 盤[0], [1] の異なる箇所の数
error board 0 1 2 3

location 0
location 1
```

## Test2

```shell
# 後手馬で先手飛を取って、アンドゥするテスト（＾～＾）
position sfen 4k4/9/9/9/9/9/2+b6/3R5/4K4 w - 1
pos
do 7g6h
pos
location 0
undo
pos
location 0
```

## Test3

break文抜け修正（＾～＾）

```shell
position startpos moves 7g7f 3c3d
pos
location 0
do 8h2b
pos
location 0
undo
location 0
```

## Test4

ソート入れて修正（＾～＾）  

```shell
# position startpos moves 1g1f 4a3b 1f1e 6a5b 2g2f 2c2d 3g3f 5a4b 3f3e 4c4d 4g4f 5b4c 5g5f 7c7d 5f5e 7d7e 6g6f 9c9d 6f6e 9d9e 8g8f 8c8d 2h4h 5c5d 1e1d 1c1d
position startpos moves 1g1f 4a3b 1f1e 6a5b 2g2f 2c2d 3g3f 5a4b 3f3e 4c4d 4g4f 5b4c 5g5f 7c7d 5f5e 7d7e 6g6f 9c9d 6f6e 9d9e 8g8f 8c8d 2h4h 5c5d 1e1d 1c1d 1i1d

pos

location 0

do 1a1d

pos

location 0

# 強制終了（＾～＾）
# go

undo
# pawnが減ってない？
pos
location 0

# 後手香車で、先手香車を取ると不具合（＾～＾）？
# 探索をやっていると、位置が変わることがあるようだ（＾～＾）？
#
#  K   k      R          B          L
# +---+---+  +---+---+  +---+---+  +---+---+---+---+
# | 59| 42|  | 48| 82|  | 22| 88|  | 14| 11| 91| 99|
# +---+---+  +---+---+  +---+---+  +---+---+---+---+
#
#  K   k      R          B          L
# +---+---+  +---+---+  +---+---+  +---+---+---+---+
# | 59| 42|  | 48| 82|  | 22| 88|  | 11| 14| 91| 99|
# +---+---+  +---+---+  +---+---+  +---+---+---+---+
```

## Test5

取った駒のアンドゥ時、フェーズ 0 を return してたんで、それを除去して修正（＾～＾）

```shell
position startpos moves 7g7f 7a6b 6i7h 5c5d 7i6h 6a7a 6h7g 4a4b 5i6h 8b9b 4i5h 1c1d 3i3h 2b1c 1g1f 9c9d 1f1e 1d1e 1i1e 5d5e 1e1c 2a1c P*1d 1c2e 2g2f L*6a 2f2e P*1e N*5d 4b4a 5d6b+ 5a6b 1d1c+ 1a1c B*3e 4c4d 3e1c+ 3a3b 1c4f 9b8b 4f5e N*9e 5e4d P*5c P*1b

pos

# 強制終了（玉を取ってしまうらしい）
go

position sfen 4k4/9/9/9/9/9/9/4p4/4K4 w - 1
pos
location 0
# らいおんきゃっち（相手玉を取ったから自分の駒台に玉が乗る）
do 5h5i
pos
location 0
record
undo
# 玉が戻ってない
pos
location 0
```

## Test6

```shell
position startpos moves 5i5h 4a3b 8g8f 6a5b 2g2f 5a4b 6g6f 3c3d 5h4h 2b6f

# 強制終了（＾～＾） 後手玉の位置が、先手玉の位置になってる（＾～＾）
go

position sfen 4k4/4P4/9/9/9/9/9/9/4K4 b - 1
pos
do 5b5a
location 0
pos
undo
pos
location 0
```

## Test7

```shell
position startpos moves 7g7f 9c9d 6i7h 9d9e 7i6h 9e9f 9g9f 1c1d 9f9e 8b5b 3i4h 7a7b 9e9d 4a4b 5i6i 5b6b 6h7g 2a1c 4i5h 3c3d 3g3f 2b5e 4h3g 1d1e 1g1f 6a7a 1f1e 6b5b 5g5f 5e2b 2g2f 2b4d 1e1d 6c6d 1d1c+ 2c2d P*1b 7b6a 1b1a+ 4d1a N*2c

pos
location 0
record
count
movelist

# >2:go btime 81000 wtime 58000 binc 2000 winc 2000
```
