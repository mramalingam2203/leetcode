// https://leetcode.com/problems/max-points-on-a-line/
/*

   Faster Sorting based Algorithm (More Efficient one)
   Algorithm for Faster Sorting Based Algorithm

   Initally consider a set of all points S={q1,q2,……qn}

   1. Sort the points S. i.e S={q2,q1,q4…qn} arbitrary order
   2. Lets take point q as origin (q ∈ S) … i.e (q=q2);
   3. With respect to q find slope to {other points} in 2d plane.i.e(other points={S}-q)
   4. Sort the all the points with the slope they make with respect to q
   5. if there are more than 2 adjacent points together such that their slopes are equal with respect to q then the points adjacent are said to be collinear to the point q
   6. if(!all points q are take as origin)
       {
           repeat from step 2 to step 6 by taking q = next q // i.e q=q1;
       }else
       {
           exit(1);
       }
*/

package main

import (
	"fmt"
	_ "os"
)

func slope(points [][]int) float64 {
	fmt.Print(points)
	return float64(points[1][1]-points[0][1]) / float64(points[1][0]-points[0][0])
}

func maxPoints(points [][]int) int {
	return 0
}

func main() {
	points := make([][]int, 2)
	for i := 0; i < 2; i++ {
		points[i] = make([]int, 2)
	}
	points[0][0] = 1
	points[0][1] = 2
	points[1][0] = 4
	points[1][1] = 3

	fmt.Println(slope(points))

}
