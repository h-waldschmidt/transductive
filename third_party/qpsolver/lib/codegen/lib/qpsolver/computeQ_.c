/*
 * File: computeQ_.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "computeQ_.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : e_struct_T *obj
 *                int nrows
 * Return Type  : void
 */
void computeQ_(e_struct_T *obj, int nrows)
{
  double work[4];
  int b_i;
  int i;
  int iQR0;
  int ia;
  int idx;
  int lda;
  int m;
  int n;
  i = obj->minRowCol;
  for (idx = 0; idx < i; idx++) {
    iQR0 = obj->ldq * idx + idx;
    n = obj->mrows - idx;
    if (n - 2 >= 0) {
      memcpy(&obj->Q[iQR0 + 1], &obj->QR[iQR0 + 1],
             (((n + iQR0) - iQR0) - 1) * sizeof(double));
    }
  }
  m = obj->mrows;
  n = obj->minRowCol;
  lda = obj->ldq;
  if (nrows >= 1) {
    int i1;
    int itau;
    i = nrows - 1;
    for (idx = n; idx <= i; idx++) {
      ia = idx * lda;
      i1 = m - 1;
      if (i1 >= 0) {
        memset(&obj->Q[ia], 0, (((i1 + ia) - ia) + 1) * sizeof(double));
      }
      obj->Q[ia + idx] = 1.0;
    }
    itau = obj->minRowCol - 1;
    work[0] = 0.0;
    work[1] = 0.0;
    work[2] = 0.0;
    work[3] = 0.0;
    for (b_i = obj->minRowCol; b_i >= 1; b_i--) {
      int iaii;
      iaii = (b_i + (b_i - 1) * lda) - 1;
      if (b_i < nrows) {
        int jA;
        int lastc;
        int lastv;
        obj->Q[iaii] = 1.0;
        jA = (iaii + lda) + 1;
        if (obj->tau[itau] != 0.0) {
          boolean_T exitg2;
          lastv = m - b_i;
          iQR0 = (iaii + m) - b_i;
          while ((lastv + 1 > 0) && (obj->Q[iQR0] == 0.0)) {
            lastv--;
            iQR0--;
          }
          lastc = (nrows - b_i) - 1;
          exitg2 = false;
          while ((!exitg2) && (lastc + 1 > 0)) {
            int exitg1;
            iQR0 = jA + lastc * lda;
            ia = iQR0;
            do {
              exitg1 = 0;
              if (ia <= iQR0 + lastv) {
                if (obj->Q[ia - 1] != 0.0) {
                  exitg1 = 1;
                } else {
                  ia++;
                }
              } else {
                lastc--;
                exitg1 = 2;
              }
            } while (exitg1 == 0);
            if (exitg1 == 1) {
              exitg2 = true;
            }
          }
        } else {
          lastv = -1;
          lastc = -1;
        }
        if (lastv + 1 > 0) {
          double c;
          if (lastc + 1 != 0) {
            if (lastc >= 0) {
              memset(&work[0], 0, (lastc + 1) * sizeof(double));
            }
            iQR0 = 0;
            i = jA + lda * lastc;
            for (n = jA; lda < 0 ? n >= i : n <= i; n += lda) {
              c = 0.0;
              i1 = n + lastv;
              for (ia = n; ia <= i1; ia++) {
                c += obj->Q[ia - 1] * obj->Q[(iaii + ia) - n];
              }
              work[iQR0] += c;
              iQR0++;
            }
          }
          if (!(-obj->tau[itau] == 0.0)) {
            for (idx = 0; idx <= lastc; idx++) {
              c = work[idx];
              if (c != 0.0) {
                c *= -obj->tau[itau];
                i = lastv + jA;
                for (iQR0 = jA; iQR0 <= i; iQR0++) {
                  obj->Q[iQR0 - 1] += obj->Q[(iaii + iQR0) - jA] * c;
                }
              }
              jA += lda;
            }
          }
        }
      }
      if (b_i < m) {
        iQR0 = iaii + 2;
        i = ((iaii + m) - b_i) + 1;
        for (n = iQR0; n <= i; n++) {
          obj->Q[n - 1] *= -obj->tau[itau];
        }
      }
      obj->Q[iaii] = 1.0 - obj->tau[itau];
      for (idx = 0; idx <= b_i - 2; idx++) {
        obj->Q[(iaii - idx) - 1] = 0.0;
      }
      itau--;
    }
  }
}

/*
 * File trailer for computeQ_.c
 *
 * [EOF]
 */
