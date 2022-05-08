/*
 * File: PresolveWorkingSet.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef PRESOLVEWORKINGSET_H
#define PRESOLVEWORKINGSET_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
void PresolveWorkingSet(struct_T *solution, d_struct_T *memspace,
                        f_struct_T *workingset, e_struct_T *qrmanager);

void b_PresolveWorkingSet(struct_T *solution, d_struct_T *memspace,
                          f_struct_T *workingset, e_struct_T *qrmanager,
                          const g_struct_T *options);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for PresolveWorkingSet.h
 *
 * [EOF]
 */
