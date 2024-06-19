// problem
// today you're distributing bread. Your subjects are in a line, and some of them already have some loaves. 
// Times are hard and your food stocks are dwindling, 
// so you must distribute as few loaves as possible according to the following rules:

// Every time you give a loaf of bread to some person , 
// you must also give a loaf of bread to the person immediately in front of or behind them in the line 
// (i.e., [i+1] persons  or [i-1]).
// After all the bread is distributed, each person must have an even number of loaves.
// Given the number of loaves already held by each citizen, 
// find and print the minimum number of loaves you must distribute to satisfy the two rules above. 
// If this is not possible, print NO.

// Example
// B = [4,5,6,7]

// We can first give a loaf to i=3  and  so B=[4,5,7,8].
// Next we give a loaf to i=2  and i=3 and have B=[4,6,8,8] which satisfies our conditions.
// All of the counts are now even numbers. We had to distribute 4 loaves.

// Function Description

// fairRations has the following parameter(s):
// int B[N]: the numbers of loaves each persons starts with

// Returns
// string: the minimum number of loaves required, cast as a string, or 'NO'

// Input Format
// The first line contains an integer N, the number of subjects in the bread line.

// The second line contains N space-separated integers B[i].

// Constraints
// 2<= N <= 1000
// 1 <= B[i] <= 10, where 1 <= i <= N

// Output Format

// Sample Input 0
// STDIN       Function
// -----       --------
// 5           B[] size N = 5
// 2 3 4 5 6   B = [2, 3, 4, 5, 6]   

// Sample Output 0
// 4

// Explanation 0
// The initial distribution is [2,3,4,5,6]. The requirements can be met as follows:

// Give 1 loaf of bread each to the second and third people so that the distribution becomes [2,4,5,5,6].
// Give 1 loaf of bread each to the third and fourth people so that the distribution becomes [2,4,6,6,6].
// Each of the  subjects has an even number of loaves after  loaves were distributed.

// Sample Input 1
// 2
// 1 2

// Sample Output 1
// NO
// Explanation 1

// The initial distribution is [1,2]. 
// As there are only 2 people in the line, 
// any time you give one person a loaf you must always give the other person a loaf. 
// Because the first person has an odd number of loaves and the second person has an even number of loaves, 
// no amount of distributed loaves will ever result in both subjects having an even number of loaves.

package main

import (
	
	"fmt"
	
)

/*
* Complete the 'fairRations' function below.
*
* The function is expected to return a STRING.
* The function accepts INTEGER_ARRAY B as parameter.
*/

func fairRations(B []int32) string {
	// Write your code here
	var total int32 = 0 // ini untuk total roti yang dibagi

	n := len(B)

	for i:=0; i < n-1; i++{
		if(B[i]%2 != 0){
			B[i]++
			B[i+1]++
			total +=2
		}
	}

	if B[n-1]%2 !=0{
		return "NO"
	}

	return fmt.Sprintf("%d", total)
	
}

func main() {
	fmt.Println(fairRations([]int32{2,3,4,5,6}))
}