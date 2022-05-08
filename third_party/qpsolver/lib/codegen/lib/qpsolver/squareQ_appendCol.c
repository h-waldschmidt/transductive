/*
 * File: squareQ_appendCol.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "squareQ_appendCol.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "xrotg.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : e_struct_T *obj
 *                const double vec[32]
 *                int iv0
 * Return Type  : void
 */
void squareQ_appendCol(e_struct_T *obj, const double vec[32], int iv0)
{
  double c;
  double s;
  double temp_tmp;
  int ia;
  int iac;
  int idx;
  int idxRotGCol;
  int iy;
  int iyend;
  int k;
  iyend = obj->mrows;
  idxRotGCol = obj->ncols + 1;
  if (iyend <= idxRotGCol) {
    idxRotGCol = iyend;
  }
  obj->minRowCol = idxRotGCol;
  iy = obj->ldq * obj->ncols;
  idxRotGCol = obj->ldq;
  if (obj->mrows != 0) {
    iyend = iy + obj->mrows;
    if (iy + 1 <= iyend) {
      memset(&obj->QR[iy], 0, (iyend - iy) * sizeof(double));
    }
    k = obj->ldq * (obj->mrows - 1) + 1;
    for (iac = 1; idxRotGCol < 0 ? iac >= k : iac <= k; iac += idxRotGCol) {
      c = 0.0;
      iyend = (iac + obj->mrows) - 1;
      for (ia = iac; ia <= iyend; ia++) {
        c += obj->Q[ia - 1] * vec[((iv0 + ia) - iac) - 1];
      }
      obj->QR[iy] += c;
      iy++;
    }
  }
  obj->ncols++;
  obj->jpvt[obj->ncols - 1] = obj->ncols;
  for (idx = obj->mrows - 2; idx + 2 > obj->ncols; idx--) {
    idxRotGCol = obj->ldq * (obj->ncols - 1);
    k = (idx + idxRotGCol) + 1;
    temp_tmp = obj->QR[k];
    xrotg(&obj->QR[idx + idxRotGCol], &temp_tmp, &c, &s);
    obj->QR[k] = temp_tmp;
    iyend = obj->ldq * idx;
    idxRotGCol = obj->mrows;
    if (obj->mrows >= 1) {
      iy = obj->ldq + iyend;
      for (k = 0; k < idxRotGCol; k++) {
        double b_temp_tmp;
        iac = iy + k;
        temp_tmp = obj->Q[iac];
        ia = iyend + k;
        b_temp_tmp = obj->Q[ia];
        obj->Q[iac] = c * temp_tmp - s * b_temp_tmp;
        obj->Q[ia] = c * b_temp_tmp + s * temp_tmp;
      }
    }
  }
}

/*
 * File trailer for squareQ_appendCol.c
 *
 * [EOF]
 */
