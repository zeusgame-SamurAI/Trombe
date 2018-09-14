package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"./debug"
)

type player struct {
	x  int
	y  int
	Vx int
	Vy int
}

var sc = bufio.NewScanner(os.Stdin)

var thinkTime int      //考慮時間
var maxStep int        //制限ステップ数
var courseW int        //コースサイズ幅
var courseL int        //コースサイズ長さ
var viewCnt int        //視界距離
var course = [][]int{} //コース
var stepNumber int     //ステップ数
var timeLeft int       //残り考慮時間
var p1 player          //プレイヤー１情報（自キャラ）
var p2 player          //プレイヤー２情報（敵キャラ）

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func GetFirstInput() {
	sc.Split(bufio.ScanWords)       //スペース区切りで文字を入力
	log.SetFlags(log.Lmicroseconds) //logをミリ秒表示

	thinkTime = nextInt()
	maxStep = nextInt()
	courseW = nextInt()
	courseL = nextInt()
	viewCnt = nextInt()

	debug.Println("考慮時間　　　　　：", thinkTime)
	debug.Println("制限ステップ数　　：", maxStep)
	debug.Println("コースサイズ幅　　：", courseW)
	debug.Println("コースサイズ長さ　：", courseL)
	debug.Println("視界距離　　　　　：", viewCnt)
	debug.Println("-----------------------------")
	course = make([][]int, courseL)
	for i := 0; i < courseL; i++ {
		course[i] = make([]int, courseW)
	}
}

func GetInput() {
	stepNumber = nextInt()
	timeLeft = nextInt()
	p1 = player{}
	p1.x = nextInt()
	p1.y = nextInt()
	p1.Vx = nextInt()
	p1.Vy = nextInt()
	p2 = player{}
	p2.x = nextInt()
	p2.y = nextInt()
	p2.Vx = nextInt()
	p2.Vy = nextInt()

	for y := 0; y < courseL; y++ {
		for x := 0; x < courseW; x++ {
			course[y][x] = nextInt()
		}
	}

	debug.Println("該当ステップ数：", stepNumber)
	debug.Println("残り考慮時間  ：", timeLeft)
	debug.Println("プレイヤー1   ：", p1.x, p1.y, p1.Vx, p1.Vy)
	debug.Println("プレイヤー2   ：", p2.x, p2.y, p2.Vx, p2.Vy)
	debug.Println("コース        ：")
	for y := 0; y < len(course); y++ {
		for x := 0; x < len(course[y]); x++ {
			debug.Print(course[y][x])
		}
		debug.Print("\n")
	}
	debug.Println("-----------------------------")
}

func main() {
	GetFirstInput()
	fmt.Println(0)

	for true {
		//ゲームからの入力を受け取る。
		GetInput()

		//ゴール一覧を作る
		var goal_y int
		if p1.y > p2.y {
			goal_y = 2 + 1
		} else {
			goal_y = p2.y + viewCnt
		}
		debug.Println("今、一番見えるコースの最深部距離：", goal_y)

		goal_x := []int{}
		for x := 0; x < courseW; x++ {
			if course[goal_y][x] != 2 {
				goal_x = append(goal_x, x)
			}
		}
		debug.Println("最深部の山以外の位置:", goal_x)

		//各ゴールごとの最短距離（最短カーブ数）を求める
		route := [][][]int{}
		for routeIdx := 0; routeIdx > 0; routeIdx++ {
			route[routeIdx] = [][]int{}
			for routeStep := 0; routeStep > 0; routeStep++ {
				route[routeIdx][routeStep] = [2]int{}
			}
		}

		//最短で行けるゴールを決める

		//ゴールまでの行き方を求める

		fmt.Println(0, 1)
	}
}
