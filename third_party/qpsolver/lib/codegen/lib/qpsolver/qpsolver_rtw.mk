###########################################################################
## Makefile generated for component 'qpsolver'. 
## 
## Makefile     : qpsolver_rtw.mk
## Generated on : Sat May 07 19:39:14 2022
## Final product: ./qpsolver.a
## Product type : static-library
## 
###########################################################################

###########################################################################
## MACROS
###########################################################################

# Macro Descriptions:
# PRODUCT_NAME            Name of the system to build
# MAKEFILE                Name of this makefile
# MODELLIB                Static library target

PRODUCT_NAME              = qpsolver
MAKEFILE                  = qpsolver_rtw.mk
MATLAB_ROOT               = /MATLAB
MATLAB_BIN                = /MATLAB/bin
MATLAB_ARCH_BIN           = $(MATLAB_BIN)/glnxa64
START_DIR                 = /MATLAB\ Drive
TGT_FCN_LIB               = ISO_C
SOLVER_OBJ                = 
CLASSIC_INTERFACE         = 0
MODEL_HAS_DYNAMICALLY_LOADED_SFCNS = 
RELATIVE_PATH_TO_ANCHOR   = ../../..
C_STANDARD_OPTS           = -fwrapv
CPP_STANDARD_OPTS         = -fwrapv
MODELLIB                  = qpsolver.a

###########################################################################
## TOOLCHAIN SPECIFICATIONS
###########################################################################

# Toolchain Name:          GNU gcc/g++ | gmake (64-bit Linux)
# Supported Version(s):    
# ToolchainInfo Version:   2022a
# Specification Revision:  1.0
# 
#-------------------------------------------
# Macros assumed to be defined elsewhere
#-------------------------------------------

# C_STANDARD_OPTS
# CPP_STANDARD_OPTS

#-----------
# MACROS
#-----------

WARN_FLAGS         = -Wall -W -Wwrite-strings -Winline -Wstrict-prototypes -Wnested-externs -Wpointer-arith -Wcast-align
WARN_FLAGS_MAX     = $(WARN_FLAGS) -Wcast-qual -Wshadow
CPP_WARN_FLAGS     = -Wall -W -Wwrite-strings -Winline -Wpointer-arith -Wcast-align
CPP_WARN_FLAGS_MAX = $(CPP_WARN_FLAGS) -Wcast-qual -Wshadow

TOOLCHAIN_SRCS = 
TOOLCHAIN_INCS = 
TOOLCHAIN_LIBS = 

#------------------------
# BUILD TOOL COMMANDS
#------------------------

# C Compiler: GNU C Compiler
CC = gcc

# Linker: GNU Linker
LD = g++

# C++ Compiler: GNU C++ Compiler
CPP = g++

# C++ Linker: GNU C++ Linker
CPP_LD = g++

# Archiver: GNU Archiver
AR = ar

# MEX Tool: MEX Tool
MEX_PATH = $(MATLAB_ARCH_BIN)
MEX = "$(MEX_PATH)/mex"

# Download: Download
DOWNLOAD =

# Execute: Execute
EXECUTE = $(PRODUCT)

# Builder: GMAKE Utility
MAKE_PATH = %MATLAB%/bin/glnxa64
MAKE = "$(MAKE_PATH)/gmake"


#-------------------------
# Directives/Utilities
#-------------------------

CDEBUG              = -g
C_OUTPUT_FLAG       = -o
LDDEBUG             = -g
OUTPUT_FLAG         = -o
CPPDEBUG            = -g
CPP_OUTPUT_FLAG     = -o
CPPLDDEBUG          = -g
OUTPUT_FLAG         = -o
ARDEBUG             =
STATICLIB_OUTPUT_FLAG =
MEX_DEBUG           = -g
RM                  = @rm -f
ECHO                = @echo
MV                  = @mv
RUN                 =

#--------------------------------------
# "Faster Runs" Build Configuration
#--------------------------------------

ARFLAGS              = ruvs
CFLAGS               = -c $(C_STANDARD_OPTS) -fPIC \
                       -O3 -fno-loop-optimize -fno-aggressive-loop-optimizations
CPPFLAGS             = -c $(CPP_STANDARD_OPTS) -fPIC \
                       -O3 -fno-loop-optimize -fno-aggressive-loop-optimizations
