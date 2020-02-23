// Package 5341.product-of-the-last-k-numbers - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* 5341.product-of-the-last-k-numbers - file brief introduce */
/*
modification history
-----------------------------
2020/2/16 10:21 AM, by gosyang, create
*/
/*
DESCRIPTION
总结：
1. 旧方法每次add都重复乘，导致超时
2. 发现go超时，优化得从优化核心数据结构入手了，存储倒数几个数据的乘机，最终除可以减少运算
总结来说对于业务如果请求远大，那直接map，不除还是可以的。
旧方法每次乘其实也可以
*/
package _341_product_of_the_last_k_numbers

type ProductOfNumbers struct {
	//zero int
	//length int
	//kMap map[int]int
	// TODO 这里精髓，从0:1, 1:last， 2: last 2...
	NonZeroRes []int
}

func Constructor() ProductOfNumbers {
	//return ProductOfNumbers{
	//	zero: 0,
	//	length: 0,
	//	kMap: make(map[int]int),
	//}
	return ProductOfNumbers{
		NonZeroRes: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	//if num == 0 {
	//	this.zero = 1
	//	this.length = 0
	//} else {
	//	if this.zero != 0 {
	//		this.zero += 1
	//	}
	//	this.length += 1
	//}
	//for i := this.length; i > 1; i-- {
	//	this.kMap[i] = this.kMap[i-1] * num
	//}
	//this.kMap[1] = num
	if num == 0 {
		this.NonZeroRes = []int{1}
		return
	}
	this.NonZeroRes = append(this.NonZeroRes, this.NonZeroRes[len(this.NonZeroRes)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	//if this.zero != 0 && k >= this.zero {
	//	return 0
	//}
	//return this.kMap[k]
	if k >= len(this.NonZeroRes) {
		return 0
	}
	return this.NonZeroRes[len(this.NonZeroRes)-1] / this.NonZeroRes[len(this.NonZeroRes)-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
