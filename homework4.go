package main

import (
	"fmt"
	"time"
)


func renxingxueba_aike(input map[string]int){
	var tm1,tm2,tm3,tm4,tm5 int64
	var result string
	fmt.Scan(&tm1,&tm2,&tm3,&tm4,&tm5,&result)
	t1 := time.Unix(int64(tm1), 0)
	t2 := time.Unix(int64(tm2), 0)
	t3 := time.Unix(int64(tm3), 0)
	t4 := time.Unix(int64(tm4), 0)
	t5 := time.Unix(int64(tm5), 0)
	T1:= t1.Format("2006-01-02 15:04:05")
	T2:= t2.Format("2006-01-02 15:04:05")
	T3:= t3.Format("2006-01-02 15:04:05")
	T4:= t4.Format("2006-01-02 15:04:05")
	T5:= t5.Format("2006-01-02 15:04:05")
	input[T1]=1
	input[T2]=2
	input[T3]=3
	input[T4]=4
	input[T5]=5
}

func main() {
	var t string
	m:=make(map[string]int)
	renxingxueba_aike(m)
	for k := range m {
		if _, ok := m[k];ok{
			println("input ok")
		}
	}
	println("输入777来触发时空断裂")
	fmt.Scan(&t)
	switch {
	case t=="777":
		fmt.Println("the result are :")
		for v := range m {
			fmt.Println(v)
		}
	default:
		println("赛文老祖默默离开")
	}
}
