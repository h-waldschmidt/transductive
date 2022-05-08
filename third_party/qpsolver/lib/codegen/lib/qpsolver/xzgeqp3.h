/*
 * File: xzgeqp3.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef XZGEQP3_H
#define XZGEQP3_H

/* Include Files */
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void qrf(double A[32], int m, int n, int nfxd, double tau[4]);

void xzgeqp3(double A[32], int m, int n, int jpvt[8], double tau[4]);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for xzgeqp3.h
 *
 * [EOF]
 */
