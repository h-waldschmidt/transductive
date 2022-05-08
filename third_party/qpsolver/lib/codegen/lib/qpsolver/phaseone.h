/*
 * File: phaseone.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef PHASEONE_H
#define PHASEONE_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void phaseone(struct_T *solution, d_struct_T *memspace, f_struct_T *workingset,
              e_struct_T *qrmanager, c_struct_T *cholmanager,
              b_struct_T *objective, g_struct_T *options);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for phaseone.h
 *
 * [EOF]
 */
