/*
 * File: computeFval.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "computeFval.h"
#include "linearForm_.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"

/* Function Definitions */
/*
 * Arguments    : const b_struct_T *obj
 *                double workspace[32]
 *                const double x[4]
 * Return Type  : double
 */
double computeFval(const b_struct_T *obj, double workspace[32],
                   const double x[4])
{
  double val;
  int idx;
  int k;
  val = 0.0;
  switch (obj->objtype) {
  case 5:
    val = obj->gammaScalar * x[obj->nvar - 1];
    break;
  case 3: {
    linearForm_(obj->hasLinear, obj->nvar, workspace, x);
    if (obj->nvar >= 1) {
      int ixlast;
      ixlast = obj->nvar;
      for (k = 0; k < ixlast; k++) {
        val += x[k] * workspace[k];
      }
    }
  } break;
  case 4: {
    int ixlast;
    linearForm_(obj->hasLinear, obj->nvar, workspace, x);
    ixlast = obj->nvar + 1;
    k = obj->maxVar - 1;
    for (idx = ixlast; idx <= k; idx++) {
      workspace[idx - 1] = 0.5 * obj->beta * x[idx - 1] + obj->rho;
    }
    if (k >= 1) {
      ixlast = obj->maxVar;
      for (k = 0; k <= ixlast - 2; k++) {
        val += x[k] * workspace[k];
      }
    }
  } break;
  }
  return val;
}

/*
 * File trailer for computeFval.c
 *
 * [EOF]
 */
