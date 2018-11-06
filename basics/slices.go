package main

import "golang.org/x/tour/pic"

func Pic(dx,dy int) [][]uint8 {

	//結果格納用変数( => 長さdyのslice)
	val := make([][]uint8,dy)

	//valの長さ分ぶん回している
	for y := range val {
		//val[y]の中にdx分のsliceを作成
		val[y] = make([]uint8,dx)
		//作成したval[y]内のdx分ぶん回す
		for x:= range val[y]{
			//二次元sliceの１要素ごとに描画するものを突っ込んでいるはず？
			val[y][x] = uint8((x+y)/2)
		}
	}
	return val
}

func main() {
	pic.Show(Pic)
}
