/*
 * File: phaseone.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "phaseone.h"
#include "addBoundToActiveSetMatrix_.h"
#include "computeFval.h"
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
#include "qpsolver_data.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "setProblemType.h"
#include "squareQ_appendCol.h"
#include "xnrm2.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : struct_T *solution
 *                d_struct_T *memspace
 *                f_struct_T *workingset
 *                e_struct_T *qrmanager
 *                c_struct_T *cholmanager
 *                b_struct_T *objective
 *                g_struct_T *options
 * Return Type  : void
 */
void phaseone(struct_T *solution, d_struct_T *memspace, f_struct_T *workingset,
              e_struct_T *qrmanager, c_struct_T *cholmanager,
              b_struct_T *objective, g_struct_T *options)
{
  static const char t0_SolverName[8] = {'q', 'u', 'a', 'd', 'p', 'r', 'o', 'g'};
  static const char t0_FiniteDifferenceType[7] = {'f', 'o', 'r', 'w',
                                                  'a', 'r', 'd'};
  static const char t0_Display[5] = {'f', 'i', 'n', 'a', 'l'};
  b_struct_T b_objective;
  double minLambda;
  int PROBTYPE_ORIG;
  int TYPE;
  int activeSetChangeID;
  int globalActiveConstrIdx;
  int i;
  int i1;
  int idx;
  int idxMinLambda;
  int ixlast;
  int nVar_tmp_tmp;
  boolean_T subProblemChanged;
  boolean_T updateFval;
  options->InitDamping = 0.01;
  for (i = 0; i < 7; i++) {
    options->FiniteDifferenceType[i] = t0_FiniteDifferenceType[i];
  }
  options->SpecifyObjectiveGradient = false;
  options->ScaleProblem = false;
  options->SpecifyConstraintGradient = false;
  options->NonFiniteSupport = true;
  options->IterDisplaySQP = false;
  options->FiniteDifferenceStepSize = -1.0;
  options->MaxFunctionEvaluations = -1.0;
  options->IterDisplayQP = false;
  options->PricingTolerance = 0.0;
  for (i = 0; i < 10; i++) {
    options->Algorithm[i] = cv[i];
  }
  options->ConstraintTolerance = 1.0E-8;
  options->OptimalityTolerance = 1.0E-8;
  options->MaxIterations = -1.0;
  options->FunctionTolerance = rtInf;
  for (i = 0; i < 8; i++) {
    options->SolverName[i] = t0_SolverName[i];
  }
  options->CheckGradients = false;
  options->Diagnostics[0] = 'o';
  options->Diagnostics[1] = 'f';
  options->Diagnostics[2] = 'f';
  options->DiffMaxChange = rtInf;
  options->DiffMinChange = 0.0;
  for (i = 0; i < 5; i++) {
    options->Display[i] = t0_Display[i];
  }
  options->FunValCheck[0] = 'o';
  options->FunValCheck[1] = 'f';
  options->FunValCheck[2] = 'f';
  options->UseParallel = false;
  options->LinearSolver[0] = 'a';
  options->LinearSolver[1] = 'u';
  options->LinearSolver[2] = 't';
  options->LinearSolver[3] = 'o';
  options->SubproblemAlgorithm[0] = 'c';
  options->SubproblemAlgorithm[1] = 'g';
  PROBTYPE_ORIG = workingset->probType;
  nVar_tmp_tmp = workingset->nVar;
  solution->xstar[workingset->nVar] = solution->maxConstr + 1.0;
  if (workingset->probType == 3) {
    idxMinLambda = 1;
  } else {
    idxMinLambda = 4;
  }
  setProblemType(workingset, idxMinLambda);
  ixlast = (workingset->nWConstr[0] + workingset->nWConstr[1]) + 1;
  idxMinLambda = workingset->nActiveConstr;
  for (TYPE = ixlast; TYPE <= idxMinLambda; TYPE++) {
    workingset->isActiveConstr
        [(workingset->isActiveIdx[workingset->Wid[TYPE - 1] - 1] +
          workingset->Wlocalidx[TYPE - 1]) -
         2] = false;
  }
  workingset->nWConstr[2] = 0;
  workingset->nWConstr[3] = 0;
  workingset->nWConstr[4] = 0;
  workingset->nActiveConstr = workingset->nWConstr[0] + workingset->nWConstr[1];
  objective->grad[0] = 0.0;
  objective->grad[1] = 0.0;
  objective->grad[2] = 0.0;
  objective->grad[3] = 0.0;
  objective->Hx[0] = 0.0;
  objective->Hx[1] = 0.0;
  objective->Hx[2] = 0.0;
  objective->maxVar = 4;
  objective->beta = 0.0;
  objective->rho = 0.0;
  objective->nvar = nVar_tmp_tmp + 1;
  objective->prev_objtype = 3;
  objective->prev_nvar = 3;
  objective->prev_hasLinear = true;
  objective->objtype = 5;
  objective->gammaScalar = 1.0;
  objective->hasLinear = true;
  solution->fstar =
      computeFval(objective, memspace->workspace_double, solution->xstar);
  subProblemChanged = true;
  updateFval = true;
  activeSetChangeID = 0;
  i = workingset->nVar;
  globalActiveConstrIdx = 0;
  b_objective = *objective;
  computeGrad_StoreHx(&b_objective, solution->xstar);
  solution->fstar = computeFval_ReuseHx(
      &b_objective, memspace->workspace_double, solution->xstar);
  if (solution->iterations < 100) {
    solution->state = -5;
  } else {
    solution->state = 0;
  }
  ixlast = workingset->mConstrMax;
  if (ixlast - 1 >= 0) {
    memset(&solution->lambda[0], 0, ixlast * sizeof(double));
  }
  int exitg1;
  do {
    exitg1 = 0;
    if (solution->state == -5) {
      boolean_T guard1 = false;
      boolean_T guard2 = false;
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
          factorQR(qrmanager, workingset->ATwset, i, workingset->nActiveConstr,
                   workingset->ldA);
          computeQ_(qrmanager, qrmanager->mrows);
          break;
        }
        compute_deltax(solution, memspace, qrmanager, cholmanager,
                       &b_objective);
        if (solution->state != -5) {
          exitg1 = 1;
        } else if ((b_xnrm2(i, solution->searchDir) < 1.4901161193847657E-10) ||
                   (workingset->nActiveConstr >= i)) {
          guard2 = true;
        } else {
          feasibleratiotest(solution->xstar, solution->searchDir,
                            memspace->workspace_double, workingset->nVar,
                            workingset->lb, workingset->ub, workingset->indexLB,
                            workingset->indexUB, workingset->sizes,
                            workingset->isActiveIdx, workingset->isActiveConstr,
                            workingset->nWConstr, true, 1.0E-8, &minLambda,
                            &updateFval, &i1, &ixlast);
          if (updateFval) {
            switch (i1) {
            case 3:
              workingset->nWConstr[2]++;
              workingset
                  ->isActiveConstr[(workingset->isActiveIdx[2] + ixlast) - 2] =
                  true;
              workingset->nActiveConstr++;
              workingset->Wid[workingset->nActiveConstr - 1] = 3;
              workingset->Wlocalidx[workingset->nActiveConstr - 1] = ixlast;
              /* A check that is always false is detected at compile-time.
               * Eliminating code that follows. */
              break;
            case 4:
              addBoundToActiveSetMatrix_(workingset, 4, ixlast);
              break;
            default:
              addBoundToActiveSetMatrix_(workingset, 5, ixlast);
              break;
            }
            activeSetChangeID = 1;
          } else {
            if (b_objective.objtype == 5) {
              if (b_xnrm2(b_objective.nvar, solution->searchDir) >
                  100.0 * (double)b_objective.nvar * 1.4901161193847656E-8) {
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
          if ((i >= 1) && (!(minLambda == 0.0))) {
            ixlast = i - 1;
            for (TYPE = 0; TYPE <= ixlast; TYPE++) {
              solution->xstar[TYPE] += minLambda * solution->searchDir[TYPE];
            }
          }
          computeGrad_StoreHx(&b_objective, solution->xstar);
          updateFval = true;
          guard1 = true;
        }
      } else {
        if (i - 1 >= 0) {
          memset(&solution->searchDir[0], 0, i * sizeof(double));
        }
        guard2 = true;
      }
      if (guard2) {
        compute_lambda(memspace->workspace_double, solution, &b_objective,
                       qrmanager);
        if ((solution->state != -7) || (workingset->nActiveConstr > i)) {
          idxMinLambda = -1;
          minLambda = 0.0;
          i1 = (workingset->nWConstr[0] + workingset->nWConstr[1]) + 1;
          ixlast = workingset->nActiveConstr;
          for (idx = i1; idx <= ixlast; idx++) {
            double d;
            d = solution->lambda[idx - 1];
            if (d < minLambda) {
              minLambda = d;
              idxMinLambda = idx - 1;
            }
          }
          if (idxMinLambda + 1 == 0) {
            solution->state = 1;
          } else {
            activeSetChangeID = -1;
            globalActiveConstrIdx = idxMinLambda + 1;
            subProblemChanged = true;
            TYPE = workingset->Wid[idxMinLambda] - 1;
            workingset->isActiveConstr
                [(workingset->isActiveIdx[workingset->Wid[idxMinLambda] - 1] +
                  workingset->Wlocalidx[idxMinLambda]) -
                 2] = false;
            workingset->Wid[idxMinLambda] =
                workingset->Wid[workingset->nActiveConstr - 1];
            workingset->Wlocalidx[idxMinLambda] =
                workingset->Wlocalidx[workingset->nActiveConstr - 1];
            i1 = workingset->nVar;
            for (idx = 0; idx < i1; idx++) {
              workingset->ATwset[idx + workingset->ldA * idxMinLambda] =
                  workingset->ATwset[idx + workingset->ldA *
                                               (workingset->nActiveConstr - 1)];
            }
            workingset->bwset[idxMinLambda] =
                workingset->bwset[workingset->nActiveConstr - 1];
            workingset->nActiveConstr--;
            workingset->nWConstr[TYPE]--;
            solution->lambda[idxMinLambda] = 0.0;
          }
        } else {
          i1 = workingset->nActiveConstr;
          activeSetChangeID = 0;
          globalActiveConstrIdx = workingset->nActiveConstr;
          subProblemChanged = true;
          ixlast = workingset->nActiveConstr - 1;
          idxMinLambda = workingset->Wid[ixlast] - 1;
          workingset->isActiveConstr[(workingset->isActiveIdx[idxMinLambda] +
                                      workingset->Wlocalidx[ixlast]) -
                                     2] = false;
          workingset->nActiveConstr--;
          workingset->nWConstr[idxMinLambda]--;
          solution->lambda[i1 - 1] = 0.0;
        }
        updateFval = false;
        guard1 = true;
      }
      if (guard1) {
        solution->iterations++;
        ixlast = b_objective.nvar - 1;
        if ((solution->iterations >= 100) &&
            ((solution->state != 1) || (b_objective.objtype == 5))) {
          solution->state = 0;
        }
        if (solution->iterations - solution->iterations / 50 * 50 == 0) {
          solution->maxConstr =
              maxConstraintViolation(workingset, solution->xstar);
          minLambda = solution->maxConstr;
          if (b_objective.objtype == 5) {
            minLambda =
                solution->maxConstr - solution->xstar[b_objective.nvar - 1];
          }
          if (minLambda > 3.0000000000000004E-8) {
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
        if (updateFval) {
          solution->fstar = computeFval_ReuseHx(
              &b_objective, memspace->workspace_double, solution->xstar);
          if ((solution->fstar < 3.0000000000000004E-8) &&
              ((solution->state != 0) || (b_objective.objtype != 5))) {
            solution->state = 2;
          }
        }
      }
    } else {
      if (!updateFval) {
        solution->fstar = computeFval_ReuseHx(
            &b_objective, memspace->workspace_double, solution->xstar);
      }
      exitg1 = 1;
    }
  } while (exitg1 == 0);
  if (workingset
          ->isActiveConstr[(workingset->isActiveIdx[3] + workingset->sizes[3]) -
                           2]) {
    boolean_T exitg2;
    idx = workingset->sizes[0] + workingset->sizes[1];
    exitg2 = false;
    while ((!exitg2) && (idx + 1 <= workingset->nActiveConstr)) {
      if ((workingset->Wid[idx] == 4) &&
          (workingset->Wlocalidx[idx] == workingset->sizes[3])) {
        TYPE = workingset->Wid[idx] - 1;
        workingset->isActiveConstr
            [(workingset->isActiveIdx[workingset->Wid[idx] - 1] +
              workingset->Wlocalidx[idx]) -
             2] = false;
        workingset->Wid[idx] = workingset->Wid[workingset->nActiveConstr - 1];
        workingset->Wlocalidx[idx] =
            workingset->Wlocalidx[workingset->nActiveConstr - 1];
        i = workingset->nVar;
        for (ixlast = 0; ixlast < i; ixlast++) {
          workingset->ATwset[ixlast + workingset->ldA * idx] =
              workingset->ATwset[ixlast + workingset->ldA *
                                              (workingset->nActiveConstr - 1)];
        }
        workingset->bwset[idx] =
            workingset->bwset[workingset->nActiveConstr - 1];
        workingset->nActiveConstr--;
        workingset->nWConstr[TYPE]--;
        exitg2 = true;
      } else {
        idx++;
      }
    }
  }
  idxMinLambda = workingset->nActiveConstr - 1;
  ixlast = workingset->sizes[0] + workingset->sizes[1];
  while ((idxMinLambda + 1 > ixlast) && (idxMinLambda + 1 > nVar_tmp_tmp)) {
    TYPE = workingset->Wid[idxMinLambda] - 1;
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
          workingset
              ->ATwset[idx + workingset->ldA * (workingset->nActiveConstr - 1)];
    }
    workingset->bwset[idxMinLambda] =
        workingset->bwset[workingset->nActiveConstr - 1];
    workingset->nActiveConstr--;
    workingset->nWConstr[TYPE]--;
    idxMinLambda--;
  }
  solution->maxConstr = solution->xstar[nVar_tmp_tmp];
  setProblemType(workingset, PROBTYPE_ORIG);
  *objective = b_objective;
  objective->objtype = b_objective.prev_objtype;
  objective->nvar = b_objective.prev_nvar;
  objective->hasLinear = b_objective.prev_hasLinear;
  options->ObjectiveLimit = -1.0E+20;
  options->StepTolerance = 1.0E-8;
}

/*
 * File trailer for phaseone.c
 *
 * [EOF]
 */
