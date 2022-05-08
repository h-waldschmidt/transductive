/*
 * File: _coder_qpsolver_api.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef _CODER_QPSOLVER_API_H
#define _CODER_QPSOLVER_API_H

/* Include Files */
#include "emlrt.h"
#include "tmwtypes.h"
#include <string.h>

/* Variable Declarations */
extern emlrtCTX emlrtRootTLSGlobal;
extern emlrtContext emlrtContextGlobal;

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void qpsolver(real_T x[3], real_T *fval);

void qpsolver_api(int32_T nlhs, const mxArray *plhs[2]);

void qpsolver_atexit(void);

void qpsolver_initialize(void);

void qpsolver_terminate(void);

void qpsolver_xil_shutdown(void);

void qpsolver_xil_terminate(void);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for _coder_qpsolver_api.h
 *
 * [EOF]
 */
