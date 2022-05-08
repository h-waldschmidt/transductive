/*
 * File: iterate.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "iterate.h"
#include "addBoundToActiveSetMatrix_.h"
#include "computeFirstOrderOpt.h"
#include "computeFval_ReuseHx.h"
#include "computeGrad_StoreHx.h"
#include "computeQ_.h"
#include "compute_deltax.h"
#include "compute_lambda.h"
#include "deleteColMoveEnd.h"
#include "factorQR.h"
#include "feasibleX0ForWorkingSet.h"
#include "feasibleratiotest.h"
#include "maxConstraintViolation.h"
#include "printInfo.h"
#include "qpsolver_internal_types.h"
#include "ratiotest.h"
#include "rt_nonfinite.h"
#include "squareQ_appendCol.h"
#include "xnrm2.h"
#include <stdio.h>
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : struct_T *solution
 *                d_struct_T *memspace
 *                f_struct_T *workingset
 *                e_struct_T *qrmanager
 *                c_struct_T *cholmanager
 *                b_struct_T *objective
 *                boolean_T options_IterDisplayQP
 *                double options_PricingTolerance
 *                double options_ObjectiveLimit
 *                double options_ConstraintTolerance
 *                double options_StepTolerance
 *                boolean_T runTimeOptions_RemainFeasible
 * Return Type  : void
 */