CPP_LDFLAGS          =
CPP_SHAREDLIB_LDFLAGS  = -shared -Wl,--no-undefined
DOWNLOAD_FLAGS       =
EXECUTE_FLAGS        =
LDFLAGS              =
MEX_CPPFLAGS         =
MEX_CPPLDFLAGS       =
MEX_CFLAGS           =
MEX_LDFLAGS          =
MAKE_FLAGS           = -f $(MAKEFILE)
SHAREDLIB_LDFLAGS    = -shared -Wl,--no-undefined



###########################################################################
## OUTPUT INFO
###########################################################################

PRODUCT = ./qpsolver.a
PRODUCT_TYPE = "static-library"
BUILD_TYPE = "Static Library"

###########################################################################
## INCLUDE PATHS
###########################################################################

INCLUDES_BUILDINFO = -I$(START_DIR)/codegen/lib/qpsolver -I$(START_DIR) -I$(MATLAB_ROOT)/extern/include

INCLUDES = $(INCLUDES_BUILDINFO)

###########################################################################
## DEFINES
###########################################################################

DEFINES_CUSTOM = 
DEFINES_STANDARD = -DMODEL=qpsolver

DEFINES = $(DEFINES_CUSTOM) $(DEFINES_STANDARD)

###########################################################################
## SOURCE FILES
###########################################################################

SRCS = $(START_DIR)/codegen/lib/qpsolver/qpsolver_data.c $(START_DIR)/codegen/lib/qpsolver/rt_nonfinite.c $(START_DIR)/codegen/lib/qpsolver/rtGetNaN.c $(START_DIR)/codegen/lib/qpsolver/rtGetInf.c $(START_DIR)/codegen/lib/qpsolver/qpsolver_initialize.c $(START_DIR)/codegen/lib/qpsolver/qpsolver_terminate.c $(START_DIR)/codegen/lib/qpsolver/qpsolver.c $(START_DIR)/codegen/lib/qpsolver/xnrm2.c $(START_DIR)/codegen/lib/qpsolver/ixamax.c $(START_DIR)/codegen/lib/qpsolver/printInfo.c $(START_DIR)/codegen/lib/qpsolver/xzlarfg.c $(START_DIR)/codegen/lib/qpsolver/xzlarf.c $(START_DIR)/codegen/lib/qpsolver/PresolveWorkingSet.c $(START_DIR)/codegen/lib/qpsolver/computeQ_.c $(START_DIR)/codegen/lib/qpsolver/countsort.c $(START_DIR)/codegen/lib/qpsolver/removeEqConstr.c $(START_DIR)/codegen/lib/qpsolver/RemoveDependentIneq_.c $(START_DIR)/codegen/lib/qpsolver/feasibleX0ForWorkingSet.c $(START_DIR)/codegen/lib/qpsolver/factorQR.c $(START_DIR)/codegen/lib/qpsolver/xgemv.c $(START_DIR)/codegen/lib/qpsolver/maxConstraintViolation.c $(START_DIR)/codegen/lib/qpsolver/setProblemType.c $(START_DIR)/codegen/lib/qpsolver/modifyOverheadPhaseOne_.c $(START_DIR)/codegen/lib/qpsolver/phaseone.c $(START_DIR)/codegen/lib/qpsolver/computeFval.c $(START_DIR)/codegen/lib/qpsolver/linearForm_.c $(START_DIR)/codegen/lib/qpsolver/computeGrad_StoreHx.c $(START_DIR)/codegen/lib/qpsolver/computeFval_ReuseHx.c $(START_DIR)/codegen/lib/qpsolver/squareQ_appendCol.c $(START_DIR)/codegen/lib/qpsolver/xrotg.c $(START_DIR)/codegen/lib/qpsolver/deleteColMoveEnd.c $(START_DIR)/codegen/lib/qpsolver/compute_deltax.c $(START_DIR)/codegen/lib/qpsolver/compute_lambda.c $(START_DIR)/codegen/lib/qpsolver/feasibleratiotest.c $(START_DIR)/codegen/lib/qpsolver/addBoundToActiveSetMatrix_.c $(START_DIR)/codegen/lib/qpsolver/computeFirstOrderOpt.c $(START_DIR)/codegen/lib/qpsolver/fullColLDL2_.c $(START_DIR)/codegen/lib/qpsolver/iterate.c $(START_DIR)/codegen/lib/qpsolver/ratiotest.c $(START_DIR)/codegen/lib/qpsolver/xzgeqp3.c

