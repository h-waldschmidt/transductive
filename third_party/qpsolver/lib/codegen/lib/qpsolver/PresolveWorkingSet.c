/*
 * File: PresolveWorkingSet.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "PresolveWorkingSet.h"
#include "RemoveDependentIneq_.h"
#include "computeQ_.h"
#include "countsort.h"
#include "feasibleX0ForWorkingSet.h"
#include "maxConstraintViolation.h"
#include "qpsolver_internal_types.h"
#include "removeEqConstr.h"
#include "rt_nonfinite.h"
#include "xzgeqp3.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : struct_T *solution
 *                d_struct_T *memspace
 *                f_struct_T *workingset
 *                e_struct_T *qrmanager
 * Return Type  : void
 */
void PresolveWorkingSet(struct_T *solution, d_struct_T *memspace,
                        f_struct_T *workingset, e_struct_T *qrmanager)
{
  static const double t3_bwset[8] = {0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0};
  static const signed char t3_ATwset[32] = {1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
                                            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                                            0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
  static const signed char x[32] = {1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
                                    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                                    0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
  static const signed char t3_Wid[8] = {2, 0, 0, 0, 0, 0, 0, 0};
  static const signed char t3_Wlocalidx[8] = {1, 0, 0, 0, 0, 0, 0, 0};
  static const signed char t3_isActiveIdx[6] = {1, 1, 2, 2, 5, 8};
  static const signed char t3_isActiveIdxNormal[6] = {1, 1, 2, 2, 5, 8};
  static const signed char t3_isActiveIdxPhaseOne[6] = {1, 1, 2, 2, 6, 9};
  static const signed char t3_isActiveIdxRegPhaseOne[6] = {1, 1, 2, 2, 8, 11};
  static const signed char t3_isActiveIdxRegularized[6] = {1, 1, 2, 2, 7, 10};
  static const signed char t3_nWConstr[5] = {0, 1, 0, 0, 0};
  static const signed char t3_sizes[5] = {0, 1, 0, 3, 3};
  static const signed char t3_sizesNormal[5] = {0, 1, 0, 3, 3};
  static const signed char t3_sizesPhaseOne[5] = {0, 1, 0, 4, 3};
  static const signed char t3_sizesRegPhaseOne[5] = {0, 1, 0, 6, 3};
  static const signed char t3_sizesRegularized[5] = {0, 1, 0, 5, 3};
  static const boolean_T t3_isActiveConstr[8] = {true,  false, false, false,
                                                 false, false, false, false};
  int i;
  int idx_global;
  solution->state = 82;
  qrmanager->ldq = 4;
  memset(&qrmanager->QR[0], 0, 32U * sizeof(double));
  memset(&qrmanager->Q[0], 0, 16U * sizeof(double));
  for (i = 0; i < 8; i++) {
    qrmanager->jpvt[i] = 0;
  }
  workingset->mConstr = 7;
  workingset->mConstrOrig = 7;
  workingset->mConstrMax = 8;
  workingset->nVar = 3;
  workingset->nVarOrig = 3;
  workingset->nVarMax = 4;
  workingset->ldA = 4;
  workingset->beq = 0.5;
  qrmanager->tau[0] = 0.0;
  workingset->Aeq[0] = 1.0;
  workingset->lb[0] = 0.0;
  workingset->ub[0] = 1.0;
  workingset->indexLB[0] = 1;
  workingset->indexUB[0] = 1;
  workingset->indexFixed[0] = 0;
  qrmanager->tau[1] = 0.0;
  workingset->Aeq[1] = 1.0;
  workingset->lb[1] = 0.0;
  workingset->ub[1] = 1.0;
  workingset->indexLB[1] = 2;
  workingset->indexUB[1] = 2;
  workingset->indexFixed[1] = 0;
  qrmanager->tau[2] = 0.0;
  workingset->Aeq[2] = 1.0;
  workingset->lb[2] = 0.0;
  workingset->ub[2] = 1.0;
  workingset->indexLB[2] = 3;
  workingset->indexUB[2] = 3;
  workingset->indexFixed[2] = 0;
  qrmanager->tau[3] = 0.0;
  workingset->Aeq[3] = 0.0;
  workingset->lb[3] = 0.0;
  workingset->ub[3] = 0.0;
  workingset->indexLB[3] = 0;
  workingset->indexUB[3] = 0;
  workingset->indexFixed[3] = 0;
  workingset->mEqRemoved = 0;
  workingset->indexEqRemoved = 0;
  for (i = 0; i < 32; i++) {
    workingset->ATwset[i] = t3_ATwset[i];
  }
  workingset->nActiveConstr = 1;
  for (i = 0; i < 8; i++) {
    workingset->bwset[i] = t3_bwset[i];
    workingset->maxConstrWorkspace[i] = 0.0;
  }
  for (i = 0; i < 5; i++) {
    workingset->sizes[i] = t3_sizes[i];
    workingset->sizesNormal[i] = t3_sizesNormal[i];
    workingset->sizesPhaseOne[i] = t3_sizesPhaseOne[i];
    workingset->sizesRegularized[i] = t3_sizesRegularized[i];
    workingset->sizesRegPhaseOne[i] = t3_sizesRegPhaseOne[i];
  }
  for (i = 0; i < 6; i++) {
    workingset->isActiveIdx[i] = t3_isActiveIdx[i];
    workingset->isActiveIdxNormal[i] = t3_isActiveIdxNormal[i];
    workingset->isActiveIdxPhaseOne[i] = t3_isActiveIdxPhaseOne[i];
    workingset->isActiveIdxRegularized[i] = t3_isActiveIdxRegularized[i];
    workingset->isActiveIdxRegPhaseOne[i] = t3_isActiveIdxRegPhaseOne[i];
  }
  for (i = 0; i < 8; i++) {
    workingset->isActiveConstr[i] = t3_isActiveConstr[i];
    workingset->Wid[i] = t3_Wid[i];
    workingset->Wlocalidx[i] = t3_Wlocalidx[i];
  }
  for (i = 0; i < 5; i++) {
    workingset->nWConstr[i] = t3_nWConstr[i];
  }
  workingset->probType = 3;
  workingset->SLACK0 = 0.0;
  i = 0;
  qrmanager->QR[0] = x[0];
  qrmanager->jpvt[0] = 0;
  qrmanager->QR[4] = x[1];
  qrmanager->jpvt[1] = 0;
  qrmanager->QR[8] = x[2];
  qrmanager->jpvt[2] = 0;
  qrmanager->usedPivoting = true;
  qrmanager->mrows = 1;
  qrmanager->ncols = 3;
  qrmanager->minRowCol = 1;
  xzgeqp3(qrmanager->QR, 1, 3, qrmanager->jpvt, qrmanager->tau);
  if (fabs(qrmanager->QR[0]) < 6.6613381477509392E-14) {
    i = 1;
    computeQ_(qrmanager, 1);
    if (fabs(qrmanager->Q[0] * 0.5) >= 6.6613381477509392E-14) {
      i = -1;
    }
  }
  if (i > 0) {
    qrmanager->QR[0] = x[0];
    qrmanager->QR[1] = x[1];
    qrmanager->QR[2] = x[2];
    qrmanager->jpvt[0] = 0;
    qrmanager->usedPivoting = true;
    qrmanager->mrows = 3;
    qrmanager->ncols = 1;
    qrmanager->minRowCol = 1;
    xzgeqp3(qrmanager->QR, 3, 1, qrmanager->jpvt, qrmanager->tau);
    memspace->workspace_int[0] = qrmanager->jpvt[0];
    countsort(memspace->workspace_int, 1, memspace->workspace_sort, 1, 1);
    removeEqConstr(workingset, memspace->workspace_int[0]);
  }
  if ((i != -1) && (workingset->nActiveConstr <= qrmanager->ldq)) {
    boolean_T guard1 = false;
    boolean_T okWorkingSet;
    RemoveDependentIneq_(workingset, qrmanager, memspace, 100.0);
    okWorkingSet = feasibleX0ForWorkingSet(
        memspace->workspace_double, solution->xstar, workingset, qrmanager);
    guard1 = false;
    if (!okWorkingSet) {
      RemoveDependentIneq_(workingset, qrmanager, memspace, 1000.0);
      okWorkingSet = feasibleX0ForWorkingSet(
          memspace->workspace_double, solution->xstar, workingset, qrmanager);
      if (!okWorkingSet) {
        solution->state = -7;
      } else {
        guard1 = true;
      }
    } else {
      guard1 = true;
    }
    if (guard1 && (workingset->nWConstr[0] + workingset->nWConstr[1] ==
                   workingset->nVar)) {
      double constrViolation;
      constrViolation = maxConstraintViolation(workingset, solution->xstar);
      if (constrViolation > 1.0E-8) {
        solution->state = -2;
      }
    }
  } else {
    int idxEndIneq;
    solution->state = -3;
    i = (workingset->nWConstr[0] + workingset->nWConstr[1]) + 1;
    idxEndIneq = workingset->nActiveConstr;
    for (idx_global = i; idx_global <= idxEndIneq; idx_global++) {
      workingset->isActiveConstr
          [(workingset->isActiveIdx[workingset->Wid[idx_global - 1] - 1] +
            workingset->Wlocalidx[idx_global - 1]) -
           2] = false;
    }
    workingset->nWConstr[2] = 0;
    workingset->nWConstr[3] = 0;
    workingset->nWConstr[4] = 0;
    workingset->nActiveConstr =
        workingset->nWConstr[0] + workingset->nWConstr[1];
  }
}

/*
 * Arguments    : struct_T *solution
 *                d_struct_T *memspace
 *                f_struct_T *workingset
 *                e_struct_T *qrmanager
 *                const g_struct_T *options
 * Return Type  : void
 */
void b_PresolveWorkingSet(struct_T *solution, d_struct_T *memspace,
                          f_struct_T *workingset, e_struct_T *qrmanager,
                          const g_struct_T *options)
{
  double tol;
  int idxDiag;
  int ix;
  int ix0;
  int k;
  int mTotalWorkingEq_tmp_tmp;
  int mWorkingFixed;
  int nDepInd;
  int nVar;
  solution->state = 82;
  nVar = workingset->nVar - 1;
  mWorkingFixed = workingset->nWConstr[0];
  mTotalWorkingEq_tmp_tmp = workingset->nWConstr[0] + workingset->nWConstr[1];
  nDepInd = 0;
  if (mTotalWorkingEq_tmp_tmp > 0) {
    int i;
    int u0;
    for (idxDiag = 0; idxDiag < mTotalWorkingEq_tmp_tmp; idxDiag++) {
      for (ix = 0; ix <= nVar; ix++) {
        qrmanager->QR[idxDiag + qrmanager->ldq * ix] =
            workingset->ATwset[ix + workingset->ldA * idxDiag];
      }
    }
    nDepInd = mTotalWorkingEq_tmp_tmp - workingset->nVar;
    if (nDepInd <= 0) {
      nDepInd = 0;
    }
    if (nVar >= 0) {
      memset(&qrmanager->jpvt[0], 0, (nVar + 1) * sizeof(int));
    }
    i = mTotalWorkingEq_tmp_tmp * workingset->nVar;
    if (i == 0) {
      qrmanager->mrows = mTotalWorkingEq_tmp_tmp;
      qrmanager->ncols = workingset->nVar;
      qrmanager->minRowCol = 0;
    } else {
      qrmanager->usedPivoting = true;
      qrmanager->mrows = mTotalWorkingEq_tmp_tmp;
      qrmanager->ncols = workingset->nVar;
      idxDiag = workingset->nVar;
      if (mTotalWorkingEq_tmp_tmp <= idxDiag) {
        idxDiag = mTotalWorkingEq_tmp_tmp;
      }
      qrmanager->minRowCol = idxDiag;
      xzgeqp3(qrmanager->QR, mTotalWorkingEq_tmp_tmp, workingset->nVar,
              qrmanager->jpvt, qrmanager->tau);
    }
    tol = 100.0 * (double)workingset->nVar * 2.2204460492503131E-16;
    u0 = workingset->nVar;
    if (u0 > mTotalWorkingEq_tmp_tmp) {
      u0 = mTotalWorkingEq_tmp_tmp;
    }
    idxDiag = u0 + qrmanager->ldq * (u0 - 1);
    while ((idxDiag > 0) && (fabs(qrmanager->QR[idxDiag - 1]) < tol)) {
      idxDiag = (idxDiag - qrmanager->ldq) - 1;
      nDepInd++;
    }
    if (nDepInd > 0) {
      boolean_T exitg1;
      computeQ_(qrmanager, qrmanager->mrows);
      idxDiag = 0;
      exitg1 = false;
      while ((!exitg1) && (idxDiag <= nDepInd - 1)) {
        double qtb;
        ix = qrmanager->ldq * ((mTotalWorkingEq_tmp_tmp - idxDiag) - 1);
        qtb = 0.0;
        for (k = 0; k < mTotalWorkingEq_tmp_tmp; k++) {
          qtb += qrmanager->Q[ix + k] * workingset->bwset[k];
        }
        if (fabs(qtb) >= tol) {
          nDepInd = -1;
          exitg1 = true;
        } else {
          idxDiag++;
        }
      }
    }
    if (nDepInd > 0) {
      for (ix = 0; ix < mTotalWorkingEq_tmp_tmp; ix++) {
        idxDiag = qrmanager->ldq * ix;
        ix0 = workingset->ldA * ix;
        for (k = 0; k <= nVar; k++) {
          qrmanager->QR[idxDiag + k] = workingset->ATwset[ix0 + k];
        }
      }
      for (idxDiag = 0; idxDiag < mWorkingFixed; idxDiag++) {
        qrmanager->jpvt[idxDiag] = 1;
      }
      idxDiag = workingset->nWConstr[0] + 1;
      if (idxDiag <= mTotalWorkingEq_tmp_tmp) {
        memset(&qrmanager->jpvt[idxDiag + -1], 0,
               ((mTotalWorkingEq_tmp_tmp - idxDiag) + 1) * sizeof(int));
      }
      if (i == 0) {
        qrmanager->mrows = workingset->nVar;
        qrmanager->ncols = mTotalWorkingEq_tmp_tmp;
        qrmanager->minRowCol = 0;
      } else {
        qrmanager->usedPivoting = true;
        qrmanager->mrows = workingset->nVar;
        qrmanager->ncols = mTotalWorkingEq_tmp_tmp;
        qrmanager->minRowCol = u0;
        xzgeqp3(qrmanager->QR, workingset->nVar, mTotalWorkingEq_tmp_tmp,
                qrmanager->jpvt, qrmanager->tau);
      }
      for (idxDiag = 0; idxDiag < nDepInd; idxDiag++) {
        memspace->workspace_int[idxDiag] =
            qrmanager->jpvt[(mTotalWorkingEq_tmp_tmp - nDepInd) + idxDiag];
      }
      countsort(memspace->workspace_int, nDepInd, memspace->workspace_sort, 1,
                mTotalWorkingEq_tmp_tmp);
      for (idxDiag = nDepInd; idxDiag >= 1; idxDiag--) {
        removeEqConstr(workingset, memspace->workspace_int[idxDiag - 1]);
      }
    }
  }
  if ((nDepInd != -1) && (workingset->nActiveConstr <= qrmanager->ldq)) {
    boolean_T guard1 = false;
    boolean_T okWorkingSet;
    RemoveDependentIneq_(workingset, qrmanager, memspace, 100.0);
    okWorkingSet = feasibleX0ForWorkingSet(
        memspace->workspace_double, solution->xstar, workingset, qrmanager);
    guard1 = false;
    if (!okWorkingSet) {
      RemoveDependentIneq_(workingset, qrmanager, memspace, 1000.0);
      okWorkingSet = feasibleX0ForWorkingSet(
          memspace->workspace_double, solution->xstar, workingset, qrmanager);
      if (!okWorkingSet) {
        solution->state = -7;
      } else {
        guard1 = true;
      }
    } else {
      guard1 = true;
    }
    if (guard1 && (workingset->nWConstr[0] + workingset->nWConstr[1] ==
                   workingset->nVar)) {
      tol = maxConstraintViolation(workingset, solution->xstar);
      if (tol > options->ConstraintTolerance) {
        solution->state = -2;
      }
    }
  } else {
    solution->state = -3;
    idxDiag = (workingset->nWConstr[0] + workingset->nWConstr[1]) + 1;
    ix = workingset->nActiveConstr;
    for (ix0 = idxDiag; ix0 <= ix; ix0++) {
      workingset->isActiveConstr
          [(workingset->isActiveIdx[workingset->Wid[ix0 - 1] - 1] +
            workingset->Wlocalidx[ix0 - 1]) -
           2] = false;
    }
    workingset->nWConstr[2] = 0;
    workingset->nWConstr[3] = 0;
    workingset->nWConstr[4] = 0;
    workingset->nActiveConstr =
        workingset->nWConstr[0] + workingset->nWConstr[1];
  }
}

/*
 * File trailer for PresolveWorkingSet.c
 *
 * [EOF]
 */
