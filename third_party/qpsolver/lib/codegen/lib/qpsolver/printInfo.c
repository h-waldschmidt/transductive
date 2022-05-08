/*
 * File: printInfo.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "printInfo.h"
#include "rt_nonfinite.h"
#include "rt_nonfinite.h"
#include <stdio.h>

/* Function Definitions */
/*
 * Arguments    : boolean_T newBlocking
 *                int PROBLEM_TYPE
 *                double alpha
 *                double stepNorm
 *                int activeConstrChangedType
 *                int localActiveConstrIdx
 *                int activeSetChangeID
 *                double solution_fstar
 *                double solution_firstorderopt
 *                double solution_maxConstr
 *                int solution_iterations
 *                const int workingset_indexLB[4]
 *                const int workingset_indexUB[4]
 *                int workingset_nActiveConstr
 * Return Type  : void
 */
void printInfo(boolean_T newBlocking, int PROBLEM_TYPE, double alpha,
               double stepNorm, int activeConstrChangedType,
               int localActiveConstrIdx, int activeSetChangeID,
               double solution_fstar, double solution_firstorderopt,
               double solution_maxConstr, int solution_iterations,
               const int workingset_indexLB[4], const int workingset_indexUB[4],
               int workingset_nActiveConstr)
{
  printf("%5i  %14.6e  %14.6e  %14.6e", solution_iterations, solution_fstar,
         solution_firstorderopt, solution_maxConstr);
  fflush(stdout);
  printf("  ");
  fflush(stdout);
  if (rtIsNaN(alpha)) {
    printf("       -      ");
    fflush(stdout);
  } else {
    printf("%14.6e", alpha);
    fflush(stdout);
  }
  printf("  ");
  fflush(stdout);
  printf("%14.6e", stepNorm);
  fflush(stdout);
  printf("    ");
  fflush(stdout);
  if (newBlocking || (activeSetChangeID == -1)) {
    int b_localActiveConstrIdx;
    if (newBlocking) {
      activeSetChangeID = 1;
    }
    b_localActiveConstrIdx = localActiveConstrIdx;
    switch (activeSetChangeID) {
    case -1:
      printf("-");
      fflush(stdout);
      break;
    case 1:
      printf("+");
      fflush(stdout);
      break;
    default:
      printf(" ");
      fflush(stdout);
      break;
    }
    switch (activeConstrChangedType) {
    case 3:
      printf("AINEQ");
      fflush(stdout);
      break;
    case 4:
      printf("LOWER");
      fflush(stdout);
      b_localActiveConstrIdx = workingset_indexLB[localActiveConstrIdx - 1];
      break;
    case 5:
      printf("UPPER");
      fflush(stdout);
      b_localActiveConstrIdx = workingset_indexUB[localActiveConstrIdx - 1];
      break;
    default:
      printf("SAME ");
      fflush(stdout);
      b_localActiveConstrIdx = -1;
      break;
    }
    printf("(%-5i)", b_localActiveConstrIdx);
    fflush(stdout);
  } else {
    printf(" SAME ");
    fflush(stdout);
    printf("(%-5i)", -1);
    fflush(stdout);
  }
  printf("           ");
  fflush(stdout);
  printf("%5i", workingset_nActiveConstr);
  fflush(stdout);
  printf("    ");
  fflush(stdout);
  switch (PROBLEM_TYPE) {
  case 1:
    printf("Phase One");
    fflush(stdout);
    break;
  case 2:
    printf("Regularized");
    fflush(stdout);
    break;
  case 4:
    printf("Phase One Reg");
    fflush(stdout);
    break;
  default:
    printf("Normal");
    fflush(stdout);
    break;
  }
  printf("\n");
  fflush(stdout);
}

/*
 * File trailer for printInfo.c
 *
 * [EOF]
 */
