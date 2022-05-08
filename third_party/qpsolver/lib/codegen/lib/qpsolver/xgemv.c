/*
 * File: xgemv.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "xgemv.h"
#include "rt_nonfinite.h"

/* Function Definitions */
/*
 * Arguments    : int m
 *                int n
 *                const double A[4]
 *                int lda
 *                const double x[32]
 *                double y[8]
 * Return Type  : void
 */
void b_xgemv(int m, int n, const double A[4], int lda, const double x[32],
             double y[8])
{
  int ia;
  int iac;
  int iy;
  if ((m != 0) && (n != 0)) {
    int i;
    for (iy = 0; iy < n; iy++) {
      y[iy] = -y[iy];
    }
    iy = 0;
    i = lda * (n - 1) + 1;
    for (iac = 1; lda < 0 ? iac >= i : iac <= i; iac += lda) {
      double c;
      int i1;
      c = 0.0;
      i1 = (iac + m) - 1;
      for (ia = iac; ia <= i1; ia++) {
        c += A[ia - 1] * x[(ia - iac) + 8];
      }
      y[iy] += c;
      iy++;
    }
  }
}

/*
 * Arguments    : int m
 *                int n
 *                const double A[4]
 *                int lda
 *                const double x[4]
 *                double y[8]
 * Return Type  : void
 */
void c_xgemv(int m, int n, const double A[4], int lda, const double x[4],
             double y[8])
{
  int ia;
  int iac;
  int iy;
  if ((m != 0) && (n != 0)) {
    int i;
    for (iy = 0; iy < n; iy++) {
      y[iy] = -y[iy];
    }
    iy = 0;
    i = lda * (n - 1) + 1;
    for (iac = 1; lda < 0 ? iac >= i : iac <= i; iac += lda) {
      double c;
      int i1;
      c = 0.0;
      i1 = (iac + m) - 1;
      for (ia = iac; ia <= i1; ia++) {
        c += A[ia - 1] * x[ia - iac];
      }
      y[iy] += c;
      iy++;
    }
  }
}

/*
 * Arguments    : int m
 *                int n
 *                const double A[4]
 *                int lda
 *                const double x[32]
 *                double y[8]
 * Return Type  : void
 */
void xgemv(int m, int n, const double A[4], int lda, const double x[32],
           double y[8])
{
  int ia;
  int iac;
  int iy;
  if ((m != 0) && (n != 0)) {
    int i;
    for (iy = 0; iy < n; iy++) {
      y[iy] = -y[iy];
    }
    iy = 0;
    i = lda * (n - 1) + 1;
    for (iac = 1; lda < 0 ? iac >= i : iac <= i; iac += lda) {
      double c;
      int i1;
      c = 0.0;
      i1 = (iac + m) - 1;
      for (ia = iac; ia <= i1; ia++) {
        c += A[ia - 1] * x[ia - iac];
      }
      y[iy] += c;
      iy++;
    }
  }
}

/*
 * File trailer for xgemv.c
 *
 * [EOF]
 */
