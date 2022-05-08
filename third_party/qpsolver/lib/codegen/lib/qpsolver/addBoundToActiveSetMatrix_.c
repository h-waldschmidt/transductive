/*
 * File: addBoundToActiveSetMatrix_.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "addBoundToActiveSetMatrix_.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : f_struct_T *obj
 *                int TYPE
 *                int idx_local
 * Return Type  : void
 */
void addBoundToActiveSetMatrix_(f_struct_T *obj, int TYPE, int idx_local)
{
  int colOffset;
  int i;
  int idx_bnd_local;
  obj->nWConstr[TYPE - 1]++;
  obj->isActiveConstr[(obj->isActiveIdx[TYPE - 1] + idx_local) - 2] = true;
  obj->nActiveConstr++;
  obj->Wid[obj->nActiveConstr - 1] = TYPE;
  obj->Wlocalidx[obj->nActiveConstr - 1] = idx_local;
  colOffset = obj->ldA * (obj->nActiveConstr - 1) - 1;
  if (TYPE == 5) {
    idx_bnd_local = obj->indexUB[idx_local - 1];
    obj->bwset[obj->nActiveConstr - 1] = obj->ub[idx_bnd_local - 1];
  } else {
    idx_bnd_local = obj->indexLB[idx_local - 1];
    obj->bwset[obj->nActiveConstr - 1] = obj->lb[idx_bnd_local - 1];
  }
  if (idx_bnd_local - 2 >= 0) {
    memset(&obj->ATwset[colOffset + 1], 0,
           (((idx_bnd_local + colOffset) - colOffset) - 1) * sizeof(double));
  }
  obj->ATwset[idx_bnd_local + colOffset] = 2.0 * (double)(TYPE == 5) - 1.0;
  idx_bnd_local++;
  i = obj->nVar;
  if (idx_bnd_local <= i) {
    memset(&obj->ATwset[idx_bnd_local + colOffset], 0,
           ((((i + colOffset) - idx_bnd_local) - colOffset) + 1) *
               sizeof(double));
  }
  switch (obj->probType) {
  case 3:
  case 2:
    break;
  default:
    obj->ATwset[obj->nVar + colOffset] = -1.0;
    break;
  }
}

/*
 * File trailer for addBoundToActiveSetMatrix_.c
 *
 * [EOF]
 */
