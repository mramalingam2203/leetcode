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

func slope(coordinates [][]int) float64 {
	return float64(coordinates[1][1] - coordinates[0][1]/coordinates[1][0] - coordinates[0][0])
}

func slopesCheck(points [][]int) float64 {
	for i := 0; i < len(points); i++ {
		slope(points[i])
	}
	return 0.0
}

func maxPoints(points [][]int) int {

	return 0
}

func main() {

	points := make([][]int, 3)
	for i := 0; i < 3; i++ {
		points[i] = make([]int, 2)
	}
	points[0][0] = 1
	points[0][1] = 2
	points[1][0] = 4
	points[1][1] = 3
	points[2][0] = 5
	points[2][1] = 2
	fmt.Print(points)
	slopesCheck(points)

}
