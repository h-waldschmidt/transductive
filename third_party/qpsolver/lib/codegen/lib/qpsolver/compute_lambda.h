/*
 * File: compute_lambda.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef COMPUTE_LAMBDA_H
#define COMPUTE_LAMBDA_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void compute_lambda(double workspace[32], struct_T *solution,
                    const b_struct_T *objective, const e_struct_T *qrmanager);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for compute_lambda.h
 *
 * [EOF]
 */
