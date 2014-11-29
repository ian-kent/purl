#include "EXTERN.h"
#include "perl.h"
#include "XSUB.h"

#include "purl.c"

XS(PurlXS);
static void RunPurl();
