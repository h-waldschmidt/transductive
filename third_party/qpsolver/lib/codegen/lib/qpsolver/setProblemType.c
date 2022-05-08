/*
 * File: setProblemType.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "setProblemType.h"
#include "modifyOverheadPhaseOne_.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : f_struct_T *obj
 *                int PROBLEM_TYPE
 * Return Type  : void
 */
void setProblemType(f_struct_T *obj, int PROBLEM_TYPE)
{
  int i;
  int idx;
  int idx_col;
  switch (PROBLEM_TYPE) {
  case 3:
    obj->nVar = obj->nVarOrig;
    obj->mConstr = obj->mConstrOrig;
    if (obj->nWConstr[4] > 0) {
      i = obj->sizesNormal[4];
      for (idx = 0; idx < i; idx++) {
        obj->isActiveConstr[(obj->isActiveIdxNormal[4] + idx) - 1] =
            obj->isActiveConstr[(obj->isActiveIdx[4] + idx) - 1];
      }
    }
    for (i = 0; i < 5; i++) {
      obj->sizes[i] = obj->sizesNormal[i];
    }
    for (i = 0; i < 6; i++) {
      obj->isActiveIdx[i] = obj->isActiveIdxNormal[i];
    }
    break;
  case 1:
    obj->nVar = obj->nVarOrig + 1;
    obj->mConstr = obj->mConstrOrig + 1;
    for (i = 0; i < 5; i++) {
      obj->sizes[i] = obj->sizesPhaseOne[i];
    }
    for (i = 0; i < 6; i++) {
      obj->isActiveIdx[i] = obj->isActiveIdxPhaseOne[i];
    }
    modifyOverheadPhaseOne_(obj);
    break;
  case 2: {
    obj->nVar = obj->nVarMax - 1;
    obj->mConstr = obj->mConstrMax - 1;
    for (i = 0; i < 5; i++) {
      obj->sizes[i] = obj->sizesRegularized[i];
    }
    if (obj->probType != 4) {
      int colOffsetATw;
      int i1;
      int i2;
      int mEq;
      int offsetEq2;
      int offsetIneq_tmp_tmp;
      mEq = obj->sizes[1];
      offsetIneq_tmp_tmp = obj->nVarOrig + 1;
      idx = obj->nVarOrig + obj->sizes[2];
      offsetEq2 = idx + obj->sizes[1];
      i = obj->sizes[0];
      for (idx_col = 0; idx_col < i; idx_col++) {
        colOffsetATw = obj->ldA * idx_col;
        i1 = obj->nVarOrig + 1;
        i2 = obj->nVar;
        if (i1 <= i2) {
          memset(&obj->ATwset[(i1 + colOffsetATw) + -1], 0,
                 ((((i2 + colOffsetATw) - i1) - colOffsetATw) + 1) *
                     sizeof(double));
        }
      }
      for (idx_col = 0; idx_col < mEq; idx_col++) {
        int colOffsetAeq;
        colOffsetAeq = obj->ldA * idx_col - 1;
        colOffsetATw = colOffsetAeq + obj->ldA * (obj->isActiveIdx[1] - 1);
        if (offsetIneq_tmp_tmp <= idx) {
          memset(&obj->Aeq[offsetIneq_tmp_tmp + colOffsetAeq], 0,
                 ((((idx + colOffsetAeq) - offsetIneq_tmp_tmp) - colOffsetAeq) +
                  1) *
                     sizeof(double));
          memset(&obj->ATwset[offsetIneq_tmp_tmp + colOffsetATw], 0,
                 ((((idx + colOffsetATw) - offsetIneq_tmp_tmp) - colOffsetATw) +
                  1) *
                     sizeof(double));
        }
        i = idx + 1;
        i1 = (idx + idx_col) + 1;
        i2 = i1 - 1;
        if (i <= i2) {
          memset(&obj->Aeq[i + colOffsetAeq], 0,
                 ((((i2 + colOffsetAeq) - i) - colOffsetAeq) + 1) *
                     sizeof(double));
          memset(&obj->ATwset[i + colOffsetATw], 0,
                 ((((i2 + colOffsetATw) - i) - colOffsetATw) + 1) *
                     sizeof(double));
        }
        obj->Aeq[i1 + colOffsetAeq] = -1.0;
        obj->ATwset[i1 + colOffsetATw] = -1.0;
        i = i1 + 1;
        if (i <= offsetEq2) {
          memset(&obj->Aeq[i + colOffsetAeq], 0,
                 ((((offsetEq2 + colOffsetAeq) - i) - colOffsetAeq) + 1) *
                     sizeof(double));
          memset(&obj->ATwset[i + colOffsetATw], 0,
                 ((((offsetEq2 + colOffsetATw) - i) - colOffsetATw) + 1) *
                     sizeof(double));
        }
        i = offsetEq2 + 1;
        i1 = (offsetEq2 + idx_col) + 1;
        i2 = i1 - 1;
        if (i <= i2) {
          memset(&obj->Aeq[i + colOffsetAeq], 0,
                 ((((i2 + colOffsetAeq) - i) - colOffsetAeq) + 1) *
                     sizeof(double));
          memset(&obj->ATwset[i + colOffsetATw], 0,
                 ((((i2 + colOffsetATw) - i) - colOffsetATw) + 1) *
                     sizeof(double));
        }
        obj->Aeq[i1 + colOffsetAeq] = 1.0;
        obj->ATwset[i1 + colOffsetATw] = 1.0;
        i = i1 + 1;
        i1 = obj->nVar;
        if (i <= i1) {
          memset(&obj->Aeq[i + colOffsetAeq], 0,
                 ((((i1 + colOffsetAeq) - i) - colOffsetAeq) + 1) *
                     sizeof(double));
          memset(&obj->ATwset[i + colOffsetATw], 0,
                 ((((i1 + colOffsetATw) - i) - colOffsetATw) + 1) *
                     sizeof(double));
        }
      }
      mEq = obj->nVarOrig;
      i = obj->sizesNormal[3] + 1;
      i1 = obj->sizesRegularized[3];
      for (idx = i; idx <= i1; idx++) {
        mEq++;
        obj->indexLB[idx - 1] = mEq;
      }
      if (obj->nWConstr[4] > 0) {
        i = obj->sizesRegularized[4];
        for (idx = 0; idx < i; idx++) {
          obj->isActiveConstr[obj->isActiveIdxRegularized[4] + idx] =
              obj->isActiveConstr[(obj->isActiveIdx[4] + idx) - 1];
        }
      }
      i = obj->isActiveIdx[4];
      i1 = obj->isActiveIdxRegularized[4] - 1;
      if (i <= i1) {
        memset(&obj->isActiveConstr[i + -1], 0,
               ((i1 - i) + 1) * sizeof(boolean_T));
      }
      i = obj->nVarOrig + 1;
      i1 = (obj->nVarOrig + obj->sizes[2]) + (obj->sizes[1] << 1);
      if (i <= i1) {
        memset(&obj->lb[i + -1], 0, ((i1 - i) + 1) * sizeof(double));
      }
      mEq = obj->isActiveIdx[2];
      i = obj->nActiveConstr;
      for (idx_col = mEq; idx_col <= i; idx_col++) {
        colOffsetATw = obj->ldA * (idx_col - 1) - 1;
        if (obj->Wid[idx_col - 1] == 3) {
          i1 = (offsetIneq_tmp_tmp + obj->Wlocalidx[idx_col - 1]) - 2;
          if (offsetIneq_tmp_tmp <= i1) {
            memset(
                &obj->ATwset[offsetIneq_tmp_tmp + colOffsetATw], 0,
                ((((i1 + colOffsetATw) - offsetIneq_tmp_tmp) - colOffsetATw) +
                 1) *
                    sizeof(double));
          }
          obj->ATwset[((offsetIneq_tmp_tmp + obj->Wlocalidx[idx_col - 1]) +
                       colOffsetATw) -
                      1] = -1.0;
          i1 = offsetIneq_tmp_tmp + obj->Wlocalidx[idx_col - 1];
          i2 = obj->nVar;
          if (i1 <= i2) {
            memset(&obj->ATwset[i1 + colOffsetATw], 0,
                   ((((i2 + colOffsetATw) - i1) - colOffsetATw) + 1) *
                       sizeof(double));
          }
        } else {
          i1 = obj->nVar;
          if (offsetIneq_tmp_tmp <= i1) {
            memset(
                &obj->ATwset[offsetIneq_tmp_tmp + colOffsetATw], 0,
                ((((i1 + colOffsetATw) - offsetIneq_tmp_tmp) - colOffsetATw) +
                 1) *
                    sizeof(double));
          }
        }
      }
    }
    for (i = 0; i < 6; i++) {
      obj->isActiveIdx[i] = obj->isActiveIdxRegularized[i];
    }
  } break;
  default:
    obj->nVar = obj->nVarMax;
    obj->mConstr = obj->mConstrMax;
    for (i = 0; i < 5; i++) {
      obj->sizes[i] = obj->sizesRegPhaseOne[i];
    }
    for (i = 0; i < 6; i++) {
      obj->isActiveIdx[i] = obj->isActiveIdxRegPhaseOne[i];
    }
    modifyOverheadPhaseOne_(obj);
    break;
  }
  obj->probType = PROBLEM_TYPE;
}

/*
 * File trailer for setProblemType.c
 *
 * [EOF]
 */
