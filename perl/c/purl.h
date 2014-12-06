#include "EXTERN.h"
#include "perl.h"
#include "XSUB.h"
#include "embed.h"

XS(XSHook);
XS(PurlXSTest);
static char* EvalPerl();

#include "purl.c"
