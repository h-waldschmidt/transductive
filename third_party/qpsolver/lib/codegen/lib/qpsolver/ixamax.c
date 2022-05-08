/*
 * File: ixamax.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "ixamax.h"
#include "rt_nonfinite.h"
#include <math.h>

/* Function Definitions */
/*
 * Arguments    : int n
 *                const double x[16]
 *                int incx
 * Return Type  : int
 */
int ixamax(int n, const double x[16], int incx)
{
  int idxmax;
  int k;
  if ((n < 1) || (incx < 1)) {
    idxmax = 0;
  } else {
    idxmax = 1;
    if (n > 1) {
      double smax;
      smax = fabs(x[0]);
      for (k = 2; k <= n; k++) {
        double s;
        s = fabs(x[(k - 1) * incx]);
        if (s > smax) {
          idxmax = k;
          smax = s;
        }
      }
    }
  }
  return idxmax;
}

/*
 * File trailer for ixamax.c
 *
 * [EOF]
 */
