/*
 * File: linearForm_.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "linearForm_.h"
#include "qpsolver_data.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : boolean_T obj_hasLinear
 *                int obj_nvar
 *                double workspace[32]
 *                const double x[4]
 * Return Type  : void
 */
void linearForm_(boolean_T obj_hasLinear, int obj_nvar, double workspace[32],
                 const double x[4])
{
  static const signed char f[3] = {2, -3, 1};
  int ia;
  int iac;
  int ix;
  ix = 0;
  if (obj_hasLinear) {
    for (ix = 0; ix < obj_nvar; ix++) {
      workspace[ix] = f[ix];
    }
    ix = 1;
  }
  if (obj_nvar != 0) {
    int i;
    if ((ix != 1) && (obj_nvar - 1 >= 0)) {
      memset(&workspace[0], 0, obj_nvar * sizeof(double));
    }
    ix = 0;
    i = obj_nvar * (obj_nvar - 1) + 1;
    for (iac = 1; obj_nvar < 0 ? iac >= i : iac <= i; iac += obj_nvar) {
      double c;
      int i1;
      c = 0.5 * x[ix];
      i1 = (iac + obj_nvar) - 1;
      for (ia = iac; ia <= i1; ia++) {
        int i2;
        i2 = ia - iac;
        workspace[i2] += (double)iv[ia - 1] * c;
      }
      ix++;
    }
  }
}

/*
 * File trailer for linearForm_.c
 *
 * [EOF]
 */
