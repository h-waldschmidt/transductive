/*
 * File: factorQR.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef FACTORQR_H
#define FACTORQR_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void factorQR(e_struct_T *obj, const double A[32], int mrows, int ncols,
              int ldA);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for factorQR.h
 *
 * [EOF]
 */
