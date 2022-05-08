/*
 * File: computeFval_ReuseHx.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "computeFval_ReuseHx.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"

/* Function Definitions */
/*
 * Arguments    : const b_struct_T *obj
 *                double workspace[32]
 *                const double x[4]
 * Return Type  : double
 */
double computeFval_ReuseHx(const b_struct_T *obj, double workspace[32],
                           const double x[4])
{
  static const signed char f[3] = {2, -3, 1};
  double val;
  int k;
  val = 0.0;
  switch (obj->objtype) {
  case 5:
    val = obj->gammaScalar * x[obj->nvar - 1];
    break;
  case 3: {
    if (obj->hasLinear) {
      int ixlast;
      ixlast = obj->nvar;
      for (k = 0; k < ixlast; k++) {
        workspace[k] = 0.5 * obj->Hx[k] + (double)f[k];
      }
      if (obj->nvar >= 1) {
        ixlast = obj->nvar;
        for (k = 0; k < ixlast; k++) {
          val += x[k] * workspace[k];
        }
      }
    } else {
      if (obj->nvar >= 1) {
        int ixlast;
        ixlast = obj->nvar;
        for (k = 0; k < ixlast; k++) {
          val += x[k] * obj->Hx[k];
        }
      }
      val *= 0.5;
    }
  } break;
  case 4: {
    int maxRegVar_tmp;
    maxRegVar_tmp = obj->maxVar - 1;
    if (obj->hasLinear) {
      int ixlast;
      ixlast = obj->nvar;
      for (k = 0; k < ixlast; k++) {
        workspace[k] = f[k];
      }
      ixlast = obj->maxVar - obj->nvar;
      for (k = 0; k <= ixlast - 2; k++) {
        workspace[obj->nvar + k] = obj->rho;
      }
      for (k = 0; k < maxRegVar_tmp; k++) {
        workspace[k] += 0.5 * obj->Hx[k];
      }
      if (maxRegVar_tmp >= 1) {
        ixlast = obj->maxVar;
        for (k = 0; k <= ixlast - 2; k++) {
          val += x[k] * workspace[k];
        }
      }
    } else {
      int ixlast;
      if (maxRegVar_tmp >= 1) {
        ixlast = obj->maxVar;
        for (k = 0; k <= ixlast - 2; k++) {
          val += x[k] * obj->Hx[k];
        }
      }
      val *= 0.5;
      ixlast = obj->nvar + 1;
      for (k = ixlast; k <= maxRegVar_tmp; k++) {
        val += x[k - 1] * obj->rho;
      }
    }
  } break;
  }
  return val;
}

/*
 * File trailer for computeFval_ReuseHx.c
 *
 * [EOF]
 */
