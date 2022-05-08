/*
 * File: compute_lambda.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "compute_lambda.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : double workspace[32]
 *                struct_T *solution
 *                const b_struct_T *objective
 *                const e_struct_T *qrmanager
 * Return Type  : void
 */
void compute_lambda(double workspace[32], struct_T *solution,
                    const b_struct_T *objective, const e_struct_T *qrmanager)
{
  int ia;
  int iac;
  int idx;
  int j;
  int nActiveConstr_tmp;
  nActiveConstr_tmp = qrmanager->ncols;
  if (qrmanager->ncols > 0) {
    double c;
    int idxQR;
    boolean_T guard1 = false;
    guard1 = false;
    if (objective->objtype != 4) {
      boolean_T nonDegenerate;
      c = 100.0 * (double)qrmanager->mrows * 2.2204460492503131E-16;
      if ((qrmanager->mrows > 0) && (qrmanager->ncols > 0)) {
        nonDegenerate = true;
      } else {
        nonDegenerate = false;
      }
      if (nonDegenerate) {
        boolean_T guard2 = false;
        idx = nActiveConstr_tmp;
        guard2 = false;
        if (qrmanager->mrows < qrmanager->ncols) {
          idxQR = qrmanager->mrows + qrmanager->ldq * (qrmanager->ncols - 1);
          while ((idx > qrmanager->mrows) &&
                 (fabs(qrmanager->QR[idxQR - 1]) >= c)) {
            idx--;
            idxQR -= qrmanager->ldq;
          }
          nonDegenerate = (idx == qrmanager->mrows);
          if (nonDegenerate) {
            guard2 = true;
          }
        } else {
          guard2 = true;
        }
        if (guard2) {
          idxQR = idx + qrmanager->ldq * (idx - 1);
          while ((idx >= 1) && (fabs(qrmanager->QR[idxQR - 1]) >= c)) {
            idx--;
            idxQR = (idxQR - qrmanager->ldq) - 1;
          }
          nonDegenerate = (idx == 0);
        }
      }
      if (!nonDegenerate) {
        solution->state = -7;
      } else {
        guard1 = true;
      }
    } else {
      guard1 = true;
    }
    if (guard1) {
      int ldq;
      ldq = qrmanager->ldq;
      if (qrmanager->mrows != 0) {
        memset(&workspace[0], 0, nActiveConstr_tmp * sizeof(double));
        idxQR = 0;
        idx = qrmanager->ldq * (qrmanager->ncols - 1) + 1;
        for (iac = 1; ldq < 0 ? iac >= idx : iac <= idx; iac += ldq) {
          c = 0.0;
          j = (iac + qrmanager->mrows) - 1;
          for (ia = iac; ia <= j; ia++) {
            c += qrmanager->Q[ia - 1] * objective->grad[ia - iac];
          }
          workspace[idxQR] += c;
          idxQR++;
        }
      }
      for (j = nActiveConstr_tmp; j >= 1; j--) {
        idxQR = (j + (j - 1) * ldq) - 1;
        workspace[j - 1] /= qrmanager->QR[idxQR];
        for (iac = 0; iac <= j - 2; iac++) {
          idx = (j - iac) - 2;
          workspace[idx] -= workspace[j - 1] * qrmanager->QR[(idxQR - iac) - 1];
        }
      }
      for (idx = 0; idx < nActiveConstr_tmp; idx++) {
        solution->lambda[idx] = -workspace[idx];
      }
    }
  }
}

/*
 * File trailer for compute_lambda.c
 *
 * [EOF]
 */
