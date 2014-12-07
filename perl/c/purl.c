static void xs_init (pTHX);
static PerlInterpreter *my_perl;

static int dummy_argc = 3;
static char** dummy_argv;
static char** dummy_env;

typedef void (*purlXSCallback) ();

extern void PurlTest();
extern SV* PurlXSHook();

XS(XSTest)
{
  dXSARGS;
  PERL_UNUSED_VAR(cv);
  PurlTest();
  XSRETURN_EMPTY;
}

XS(XSInvoke)
{
  dXSARGS;

  if (items < 2) {
    croak("expected invocant and delegate");
  }

  int n_perl_args = items - 2;
  SV** perl_args = malloc(sizeof(SV*) * n_perl_args);

  int i;
  for (i = 0; i < n_perl_args; i++) {
    perl_args[i] = ST(i+2);
    SvREFCNT_inc(perl_args[i]);
  }

  SV* scalar_return = 0;
  printf("%s\n", SvPV_nolen(ST(1)));
  scalar_return = PurlXSHook(ST(0), SvPV_nolen(ST(1)), n_perl_args, perl_args);

  free(perl_args);

  if (!scalar_return) {
    ST(0) = &PL_sv_undef;
  } else {
    ST(0) = sv_2mortal(scalar_return);
  }

  XSRETURN(1);
}

EXTERN_C void
xs_init(pTHX)
{
  char *file = __FILE__;
  newXS("Purl::Test", XSTest, file);
  newXS("Purl::XS::Invoke", XSInvoke, file);
}

static void
PurlInit()
{
  PL_origalen = 1;

  dummy_argv = malloc(sizeof(char*) * 3);
  dummy_argv[0] = "purl";
  dummy_argv[1] = "-e";
  dummy_argv[2] = "0";

  PERL_SYS_INIT3(&dummy_argc,&dummy_argv,&dummy_env);

  my_perl = perl_alloc();
  perl_construct(my_perl);
  PL_exit_flags |= PERL_EXIT_DESTRUCT_END;
  perl_parse(my_perl, xs_init, dummy_argc, dummy_argv, (char **)NULL);
  //perl_run(my_perl);
}

static void
PurlDestroy()
{
  perl_destruct(my_perl);
  perl_free(my_perl);

  PERL_SYS_TERM();
}

static char*
GetSVString(SV *sv)
{
  return SvPV_nolen(sv);
}

static char*
EvalPerl(char *src)
{
  char* pv;

  SV *val = eval_pv(src, TRUE);
  if(SvOK(val)) {
    pv = SvPV_nolen(val);
  }

  return pv;
}
