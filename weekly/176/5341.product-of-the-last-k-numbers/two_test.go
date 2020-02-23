// Package 5341.product-of-the-last-k-numbers - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* two_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/16 10:21 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _341_product_of_the_last_k_numbers

import (
	"fmt"
	"testing"
)

func TestProductOfNumbers(t *testing.T) {
	obj := Constructor()
	obj.Add(7)
	fmt.Println(obj.GetProduct(1))
	fmt.Println(obj.GetProduct(1))
	obj.Add(4)
	obj.Add(5)
	fmt.Println(obj.GetProduct(3))
	obj.Add(4)
	fmt.Println(obj.GetProduct(4))
	obj.Add(3)
	fmt.Println(obj.GetProduct(4))
	obj.Add(8)
	fmt.Println(obj.GetProduct(1))
	fmt.Println(obj.GetProduct(6))
	obj.Add(2)
	fmt.Println(obj.GetProduct(3))
}
