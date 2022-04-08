//! Position と Record を疎結合にするための仕掛け。両方から参照されるもの（＾～＾）
package take16

// 指し手
//
// 15bit で表せるはず（＾～＾）
// .pdd dddd dsss ssss
//
// 1～7bit: 移動元(0～127)
// 8～14bit: 移動先(0～127)
// 15bit: 成(0～1)
type Move uint16
