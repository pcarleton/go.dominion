package dominion

import (
  "testing"
  "sort"
)


func TestScoreSorter(t *testing.T) {

  scores := make([]Score, 3)
  scores[0] = Score{VP:20}
  scores[1] = Score{VP:10}
  scores[2] = Score{VP:30}

  sort.Sort(&scoreSorter{ss:scores})
  if scores[0].VP != 10 || scores[1].VP != 20 || scores[2].VP != 30 {
    t.Errorf("Order not as expected %+v", scores)
  }

}
