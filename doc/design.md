# Design

## Square number

![20210627shogi38.png](./img/20210627shogi38.png)  
![20210627shogi39.png](./img/20210627shogi39.png)  

* `A` - Board
* `B` - Hand

## Hand piece type

![20210627shogi40.png](./img/20210627shogi40.png)  

## l03.Move

![20210627shogi33a1.png](./img/20210627shogi33a1.png)  

* `A` - 1～7bit: 移動元(0～127)
* `B` - 8～14bit: 移動先(0～127)
* `C` - 15bit: 成(0～1)
* `move = 0` - 投了
