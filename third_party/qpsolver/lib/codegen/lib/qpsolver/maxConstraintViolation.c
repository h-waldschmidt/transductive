/*
 * File: maxConstraintViolation.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "maxConstraintViolation.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "xgemv.h"
#include <math.h>

/* Function Definitions */
/*
 * Arguments    : f_struct_T *obj
 *                const double x[4]
 * Return Type  : double
 */
double maxConstraintViolation(f_struct_T *obj, const double x[4])
{
  double v;
  int idx;
  int mFixed;
  int mLB;
  int mUB;
  int offsetEq1;
  mLB = obj->sizes[3];
  mUB = obj->sizes[4];
  mFixed = obj->sizes[0];
  if (obj->probType == 2) {
    int mEq;
    int offsetEq2;
    v = 0.0;
    mEq = obj->sizes[1] - 1;
    for (offsetEq1 = 0; offsetEq1 <= mEq; offsetEq1++) {
      obj->maxConstrWorkspace[offsetEq1] = obj->beq;
    }
    c_xgemv(obj->nVarOrig, obj->sizes[1], obj->Aeq, obj->ldA, x,
            obj->maxConstrWorkspace);
    offsetEq1 = obj->nVarOrig + obj->sizes[2];
    offsetEq2 = offsetEq1 + obj->sizes[1];
    for (idx = 0; idx <= mEq; idx++) {
      obj->maxConstrWorkspace[idx] =
          (obj->maxConstrWorkspace[idx] - x[offsetEq1 + idx]) +
          x[offsetEq2 + idx];
      v = fmax(v, fabs(obj->maxConstrWorkspace[idx]));
    }
  } else {
    int mEq;
    v = 0.0;
    mEq = obj->sizes[1] - 1;
    for (offsetEq1 = 0; offsetEq1 <= mEq; offsetEq1++) {
      obj->maxConstrWorkspace[offsetEq1] = obj->beq;
    }
    c_xgemv(obj->nVar, obj->sizes[1], obj->Aeq, obj->ldA, x,
            obj->maxConstrWorkspace);
    for (idx = 0; idx <= mEq; idx++) {
      v = fmax(v, fabs(obj->maxConstrWorkspace[idx]));
    }
  }
  if (obj->sizes[3] > 0) {
    for (idx = 0; idx < mLB; idx++) {
      offsetEq1 = obj->indexLB[idx] - 1;
      v = fmax(v, -x[offsetEq1] - obj->lb[offsetEq1]);
    }
  }
  if (obj->sizes[4] > 0) {
    for (idx = 0; idx < mUB; idx++) {
      offsetEq1 = obj->indexUB[idx] - 1;
      v = fmax(v, x[offsetEq1] - obj->ub[offsetEq1]);
    }
  }
  if (obj->sizes[0] > 0) {
    for (idx = 0; idx < mFixed; idx++) {
      v = fmax(v, fabs(x[obj->indexFixed[idx] - 1] -
                       obj->ub[obj->indexFixed[idx] - 1]));
    }
  }
  return v;
}

/*
 * File trailer for maxConstraintViolation.c
 *
 * [EOF]
 */
