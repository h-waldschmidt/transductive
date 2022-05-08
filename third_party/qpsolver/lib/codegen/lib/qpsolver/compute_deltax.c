/*
 * File: compute_deltax.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "compute_deltax.h"
#include "fullColLDL2_.h"
#include "ixamax.h"
#include "qpsolver_data.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : struct_T *solution
 *                d_struct_T *memspace
 *                const e_struct_T *qrmanager
 *                c_struct_T *cholmanager
 *                const b_struct_T *objective
 * Return Type  : void
 */
void compute_deltax(struct_T *solution, d_struct_T *memspace,
                    const e_struct_T *qrmanager, c_struct_T *cholmanager,
                    const b_struct_T *objective)
{
  int A_maxDiag_idx;
  int br;
  int cr;
  int ib;
  int ic;
  int ix;
  int ldQ;
  int mNull_tmp;
  int nVar_tmp;
  int nVars;
  nVar_tmp = qrmanager->mrows - 1;
  mNull_tmp = qrmanager->mrows - qrmanager->ncols;
  if (mNull_tmp <= 0) {
    if (nVar_tmp >= 0) {
      memset(&solution->searchDir[0], 0, (nVar_tmp + 1) * sizeof(double));
    }
  } else {
    for (ix = 0; ix <= nVar_tmp; ix++) {
      solution->searchDir[ix] = -objective->grad[ix];
    }
    if (qrmanager->ncols <= 0) {
      switch (objective->objtype) {
      case 5:
      case 4:
        break;
      case 3: {
        double temp;
        temp = 1.4901161193847656E-8 * cholmanager->scaleFactor *
               (double)qrmanager->mrows;
        cholmanager->ndims = qrmanager->mrows;
        for (ix = 0; ix <= nVar_tmp; ix++) {
          A_maxDiag_idx = (nVar_tmp + 1) * ix;
          ldQ = cholmanager->ldm * ix;
          for (nVars = 0; nVars <= nVar_tmp; nVars++) {
            cholmanager->FMat[ldQ + nVars] = iv[A_maxDiag_idx + nVars];
          }
        }
        A_maxDiag_idx =
            ixamax(qrmanager->mrows, cholmanager->FMat, cholmanager->ldm + 1) -
            1;
        cholmanager->regTol_ = fmax(
            fabs(cholmanager
                     ->FMat[A_maxDiag_idx + cholmanager->ldm * A_maxDiag_idx]) *
                2.2204460492503131E-16,
            fabs(temp));
        fullColLDL2_(cholmanager, qrmanager->mrows, temp);
        if (cholmanager->ConvexCheck) {
          ix = 0;
          int exitg1;
          do {
            exitg1 = 0;
            if (ix <= nVar_tmp) {
              if (cholmanager->FMat[ix + cholmanager->ldm * ix] <= 0.0) {
                cholmanager->info = -ix - 1;
                exitg1 = 1;
              } else {
                ix++;
              }
            } else {
              cholmanager->ConvexCheck = false;
              exitg1 = 1;
            }
          } while (exitg1 == 0);
        }
        if (cholmanager->info != 0) {
          solution->state = -6;
        } else {
          int i;
          ldQ = cholmanager->ndims - 2;
          if (cholmanager->ndims != 0) {
            for (nVars = 0; nVars <= ldQ + 1; nVars++) {
              A_maxDiag_idx = nVars + nVars * cholmanager->ldm;
              i = ldQ - nVars;
              for (br = 0; br <= i; br++) {
                ix = (nVars + br) + 1;
                solution->searchDir[ix] -=
                    solution->searchDir[nVars] *
                    cholmanager->FMat[(A_maxDiag_idx + br) + 1];
              }
            }
          }
          i = cholmanager->ndims;
          for (ix = 0; ix < i; ix++) {
            solution->searchDir[ix] /=
                cholmanager->FMat[ix + cholmanager->ldm * ix];
          }
          ldQ = cholmanager->ndims;
          if (cholmanager->ndims != 0) {
            for (nVars = ldQ; nVars >= 1; nVars--) {
              A_maxDiag_idx = (nVars - 1) * cholmanager->ldm;
              temp = solution->searchDir[nVars - 1];
              i = nVars + 1;
              for (br = ldQ; br >= i; br--) {
                temp -= cholmanager->FMat[(A_maxDiag_idx + br) - 1] *
                        solution->searchDir[br - 1];
              }
              solution->searchDir[nVars - 1] = temp;
            }
          }
        }
      } break;
      }
    } else {
      int nullStartIdx;
      int nullStartIdx_tmp;
      nullStartIdx_tmp = qrmanager->ldq * qrmanager->ncols;
      nullStartIdx = nullStartIdx_tmp + 1;
      if (objective->objtype == 5) {
        for (ix = 0; ix < mNull_tmp; ix++) {
          memspace->workspace_double[ix] =
              -qrmanager
                   ->Q[nVar_tmp + qrmanager->ldq * (qrmanager->ncols + ix)];
        }
        nVars = qrmanager->ldq;
        if (qrmanager->mrows != 0) {
          int i;
          if (nVar_tmp >= 0) {
            memset(&solution->searchDir[0], 0, (nVar_tmp + 1) * sizeof(double));
          }
          ix = 0;
          i = (nullStartIdx_tmp + qrmanager->ldq * (mNull_tmp - 1)) + 1;
          for (ldQ = nullStartIdx; nVars < 0 ? ldQ >= i : ldQ <= i;
               ldQ += nVars) {
            int i1;
            i1 = ldQ + nVar_tmp;
            for (br = ldQ; br <= i1; br++) {
              int i2;
              i2 = br - ldQ;
              solution->searchDir[i2] +=
                  qrmanager->Q[br - 1] * memspace->workspace_double[ix];
            }
            ix++;
          }
        }
      } else {
        double temp;
        int i;
        int i1;
        int i2;
        if (objective->objtype == 3) {
          int ar;
          int lastColC;
          nVars = qrmanager->mrows;
          ix = cholmanager->ldm;
          ldQ = qrmanager->ldq;
          if (qrmanager->mrows != 0) {
            br = nullStartIdx_tmp;
            lastColC = (mNull_tmp - 1) << 3;
            for (cr = 0; cr <= lastColC; cr += 8) {
              i = cr + 1;
              i1 = cr + nVars;
              if (i <= i1) {
                memset(&memspace->workspace_double[i + -1], 0,
                       ((i1 - i) + 1) * sizeof(double));
              }
            }
            for (cr = 0; cr <= lastColC; cr += 8) {
              ar = -1;
              i = br + 1;
              i1 = br + nVars;
              for (ib = i; ib <= i1; ib++) {
                i2 = cr + 1;
                A_maxDiag_idx = cr + nVars;
                for (ic = i2; ic <= A_maxDiag_idx; ic++) {
                  memspace->workspace_double[ic - 1] +=
                      qrmanager->Q[ib - 1] * (double)iv[(ar + ic) - cr];
                }
                ar += nVars;
              }
              br += ldQ;
            }
          }
          lastColC = cholmanager->ldm * (mNull_tmp - 1);
          for (cr = 0; ix < 0 ? cr >= lastColC : cr <= lastColC; cr += ix) {
            i = cr + 1;
            i1 = cr + mNull_tmp;
            if (i <= i1) {
              memset(&cholmanager->FMat[i + -1], 0,
                     ((i1 - i) + 1) * sizeof(double));
            }
          }
          br = -1;
          for (cr = 0; ix < 0 ? cr >= lastColC : cr <= lastColC; cr += ix) {
            ar = nullStartIdx_tmp;
            i = cr + 1;
            i1 = cr + mNull_tmp;
            for (ic = i; ic <= i1; ic++) {
              temp = 0.0;
              for (A_maxDiag_idx = 0; A_maxDiag_idx < nVars; A_maxDiag_idx++) {
                temp += qrmanager->Q[A_maxDiag_idx + ar] *
                        memspace->workspace_double[(A_maxDiag_idx + br) + 1];
              }
              cholmanager->FMat[ic - 1] += temp;
              ar += ldQ;
            }
            br += 8;
          }
        }
        temp = 1.4901161193847656E-8 * cholmanager->scaleFactor *
               (double)mNull_tmp;
        cholmanager->ndims = mNull_tmp;
        A_maxDiag_idx =
            ixamax(mNull_tmp, cholmanager->FMat, cholmanager->ldm + 1) - 1;
        cholmanager->regTol_ = fmax(
            fabs(cholmanager
                     ->FMat[A_maxDiag_idx + cholmanager->ldm * A_maxDiag_idx]) *
                2.2204460492503131E-16,
            fabs(temp));
        fullColLDL2_(cholmanager, mNull_tmp, temp);
        if (cholmanager->ConvexCheck) {
          ix = 0;
          int exitg1;
          do {
            exitg1 = 0;
            if (ix <= mNull_tmp - 1) {
              if (cholmanager->FMat[ix + cholmanager->ldm * ix] <= 0.0) {
                cholmanager->info = -ix - 1;
                exitg1 = 1;
              } else {
                ix++;
              }
            } else {
              cholmanager->ConvexCheck = false;
              exitg1 = 1;
            }
          } while (exitg1 == 0);
        }
        if (cholmanager->info != 0) {
          solution->state = -6;
        } else {
          nVars = qrmanager->ldq;
          if (qrmanager->mrows != 0) {
            memset(&memspace->workspace_double[0], 0,
                   mNull_tmp * sizeof(double));
            A_maxDiag_idx = 0;
            i = (nullStartIdx_tmp + qrmanager->ldq * (mNull_tmp - 1)) + 1;
            for (ldQ = nullStartIdx; nVars < 0 ? ldQ >= i : ldQ <= i;
                 ldQ += nVars) {
              temp = 0.0;
              i1 = ldQ + nVar_tmp;
              for (br = ldQ; br <= i1; br++) {
                temp += qrmanager->Q[br - 1] * objective->grad[br - ldQ];
              }
              memspace->workspace_double[A_maxDiag_idx] += -temp;
              A_maxDiag_idx++;
            }
          }
          ldQ = cholmanager->ndims - 2;
          if (cholmanager->ndims != 0) {
            for (nVars = 0; nVars <= ldQ + 1; nVars++) {
              A_maxDiag_idx = nVars + nVars * cholmanager->ldm;
              i = ldQ - nVars;
              for (br = 0; br <= i; br++) {
                ix = (nVars + br) + 1;
                memspace->workspace_double[ix] -=
                    memspace->workspace_double[nVars] *
                    cholmanager->FMat[(A_maxDiag_idx + br) + 1];
              }
            }
          }
          i = cholmanager->ndims;
          for (ix = 0; ix < i; ix++) {
            memspace->workspace_double[ix] /=
                cholmanager->FMat[ix + cholmanager->ldm * ix];
          }
          ldQ = cholmanager->ndims;
          if (cholmanager->ndims != 0) {
            for (nVars = ldQ; nVars >= 1; nVars--) {
              A_maxDiag_idx = (nVars - 1) * cholmanager->ldm;
              temp = memspace->workspace_double[nVars - 1];
              i = nVars + 1;
              for (br = ldQ; br >= i; br--) {
                temp -= cholmanager->FMat[(A_maxDiag_idx + br) - 1] *
                        memspace->workspace_double[br - 1];
              }
              memspace->workspace_double[nVars - 1] = temp;
            }
          }
          nVars = qrmanager->ldq;
          if (qrmanager->mrows != 0) {
            if (nVar_tmp >= 0) {
              memset(&solution->searchDir[0], 0,
                     (nVar_tmp + 1) * sizeof(double));
            }
            ix = 0;
            i = (nullStartIdx_tmp + qrmanager->ldq * (mNull_tmp - 1)) + 1;
            for (ldQ = nullStartIdx; nVars < 0 ? ldQ >= i : ldQ <= i;
                 ldQ += nVars) {
              i1 = ldQ + nVar_tmp;
              for (br = ldQ; br <= i1; br++) {
                i2 = br - ldQ;
                solution->searchDir[i2] +=
                    qrmanager->Q[br - 1] * memspace->workspace_double[ix];
              }
              ix++;
            }
          }
        }
      }
    }
  }
}

/*
 * File trailer for compute_deltax.c
 *
 * [EOF]
 */
