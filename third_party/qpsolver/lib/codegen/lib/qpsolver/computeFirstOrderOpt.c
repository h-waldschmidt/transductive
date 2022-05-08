/*
 * File: computeFirstOrderOpt.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "computeFirstOrderOpt.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : struct_T *solution
 *                const b_struct_T *objective
 *                int workingset_nVar
 *                int workingset_ldA
 *                const double workingset_ATwset[32]
 *                int workingset_nActiveConstr
 *                double workspace[32]
 * Return Type  : void
 */
void computeFirstOrderOpt(struct_T *solution, const b_struct_T *objective,
                          int workingset_nVar, int workingset_ldA,
                          const double workingset_ATwset[32],
                          int workingset_nActiveConstr, double workspace[32])
{
  int ia;
  int iac;
  int ix;
  int k;
  if (workingset_nVar - 1 >= 0) {
    memcpy(&workspace[0], &objective->grad[0],
           workingset_nVar * sizeof(double));
  }
  if ((workingset_nVar != 0) && (workingset_nActiveConstr != 0)) {
    ix = 0;
    k = workingset_ldA * (workingset_nActiveConstr - 1) + 1;
    for (iac = 1; workingset_ldA < 0 ? iac >= k : iac <= k;
         iac += workingset_ldA) {
      int i;
      i = (iac + workingset_nVar) - 1;
      for (ia = iac; ia <= i; ia++) {
        int i1;
        i1 = ia - iac;
        workspace[i1] += workingset_ATwset[ia - 1] * solution->lambda[ix];
      }
      ix++;
    }
  }
  if (workingset_nVar < 1) {
    ix = 0;
  } else {
    ix = 1;
    if (workingset_nVar > 1) {
      double smax;
      smax = fabs(workspace[0]);
      for (k = 2; k <= workingset_nVar; k++) {
        double s;
        s = fabs(workspace[k - 1]);
        if (s > smax) {
          ix = k;
          smax = s;
        }
      }
    }
  }
  solution->firstorderopt = fabs(workspace[ix - 1]);
}

/*
 * File trailer for computeFirstOrderOpt.c
 *
 * [EOF]
 */