ALL_SRCS = $(SRCS)

###########################################################################
## OBJECTS
###########################################################################

OBJS = qpsolver_data.o rt_nonfinite.o rtGetNaN.o rtGetInf.o qpsolver_initialize.o qpsolver_terminate.o qpsolver.o xnrm2.o ixamax.o printInfo.o xzlarfg.o xzlarf.o PresolveWorkingSet.o computeQ_.o countsort.o removeEqConstr.o RemoveDependentIneq_.o feasibleX0ForWorkingSet.o factorQR.o xgemv.o maxConstraintViolation.o setProblemType.o modifyOverheadPhaseOne_.o phaseone.o computeFval.o linearForm_.o computeGrad_StoreHx.o computeFval_ReuseHx.o squareQ_appendCol.o xrotg.o deleteColMoveEnd.o compute_deltax.o compute_lambda.o feasibleratiotest.o addBoundToActiveSetMatrix_.o computeFirstOrderOpt.o fullColLDL2_.o iterate.o ratiotest.o xzgeqp3.o

ALL_OBJS = $(OBJS)

###########################################################################
## PREBUILT OBJECT FILES
###########################################################################

PREBUILT_OBJS = 

###########################################################################
## LIBRARIES
###########################################################################

LIBS = 

###########################################################################
## SYSTEM LIBRARIES
###########################################################################

SYSTEM_LIBS =  -lm

###########################################################################
## ADDITIONAL TOOLCHAIN FLAGS
###########################################################################

#---------------
# C Compiler
#---------------

CFLAGS_BASIC = $(DEFINES) $(INCLUDES)

CFLAGS += $(CFLAGS_BASIC)

#-----------------
# C++ Compiler
#-----------------

CPPFLAGS_BASIC = $(DEFINES) $(INCLUDES)

CPPFLAGS += $(CPPFLAGS_BASIC)

###########################################################################
## INLINED COMMANDS
###########################################################################

###########################################################################
## PHONY TARGETS
###########################################################################

.PHONY : all build clean info prebuild download execute


all : build
	@echo "### Successfully generated all binary outputs."


build : prebuild $(PRODUCT)


prebuild : 


download : $(PRODUCT)


execute : download


###########################################################################
## FINAL TARGET
###########################################################################

#---------------------------------
# Create a static library         
#---------------------------------

$(PRODUCT) : $(OBJS) $(PREBUILT_OBJS)
	@echo "### Creating static library "$(PRODUCT)" ..."
	$(AR) $(ARFLAGS)  $(PRODUCT) $(OBJS)
	@echo "### Created: $(PRODUCT)"


###########################################################################
## INTERMEDIATE TARGETS
###########################################################################

#---------------------
# SOURCE-TO-OBJECT
#---------------------

%.o : %.c
	$(CC) $(CFLAGS) -o "$@" "$<"


%.o : %.cpp
	$(CPP) $(CPPFLAGS) -o "$@" "$<"


%.o : $(RELATIVE_PATH_TO_ANCHOR)/%.c
	$(CC) $(CFLAGS) -o "$@" "$<"


%.o : $(RELATIVE_PATH_TO_ANCHOR)/%.cpp
	$(CPP) $(CPPFLAGS) -o "$@" "$<"


%.o : $(START_DIR)/codegen/lib/qpsolver/%.c
	$(CC) $(CFLAGS) -o "$@" "$<"


%.o : $(START_DIR)/codegen/lib/qpsolver/%.cpp
	$(CPP) $(CPPFLAGS) -o "$@" "$<"


%.o : $(START_DIR)/%.c
	$(CC) $(CFLAGS) -o "$@" "$<"


%.o : $(START_DIR)/%.cpp
	$(CPP) $(CPPFLAGS) -o "$@" "$<"


qpsolver_data.o : $(START_DIR)/codegen/lib/qpsolver/qpsolver_data.c
	$(CC) $(CFLAGS) -o "$@" "$<"


rt_nonfinite.o : $(START_DIR)/codegen/lib/qpsolver/rt_nonfinite.c
	$(CC) $(CFLAGS) -o "$@" "$<"


rtGetNaN.o : $(START_DIR)/codegen/lib/qpsolver/rtGetNaN.c
	$(CC) $(CFLAGS) -o "$@" "$<"


