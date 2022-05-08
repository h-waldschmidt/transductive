/*
 * File: computeFirstOrderOpt.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef COMPUTEFIRSTORDEROPT_H
#define COMPUTEFIRSTORDEROPT_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void computeFirstOrderOpt(struct_T *solution, const b_struct_T *objective,
                          int workingset_nVar, int workingset_ldA,
                          const double workingset_ATwset[32],
                          int workingset_nActiveConstr, double workspace[32]);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for computeFirstOrderOpt.h
 *
 * [EOF]
 */
