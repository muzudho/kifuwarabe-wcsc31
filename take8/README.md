# take8

玉が利きに飛び込むの止めてほしいなあ（＾～＾）  
利きを調べるかな（＾～＾）  
利きテーブル作るか～（＾～＾）  
まず、簡単な指し手は生成したぜ（＾～＾）  

## Test

```plain
position startpos moves 3i3h 4a3b 5i6h 6a5b 4g4f 5a4b 5g5f 4c4d 6g6f 5b4c 6h5i 1c1d 4i3i 2b1c 1g1f 1c4f 3i4h 5c5d 3g3f 4f2h+ 6i7h
pos

[1 -> 22 moves / Second / ? repeats]

  r  b  g  s  n  l  p
+--+--+--+--+--+--+--+
| 1| 0| 0| 0| 0| 0| 1|
+--+--+--+--+--+--+--+

  9  8  7  6  5  4  3  2  1
+--+--+--+--+--+--+--+--+--+
| l| n| s|  |  |  | s| n| l| a
+--+--+--+--+--+--+--+--+--+
|  | r|  |  |  | k| g|  |  | b
+--+--+--+--+--+--+--+--+--+
| p| p| p| p|  | g| p| p|  | c
+--+--+--+--+--+--+--+--+--+
|  |  |  |  | p| p|  |  | p| d
+--+--+--+--+--+--+--+--+--+
|  |  |  |  |  |  |  |  |  | e
+--+--+--+--+--+--+--+--+--+
|  |  |  | P| P|  | P|  | P| f
+--+--+--+--+--+--+--+--+--+
| P| P| P|  |  |  |  | P|  | g
+--+--+--+--+--+--+--+--+--+
|  | B| G|  |  | G| S| b|  | h
+--+--+--+--+--+--+--+--+--+
| L| N| S|  | K|  |  | N| L| i
+--+--+--+--+--+--+--+--+--+

        R  B  G  S  N  L  P
      +--+--+--+--+--+--+--+
      | 0| 0| 0| 0| 0| 0| 0|
      +--+--+--+--+--+--+--+

# 28 にある駒の利きを調べようとするが、そこに駒は無いぜ（＾～＾）
position startpos moves 3i3h 4a3b 5i6h 6a5b 4g4f 5a4b 5g5f 4c4d 6g6f 5b4c 6h5i 1c1d 4i3i 2b1c 1g1f 1c4f 3i4h 5c5d 3g3f 4f2h+ 6i7h 2h1i
go btime 71000 wtime 71000 binc 2000 winc 2000

# 長い利きの駒である 馬 が 1i に動いたとき、角と馬を同一視できてない（＾～＾）？
# PieceType の B には RB は含んでないしな（＾～＾）
```

## Test2

```plain
# 平手初期局面を sfen で指定（＾～＾）
position sfen lnsgkgsnl/1r5b1/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL b - 1
pos
control
control diff

# 玉と歩だけ（＾～＾）
position sfen 4k4/9/ppppppppp/9/9/9/PPPPPPPPP/9/4K4 b - 1
pos
control
control diff

# 玉と歩と飛だけ（＾～＾）
position sfen 4k4/1r7/ppppppppp/9/9/9/PPPPPPPPP/7R1/4K4 b - 1
pos
control
control diff

# 玉だけ（＾～＾）
position sfen 4k4/9/9/9/9/9/9/9/4K4 b - 1
pos
control
control diff

# 香を１個にして利きの様子を見たろ（＾～＾）
position sfen 4k4/9/9/9/9/9/9/9/4K3L b - 1
pos
control diff

# 歩を置いたろ（＾～＾）
position sfen 4k4/9/9/9/9/9/8P/9/4K3L b - 1
pos
control diff

# 飛車の利きをテストするとき（＾～＾）
position sfen 4k4/9/9/5R3/2P3P2/9/9/9/4K4 b - 1
pos
control diff
do 4d4e
```

## Test3

```plain
# 歩が成られたところで、強制終了した（＾～＾） --> dropのとこ直した
position startpos moves 2g2f 1c1d 6i7h 9a9b 5i6h 4a3b 4i5h 5c5d 6g6f 7a6b 5h6g 6a7b 2f2e 2a1c 2e2d 7c7d 2d2c+

# 飛打で王手されたところで、強制終了した（＾～＾）
# 王手回避を考えてるときに、盤のUndoがちゃんとできてないようだぜ（＾～＾）
position startpos moves 7g7f 4a5b 6i7h 6a6b 7i6h 5b4b 6h7g 4b5b 5i6h 5a4a 4i5h 2c2d 6g6f 7c7d 5h6g 3a3b 9g9f 6b7c 8i9g 7a6b 6f6e 5c5d 9g8e 5b4b 8e7c+ 2d2e 7c8b 1c1d R*6a

# 逃げなかった（＾～＾）
position startpos moves 7g7f 3c3d 8h5e 2b5e 7i7h 5e9i+ 3i4h L*7d 7f7e 7d7e 4i5h B*9h 5h6h 7e7h+ 1i1h 7h6i
```

## Test4

```plain
# 王手されている局面（タダの頭金）
position sfen 4k4/9/9/9/9/9/9/4g4/4K4 b - 1
movelist
MoveList
--------
(0) 5i5h

# 詰んでいる局面（頭金）
position sfen 4k4/9/9/9/9/9/4p4/4g4/4K4 b - 1
movelist
# 指し手なし

# 王手されている局面（頭金、逃げ場所あり）
position sfen 4k4/9/9/9/9/4p4/4g4/4K4/9 b - 1
movelist
MoveList
--------
(0) 5h5i
(1) 5h6i
(2) 5h4i

# 長い利きで王手（１間、逃げ場所あり）
position sfen 4k4/9/9/9/9/4l4/9/4K4/9 b - 1
movelist
MoveList
--------
(0) 5h6g
(1) 5h4g
(2) 5h6i
(3) 5h4i
(4) 5h6h
(5) 5h4h
# 5i に逃げれないことが分かっていればOK

# 長い利きで王手（１間打、逃げ場所あり）
position sfen 4k4/9/9/9/9/9/9/4K4/9 w l 1 moves L*5f
movelist

# 桂馬で王手（逃げ場所あり、他の自駒あり）
position sfen 4k4/9/9/9/3n5/9/4K4/9/7R1 b - 1
movelist

# 桂打で王手（逃げ場所あり、他の自駒あり）
position sfen 4k4/9/9/9/9/9/4K4/9/7R1 w - 1 moves N*6e
movelist

# 杏が寄ってきているが王手ではない（逃げ場所あり、他の自駒あり）
position sfen 4k4/9/9/9/9/9/9/2+lS5/5K3 w - 1 moves 7h6i
movelist

# 杏が寄ってきているが王手ではない（逃げ場所あり、他の自駒あり）
position sfen 4k4/9/9/9/9/9/9/2+l1S4/5K3 w - 1 moves 7h6i
movelist

# 杏が寄ってきて王手（逃げ場所あり、他の自駒あり）
position sfen 4k4/9/9/9/9/9/9/2+l2G3/4K4 w - 1 moves 7h6i
movelist

# 杏が寄ってきて王手（逃げ場所あり、他の自駒あり）
# なぜか玉が逃げない（＾～＾）
position sfen 4k4/9/9/9/9/9/9/2+lS5/4K4 w - 1 moves 7h6i
movelist

# 杏が寄ってきて王手（逃げ場所あり、他の自駒あり）
position sfen 4k4/9/9/9/9/9/9/2+lS1G3/4K4 w - 1 moves 7h6i
movelist
```