void iterate(struct_T *solution, d_struct_T *memspace, f_struct_T *workingset,
             e_struct_T *qrmanager, c_struct_T *cholmanager,
             b_struct_T *objective, boolean_T options_IterDisplayQP,
             double options_PricingTolerance, double options_ObjectiveLimit,
             double options_ConstraintTolerance, double options_StepTolerance,
             boolean_T runTimeOptions_RemainFeasible)
{
  static const char b_cv[13] = {'N', 'o', 'r', 'm', 'a', 'l', ' ',
                                ' ', ' ', ' ', ' ', ' ', ' '};
  static const char cv1[13] = {'P', 'h', 'a', 's', 'e', ' ', 'O',
                               'n', 'e', ' ', ' ', ' ', ' '};
  static const char cv2[13] = {'R', 'e', 'g', 'u', 'l', 'a', 'r',
                               'i', 'z', 'e', 'd', ' ', ' '};
  static const char cv3[13] = {'P', 'h', 'a', 's', 'e', ' ', 'O',
                               'n', 'e', ' ', 'R', 'e', 'g'};
  double alpha;
  double normDelta;
  double tolDelta;
  int TYPE;
  int activeConstrChangedType;
  int activeSetChangeID;
  int globalActiveConstrIdx;
  int i;
  int idx;
  int idxMinLambda;
  int ixlast;
  int localActiveConstrIdx;
  int nVar_tmp_tmp;
  boolean_T newBlocking;
  boolean_T subProblemChanged;
  boolean_T updateFval;
  subProblemChanged = true;
  updateFval = true;
  activeSetChangeID = 0;
  TYPE = objective->objtype;
  tolDelta = 6.7434957617430445E-7;
  nVar_tmp_tmp = workingset->nVar;
  activeConstrChangedType = 1;
  localActiveConstrIdx = 0;
  globalActiveConstrIdx = 0;
  computeGrad_StoreHx(objective, solution->xstar);
  solution->fstar = computeFval_ReuseHx(objective, memspace->workspace_double,
                                        solution->xstar);
  if (solution->iterations < 100) {
    solution->state = -5;
  } else {
    solution->state = 0;
  }
  ixlast = workingset->mConstrMax;
  if (ixlast - 1 >= 0) {
    memset(&solution->lambda[0], 0, ixlast * sizeof(double));
  }
  if ((solution->iterations == 0) && options_IterDisplayQP) {
    char varargin_4[14];
    char stepType_str[13];
    printf("                          First-order                              "
           "                                            Active\n");
    fflush(stdout);
    printf(" Iter            Fval      Optimality     Feasibility           "
           "alpha    Norm of step           Action     Constraints    Step T"
           "ype\n");
    fflush(stdout);
    printf("\n");
    fflush(stdout);
    switch (workingset->probType) {
    case 1:
      for (i = 0; i < 13; i++) {
        stepType_str[i] = cv1[i];
      }
      break;
    case 2:
      for (i = 0; i < 13; i++) {
        stepType_str[i] = cv2[i];
      }
      break;
    case 4:
      for (i = 0; i < 13; i++) {
        stepType_str[i] = cv3[i];
      }
      break;
    default:
      for (i = 0; i < 13; i++) {
        stepType_str[i] = b_cv[i];
      }
      break;
    }
    for (i = 0; i < 13; i++) {
      varargin_4[i] = stepType_str[i];
    }
    varargin_4[13] = '\x00';
    printf("%5i  %14.6e                                                        "
           "                                    %5i    %s",
           0, solution->fstar, workingset->nActiveConstr, &varargin_4[0]);
    fflush(stdout);
    printf("\n");
    fflush(stdout);
  }
  int exitg1;
  do {
    exitg1 = 0;
    if (solution->state == -5) {
      double minLambda;
      boolean_T guard1 = false;
      boolean_T guard2 = false;
      newBlocking = false;
      guard1 = false;
      guard2 = false;
      if (subProblemChanged) {
        switch (activeSetChangeID) {
        case 1:
          squareQ_appendCol(qrmanager, workingset->ATwset,
                            workingset->ldA * (workingset->nActiveConstr - 1) +
                                1);
          break;
        case -1:
          deleteColMoveEnd(qrmanager, globalActiveConstrIdx);
          break;
        default:
          factorQR(qrmanager, workingset->ATwset, nVar_tmp_tmp,
                   workingset->nActiveConstr, workingset->ldA);
          computeQ_(qrmanager, qrmanager->mrows);
          break;
        }
        compute_deltax(solution, memspace, qrmanager, cholmanager, objective);
        if (solution->state != -5) {
          exitg1 = 1;
        } else {
          normDelta = b_xnrm2(nVar_tmp_tmp, solution->searchDir);
          if ((normDelta < options_StepTolerance) ||
              (workingset->nActiveConstr >= nVar_tmp_tmp)) {
            guard2 = true;
          } else {
            updateFval = (TYPE == 5);
            if (updateFval || runTimeOptions_RemainFeasible) {
              feasibleratiotest(
                  solution->xstar, solution->searchDir,
                  memspace->workspace_double, workingset->nVar, workingset->lb,
                  workingset->ub, workingset->indexLB, workingset->indexUB,
                  workingset->sizes, workingset->isActiveIdx,
                  workingset->isActiveConstr, workingset->nWConstr, updateFval,
                  options_ConstraintTolerance, &alpha, &newBlocking,
                  &activeConstrChangedType, &localActiveConstrIdx);
            } else {
              ratiotest(solution->xstar, solution->searchDir,
                        memspace->workspace_double, workingset->nVar,
                        workingset->lb, workingset->ub, workingset->indexLB,
                        workingset->indexUB, workingset->sizes,
                        workingset->isActiveIdx, workingset->isActiveConstr,
                        workingset->nWConstr, options_ConstraintTolerance,
                        &tolDelta, &alpha, &newBlocking,
                        &activeConstrChangedType, &localActiveConstrIdx);
            }
            if (newBlocking) {
              switch (activeConstrChangedType) {
              case 3:
                workingset->nWConstr[2]++;
                workingset->isActiveConstr[(workingset->isActiveIdx[2] +
                                            localActiveConstrIdx) -
                                           2] = true;
                workingset->nActiveConstr++;
                workingset->Wid[workingset->nActiveConstr - 1] = 3;
                workingset->Wlocalidx[workingset->nActiveConstr - 1] =
                    localActiveConstrIdx;
                /* A check that is always false is detected at compile-time.
                 * Eliminating code that follows. */
                break;
              case 4:
                addBoundToActiveSetMatrix_(workingset, 4, localActiveConstrIdx);
                break;
              default:
                addBoundToActiveSetMatrix_(workingset, 5, localActiveConstrIdx);
                break;
              }
              activeSetChangeID = 1;
            } else {
              if (objective->objtype == 5) {
                if (b_xnrm2(objective->nvar, solution->searchDir) >
                    100.0 * (double)objective->nvar * 1.4901161193847656E-8) {
                  solution->state = 3;
                } else {
                  solution->state = 4;
                }
              }
              subProblemChanged = false;
              if (workingset->nActiveConstr == 0) {
                solution->state = 1;
              }
            }
            if ((nVar_tmp_tmp >= 1) && (!(alpha == 0.0))) {
              ixlast = nVar_tmp_tmp - 1;
              for (idxMinLambda = 0; idxMinLambda <= ixlast; idxMinLambda++) {
                solution->xstar[idxMinLambda] +=
                    alpha * solution->searchDir[idxMinLambda];
              }
            }
            computeGrad_StoreHx(objective, solution->xstar);
            updateFval = true;
            guard1 = true;
          }
        }
      } else {
        if (nVar_tmp_tmp - 1 >= 0) {
          memset(&solution->searchDir[0], 0, nVar_tmp_tmp * sizeof(double));
        }
        normDelta = 0.0;
        guard2 = true;
      }
      if (guard2) {
        compute_lambda(memspace->workspace_double, solution, objective,
                       qrmanager);
        if ((solution->state != -7) ||
            (workingset->nActiveConstr > nVar_tmp_tmp)) {
          idxMinLambda = -1;
          minLambda = options_PricingTolerance * 7.0 * (double)(TYPE != 5);
          i = (workingset->nWConstr[0] + workingset->nWConstr[1]) + 1;
          ixlast = workingset->nActiveConstr;
          for (idx = i; idx <= ixlast; idx++) {
            alpha = solution->lambda[idx - 1];
            if (alpha < minLambda) {
              minLambda = alpha;
              idxMinLambda = idx - 1;
            }
          }
          if (idxMinLambda + 1 == 0) {
            solution->state = 1;
          } else {
            activeSetChangeID = -1;
            globalActiveConstrIdx = idxMinLambda + 1;
            subProblemChanged = true;
            activeConstrChangedType = workingset->Wid[idxMinLambda];
            localActiveConstrIdx = workingset->Wlocalidx[idxMinLambda];
            ixlast = workingset->Wid[idxMinLambda] - 1;
            workingset->isActiveConstr
                [(workingset->isActiveIdx[workingset->Wid[idxMinLambda] - 1] +
                  workingset->Wlocalidx[idxMinLambda]) -
                 2] = false;
            workingset->Wid[idxMinLambda] =
                workingset->Wid[workingset->nActiveConstr - 1];
            workingset->Wlocalidx[idxMinLambda] =
                workingset->Wlocalidx[workingset->nActiveConstr - 1];
            i = workingset->nVar;
            for (idx = 0; idx < i; idx++) {
              workingset->ATwset[idx + workingset->ldA * idxMinLambda] =
                  workingset->ATwset[idx + workingset->ldA *
                                               (workingset->nActiveConstr - 1)];
            }
            workingset->bwset[idxMinLambda] =
                workingset->bwset[workingset->nActiveConstr - 1];
            workingset->nActiveConstr--;
            workingset->nWConstr[ixlast]--;
            solution->lambda[idxMinLambda] = 0.0;
          }
        } else {
          idxMinLambda = workingset->nActiveConstr;
          activeSetChangeID = 0;
          globalActiveConstrIdx = workingset->nActiveConstr;
          subProblemChanged = true;
          ixlast = workingset->nActiveConstr - 1;
          activeConstrChangedType = workingset->Wid[ixlast];
          localActiveConstrIdx = workingset->Wlocalidx[ixlast];
          workingset->isActiveConstr
              [(workingset->isActiveIdx[activeConstrChangedType - 1] +
                localActiveConstrIdx) -
               2] = false;
          workingset->nActiveConstr--;
          workingset->nWConstr[activeConstrChangedType - 1]--;
          solution->lambda[idxMinLambda - 1] = 0.0;
        }
        updateFval = false;
        alpha = rtNaN;
        guard1 = true;
      }
      if (guard1) {
        solution->iterations++;
        ixlast = objective->nvar - 1;
        if ((solution->iterations >= 100) &&
            ((solution->state != 1) || (objective->objtype == 5))) {
          solution->state = 0;
        }
        if (solution->iterations - solution->iterations / 50 * 50 == 0) {
          solution->maxConstr =
              maxConstraintViolation(workingset, solution->xstar);
          minLambda = solution->maxConstr;
          if (objective->objtype == 5) {
            minLambda = solution->maxConstr - solution->xstar[ixlast];
          }
          if (minLambda > options_ConstraintTolerance * 3.0) {
            boolean_T nonDegenerateWset;
            if (ixlast >= 0) {
              memcpy(&solution->searchDir[0], &solution->xstar[0],
                     (ixlast + 1) * sizeof(double));
            }
            nonDegenerateWset = feasibleX0ForWorkingSet(
                memspace->workspace_double, solution->searchDir, workingset,
                qrmanager);
            if ((!nonDegenerateWset) && (solution->state != 0)) {
              solution->state = -2;
            }
            activeSetChangeID = 0;
            minLambda = maxConstraintViolation(workingset, solution->searchDir);
            if (minLambda < solution->maxConstr) {
              if (ixlast >= 0) {
                memcpy(&solution->xstar[0], &solution->searchDir[0],
                       (ixlast + 1) * sizeof(double));
              }
              solution->maxConstr = minLambda;
            }
          }
        }
        if (updateFval &&
            ((options_ObjectiveLimit > rtMinusInf) || options_IterDisplayQP)) {
          solution->fstar = computeFval_ReuseHx(
              objective, memspace->workspace_double, solution->xstar);
          if ((options_ObjectiveLimit > rtMinusInf) &&
              (solution->fstar < options_ObjectiveLimit) &&
              ((solution->state != 0) || (objective->objtype != 5))) {
            solution->state = 2;
          }
        }
        if (options_IterDisplayQP) {
          if (solution->iterations - solution->iterations / 50 * 50 == 0) {
            printf(
                "                          First-order                         "
                "                                                 Active\n");
            fflush(stdout);
            printf(" Iter            Fval      Optimality     Feasibility      "
                   "     alpha    Norm of step           Action     "
                   "Constraints    Step T"
                   "ype\n");
            fflush(stdout);
            printf("\n");
            fflush(stdout);
          } else {
            solution->maxConstr =
                maxConstraintViolation(workingset, solution->xstar);
          }
          computeFirstOrderOpt(solution, objective, workingset->nVar,
                               workingset->ldA, workingset->ATwset,
                               workingset->nActiveConstr,
                               memspace->workspace_double);
          printInfo(newBlocking, workingset->probType, alpha, normDelta,
                    activeConstrChangedType, localActiveConstrIdx,
                    activeSetChangeID, solution->fstar, solution->firstorderopt,
                    solution->maxConstr, solution->iterations,
                    workingset->indexLB, workingset->indexUB,
                    workingset->nActiveConstr);
        }
      }
    } else {
      if (!updateFval) {
        solution->fstar = computeFval_ReuseHx(
            objective, memspace->workspace_double, solution->xstar);
      }
      exitg1 = 1;
    }
  } while (exitg1 == 0);
}

/*
 * File trailer for iterate.c
 *
 * [EOF]
 */
