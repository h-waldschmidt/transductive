/*
 * File: computeGrad_StoreHx.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "computeGrad_StoreHx.h"
#include "qpsolver_data.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : b_struct_T *obj
 *                const double x[4]
 * Return Type  : void
 */
void computeGrad_StoreHx(b_struct_T *obj, const double x[4])
{
  static const signed char b_x[3] = {2, -3, 1};
  int ia;
  int iac;
  int m_tmp_tmp;
  switch (obj->objtype) {
  case 5: {
    int i;
    i = obj->nvar;
    if (i - 2 >= 0) {
      memset(&obj->grad[0], 0, (i - 1) * sizeof(double));
    }
    obj->grad[obj->nvar - 1] = obj->gammaScalar;
  } break;
  case 3: {
    int i;
    int iy;
    m_tmp_tmp = obj->nvar - 1;
    iy = obj->nvar;
    if (obj->nvar != 0) {
      int ix;
      if (m_tmp_tmp >= 0) {
        memset(&obj->Hx[0], 0, (m_tmp_tmp + 1) * sizeof(double));
      }
      ix = 0;
      i = obj->nvar * (obj->nvar - 1) + 1;
      for (iac = 1; iy < 0 ? iac >= i : iac <= i; iac += iy) {
        int i1;
        i1 = iac + m_tmp_tmp;
        for (ia = iac; ia <= i1; ia++) {
          int i2;
          i2 = ia - iac;
          obj->Hx[i2] += (double)iv[ia - 1] * x[ix];
        }
        ix++;
      }
    }
    i = obj->nvar;
    if (i - 1 >= 0) {
      memcpy(&obj->grad[0], &obj->Hx[0], i * sizeof(double));
    }
    if (obj->hasLinear && (obj->nvar >= 1)) {
      i = obj->nvar - 1;
      for (m_tmp_tmp = 0; m_tmp_tmp <= i; m_tmp_tmp++) {
        obj->grad[m_tmp_tmp] += (double)b_x[m_tmp_tmp];
      }
    }
  } break;
  case 4: {
    int i;
    int i1;
    int iy;
    int maxRegVar;
    maxRegVar = obj->maxVar - 1;
    m_tmp_tmp = obj->nvar - 1;
    iy = obj->nvar;
    if (obj->nvar != 0) {
      int ix;
      if (m_tmp_tmp >= 0) {
        memset(&obj->Hx[0], 0, (m_tmp_tmp + 1) * sizeof(double));
      }
      ix = 0;
      i = obj->nvar * (obj->nvar - 1) + 1;
      for (iac = 1; iy < 0 ? iac >= i : iac <= i; iac += iy) {
        i1 = iac + m_tmp_tmp;
        for (ia = iac; ia <= i1; ia++) {
          int i2;
          i2 = ia - iac;
          obj->Hx[i2] += (double)iv[ia - 1] * x[ix];
        }
        ix++;
      }
    }
    i = obj->nvar + 1;
    for (m_tmp_tmp = i; m_tmp_tmp <= maxRegVar; m_tmp_tmp++) {
      obj->Hx[m_tmp_tmp - 1] = obj->beta * x[m_tmp_tmp - 1];
    }
    if (maxRegVar - 1 >= 0) {
      memcpy(&obj->grad[0], &obj->Hx[0], maxRegVar * sizeof(double));
    }
    if (obj->hasLinear && (obj->nvar >= 1)) {
      i = obj->nvar - 1;
      for (m_tmp_tmp = 0; m_tmp_tmp <= i; m_tmp_tmp++) {
        obj->grad[m_tmp_tmp] += (double)b_x[m_tmp_tmp];
      }
    }
    m_tmp_tmp = (obj->maxVar - obj->nvar) - 1;
    if (m_tmp_tmp >= 1) {
      iy = obj->nvar;
      i = m_tmp_tmp - 1;
      for (m_tmp_tmp = 0; m_tmp_tmp <= i; m_tmp_tmp++) {
        i1 = iy + m_tmp_tmp;
        obj->grad[i1] += obj->rho;
      }
    }
  } break;
  }
}

/*
 * File trailer for computeGrad_StoreHx.c
 *
 * [EOF]
 */
