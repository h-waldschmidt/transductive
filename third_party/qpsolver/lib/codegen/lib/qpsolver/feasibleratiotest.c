/*
 * File: feasibleratiotest.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "feasibleratiotest.h"
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
 *                boolean_T isPhaseOne
 *                double tolcon
 *                double *alpha
 *                boolean_T *newBlocking
 *                int *constrType
 *                int *constrIdx
 * Return Type  : void
 */
void feasibleratiotest(
    const double solution_xstar[4], const double solution_searchDir[4],
    const double workspace[32], int workingset_nVar,
    const double workingset_lb[4], const double workingset_ub[4],
    const int workingset_indexLB[4], const int workingset_indexUB[4],
    const int workingset_sizes[5], const int workingset_isActiveIdx[6],
    const boolean_T workingset_isActiveConstr[8],
    const int workingset_nWConstr[5], boolean_T isPhaseOne, double tolcon,
    double *alpha, boolean_T *newBlocking, int *constrType, int *constrIdx)
{
  double alphaTemp;
  double denomTol;
  double phaseOneCorrectionP;
  double phaseOneCorrectionX;
  double ratio;
  int idx;
  int totalIneq;
  int totalUB;
  totalIneq = workingset_sizes[2];
  totalUB = workingset_sizes[4];
  *alpha = 1.0E+30;
  *newBlocking = false;
  *constrType = 0;
  *constrIdx = 0;
  denomTol =
      2.2204460492503131E-13 * b_xnrm2(workingset_nVar, solution_searchDir);
  if (workingset_nWConstr[2] < workingset_sizes[2]) {
    for (idx = 0; idx < totalIneq; idx++) {
      alphaTemp = workspace[idx + 8];
      if ((alphaTemp > denomTol) &&
          (!workingset_isActiveConstr[(workingset_isActiveIdx[2] + idx) - 1])) {
        alphaTemp =
            fmin(fabs(workspace[idx]), tolcon - workspace[idx]) / alphaTemp;
        if (alphaTemp < *alpha) {
          *alpha = alphaTemp;
          *constrType = 3;
          *constrIdx = idx + 1;
          *newBlocking = true;
        }
      }
    }
  }
  if (workingset_nWConstr[3] < workingset_sizes[3]) {
    phaseOneCorrectionX =
        (double)isPhaseOne * solution_xstar[workingset_nVar - 1];
    phaseOneCorrectionP =
        (double)isPhaseOne * solution_searchDir[workingset_nVar - 1];
    totalIneq = workingset_sizes[3];
    for (idx = 0; idx <= totalIneq - 2; idx++) {
      int i;
      i = workingset_indexLB[idx];
      alphaTemp = -solution_searchDir[i - 1] - phaseOneCorrectionP;
      if ((alphaTemp > denomTol) &&
          (!workingset_isActiveConstr[(workingset_isActiveIdx[3] + idx) - 1])) {
        ratio = (-solution_xstar[i - 1] - workingset_lb[i - 1]) -
                phaseOneCorrectionX;
        alphaTemp = fmin(fabs(ratio), tolcon - ratio) / alphaTemp;
        if (alphaTemp < *alpha) {
          *alpha = alphaTemp;
          *constrType = 4;
          *constrIdx = idx + 1;
          *newBlocking = true;
        }
      }
    }
    totalIneq = workingset_indexLB[workingset_sizes[3] - 1] - 1;
    alphaTemp = -solution_searchDir[totalIneq];
    if ((alphaTemp > denomTol) &&
        (!workingset_isActiveConstr
             [(workingset_isActiveIdx[3] + workingset_sizes[3]) - 2])) {
      ratio = -solution_xstar[totalIneq] - workingset_lb[totalIneq];
      alphaTemp = fmin(fabs(ratio), tolcon - ratio) / alphaTemp;
      if (alphaTemp < *alpha) {
        *alpha = alphaTemp;
        *constrType = 4;
        *constrIdx = workingset_sizes[3];
        *newBlocking = true;
      }
    }
  }
  if (workingset_nWConstr[4] < workingset_sizes[4]) {
    phaseOneCorrectionX =
        (double)isPhaseOne * solution_xstar[workingset_nVar - 1];
    phaseOneCorrectionP =
        (double)isPhaseOne * solution_searchDir[workingset_nVar - 1];
    for (idx = 0; idx < totalUB; idx++) {
      totalIneq = workingset_indexUB[idx];
      alphaTemp = solution_searchDir[totalIneq - 1] - phaseOneCorrectionP;
      if ((alphaTemp > denomTol) &&
          (!workingset_isActiveConstr[(workingset_isActiveIdx[4] + idx) - 1])) {
        ratio = (solution_xstar[totalIneq - 1] - workingset_ub[totalIneq - 1]) -
                phaseOneCorrectionX;
        alphaTemp = fmin(fabs(ratio), tolcon - ratio) / alphaTemp;
        if (alphaTemp < *alpha) {
          *alpha = alphaTemp;
          *constrType = 5;
          *constrIdx = idx + 1;
          *newBlocking = true;
        }
      }
    }
  }
  if (!isPhaseOne) {
    if ((*newBlocking) && (*alpha > 1.0)) {
      *newBlocking = false;
    }
    *alpha = fmin(*alpha, 1.0);
  }
}

/*
 * File trailer for feasibleratiotest.c
 *
 * [EOF]
 */
