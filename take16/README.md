# take16

利きボードを使った 利き分布を利用した評価関数を作ってる途中だぜ（＾～＾）  

次は深さ２を目指すかだぜ（＾～＾）  
１手詰めも欲しいけど（＾～＾）  
処理時間を計測したいぜ（＾～＾）  
ベータカットも欲しいなあ（＾～＾）  
moveのうち、移動先と成りだけの move_end を作って genmove で返そ（＾～＾）  

# Build

```shell
go build
```

# Run

```shell
kifuwarabe-wcsc31 lesson16
```

## 計測

```plain
usi
position startpos

dev
playout
# depthEnd=2のとき 36.898604 seconds
```

## Test

```plain
position startpos
pos
control layer 0
control layer 10

# 先手から見て
watercolor 0 10 26 27 28
control layer 26
control layer 28

# 後手から見て
watercolor 10 0 26 27 28
control layer 26
control layer 28
```
