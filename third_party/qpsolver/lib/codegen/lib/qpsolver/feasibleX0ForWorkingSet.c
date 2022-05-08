/*
 * File: feasibleX0ForWorkingSet.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "feasibleX0ForWorkingSet.h"
#include "computeQ_.h"
#include "factorQR.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "xgemv.h"
#include "xzgeqp3.h"
#include "rt_nonfinite.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : double workspace[32]
 *                double xCurrent[4]
 *                f_struct_T *workingset
 *                e_struct_T *qrmanager
 * Return Type  : boolean_T
 */
boolean_T feasibleX0ForWorkingSet(double workspace[32], double xCurrent[4],
                                  f_struct_T *workingset, e_struct_T *qrmanager)
{
  double B[32];
  int br;
  int idx;
  int iy;
  int jBcol;
  int mFixed;
  int mWConstr;
  int nVar;
  boolean_T nonDegenerateWset;
  mWConstr = workingset->nActiveConstr;
  nVar = workingset->nVar;
  nonDegenerateWset = true;
  if (mWConstr != 0) {
    double c;
    int ar;
    int i;
    int i1;
    int iAcol;
    for (idx = 0; idx < mWConstr; idx++) {
      workspace[idx] = workingset->bwset[idx];
      workspace[idx + 8] = workingset->bwset[idx];
    }
    iAcol = workingset->ldA;
    if ((nVar != 0) && (mWConstr != 0)) {
      iy = 0;
      i = workingset->ldA * (mWConstr - 1) + 1;
      for (jBcol = 1; iAcol < 0 ? jBcol >= i : jBcol <= i; jBcol += iAcol) {
        c = 0.0;
        i1 = (jBcol + nVar) - 1;
        for (br = jBcol; br <= i1; br++) {
          c += workingset->ATwset[br - 1] * xCurrent[br - jBcol];
        }
        workspace[iy] += -c;
        iy++;
      }
    }
    if (mWConstr >= nVar) {
      int ldq;
      for (iy = 0; iy < nVar; iy++) {
        iAcol = qrmanager->ldq * iy;
        for (jBcol = 0; jBcol < mWConstr; jBcol++) {
          qrmanager->QR[jBcol + iAcol] =
              workingset->ATwset[iy + workingset->ldA * jBcol];
        }
      }
      if (mWConstr * nVar == 0) {
        qrmanager->mrows = mWConstr;
        qrmanager->ncols = nVar;
        qrmanager->minRowCol = 0;
      } else {
        qrmanager->usedPivoting = false;
        qrmanager->mrows = mWConstr;
        qrmanager->ncols = nVar;
        for (idx = 0; idx < nVar; idx++) {
          qrmanager->jpvt[idx] = idx + 1;
        }
        if (mWConstr <= nVar) {
          i = mWConstr;
        } else {
          i = nVar;
        }
        qrmanager->minRowCol = i;
        memcpy(&B[0], &qrmanager->QR[0], 32U * sizeof(double));
        qrmanager->tau[0] = 0.0;
        qrmanager->tau[1] = 0.0;
        qrmanager->tau[2] = 0.0;
        qrmanager->tau[3] = 0.0;
        if (i >= 1) {
          qrf(B, mWConstr, nVar, i, qrmanager->tau);
        }
        memcpy(&qrmanager->QR[0], &B[0], 32U * sizeof(double));
      }
      computeQ_(qrmanager, qrmanager->mrows);
      ldq = qrmanager->ldq;
      memcpy(&B[0], &workspace[0], 32U * sizeof(double));
      if (nVar != 0) {
        for (jBcol = 0; jBcol <= 8; jBcol += 8) {
          i = jBcol + 1;
          i1 = jBcol + nVar;
          if (i <= i1) {
            memset(&workspace[i + -1], 0, ((i1 - i) + 1) * sizeof(double));
          }
        }
        br = -1;
        for (jBcol = 0; jBcol <= 8; jBcol += 8) {
          ar = -1;
          i = jBcol + 1;
          i1 = jBcol + nVar;
          for (idx = i; idx <= i1; idx++) {
            c = 0.0;
            for (iy = 0; iy < mWConstr; iy++) {
              c += qrmanager->Q[(iy + ar) + 1] * B[(iy + br) + 1];
            }
            workspace[idx - 1] += c;
            ar += ldq;
          }
          br += 8;
        }
      }
      for (mFixed = 0; mFixed < 2; mFixed++) {
        jBcol = (mFixed << 3) - 1;
        for (idx = nVar; idx >= 1; idx--) {
          iy = ldq * (idx - 1) - 1;
          i = idx + jBcol;
          c = workspace[i];
          if (c != 0.0) {
            workspace[i] = c / qrmanager->QR[idx + iy];
            for (br = 0; br <= idx - 2; br++) {
              i1 = (br + jBcol) + 1;
              workspace[i1] -= workspace[i] * qrmanager->QR[(br + iy) + 1];
            }
          }
        }
      }
    } else {
      int ldq;
      factorQR(qrmanager, workingset->ATwset, nVar, mWConstr, workingset->ldA);
      computeQ_(qrmanager, qrmanager->minRowCol);
      ldq = qrmanager->ldq;
      for (mFixed = 0; mFixed < 2; mFixed++) {
        jBcol = mFixed << 3;
        for (br = 0; br < mWConstr; br++) {
          iAcol = ldq * br;
          iy = br + jBcol;
          c = workspace[iy];
          for (idx = 0; idx < br; idx++) {
            c -= qrmanager->QR[idx + iAcol] * workspace[idx + jBcol];
          }
          workspace[iy] = c / qrmanager->QR[br + iAcol];
        }
      }
      memcpy(&B[0], &workspace[0], 32U * sizeof(double));
      if (nVar != 0) {
        for (jBcol = 0; jBcol <= 8; jBcol += 8) {
          i = jBcol + 1;
          i1 = jBcol + nVar;
          if (i <= i1) {
            memset(&workspace[i + -1], 0, ((i1 - i) + 1) * sizeof(double));
          }
        }
        br = 0;
        for (jBcol = 0; jBcol <= 8; jBcol += 8) {
          ar = -1;
          i = br + 1;
          i1 = br + mWConstr;
          for (mFixed = i; mFixed <= i1; mFixed++) {
            iy = jBcol + 1;
            iAcol = jBcol + nVar;
            for (idx = iy; idx <= iAcol; idx++) {
              workspace[idx - 1] +=
                  B[mFixed - 1] * qrmanager->Q[(ar + idx) - jBcol];
            }
            ar += ldq;
          }
          br += 8;
        }
      }
    }
    idx = 0;
    int exitg1;
    do {
      exitg1 = 0;
      if (idx <= nVar - 1) {
        if (rtIsInf(workspace[idx]) || rtIsNaN(workspace[idx])) {
          nonDegenerateWset = false;
          exitg1 = 1;
        } else {
          c = workspace[idx + 8];
          if (rtIsInf(c) || rtIsNaN(c)) {
            nonDegenerateWset = false;
            exitg1 = 1;
          } else {
            idx++;
          }
        }
      } else {
        double v;
        if (nVar >= 1) {
          iAcol = nVar - 1;
          for (idx = 0; idx <= iAcol; idx++) {
            workspace[idx] += xCurrent[idx];
          }
        }
        br = workingset->sizes[3];
        ar = workingset->sizes[4];
        mFixed = workingset->sizes[0];
        if (workingset->probType == 2) {
          c = 0.0;
          jBcol = workingset->sizes[1] - 1;
          for (idx = 0; idx <= jBcol; idx++) {
            workingset->maxConstrWorkspace[idx] = workingset->beq;
          }
          xgemv(workingset->nVarOrig, workingset->sizes[1], workingset->Aeq,
                workingset->ldA, workspace, workingset->maxConstrWorkspace);
          iAcol = workingset->nVarOrig + workingset->sizes[2];
          iy = iAcol + workingset->sizes[1];
          for (idx = 0; idx <= jBcol; idx++) {
            workingset->maxConstrWorkspace[idx] =
                (workingset->maxConstrWorkspace[idx] - workspace[iAcol + idx]) +
                workspace[iy + idx];
            c = fmax(c, fabs(workingset->maxConstrWorkspace[idx]));
          }
        } else {
          c = 0.0;
          jBcol = workingset->sizes[1] - 1;
          for (idx = 0; idx <= jBcol; idx++) {
            workingset->maxConstrWorkspace[idx] = workingset->beq;
          }
          xgemv(workingset->nVar, workingset->sizes[1], workingset->Aeq,
                workingset->ldA, workspace, workingset->maxConstrWorkspace);
          for (idx = 0; idx <= jBcol; idx++) {
            c = fmax(c, fabs(workingset->maxConstrWorkspace[idx]));
          }
        }
        if (workingset->sizes[3] > 0) {
          for (idx = 0; idx < br; idx++) {
            iAcol = workingset->indexLB[idx] - 1;
            c = fmax(c, -workspace[iAcol] - workingset->lb[iAcol]);
          }
        }
        if (workingset->sizes[4] > 0) {
          for (idx = 0; idx < ar; idx++) {
            iAcol = workingset->indexUB[idx] - 1;
            c = fmax(c, workspace[iAcol] - workingset->ub[iAcol]);
          }
        }
        if (workingset->sizes[0] > 0) {
          for (idx = 0; idx < mFixed; idx++) {
            c = fmax(c, fabs(workspace[workingset->indexFixed[idx] - 1] -
                             workingset->ub[workingset->indexFixed[idx] - 1]));
          }
        }
        br = workingset->sizes[3];
        ar = workingset->sizes[4];
        mFixed = workingset->sizes[0];
        if (workingset->probType == 2) {
          v = 0.0;
          jBcol = workingset->sizes[1] - 1;
          for (idx = 0; idx <= jBcol; idx++) {
            workingset->maxConstrWorkspace[idx] = workingset->beq;
          }
          b_xgemv(workingset->nVarOrig, workingset->sizes[1], workingset->Aeq,
                  workingset->ldA, workspace, workingset->maxConstrWorkspace);
          iAcol = (workingset->nVarOrig + workingset->sizes[2]) + 7;
          iy = (iAcol + workingset->sizes[1]) + 1;
          for (idx = 0; idx <= jBcol; idx++) {
            workingset->maxConstrWorkspace[idx] =
                (workingset->maxConstrWorkspace[idx] -
                 workspace[(iAcol + idx) + 1]) +
                workspace[iy + idx];
            v = fmax(v, fabs(workingset->maxConstrWorkspace[idx]));
          }
        } else {
          v = 0.0;
          jBcol = workingset->sizes[1] - 1;
          for (idx = 0; idx <= jBcol; idx++) {
            workingset->maxConstrWorkspace[idx] = workingset->beq;
          }
          b_xgemv(workingset->nVar, workingset->sizes[1], workingset->Aeq,
                  workingset->ldA, workspace, workingset->maxConstrWorkspace);
          for (idx = 0; idx <= jBcol; idx++) {
            v = fmax(v, fabs(workingset->maxConstrWorkspace[idx]));
          }
        }
        if (workingset->sizes[3] > 0) {
          for (idx = 0; idx < br; idx++) {
            iAcol = workingset->indexLB[idx];
            v = fmax(v, -workspace[iAcol + 7] - workingset->lb[iAcol - 1]);
          }
        }
        if (workingset->sizes[4] > 0) {
          for (idx = 0; idx < ar; idx++) {
            iAcol = workingset->indexUB[idx];
            v = fmax(v, workspace[iAcol + 7] - workingset->ub[iAcol - 1]);
          }
        }
        if (workingset->sizes[0] > 0) {
          for (idx = 0; idx < mFixed; idx++) {
            v = fmax(v, fabs(workspace[workingset->indexFixed[idx] + 7] -
                             workingset->ub[workingset->indexFixed[idx] - 1]));
          }
        }
        if ((c <= 2.2204460492503131E-16) || (c < v)) {
          if (nVar - 1 >= 0) {
            memcpy(&xCurrent[0], &workspace[0], nVar * sizeof(double));
          }
        } else if (nVar - 1 >= 0) {
          memcpy(&xCurrent[0], &workspace[8], nVar * sizeof(double));
        }
        exitg1 = 1;
      }
    } while (exitg1 == 0);
  }
  return nonDegenerateWset;
}

/*
 * File trailer for feasibleX0ForWorkingSet.c
 *
 * [EOF]
 */
