// 流水线
package concurrent

import (
	"fmt"
	"sync"
)

// 采购
func buy(count int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- fmt.Sprintf("采购%d", i)
		}
	}()
	return out
}

// 组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := range in {
			out <- fmt.Sprintf("组装（%s）", i)
		}
	}()
	return out
}

// 扇入
func merge(ins ...<-chan string) <-chan string {
	out := make(chan string)
	wg := sync.WaitGroup{}
	pro := func(in <-chan string) {
		defer wg.Done()
		for i := range in {
			out <- fmt.Sprintf("组装（%s)", i)
		}
	}

	wg.Add(len(ins))
	for _, in := range ins {
		go pro(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := range in {
			out <- fmt.Sprintf("打包（%s）", i)
		}
	}()
	return out
}

// 生产线1
func PipelineExample1() {
	coms := buy(50)
	prods := build(coms)
	packs := pack(prods)
	for i := range packs {
		fmt.Println(i)
	}
}

// 生产线2
func PipelineExample2() {
	coms := buy(50)
	prod1 := build(coms)
	prod2 := build(coms)
	prod3 := build(coms)
	prods := merge(prod1, prod2, prod3)
	packs := pack(prods)
	for i := range packs {
		fmt.Println(i)
	}
}
