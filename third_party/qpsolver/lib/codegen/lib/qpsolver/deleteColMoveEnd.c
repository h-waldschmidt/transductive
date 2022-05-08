/*
 * File: deleteColMoveEnd.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "deleteColMoveEnd.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "xrotg.h"

/* Function Definitions */
/*
 * Arguments    : e_struct_T *obj
 *                int idx
 * Return Type  : void
 */
void deleteColMoveEnd(e_struct_T *obj, int idx)
{
  double c;
  double s;
  double temp;
  int b_k;
  int i;
  int k;
  if (obj->usedPivoting) {
    i = 1;
    while ((i <= obj->ncols) && (obj->jpvt[i - 1] != idx)) {
      i++;
    }
    idx = i;
  }
  if (idx >= obj->ncols) {
    obj->ncols--;
  } else {
    int b_i;
    int ix;
    obj->jpvt[idx - 1] = obj->jpvt[obj->ncols - 1];
    b_i = obj->minRowCol;
    for (k = 0; k < b_i; k++) {
      obj->QR[k + obj->ldq * (idx - 1)] =
          obj->QR[k + obj->ldq * (obj->ncols - 1)];
    }
    obj->ncols--;
    ix = obj->mrows;
    i = obj->ncols;
    if (ix <= i) {
      i = ix;
    }
    obj->minRowCol = i;
    if (idx < obj->mrows) {
      double c_temp_tmp;
      int b_temp_tmp;
      int endIdx;
      int idxRotGCol;
      int n;
      int temp_tmp;
      ix = obj->mrows - 1;
      endIdx = obj->ncols;
      if (ix <= endIdx) {
        endIdx = ix;
      }
      k = endIdx;
      idxRotGCol = obj->ldq * (idx - 1);
      while (k >= idx) {
        b_i = k + idxRotGCol;
        temp = obj->QR[b_i];
        xrotg(&obj->QR[(k + idxRotGCol) - 1], &temp, &c, &s);
        obj->QR[b_i] = temp;
        obj->QR[k + obj->ldq * (k - 1)] = 0.0;
        i = k + obj->ldq * idx;
        n = obj->ncols - idx;
        if (n >= 1) {
          ix = i - 1;
          for (b_k = 0; b_k < n; b_k++) {
            temp = c * obj->QR[ix] + s * obj->QR[i];
            obj->QR[i] = c * obj->QR[i] - s * obj->QR[ix];
            obj->QR[ix] = temp;
            i += obj->ldq;
            ix += obj->ldq;
          }
        }
        b_i = obj->ldq * (k - 1);
        i = obj->ldq + b_i;
        n = obj->mrows;
        for (b_k = 0; b_k < n; b_k++) {
          temp_tmp = i + b_k;
          temp = obj->Q[temp_tmp];
          b_temp_tmp = b_i + b_k;
          c_temp_tmp = obj->Q[b_temp_tmp];
          obj->Q[temp_tmp] = c * temp - s * c_temp_tmp;
          obj->Q[b_temp_tmp] = c * c_temp_tmp + s * temp;
        }
        k--;
      }
      b_i = idx + 1;
      for (k = b_i; k <= endIdx; k++) {
        i = obj->ldq * (k - 1);
        ix = k + i;
        temp = obj->QR[ix];
        xrotg(&obj->QR[(k + i) - 1], &temp, &c, &s);
        obj->QR[ix] = temp;
        i = k * (obj->ldq + 1);
        n = obj->ncols - k;
        if (n >= 1) {
          ix = i - 1;
          for (b_k = 0; b_k < n; b_k++) {
            temp = c * obj->QR[ix] + s * obj->QR[i];
            obj->QR[i] = c * obj->QR[i] - s * obj->QR[ix];
            obj->QR[ix] = temp;
            i += obj->ldq;
            ix += obj->ldq;
          }
        }
        ix = obj->ldq * (k - 1);
        i = obj->ldq + ix;
        n = obj->mrows;
        for (b_k = 0; b_k < n; b_k++) {
          temp_tmp = i + b_k;
          temp = obj->Q[temp_tmp];
          b_temp_tmp = ix + b_k;
          c_temp_tmp = obj->Q[b_temp_tmp];
          obj->Q[temp_tmp] = c * temp - s * c_temp_tmp;
          obj->Q[b_temp_tmp] = c * c_temp_tmp + s * temp;
        }
      }
    }
  }
}

/*
 * File trailer for deleteColMoveEnd.c
 *
 * [EOF]
 */
