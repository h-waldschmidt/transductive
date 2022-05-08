/*
 * File: feasibleX0ForWorkingSet.h
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

#ifndef FEASIBLEX0FORWORKINGSET_H
#define FEASIBLEX0FORWORKINGSET_H

/* Include Files */
#include "qpsolver_internal_types.h"
#include "rtwtypes.h"
#include <stddef.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Function Declarations */
boolean_T feasibleX0ForWorkingSet(double workspace[32], double xCurrent[4],
                                  f_struct_T *workingset,
                                  e_struct_T *qrmanager);

#ifdef __cplusplus
}
#endif

#endif
/*
 * File trailer for feasibleX0ForWorkingSet.h
 *
 * [EOF]
 */
