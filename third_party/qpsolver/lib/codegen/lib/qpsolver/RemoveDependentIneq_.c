/*
 * File: RemoveDependentIneq_.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "RemoveDependentIneq_.h"
#include "countsort.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "xzgeqp3.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : f_struct_T *workingset
 *                e_struct_T *qrmanager
 *                d_struct_T *memspace
 *                double tolfactor
 * Return Type  : void
 */
void RemoveDependentIneq_(f_struct_T *workingset, e_struct_T *qrmanager,
                          d_struct_T *memspace, double tolfactor)
{
  int idx;
  int idx_col;
  int k;
  int nDepIneq;
  int nFixedConstr;
  int nVar_tmp_tmp;
  nDepIneq = workingset->nActiveConstr;
  nFixedConstr = workingset->nWConstr[0] + workingset->nWConstr[1];
  nVar_tmp_tmp = workingset->nVar;
  if ((workingset->nWConstr[2] + workingset->nWConstr[3]) +
          workingset->nWConstr[4] >
      0) {
    double tol;
    int idxDiag;
    int iy0;
    tol = tolfactor * (double)workingset->nVar * 2.2204460492503131E-16;
    for (idx = 0; idx < nFixedConstr; idx++) {
      qrmanager->jpvt[idx] = 1;
    }
    idx_col = nFixedConstr + 1;
    if (idx_col <= nDepIneq) {
      memset(&qrmanager->jpvt[idx_col + -1], 0,
             ((nDepIneq - idx_col) + 1) * sizeof(int));
    }
    for (idx_col = 0; idx_col < nDepIneq; idx_col++) {
      iy0 = qrmanager->ldq * idx_col;
      idxDiag = workingset->ldA * idx_col;
      for (k = 0; k < nVar_tmp_tmp; k++) {
        qrmanager->QR[iy0 + k] = workingset->ATwset[idxDiag + k];
      }
    }
    if (workingset->nVar * workingset->nActiveConstr == 0) {
      qrmanager->mrows = workingset->nVar;
      qrmanager->ncols = workingset->nActiveConstr;
      qrmanager->minRowCol = 0;
    } else {
      qrmanager->usedPivoting = true;
      qrmanager->mrows = workingset->nVar;
      qrmanager->ncols = workingset->nActiveConstr;
      idxDiag = workingset->nVar;
      iy0 = workingset->nActiveConstr;
      if (idxDiag <= iy0) {
        iy0 = idxDiag;
      }
      qrmanager->minRowCol = iy0;
      xzgeqp3(qrmanager->QR, workingset->nVar, workingset->nActiveConstr,
              qrmanager->jpvt, qrmanager->tau);
    }
    nDepIneq = 0;
    for (idx = workingset->nActiveConstr - 1; idx + 1 > nVar_tmp_tmp; idx--) {
      nDepIneq++;
      memspace->workspace_int[nDepIneq - 1] = qrmanager->jpvt[idx];
    }
    if (idx + 1 <= workingset->nVar) {
      idxDiag = idx + qrmanager->ldq * idx;
      while ((idx + 1 > nFixedConstr) && (fabs(qrmanager->QR[idxDiag]) < tol)) {
        nDepIneq++;
        memspace->workspace_int[nDepIneq - 1] = qrmanager->jpvt[idx];
        idx--;
        idxDiag = (idxDiag - qrmanager->ldq) - 1;
      }
    }
    countsort(memspace->workspace_int, nDepIneq, memspace->workspace_sort,
              nFixedConstr + 1, workingset->nActiveConstr);
    for (idx = nDepIneq; idx >= 1; idx--) {
      iy0 = memspace->workspace_int[idx - 1] - 1;
      idxDiag = workingset->Wid[iy0] - 1;
      workingset->isActiveConstr[(workingset->isActiveIdx[idxDiag] +
                                  workingset->Wlocalidx[iy0]) -
                                 2] = false;
      workingset->Wid[iy0] = workingset->Wid[workingset->nActiveConstr - 1];
      workingset->Wlocalidx[iy0] =
          workingset->Wlocalidx[workingset->nActiveConstr - 1];
      idx_col = workingset->nVar;
      for (k = 0; k < idx_col; k++) {
        workingset->ATwset[k + workingset->ldA * iy0] =
            workingset
                ->ATwset[k + workingset->ldA * (workingset->nActiveConstr - 1)];
      }
      workingset->bwset[iy0] = workingset->bwset[workingset->nActiveConstr - 1];
      workingset->nActiveConstr--;
      workingset->nWConstr[idxDiag]--;
    }
  }
}

/*
 * File trailer for RemoveDependentIneq_.c
 *
 * [EOF]
 */
