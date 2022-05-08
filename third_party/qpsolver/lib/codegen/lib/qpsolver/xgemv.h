/*
 * File: xgemv.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef XGEMV_H
#define XGEMV_H

/* Include Files */
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void b_xgemv(int m, int n, const double A[4], int lda, const double x[32],
             double y[8]);

void c_xgemv(int m, int n, const double A[4], int lda, const double x[4],
             double y[8]);

void xgemv(int m, int n, const double A[4], int lda, const double x[32],
           double y[8]);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for xgemv.h
 *
 * [EOF]
 */
