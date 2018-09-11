package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"./debug"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type player struct {
	x  int
	y  int
	Vx int
	Vy int
}

func main() {
	//初期化時の入力
	sc.Split(bufio.ScanWords)       //スペース区切りで文字を入力
	log.SetFlags(log.Lmicroseconds) //logをミリ秒表示

	thinkTime := nextInt()
	maxStep := nextInt()
	courseW := nextInt()
	courseL := nextInt()
	viewCnt := nextInt()

	debug.Println("考慮時間　　　　　：", thinkTime)
	debug.Println("制限ステップ数　　：", maxStep)
	debug.Println("コースサイズ幅　　：", courseW)
	debug.Println("コースサイズ長さ　：", courseL)
	debug.Println("視界距離　　　　　：", viewCnt)
	debug.Println("-----------------------------")
	course := make([][]int, courseL)
	for i := 0; i < courseL; i++ {
		course[i] = make([]int, courseW)
	}
	fmt.Println(0) //準備完了！！

	//------------------------------------------------------------
	//ステップごとの入力
	stepNumber := nextInt()
	timeLeft := nextInt()
	p1 := player{}
	p1.x = nextInt()
	p1.y = nextInt()
	p1.Vx = nextInt()
	p1.Vy = nextInt()
	p2 := player{}
	p2.x = nextInt()
	p2.y = nextInt()
	p2.Vx = nextInt()
	p2.Vy = nextInt()

	for x := 0; x < courseL; x++ {
		for y := 0; y < courseW; y++ {
			course[x][y] = nextInt()
		}
	}

	debug.Println("該当ステップ数：", stepNumber)
	debug.Println("残り考慮時間  ：", timeLeft)
	debug.Println("プレイヤー1   ：", p1.x, p1.y, p1.Vx, p1.Vy)
	debug.Println("プレイヤー2   ：", p2.x, p2.y, p2.Vx, p2.Vy)
	debug.Println("コース        ：")
	for x := 0; x < len(course); x++ {
		for y := 0; y < len(course[x]); y++ {
			debug.Print(course[x][y])
		}
		debug.Print("\n")
	}
	debug.Println("-----------------------------")

	fmt.Println(1, 0)
}
