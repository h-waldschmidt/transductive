/*
 * File: xzgeqp3.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "xzgeqp3.h"
#include "rt_nonfinite.h"
#include "xnrm2.h"
#include "xzlarf.h"
#include "xzlarfg.h"
#include <math.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : double A[32]
 *                int m
 *                int n
 *                int nfxd
 *                double tau[4]
 * Return Type  : void
 */
void qrf(double A[32], int m, int n, int nfxd, double tau[4])
{
  double work[8];
  double atmp;
  int i;
  tau[0] = 0.0;
  tau[1] = 0.0;
  tau[2] = 0.0;
  tau[3] = 0.0;
  memset(&work[0], 0, 8U * sizeof(double));
  for (i = 0; i < nfxd; i++) {
    double d;
    int ii;
    int mmi;
    ii = (i << 2) + i;
    mmi = m - i;
    if (i + 1 < m) {
      atmp = A[ii];
      d = xzlarfg(mmi, &atmp, A, ii + 2);
      tau[i] = d;
      A[ii] = atmp;
    } else {
      d = 0.0;
      tau[i] = 0.0;
    }
    if (i + 1 < n) {
      atmp = A[ii];
      A[ii] = 1.0;
      xzlarf(mmi, (n - i) - 1, ii + 1, d, A, ii + 5, work);
      A[ii] = atmp;
    }
  }
}

/*
 * Arguments    : double A[32]
 *                int m
 *                int n
 *                int jpvt[8]
 *                double tau[4]
 * Return Type  : void
 */
void xzgeqp3(double A[32], int m, int n, int jpvt[8], double tau[4])
{
  double vn1[8];
  double vn2[8];
  double work[8];
  double temp;
  int b_i;
  int k;
  int minmn_tmp;
  int pvt;
  if (m <= n) {
    minmn_tmp = m;
  } else {
    minmn_tmp = n;
  }
  tau[0] = 0.0;
  tau[1] = 0.0;
  tau[2] = 0.0;
  tau[3] = 0.0;
  if (minmn_tmp < 1) {
    for (pvt = 0; pvt < n; pvt++) {
      jpvt[pvt] = pvt + 1;
    }
  } else {
    int i;
    int ix;
    int iy;
    int nfxd;
    int temp_tmp;
    nfxd = 0;
    for (pvt = 0; pvt < n; pvt++) {
      if (jpvt[pvt] != 0) {
        nfxd++;
        if (pvt + 1 != nfxd) {
          ix = pvt << 2;
          iy = (nfxd - 1) << 2;
          for (k = 0; k < m; k++) {
            temp_tmp = ix + k;
            temp = A[temp_tmp];
            i = iy + k;
            A[temp_tmp] = A[i];
            A[i] = temp;
          }
          jpvt[pvt] = jpvt[nfxd - 1];
          jpvt[nfxd - 1] = pvt + 1;
        } else {
          jpvt[pvt] = pvt + 1;
        }
      } else {
        jpvt[pvt] = pvt + 1;
      }
    }
    if (nfxd > minmn_tmp) {
      nfxd = minmn_tmp;
    }
    qrf(A, m, n, nfxd, tau);
    if (nfxd < minmn_tmp) {
      double d;
      int ia1_tmp;
      memset(&work[0], 0, 8U * sizeof(double));
      memset(&vn1[0], 0, 8U * sizeof(double));
      memset(&vn2[0], 0, 8U * sizeof(double));
      ia1_tmp = nfxd + 1;
      for (pvt = ia1_tmp; pvt <= n; pvt++) {
        d = xnrm2(m - nfxd, A, (nfxd + ((pvt - 1) << 2)) + 1);
        vn1[pvt - 1] = d;
        vn2[pvt - 1] = d;
      }
      for (b_i = ia1_tmp; b_i <= minmn_tmp; b_i++) {
        double s;
        int ii;
        int ip1;
        int mmi;
        int nmi;
        ip1 = b_i + 1;
        nfxd = ((b_i - 1) << 2) + 1;
        ii = (nfxd + b_i) - 2;
        nmi = (n - b_i) + 1;
        mmi = m - b_i;
        if (nmi < 1) {
          iy = -2;
        } else {
          iy = -1;
          if (nmi > 1) {
            temp = fabs(vn1[b_i - 1]);
            for (k = 2; k <= nmi; k++) {
              s = fabs(vn1[(b_i + k) - 2]);
              if (s > temp) {
                iy = k - 2;
                temp = s;
              }
            }
          }
        }
        pvt = b_i + iy;
        if (pvt + 1 != b_i) {
          ix = pvt << 2;
          for (k = 0; k < m; k++) {
            temp_tmp = ix + k;
            temp = A[temp_tmp];
            i = (nfxd + k) - 1;
            A[temp_tmp] = A[i];
            A[i] = temp;
          }
          iy = jpvt[pvt];
          jpvt[pvt] = jpvt[b_i - 1];
          jpvt[b_i - 1] = iy;
          vn1[pvt] = vn1[b_i - 1];
          vn2[pvt] = vn2[b_i - 1];
        }
        if (b_i < m) {
          temp = A[ii];
          d = xzlarfg(mmi + 1, &temp, A, ii + 2);
          tau[b_i - 1] = d;
          A[ii] = temp;
        } else {
          d = 0.0;
          tau[b_i - 1] = 0.0;
        }
        if (b_i < n) {
          temp = A[ii];
          A[ii] = 1.0;
          xzlarf(mmi + 1, nmi - 1, ii + 1, d, A, ii + 5, work);
          A[ii] = temp;
        }
        for (pvt = ip1; pvt <= n; pvt++) {
          iy = (b_i + ((pvt - 1) << 2)) + 1;
          d = vn1[pvt - 1];
          if (d != 0.0) {
            temp = fabs(A[iy - 2]) / d;
            temp = 1.0 - temp * temp;
            if (temp < 0.0) {
              temp = 0.0;
            }
            s = d / vn2[pvt - 1];
            s = temp * (s * s);
            if (s <= 1.4901161193847656E-8) {
              if (b_i < m) {
                d = xnrm2(mmi, A, iy);
                vn1[pvt - 1] = d;
                vn2[pvt - 1] = d;
              } else {
                vn1[pvt - 1] = 0.0;
                vn2[pvt - 1] = 0.0;
              }
            } else {
              vn1[pvt - 1] = d * sqrt(temp);
            }
          }
        }
      }
    }
  }
}

/*
 * File trailer for xzgeqp3.c
 *
 * [EOF]
 */
