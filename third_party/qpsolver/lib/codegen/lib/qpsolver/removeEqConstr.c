/*
 * File: removeEqConstr.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "removeEqConstr.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"

/* Function Definitions */
/*
 * Arguments    : f_struct_T *obj
 *                int idx_global
 * Return Type  : void
 */
void removeEqConstr(f_struct_T *obj, int idx_global)
{
  int idx;
  int totalEq;
  totalEq = (obj->nWConstr[0] + obj->nWConstr[1]) - 1;
  if ((totalEq + 1 != 0) && (idx_global <= totalEq + 1)) {
    if ((obj->nActiveConstr == totalEq + 1) || (idx_global == totalEq + 1)) {
      int i;
      obj->mEqRemoved++;
      obj->indexEqRemoved = obj->Wlocalidx[idx_global - 1];
      totalEq = obj->Wid[idx_global - 1] - 1;
      obj->isActiveConstr[(obj->isActiveIdx[totalEq] +
                           obj->Wlocalidx[idx_global - 1]) -
                          2] = false;
      obj->Wid[idx_global - 1] = obj->Wid[obj->nActiveConstr - 1];
      obj->Wlocalidx[idx_global - 1] = obj->Wlocalidx[obj->nActiveConstr - 1];
      i = obj->nVar;
      for (idx = 0; idx < i; idx++) {
        obj->ATwset[idx + obj->ldA * (idx_global - 1)] =
            obj->ATwset[idx + obj->ldA * (obj->nActiveConstr - 1)];
      }
      obj->bwset[idx_global - 1] = obj->bwset[obj->nActiveConstr - 1];
      obj->nActiveConstr--;
      obj->nWConstr[totalEq]--;
    } else {
      int TYPE_tmp_tmp;
      int i;
      obj->mEqRemoved++;
      TYPE_tmp_tmp = obj->Wid[idx_global - 1] - 1;
      obj->indexEqRemoved = obj->Wlocalidx[idx_global - 1];
      obj->isActiveConstr[(obj->isActiveIdx[obj->Wid[idx_global - 1] - 1] +
                           obj->Wlocalidx[idx_global - 1]) -
                          2] = false;
      obj->Wid[idx_global - 1] = obj->Wid[totalEq];
      obj->Wlocalidx[idx_global - 1] = obj->Wlocalidx[totalEq];
      i = obj->nVar;
      for (idx = 0; idx < i; idx++) {
        obj->ATwset[idx + obj->ldA * (idx_global - 1)] =
            obj->ATwset[idx + obj->ldA * totalEq];
      }
      obj->bwset[idx_global - 1] = obj->bwset[totalEq];
      obj->Wid[totalEq] = obj->Wid[obj->nActiveConstr - 1];
      obj->Wlocalidx[totalEq] = obj->Wlocalidx[obj->nActiveConstr - 1];
      i = obj->nVar;
      for (idx = 0; idx < i; idx++) {
        obj->ATwset[idx + obj->ldA * totalEq] =
            obj->ATwset[idx + obj->ldA * (obj->nActiveConstr - 1)];
      }
      obj->bwset[totalEq] = obj->bwset[obj->nActiveConstr - 1];
      obj->nActiveConstr--;
      obj->nWConstr[TYPE_tmp_tmp]--;
    }
  }
}

/*
 * File trailer for removeEqConstr.c
 *
 * [EOF]
 */