rtGetInf.o : $(START_DIR)/codegen/lib/qpsolver/rtGetInf.c
	$(CC) $(CFLAGS) -o "$@" "$<"


qpsolver_initialize.o : $(START_DIR)/codegen/lib/qpsolver/qpsolver_initialize.c
	$(CC) $(CFLAGS) -o "$@" "$<"


qpsolver_terminate.o : $(START_DIR)/codegen/lib/qpsolver/qpsolver_terminate.c
	$(CC) $(CFLAGS) -o "$@" "$<"


qpsolver.o : $(START_DIR)/codegen/lib/qpsolver/qpsolver.c
	$(CC) $(CFLAGS) -o "$@" "$<"


xnrm2.o : $(START_DIR)/codegen/lib/qpsolver/xnrm2.c
	$(CC) $(CFLAGS) -o "$@" "$<"


ixamax.o : $(START_DIR)/codegen/lib/qpsolver/ixamax.c
	$(CC) $(CFLAGS) -o "$@" "$<"


printInfo.o : $(START_DIR)/codegen/lib/qpsolver/printInfo.c
	$(CC) $(CFLAGS) -o "$@" "$<"


xzlarfg.o : $(START_DIR)/codegen/lib/qpsolver/xzlarfg.c
	$(CC) $(CFLAGS) -o "$@" "$<"


xzlarf.o : $(START_DIR)/codegen/lib/qpsolver/xzlarf.c
	$(CC) $(CFLAGS) -o "$@" "$<"


PresolveWorkingSet.o : $(START_DIR)/codegen/lib/qpsolver/PresolveWorkingSet.c
	$(CC) $(CFLAGS) -o "$@" "$<"


computeQ_.o : $(START_DIR)/codegen/lib/qpsolver/computeQ_.c
	$(CC) $(CFLAGS) -o "$@" "$<"


countsort.o : $(START_DIR)/codegen/lib/qpsolver/countsort.c
	$(CC) $(CFLAGS) -o "$@" "$<"


removeEqConstr.o : $(START_DIR)/codegen/lib/qpsolver/removeEqConstr.c
	$(CC) $(CFLAGS) -o "$@" "$<"


RemoveDependentIneq_.o : $(START_DIR)/codegen/lib/qpsolver/RemoveDependentIneq_.c
	$(CC) $(CFLAGS) -o "$@" "$<"


feasibleX0ForWorkingSet.o : $(START_DIR)/codegen/lib/qpsolver/feasibleX0ForWorkingSet.c
	$(CC) $(CFLAGS) -o "$@" "$<"


factorQR.o : $(START_DIR)/codegen/lib/qpsolver/factorQR.c
	$(CC) $(CFLAGS) -o "$@" "$<"


xgemv.o : $(START_DIR)/codegen/lib/qpsolver/xgemv.c
	$(CC) $(CFLAGS) -o "$@" "$<"


maxConstraintViolation.o : $(START_DIR)/codegen/lib/qpsolver/maxConstraintViolation.c
	$(CC) $(CFLAGS) -o "$@" "$<"


setProblemType.o : $(START_DIR)/codegen/lib/qpsolver/setProblemType.c
	$(CC) $(CFLAGS) -o "$@" "$<"


modifyOverheadPhaseOne_.o : $(START_DIR)/codegen/lib/qpsolver/modifyOverheadPhaseOne_.c
	$(CC) $(CFLAGS) -o "$@" "$<"


phaseone.o : $(START_DIR)/codegen/lib/qpsolver/phaseone.c
	$(CC) $(CFLAGS) -o "$@" "$<"


computeFval.o : $(START_DIR)/codegen/lib/qpsolver/computeFval.c
	$(CC) $(CFLAGS) -o "$@" "$<"


linearForm_.o : $(START_DIR)/codegen/lib/qpsolver/linearForm_.c
	$(CC) $(CFLAGS) -o "$@" "$<"


computeGrad_StoreHx.o : $(START_DIR)/codegen/lib/qpsolver/computeGrad_StoreHx.c
	$(CC) $(CFLAGS) -o "$@" "$<"


computeFval_ReuseHx.o : $(START_DIR)/codegen/lib/qpsolver/computeFval_ReuseHx.c
	$(CC) $(CFLAGS) -o "$@" "$<"


squareQ_appendCol.o : $(START_DIR)/codegen/lib/qpsolver/squareQ_appendCol.c
	$(CC) $(CFLAGS) -o "$@" "$<"


