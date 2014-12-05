#include "EXTERN.h"
#include "perl.h"
#include "XSUB.h"

XS(PurlXS);
static char* EvalPerl();

#include "purl.c"
