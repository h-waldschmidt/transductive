/*
 * File: feasibleratiotest.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef FEASIBLERATIOTEST_H
#define FEASIBLERATIOTEST_H

/* Include Files */
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void feasibleratiotest(
    const double solution_xstar[4], const double solution_searchDir[4],
    const double workspace[32], int workingset_nVar,
    const double workingset_lb[4], const double workingset_ub[4],
    const int workingset_indexLB[4], const int workingset_indexUB[4],
    const int workingset_sizes[5], const int workingset_isActiveIdx[6],
    const boolean_T workingset_isActiveConstr[8],
    const int workingset_nWConstr[5], boolean_T isPhaseOne, double tolcon,
    double *alpha, boolean_T *newBlocking, int *constrType, int *constrIdx);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for feasibleratiotest.h
 *
 * [EOF]
 */