xrotg.o : $(START_DIR)/codegen/lib/qpsolver/xrotg.c
	$(CC) $(CFLAGS) -o "$@" "$<"


deleteColMoveEnd.o : $(START_DIR)/codegen/lib/qpsolver/deleteColMoveEnd.c
	$(CC) $(CFLAGS) -o "$@" "$<"


compute_deltax.o : $(START_DIR)/codegen/lib/qpsolver/compute_deltax.c
	$(CC) $(CFLAGS) -o "$@" "$<"


compute_lambda.o : $(START_DIR)/codegen/lib/qpsolver/compute_lambda.c
	$(CC) $(CFLAGS) -o "$@" "$<"


feasibleratiotest.o : $(START_DIR)/codegen/lib/qpsolver/feasibleratiotest.c
	$(CC) $(CFLAGS) -o "$@" "$<"


addBoundToActiveSetMatrix_.o : $(START_DIR)/codegen/lib/qpsolver/addBoundToActiveSetMatrix_.c
	$(CC) $(CFLAGS) -o "$@" "$<"


computeFirstOrderOpt.o : $(START_DIR)/codegen/lib/qpsolver/computeFirstOrderOpt.c
	$(CC) $(CFLAGS) -o "$@" "$<"


fullColLDL2_.o : $(START_DIR)/codegen/lib/qpsolver/fullColLDL2_.c
	$(CC) $(CFLAGS) -o "$@" "$<"


iterate.o : $(START_DIR)/codegen/lib/qpsolver/iterate.c
	$(CC) $(CFLAGS) -o "$@" "$<"


ratiotest.o : $(START_DIR)/codegen/lib/qpsolver/ratiotest.c
	$(CC) $(CFLAGS) -o "$@" "$<"


xzgeqp3.o : $(START_DIR)/codegen/lib/qpsolver/xzgeqp3.c
	$(CC) $(CFLAGS) -o "$@" "$<"


###########################################################################
## DEPENDENCIES
###########################################################################

$(ALL_OBJS) : rtw_proj.tmw $(MAKEFILE)


###########################################################################
## MISCELLANEOUS TARGETS
###########################################################################

info : 
	@echo "### PRODUCT = $(PRODUCT)"
	@echo "### PRODUCT_TYPE = $(PRODUCT_TYPE)"
	@echo "### BUILD_TYPE = $(BUILD_TYPE)"
	@echo "### INCLUDES = $(INCLUDES)"
	@echo "### DEFINES = $(DEFINES)"
	@echo "### ALL_SRCS = $(ALL_SRCS)"
	@echo "### ALL_OBJS = $(ALL_OBJS)"
	@echo "### LIBS = $(LIBS)"
	@echo "### MODELREF_LIBS = $(MODELREF_LIBS)"
	@echo "### SYSTEM_LIBS = $(SYSTEM_LIBS)"
	@echo "### TOOLCHAIN_LIBS = $(TOOLCHAIN_LIBS)"
	@echo "### CFLAGS = $(CFLAGS)"
	@echo "### LDFLAGS = $(LDFLAGS)"
	@echo "### SHAREDLIB_LDFLAGS = $(SHAREDLIB_LDFLAGS)"
	@echo "### CPPFLAGS = $(CPPFLAGS)"
	@echo "### CPP_LDFLAGS = $(CPP_LDFLAGS)"
	@echo "### CPP_SHAREDLIB_LDFLAGS = $(CPP_SHAREDLIB_LDFLAGS)"
	@echo "### ARFLAGS = $(ARFLAGS)"
	@echo "### MEX_CFLAGS = $(MEX_CFLAGS)"
	@echo "### MEX_CPPFLAGS = $(MEX_CPPFLAGS)"
	@echo "### MEX_LDFLAGS = $(MEX_LDFLAGS)"
	@echo "### MEX_CPPLDFLAGS = $(MEX_CPPLDFLAGS)"
	@echo "### DOWNLOAD_FLAGS = $(DOWNLOAD_FLAGS)"
	@echo "### EXECUTE_FLAGS = $(EXECUTE_FLAGS)"
	@echo "### MAKE_FLAGS = $(MAKE_FLAGS)"


clean : 
	$(ECHO) "### Deleting all derived files..."
	$(RM) $(PRODUCT)
	$(RM) $(ALL_OBJS)
	$(ECHO) "### Deleted all derived files."


