# take9

玉は逃げてくれてるようだし、アルファベータ探索入れてみよかな（＾～＾）  
その前に駒得評価関数入れてみるかな（＾～＾）？  
駒が文字列なのは気になるな、定数にするかな（＾～＾）  

## Test1

```plain
# 何が反則手なのか（＾～＾）？
position startpos moves 7g7f 6c6d 6i7h 6a5b 7i6h 1c1d 6h7g 7a6b 5i6h 4a3b 4i5h 1d1e 6g6f 5a4a 5h6g 6b6c 7f7e 8b9b 3i3h 4c4d 6g5f 9b7b 5f5e 1a1d 5e4d 7b7a 9g9f 3c3d 4d3d 7a7b 3d3e 3b4b 3e2e 3a3b 2e1d 2b5e L*4f 4a3a 4f4b+ 3a2b G*4e 5e4f 4b5b 7b5b 4g4f
# go btime 83000 wtime 82000 binc 2000 winc 2000
# 反則手
# bestmove L+3e
```

## Test2

```plain
# 持ち駒がいっぱいある局面（＾～＾）
position sfen 4k4/9/9/9/9/9/9/9/4K4 b RBG2S2N2L2P9rbg2s2n2l2p9 1
```
