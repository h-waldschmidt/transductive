/*
 * File: fullColLDL2_.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "fullColLDL2_.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <math.h>

/* Function Definitions */
/*
 * Arguments    : c_struct_T *obj
 *                int NColsRemain
 *                double REG_PRIMAL
 * Return Type  : void
 */
void fullColLDL2_(c_struct_T *obj, int NColsRemain, double REG_PRIMAL)
{
  int LDimSizeP1;
  int ijA;
  int j;
  int jA;
  int k;
  LDimSizeP1 = obj->ldm;
  for (k = 0; k < NColsRemain; k++) {
    double alpha1;
    int LD_diagOffset;
    int i;
    int subMatrixDim;
    LD_diagOffset = (LDimSizeP1 + 1) * k;
    if (fabs(obj->FMat[LD_diagOffset]) <= obj->regTol_) {
      obj->FMat[LD_diagOffset] += REG_PRIMAL;
    }
    alpha1 = -1.0 / obj->FMat[LD_diagOffset];
    subMatrixDim = (NColsRemain - k) - 2;
    for (jA = 0; jA <= subMatrixDim; jA++) {
      obj->workspace_[jA] = obj->FMat[(LD_diagOffset + jA) + 1];
    }
    if (!(alpha1 == 0.0)) {
      jA = LD_diagOffset + LDimSizeP1;
      for (j = 0; j <= subMatrixDim; j++) {
        if (obj->workspace_[j] != 0.0) {
          double temp;
          int i1;
          temp = obj->workspace_[j] * alpha1;
          i = jA + 2;
          i1 = subMatrixDim + jA;
          for (ijA = i; ijA <= i1 + 2; ijA++) {
            obj->FMat[ijA - 1] += obj->workspace_[(ijA - jA) - 2] * temp;
          }
        }
        jA += obj->ldm;
      }
    }
    for (jA = 0; jA <= subMatrixDim; jA++) {
      i = (LD_diagOffset + jA) + 1;
      obj->FMat[i] /= obj->FMat[LD_diagOffset];
    }
  }
  jA = (obj->ldm + 1) * (NColsRemain - 1);
  if (fabs(obj->FMat[jA]) <= obj->regTol_) {
    obj->FMat[jA] += REG_PRIMAL;
  }
}

/*
 * File trailer for fullColLDL2_.c
 *
 * [EOF]
 */
