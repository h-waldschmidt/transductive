/*
 * File: qpsolver.c
 *
 * MATLAB Coder version            : 5.4
 * C/C++ source code generated on  : 07-May-2022 19:38:44
 */

/* Include Files */
#include "qpsolver.h"
#include "PresolveWorkingSet.h"
#include "computeFirstOrderOpt.h"
#include "computeFval.h"
#include "feasibleX0ForWorkingSet.h"
#include "iterate.h"
#include "maxConstraintViolation.h"
#include "phaseone.h"
#include "qpsolver_data.h"
#include "qpsolver_internal_types.h"
#include "rt_nonfinite.h"
#include "setProblemType.h"
#include <string.h>

/* Function Definitions */
/*
 * Arguments    : double x[3]
 *                double *fval
 * Return Type  : void
 */
void qpsolver(double x[3], double *fval)
{
  static const char b_cv[128] = {
      '\x00', '\x01', '\x02', '\x03', '\x04', '\x05', '\x06', '\x07', '\x08',
      '\x09', '\x0a', '\x0b', '\x0c', '\x0d', '\x0e', '\x0f', '\x10', '\x11',
      '\x12', '\x13', '\x14', '\x15', '\x16', '\x17', '\x18', '\x19', '\x1a',
      '\x1b', '\x1c', '\x1d', '\x1e', '\x1f', ' ',    '!',    '\"',   '#',
      '$',    '%',    '&',    '\'',   '(',    ')',    '*',    '+',    ',',
      '-',    '.',    '/',    '0',    '1',    '2',    '3',    '4',    '5',
      '6',    '7',    '8',    '9',    ':',    ';',    '<',    '=',    '>',
      '?',    '@',    'a',    'b',    'c',    'd',    'e',    'f',    'g',
      'h',    'i',    'j',    'k',    'l',    'm',    'n',    'o',    'p',
      'q',    'r',    's',    't',    'u',    'v',    'w',    'x',    'y',
      'z',    '[',    '\\',   ']',    '^',    '_',    '`',    'a',    'b',
      'c',    'd',    'e',    'f',    'g',    'h',    'i',    'j',    'k',
      'l',    'm',    'n',    'o',    'p',    'q',    'r',    's',    't',
      'u',    'v',    'w',    'x',    'y',    'z',    '{',    '|',    '}',
      '~',    '\x7f'};
  static const char cv1[8] = {'q', 'u', 'a', 'd', 'p', 'r', 'o', 'g'};
  static const char t6_SolverName[8] = {'q', 'u', 'a', 'd', 'p', 'r', 'o', 'g'};
  static const char t6_FiniteDifferenceType[7] = {'f', 'o', 'r', 'w',
                                                  'a', 'r', 'd'};
  static const char t6_Display[5] = {'f', 'i', 'n', 'a', 'l'};
  b_struct_T objective;
  c_struct_T CholRegManager;
  d_struct_T memspace;
  e_struct_T qrmanager;
  f_struct_T workingset;
  g_struct_T options;
  struct_T solution;
  int idxEndIneq;
  int idx_global;
  int mConstr;
  solution.fstar = 0.0;
  solution.firstorderopt = 0.0;
  memset(&solution.lambda[0], 0, 8U * sizeof(double));
  solution.state = 0;
  solution.maxConstr = 0.0;
  solution.searchDir[0] = 0.0;
  solution.searchDir[1] = 0.0;
  solution.searchDir[2] = 0.0;
  solution.searchDir[3] = 0.0;
  solution.xstar[0] = 0.0;
  solution.xstar[1] = 0.0;
  solution.xstar[2] = 0.0;
  CholRegManager.ldm = 4;
  CholRegManager.ndims = 0;
  CholRegManager.info = 0;
  CholRegManager.ConvexCheck = true;
  CholRegManager.regTol_ = 0.0;
  CholRegManager.scaleFactor = 100.0;
  objective.grad[0] = 0.0;
  objective.grad[1] = 0.0;
  objective.grad[2] = 0.0;
  objective.grad[3] = 0.0;
  objective.Hx[0] = 0.0;
  objective.Hx[1] = 0.0;
  objective.Hx[2] = 0.0;
  objective.hasLinear = true;
  objective.nvar = 3;
  objective.maxVar = 4;
  objective.beta = 0.0;
  objective.rho = 0.0;
  objective.objtype = 3;
  objective.prev_objtype = 3;
  objective.prev_nvar = 0;
  objective.prev_hasLinear = false;
  objective.gammaScalar = 0.0;
  solution.iterations = 0;
  PresolveWorkingSet(&solution, &memspace, &workingset, &qrmanager);
  options.InitDamping = 0.01;
  for (idxEndIneq = 0; idxEndIneq < 7; idxEndIneq++) {
    options.FiniteDifferenceType[idxEndIneq] =
        t6_FiniteDifferenceType[idxEndIneq];
  }
  options.SpecifyObjectiveGradient = false;
  options.ScaleProblem = false;
  options.SpecifyConstraintGradient = false;
  options.NonFiniteSupport = true;
  options.IterDisplaySQP = false;
  options.FiniteDifferenceStepSize = -1.0;
  options.MaxFunctionEvaluations = -1.0;
  options.IterDisplayQP = false;
  options.PricingTolerance = 0.0;
  for (idxEndIneq = 0; idxEndIneq < 10; idxEndIneq++) {
    options.Algorithm[idxEndIneq] = cv[idxEndIneq];
  }
  options.ObjectiveLimit = -1.0E+20;
  options.ConstraintTolerance = 1.0E-8;
  options.OptimalityTolerance = 1.0E-8;
  options.StepTolerance = 1.0E-8;
  options.MaxIterations = -1.0;
  options.FunctionTolerance = rtInf;
  for (idxEndIneq = 0; idxEndIneq < 8; idxEndIneq++) {
    options.SolverName[idxEndIneq] = t6_SolverName[idxEndIneq];
  }
  options.CheckGradients = false;
  options.Diagnostics[0] = 'o';
  options.Diagnostics[1] = 'f';
  options.Diagnostics[2] = 'f';
  options.DiffMaxChange = rtInf;
  options.DiffMinChange = 0.0;
  for (idxEndIneq = 0; idxEndIneq < 5; idxEndIneq++) {
    options.Display[idxEndIneq] = t6_Display[idxEndIneq];
  }
  options.FunValCheck[0] = 'o';
  options.FunValCheck[1] = 'f';
  options.FunValCheck[2] = 'f';
  options.UseParallel = false;
  options.LinearSolver[0] = 'a';
  options.LinearSolver[1] = 'u';
  options.LinearSolver[2] = 't';
  options.LinearSolver[3] = 'o';
  options.SubproblemAlgorithm[0] = 'c';
  options.SubproblemAlgorithm[1] = 'g';
  if (solution.state >= 0) {
    double oldObjLim;
    boolean_T guard1 = false;
    solution.iterations = 0;
    solution.maxConstr = maxConstraintViolation(&workingset, solution.xstar);
    guard1 = false;
    if (solution.maxConstr > 3.0000000000000004E-8) {
      phaseone(&solution, &memspace, &workingset, &qrmanager, &CholRegManager,
               &objective, &options);
      if (solution.state != 0) {
        solution.maxConstr =
            maxConstraintViolation(&workingset, solution.xstar);
        if (solution.maxConstr > options.ConstraintTolerance * 3.0) {
          solution.fstar = computeFval(&objective, memspace.workspace_double,
                                       solution.xstar);
          solution.state = -2;
        } else {
          if (solution.maxConstr > 0.0) {
            solution.searchDir[0] = solution.xstar[0];
            solution.searchDir[1] = solution.xstar[1];
            solution.searchDir[2] = solution.xstar[2];
            b_PresolveWorkingSet(&solution, &memspace, &workingset, &qrmanager,
                                 &options);
            oldObjLim = maxConstraintViolation(&workingset, solution.xstar);
            if (oldObjLim >= solution.maxConstr) {
              solution.maxConstr = oldObjLim;
              solution.xstar[0] = solution.searchDir[0];
              solution.xstar[1] = solution.searchDir[1];
              solution.xstar[2] = solution.searchDir[2];
            }
          }
          guard1 = true;
        }
      }
    } else {
      guard1 = true;
    }
    if (guard1) {
      boolean_T b_bool;
      iterate(&solution, &memspace, &workingset, &qrmanager, &CholRegManager,
              &objective, options.IterDisplayQP, options.PricingTolerance,
              options.ObjectiveLimit, options.ConstraintTolerance,
              options.StepTolerance, true);
      b_bool = false;
      mConstr = 0;
      int exitg1;
      do {
        exitg1 = 0;
        if (mConstr < 8) {
          if (b_cv[(unsigned char)options.SolverName[mConstr] & 127] !=
              b_cv[(int)cv1[mConstr]]) {
            exitg1 = 1;
          } else {
            mConstr++;
          }
        } else {
          b_bool = true;
          exitg1 = 1;
        }
      } while (exitg1 == 0);
      if (b_bool && (solution.state != -6)) {
        solution.maxConstr =
            maxConstraintViolation(&workingset, solution.xstar);
        computeFirstOrderOpt(&solution, &objective, workingset.nVar,
                             workingset.ldA, workingset.ATwset,
                             workingset.nActiveConstr,
                             memspace.workspace_double);
        while ((solution.iterations < 100) &&
               ((solution.state == -7) ||
                ((solution.state == 1) &&
                 ((solution.maxConstr > options.ConstraintTolerance * 3.0) ||
                  (solution.firstorderopt >
                   options.OptimalityTolerance * 7.0))))) {
          int PROBTYPE_ORIG;
          int TYPE;
          int idxStartIneq;
          int nVar_tmp_tmp;
          feasibleX0ForWorkingSet(memspace.workspace_double, solution.xstar,
                                  &workingset, &qrmanager);
          b_PresolveWorkingSet(&solution, &memspace, &workingset, &qrmanager,
                               &options);
          PROBTYPE_ORIG = workingset.probType;
          nVar_tmp_tmp = workingset.nVar;
          solution.xstar[workingset.nVar] = solution.maxConstr + 1.0;
          if (workingset.probType == 3) {
            idxStartIneq = 1;
          } else {
            idxStartIneq = 4;
          }
          setProblemType(&workingset, idxStartIneq);
          mConstr = workingset.nWConstr[0] + workingset.nWConstr[1];
          idxStartIneq = mConstr + 1;
          idxEndIneq = workingset.nActiveConstr;
          for (idx_global = idxStartIneq; idx_global <= idxEndIneq;
               idx_global++) {
            workingset.isActiveConstr
                [(workingset.isActiveIdx[workingset.Wid[idx_global - 1] - 1] +
                  workingset.Wlocalidx[idx_global - 1]) -
                 2] = false;
          }
          workingset.nWConstr[2] = 0;
          workingset.nWConstr[3] = 0;
          workingset.nWConstr[4] = 0;
          workingset.nActiveConstr = mConstr;
          objective.prev_objtype = objective.objtype;
          objective.prev_nvar = objective.nvar;
          objective.prev_hasLinear = objective.hasLinear;
          objective.objtype = 5;
          objective.nvar = nVar_tmp_tmp + 1;
          objective.gammaScalar = 1.0;
          objective.hasLinear = true;
          oldObjLim = options.ObjectiveLimit;
          options.ObjectiveLimit = options.ConstraintTolerance * 3.0;
          solution.fstar = computeFval(&objective, memspace.workspace_double,
                                       solution.xstar);
          solution.state = 5;
          iterate(&solution, &memspace, &workingset, &qrmanager,
                  &CholRegManager, &objective, options.IterDisplayQP,
                  options.PricingTolerance, options.ObjectiveLimit,
                  options.ConstraintTolerance, 1.4901161193847657E-10, false);
          if (workingset.isActiveConstr
                  [(workingset.isActiveIdx[3] + workingset.sizes[3]) - 2]) {
            boolean_T exitg2;
            idx_global = workingset.sizes[0] + workingset.sizes[1];
            exitg2 = false;
            while ((!exitg2) && (idx_global + 1 <= workingset.nActiveConstr)) {
              if ((workingset.Wid[idx_global] == 4) &&
                  (workingset.Wlocalidx[idx_global] == workingset.sizes[3])) {
                TYPE = workingset.Wid[idx_global] - 1;
                workingset.isActiveConstr
                    [(workingset.isActiveIdx[workingset.Wid[idx_global] - 1] +
                      workingset.Wlocalidx[idx_global]) -
                     2] = false;
                workingset.Wid[idx_global] =
                    workingset.Wid[workingset.nActiveConstr - 1];
                workingset.Wlocalidx[idx_global] =
                    workingset.Wlocalidx[workingset.nActiveConstr - 1];
                idxEndIneq = workingset.nVar;
                for (mConstr = 0; mConstr < idxEndIneq; mConstr++) {
                  workingset.ATwset[mConstr + workingset.ldA * idx_global] =
                      workingset
                          .ATwset[mConstr + workingset.ldA *
                                                (workingset.nActiveConstr - 1)];
                }
                workingset.bwset[idx_global] =
                    workingset.bwset[workingset.nActiveConstr - 1];
                workingset.nActiveConstr--;
                workingset.nWConstr[TYPE]--;
                exitg2 = true;
              } else {
                idx_global++;
              }
            }
          }
          mConstr = workingset.nActiveConstr - 1;
          idxStartIneq = workingset.sizes[0] + workingset.sizes[1];
          while ((mConstr + 1 > idxStartIneq) && (mConstr + 1 > nVar_tmp_tmp)) {
            TYPE = workingset.Wid[mConstr] - 1;
            workingset.isActiveConstr
                [(workingset.isActiveIdx[workingset.Wid[mConstr] - 1] +
                  workingset.Wlocalidx[mConstr]) -
                 2] = false;
            workingset.Wid[mConstr] =
                workingset.Wid[workingset.nActiveConstr - 1];
            workingset.Wlocalidx[mConstr] =
                workingset.Wlocalidx[workingset.nActiveConstr - 1];
            idxEndIneq = workingset.nVar;
            for (idx_global = 0; idx_global < idxEndIneq; idx_global++) {
              workingset.ATwset[idx_global + workingset.ldA * mConstr] =
                  workingset
                      .ATwset[idx_global +
                              workingset.ldA * (workingset.nActiveConstr - 1)];
            }
            workingset.bwset[mConstr] =
                workingset.bwset[workingset.nActiveConstr - 1];
            workingset.nActiveConstr--;
            workingset.nWConstr[TYPE]--;
            mConstr--;
          }
          solution.maxConstr = solution.xstar[nVar_tmp_tmp];
          setProblemType(&workingset, PROBTYPE_ORIG);
          objective.objtype = objective.prev_objtype;
          objective.nvar = objective.prev_nvar;
          objective.hasLinear = objective.prev_hasLinear;
          options.ObjectiveLimit = oldObjLim;
          iterate(&solution, &memspace, &workingset, &qrmanager,
                  &CholRegManager, &objective, options.IterDisplayQP,
                  options.PricingTolerance, oldObjLim,
                  options.ConstraintTolerance, options.StepTolerance, false);
          solution.maxConstr =
              maxConstraintViolation(&workingset, solution.xstar);
          computeFirstOrderOpt(&solution, &objective, workingset.nVar,
                               workingset.ldA, workingset.ATwset,
                               workingset.nActiveConstr,
                               memspace.workspace_double);
        }
      }
    }
  }
  x[0] = solution.xstar[0];
  x[1] = solution.xstar[1];
  x[2] = solution.xstar[2];
  if (solution.state > 0) {
    *fval = solution.fstar;
  } else {
    *fval = computeFval(&objective, memspace.workspace_double, solution.xstar);
  }
}

/*
 * File trailer for qpsolver.c
 *
 * [EOF]
 */
