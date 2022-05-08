/*
 * File: iterate.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef ITERATE_H
#define ITERATE_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void iterate(struct_T *solution, d_struct_T *memspace, f_struct_T *workingset,
             e_struct_T *qrmanager, c_struct_T *cholmanager,
             b_struct_T *objective, boolean_T options_IterDisplayQP,
             double options_PricingTolerance, double options_ObjectiveLimit,
             double options_ConstraintTolerance, double options_StepTolerance,
             boolean_T runTimeOptions_RemainFeasible);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for iterate.h
 *
 * [EOF]
 */
