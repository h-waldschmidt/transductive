/*
 * File: printInfo.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef PRINTINFO_H
#define PRINTINFO_H

/* Include Files */
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void printInfo(boolean_T newBlocking, int PROBLEM_TYPE, double alpha,
               double stepNorm, int activeConstrChangedType,
               int localActiveConstrIdx, int activeSetChangeID,
               double solution_fstar, double solution_firstorderopt,
               double solution_maxConstr, int solution_iterations,
               const int workingset_indexLB[4], const int workingset_indexUB[4],
               int workingset_nActiveConstr);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for printInfo.h
 *
 * [EOF]
 */
