/*
 * File: ratiotest.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "ratiotest.h"
#include "rt_nonfinite.h"
#include "xnrm2.h"
#include <math.h>

/* Function Definitions */
/*
 * Arguments    : const double solution_xstar[4]
 *                const double solution_searchDir[4]
 *                const double workspace[32]
 *                int workingset_nVar
 *                const double workingset_lb[4]
 *                const double workingset_ub[4]
 *                const int workingset_indexLB[4]
 *                const int workingset_indexUB[4]
 *                const int workingset_sizes[5]
 *                const int workingset_isActiveIdx[6]
 *                const boolean_T workingset_isActiveConstr[8]
 *                const int workingset_nWConstr[5]
 *                double tolcon
 *                double *toldelta
 *                double *alpha
 *                boolean_T *newBlocking
 *                int *constrType
 *                int *constrIdx
 * Return Type  : void
 */
void ratiotest(const double solution_xstar[4],
               const double solution_searchDir[4], const double workspace[32],
               int workingset_nVar, const double workingset_lb[4],
               const double workingset_ub[4], const int workingset_indexLB[4],
               const int workingset_indexUB[4], const int workingset_sizes[5],
               const int workingset_isActiveIdx[6],
               const boolean_T workingset_isActiveConstr[8],
               const int workingset_nWConstr[5], double tolcon,
               double *toldelta, double *alpha, boolean_T *newBlocking,
               int *constrType, int *constrIdx)
{
  double alphaTemp;
  double denomTol;
  double p_max;
  double phaseOneCorrectionP;
  double phaseOneCorrectionX;
  double pk_corrected;
  double ratio_tmp;
  int idx;
  int totalIneq;
  int totalUB;
  totalIneq = workingset_sizes[2];
  totalUB = workingset_sizes[4];
  *alpha = 1.0E+30;
  *newBlocking = false;
  *constrType = 0;
  *constrIdx = 0;
  p_max = 0.0;
  denomTol =
      2.2204460492503131E-13 * b_xnrm2(workingset_nVar, solution_searchDir);
  if (workingset_nWConstr[2] < workingset_sizes[2]) {
    for (idx = 0; idx < totalIneq; idx++) {
      phaseOneCorrectionP = workspace[idx + 8];
      if ((phaseOneCorrectionP > denomTol) &&
          (!workingset_isActiveConstr[(workingset_isActiveIdx[2] + idx) - 1])) {
        phaseOneCorrectionX = tolcon - workspace[idx];
        alphaTemp = fmin(fabs(workspace[idx] - *toldelta),
                         phaseOneCorrectionX + *toldelta) /
                    phaseOneCorrectionP;
        if ((alphaTemp <= *alpha) && (fabs(phaseOneCorrectionP) > p_max)) {
          *alpha = alphaTemp;
          *constrType = 3;
          *constrIdx = idx + 1;
          *newBlocking = true;
        }
        alphaTemp = fmin(fabs(workspace[idx]), phaseOneCorrectionX) /
                    phaseOneCorrectionP;
        if (alphaTemp < *alpha) {
          *alpha = alphaTemp;
          *constrType = 3;
          *constrIdx = idx + 1;
          *newBlocking = true;
          p_max = fabs(workspace[idx + 8]);
        }
      }
    }
  }
  if (workingset_nWConstr[3] < workingset_sizes[3]) {
    phaseOneCorrectionX = 0.0 * solution_xstar[workingset_nVar - 1];
    phaseOneCorrectionP = 0.0 * solution_searchDir[workingset_nVar - 1];
    totalIneq = workingset_sizes[3];
    for (idx = 0; idx <= totalIneq - 2; idx++) {
      int i;
      i = workingset_indexLB[idx];
      pk_corrected = -solution_searchDir[i - 1] - phaseOneCorrectionP;
      if ((pk_corrected > denomTol) &&
          (!workingset_isActiveConstr[(workingset_isActiveIdx[3] + idx) - 1])) {
        ratio_tmp = -solution_xstar[i - 1] - workingset_lb[i - 1];
        alphaTemp = (ratio_tmp - *toldelta) - phaseOneCorrectionX;
        alphaTemp = fmin(fabs(alphaTemp), tolcon - alphaTemp) / pk_corrected;
        if ((alphaTemp <= *alpha) && (fabs(pk_corrected) > p_max)) {
          *alpha = alphaTemp;
          *constrType = 4;
          *constrIdx = idx + 1;
          *newBlocking = true;
        }
        alphaTemp = ratio_tmp - phaseOneCorrectionX;
        alphaTemp = fmin(fabs(alphaTemp), tolcon - alphaTemp) / pk_corrected;
        if (alphaTemp < *alpha) {
          *alpha = alphaTemp;
          *constrType = 4;
          *constrIdx = idx + 1;
          *newBlocking = true;
          p_max = fabs(pk_corrected);
        }
      }
    }
    totalIneq = workingset_indexLB[workingset_sizes[3] - 1] - 1;
    phaseOneCorrectionP = solution_searchDir[totalIneq];
    if ((-phaseOneCorrectionP > denomTol) &&
        (!workingset_isActiveConstr
             [(workingset_isActiveIdx[3] + workingset_sizes[3]) - 2])) {
      ratio_tmp = -solution_xstar[totalIneq] - workingset_lb[totalIneq];
      alphaTemp = ratio_tmp - *toldelta;
      alphaTemp =
          fmin(fabs(alphaTemp), tolcon - alphaTemp) / -phaseOneCorrectionP;
      if ((alphaTemp <= *alpha) && (fabs(phaseOneCorrectionP) > p_max)) {
        *alpha = alphaTemp;
        *constrType = 4;
        *constrIdx = workingset_sizes[3];
        *newBlocking = true;
      }
      alphaTemp =
          fmin(fabs(ratio_tmp), tolcon - ratio_tmp) / -phaseOneCorrectionP;
      if (alphaTemp < *alpha) {
        *alpha = alphaTemp;
        *constrType = 4;
        *constrIdx = workingset_sizes[3];
        *newBlocking = true;
        p_max = fabs(solution_searchDir[totalIneq]);
      }
    }
  }
  if (workingset_nWConstr[4] < workingset_sizes[4]) {
    phaseOneCorrectionX = 0.0 * solution_xstar[workingset_nVar - 1];
    phaseOneCorrectionP = 0.0 * solution_searchDir[workingset_nVar - 1];
    for (idx = 0; idx < totalUB; idx++) {
      totalIneq = workingset_indexUB[idx];
      pk_corrected = solution_searchDir[totalIneq - 1] - phaseOneCorrectionP;
      if ((pk_corrected > denomTol) &&
          (!workingset_isActiveConstr[(workingset_isActiveIdx[4] + idx) - 1])) {
        ratio_tmp =
            solution_xstar[totalIneq - 1] - workingset_ub[totalIneq - 1];
        alphaTemp = (ratio_tmp - *toldelta) - phaseOneCorrectionX;
        alphaTemp = fmin(fabs(alphaTemp), tolcon - alphaTemp) / pk_corrected;
        if ((alphaTemp <= *alpha) && (fabs(pk_corrected) > p_max)) {
          *alpha = alphaTemp;
          *constrType = 5;
          *constrIdx = idx + 1;
          *newBlocking = true;
        }
        alphaTemp = ratio_tmp - phaseOneCorrectionX;
        alphaTemp = fmin(fabs(alphaTemp), tolcon - alphaTemp) / pk_corrected;
        if (alphaTemp < *alpha) {
          *alpha = alphaTemp;
          *constrType = 5;
          *constrIdx = idx + 1;
          *newBlocking = true;
          p_max = fabs(pk_corrected);
        }
      }
    }
  }
  *toldelta += 6.608625846508183E-7;
  if (p_max > 0.0) {
    *alpha = fmax(*alpha, 6.608625846508183E-7 / p_max);
  }
  if ((*newBlocking) && (*alpha > 1.0)) {
    *newBlocking = false;
  }
  *alpha = fmin(*alpha, 1.0);
}

/*
 * File trailer for ratiotest.c
 *
 * [EOF]
 */
