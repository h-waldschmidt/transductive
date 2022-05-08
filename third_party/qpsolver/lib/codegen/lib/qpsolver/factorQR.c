/*
 * File: factorQR.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "factorQR.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "xzgeqp3.h"

/* Function Definitions */
/*
 * Arguments    : e_struct_T *obj
 *                const double A[32]
 *                int mrows
 *                int ncols
 *                int ldA
 * Return Type  : void
 */
void factorQR(e_struct_T *obj, const double A[32], int mrows, int ncols,
              int ldA)
{
  int idx;
  int ix0;
  int k;
  boolean_T guard1 = false;
  ix0 = mrows * ncols;
  guard1 = false;
  if (ix0 > 0) {
    for (idx = 0; idx < ncols; idx++) {
      int iy0;
      ix0 = ldA * idx;
      iy0 = obj->ldq * idx;
      for (k = 0; k < mrows; k++) {
        obj->QR[iy0 + k] = A[ix0 + k];
      }
    }
    guard1 = true;
  } else if (ix0 == 0) {
    obj->mrows = mrows;
    obj->ncols = ncols;
    obj->minRowCol = 0;
  } else {
    guard1 = true;
  }
  if (guard1) {
    obj->usedPivoting = false;
    obj->mrows = mrows;
    obj->ncols = ncols;
    for (idx = 0; idx < ncols; idx++) {
      obj->jpvt[idx] = idx + 1;
    }
    if (mrows <= ncols) {
      ix0 = mrows;
    } else {
      ix0 = ncols;
    }
    obj->minRowCol = ix0;
    obj->tau[0] = 0.0;
    obj->tau[1] = 0.0;
    obj->tau[2] = 0.0;
    obj->tau[3] = 0.0;
    if (ix0 >= 1) {
      qrf(obj->QR, mrows, ncols, ix0, obj->tau);
    }
  }
}

/*
 * File trailer for factorQR.c
 *
 * [EOF]
 */
