/*
 * File: modifyOverheadPhaseOne_.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "modifyOverheadPhaseOne_.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"

/* Function Definitions */
/*
 * Arguments    : f_struct_T *obj
 * Return Type  : void
 */
void modifyOverheadPhaseOne_(f_struct_T *obj)
{
  int i;
  int idx;
  int idxEq;
  i = obj->sizes[0];
  for (idx = 0; idx < i; idx++) {
    obj->ATwset[(obj->nVar + obj->ldA * idx) - 1] = 0.0;
  }
  i = obj->sizes[1];
  for (idx = 0; idx < i; idx++) {
    idxEq = (obj->nVar + obj->ldA * idx) - 1;
    obj->Aeq[idxEq] = 0.0;
    obj->ATwset[idxEq + obj->ldA * (obj->isActiveIdx[1] - 1)] = 0.0;
  }
  obj->indexLB[obj->sizes[3] - 1] = obj->nVar;
  obj->lb[obj->nVar - 1] = obj->SLACK0;
  idxEq = obj->isActiveIdx[2];
  i = obj->nActiveConstr;
  for (idx = idxEq; idx <= i; idx++) {
    obj->ATwset[(obj->nVar + obj->ldA * (idx - 1)) - 1] = -1.0;
  }
  if (obj->nWConstr[4] > 0) {
    i = obj->sizesNormal[4];
    for (idx = 0; idx <= i; idx++) {
      obj->isActiveConstr[(obj->isActiveIdx[4] + idx) - 1] = false;
    }
  }
  obj->isActiveConstr[obj->isActiveIdx[4] - 2] = false;
}

/*
 * File trailer for modifyOverheadPhaseOne_.c
 *
 * [EOF]
 */
