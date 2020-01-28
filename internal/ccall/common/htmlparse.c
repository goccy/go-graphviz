/* A Bison parser, made by GNU Bison 2.7.  */

/* Bison implementation for Yacc-like parsers in C
   
      Copyright (C) 1984, 1989-1990, 2000-2012 Free Software Foundation, Inc.
   
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.  */

/* As a special exception, you may create a larger work that contains
   part or all of the Bison parser skeleton and distribute that work
   under terms of your choice, so long as that work isn't itself a
   parser generator using the skeleton or a modified version thereof
   as a parser skeleton.  Alternatively, if you modify or redistribute
   the parser skeleton itself, you may (at your option) remove this
   special exception, which will cause the skeleton and the resulting
   Bison output files to be licensed under the GNU General Public
   License without this special exception.
   
   This special exception was added by the Free Software Foundation in
   version 2.2 of Bison.  */

/* C LALR(1) parser skeleton written by Richard Stallman, by
   simplifying the original so-called "semantic" parser.  */

/* All symbols defined below should begin with html or YY, to avoid
   infringing on user name space.  This should be done even for local
   variables, as they might otherwise be expanded by user macros.
   There are some unavoidable exceptions within include files to
   define necessary library symbols; they are noted "INFRINGES ON
   USER NAME SPACE" below.  */

/* Identify Bison output.  */
#define YYBISON 1

/* Bison version.  */
#define YYBISON_VERSION "2.7"

/* Skeleton name.  */
#define YYSKELETON_NAME "yacc.c"

/* Pure parsers.  */
#define YYPURE 0

/* Push parsers.  */
#define YYPUSH 0

/* Pull parsers.  */
#define YYPULL 1




/* Copy the first part of user declarations.  */
/* Line 371 of yacc.c  */
#line 14 "../../lib/common/htmlparse.y"


#include "render.h"
#include "htmltable.h"
#include "htmllex.h"

extern int htmlparse(void);

typedef struct sfont_t {
    textfont_t *cfont;	
    struct sfont_t *pfont;
} sfont_t;

static struct {
  htmllabel_t* lbl;       /* Generated label */
  htmltbl_t*   tblstack;  /* Stack of tables maintained during parsing */
  Dt_t*        fitemList; /* Dictionary for font text items */
  Dt_t*        fspanList; 
  agxbuf*      str;       /* Buffer for text */
  sfont_t*     fontstack;
  GVC_t*       gvc;
} HTMLstate;

/* free_ritem:
 * Free row. This closes and frees row's list, then
 * the pitem itself is freed.
 */
static void
free_ritem(Dt_t* d, pitem* p,Dtdisc_t* ds)
{
  dtclose (p->u.rp);
  free (p);
}

/* free_item:
 * Generic Dt free. Only frees container, assuming contents
 * have been copied elsewhere.
 */
static void
free_item(Dt_t* d, void* p,Dtdisc_t* ds)
{
  free (p);
}

/* cleanTbl:
 * Clean up table if error in parsing.
 */
static void
cleanTbl (htmltbl_t* tp)
{
  dtclose (tp->u.p.rows);
  free_html_data (&tp->data);
  free (tp);
}

/* cleanCell:
 * Clean up cell if error in parsing.
 */
static void
cleanCell (htmlcell_t* cp)
{
  if (cp->child.kind == HTML_TBL) cleanTbl (cp->child.u.tbl);
  else if (cp->child.kind == HTML_TEXT) free_html_text (cp->child.u.txt);
  free_html_data (&cp->data);
  free (cp);
}

/* free_citem:
 * Free cell item during parsing. This frees cell and pitem.
 */
static void
free_citem(Dt_t* d, pitem* p,Dtdisc_t* ds)
{
  cleanCell (p->u.cp);
  free (p);
}

static Dtdisc_t rowDisc = {
    offsetof(pitem,u),
    sizeof(void*),
    offsetof(pitem,link),
    NIL(Dtmake_f),
    (Dtfree_f)free_ritem,
    NIL(Dtcompar_f),
    NIL(Dthash_f),
    NIL(Dtmemory_f),
    NIL(Dtevent_f)
};
static Dtdisc_t cellDisc = {
    offsetof(pitem,u),
    sizeof(void*),
    offsetof(pitem,link),
    NIL(Dtmake_f),
    (Dtfree_f)free_item,
    NIL(Dtcompar_f),
    NIL(Dthash_f),
    NIL(Dtmemory_f),
    NIL(Dtevent_f)
};

typedef struct {
    Dtlink_t    link;
    textspan_t  ti;
} fitem;

typedef struct {
    Dtlink_t     link;
    htextspan_t  lp;
} fspan;

static void 
free_fitem(Dt_t* d, fitem* p, Dtdisc_t* ds)
{
    if (p->ti.str)
	free (p->ti.str);
    free (p);
}

static void 
free_fspan(Dt_t* d, fspan* p, Dtdisc_t* ds)
{
    textspan_t* ti;

    if (p->lp.nitems) {
	int i;
	ti = p->lp.items;
	for (i = 0; i < p->lp.nitems; i++) {
	    if (ti->str) free (ti->str);
	    ti++;
	}
	free (p->lp.items);
    }
    free (p);
}

static Dtdisc_t fstrDisc = {
    0,
    0,
    offsetof(fitem,link),
    NIL(Dtmake_f),
    (Dtfree_f)free_item,
    NIL(Dtcompar_f),
    NIL(Dthash_f),
    NIL(Dtmemory_f),
    NIL(Dtevent_f)
};


static Dtdisc_t fspanDisc = {
    0,
    0,
    offsetof(fspan,link),
    NIL(Dtmake_f),
    (Dtfree_f)free_item,
    NIL(Dtcompar_f),
    NIL(Dthash_f),
    NIL(Dtmemory_f),
    NIL(Dtevent_f)
};

/* appendFItemList:
 * Append a new fitem to the list.
 */
static void
appendFItemList (agxbuf *ag)
{
    fitem *fi = NEW(fitem);

    fi->ti.str = strdup(agxbuse(ag));
    fi->ti.font = HTMLstate.fontstack->cfont;
    dtinsert(HTMLstate.fitemList, fi);
}	

/* appendFLineList:
 */
static void 
appendFLineList (int v)
{
    int cnt;
    fspan *ln = NEW(fspan);
    fitem *fi;
    Dt_t *ilist = HTMLstate.fitemList;

    cnt = dtsize(ilist);
    ln->lp.just = v;
    if (cnt) {
        int i = 0;
	ln->lp.nitems = cnt;
	ln->lp.items = N_NEW(cnt, textspan_t);

	fi = (fitem*)dtflatten(ilist);
	for (; fi; fi = (fitem*)dtlink(fitemList,(Dtlink_t*)fi)) {
		/* NOTE: When fitemList is closed, it uses free_item, which only frees the container,
		 * not the contents, so this copy is safe.
		 */
	    ln->lp.items[i] = fi->ti;  
	    i++;
	}
    }
    else {
	ln->lp.items = NEW(textspan_t);
	ln->lp.nitems = 1;
	ln->lp.items[0].str = strdup("");
	ln->lp.items[0].font = HTMLstate.fontstack->cfont;
    }

    dtclear(ilist);

    dtinsert(HTMLstate.fspanList, ln);
}

static htmltxt_t*
mkText(void)
{
    int cnt;
    Dt_t * ispan = HTMLstate.fspanList;
    fspan *fl ;
    htmltxt_t *hft = NEW(htmltxt_t);
    
    if (dtsize (HTMLstate.fitemList)) 
	appendFLineList (UNSET_ALIGN);

    cnt = dtsize(ispan);
    hft->nspans = cnt;
    	
    if (cnt) {
	int i = 0;
	hft->spans = N_NEW(cnt,htextspan_t);	
    	for(fl=(fspan *)dtfirst(ispan); fl; fl=(fspan *)dtnext(ispan,fl)) {
    	    hft->spans[i] = fl->lp;
    	    i++;
    	}
    }
    
    dtclear(ispan);

    return hft;
}

static pitem* lastRow (void)
{
  htmltbl_t* tbl = HTMLstate.tblstack;
  pitem*     sp = dtlast (tbl->u.p.rows);
  return sp;
}

/* addRow:
 * Add new cell row to current table.
 */
static pitem* addRow (void)
{
  Dt_t*      dp = dtopen(&cellDisc, Dtqueue);
  htmltbl_t* tbl = HTMLstate.tblstack;
  pitem*     sp = NEW(pitem);
  sp->u.rp = dp;
  if (tbl->flags & HTML_HRULE)
    sp->ruled = 1;
  dtinsert (tbl->u.p.rows, sp);
  return sp;
}

/* setCell:
 * Set cell body and type and attach to row
 */
static void setCell (htmlcell_t* cp, void* obj, int kind)
{
  pitem*     sp = NEW(pitem);
  htmltbl_t* tbl = HTMLstate.tblstack;
  pitem*     rp = (pitem*)dtlast (tbl->u.p.rows);
  Dt_t*      row = rp->u.rp;
  sp->u.cp = cp;
  dtinsert (row, sp);
  cp->child.kind = kind;
  if (tbl->flags & HTML_VRULE)
    cp->ruled = HTML_VRULE;
  
  if(kind == HTML_TEXT)
  	cp->child.u.txt = (htmltxt_t*)obj;
  else if (kind == HTML_IMAGE)
    cp->child.u.img = (htmlimg_t*)obj;
  else
    cp->child.u.tbl = (htmltbl_t*)obj;
}

/* mkLabel:
 * Create label, given body and type.
 */
static htmllabel_t* mkLabel (void* obj, int kind)
{
  htmllabel_t* lp = NEW(htmllabel_t);

  lp->kind = kind;
  if (kind == HTML_TEXT)
    lp->u.txt = (htmltxt_t*)obj;
  else
    lp->u.tbl = (htmltbl_t*)obj;
  return lp;
}

/* freeFontstack:
 * Free all stack items but the last, which is
 * put on artificially during in parseHTML.
 */
static void
freeFontstack(void)
{
    sfont_t* s;
    sfont_t* next;

    for (s = HTMLstate.fontstack; (next = s->pfont); s = next) {
	free(s);
    }
}

/* cleanup:
 * Called on error. Frees resources allocated during parsing.
 * This includes a label, plus a walk down the stack of
 * tables. Note that we use the free_citem function to actually
 * free cells.
 */
static void cleanup (void)
{
  htmltbl_t* tp = HTMLstate.tblstack;
  htmltbl_t* next;

  if (HTMLstate.lbl) {
    free_html_label (HTMLstate.lbl,1);
    HTMLstate.lbl = NULL;
  }
  cellDisc.freef = (Dtfree_f)free_citem;
  while (tp) {
    next = tp->u.p.prev;
    cleanTbl (tp);
    tp = next;
  }
  cellDisc.freef = (Dtfree_f)free_item;

  fstrDisc.freef = (Dtfree_f)free_fitem;
  dtclear (HTMLstate.fitemList);
  fstrDisc.freef = (Dtfree_f)free_item;

  fspanDisc.freef = (Dtfree_f)free_fspan;
  dtclear (HTMLstate.fspanList);
  fspanDisc.freef = (Dtfree_f)free_item;

  freeFontstack();
}

/* nonSpace:
 * Return 1 if s contains a non-space character.
 */
static int nonSpace (char* s)
{
  char   c;

  while ((c = *s++)) {
    if (c != ' ') return 1;
  }
  return 0;
}

/* pushFont:
 * Fonts are allocated in the lexer.
 */
static void
pushFont (textfont_t *fp)
{
    sfont_t *ft = NEW(sfont_t);
    textfont_t* curfont = HTMLstate.fontstack->cfont;
    textfont_t  f = *fp;

    if (curfont) {
	if (!f.color && curfont->color)
	    f.color = curfont->color;
	if ((f.size < 0.0) && (curfont->size >= 0.0))
	    f.size = curfont->size;
	if (!f.name && curfont->name)
	    f.name = curfont->name;
	if (curfont->flags)
	    f.flags |= curfont->flags;
    }

    ft->cfont = dtinsert(HTMLstate.gvc->textfont_dt, &f);
    ft->pfont = HTMLstate.fontstack;
    HTMLstate.fontstack = ft;
}

/* popFont:
 */
static void 
popFont (void)
{
    sfont_t* curfont = HTMLstate.fontstack;
    sfont_t* prevfont = curfont->pfont;

    free (curfont);
    HTMLstate.fontstack = prevfont;
}


/* Line 371 of yacc.c  */
#line 469 "y.tab.c"

# ifndef YY_NULL
#  if defined __cplusplus && 201103L <= __cplusplus
#   define YY_NULL nullptr
#  else
#   define YY_NULL 0
#  endif
# endif

/* Enabling verbose error messages.  */
#ifdef YYERROR_VERBOSE
# undef YYERROR_VERBOSE
# define YYERROR_VERBOSE 1
#else
# define YYERROR_VERBOSE 0
#endif

/* In a future release of Bison, this section will be replaced
   by #include "y.tab.h".  */
#ifndef YY_YY_Y_TAB_H_INCLUDED
# define YY_YY_Y_TAB_H_INCLUDED
/* Enabling traces.  */
#ifndef YYDEBUG
# define YYDEBUG 0
#endif
#if YYDEBUG
extern int htmldebug;
#endif

/* Tokens.  */
#ifndef YYTOKENTYPE
# define YYTOKENTYPE
   /* Put the tokens into the symbol table, so that GDB and other debuggers
      know about them.  */
   enum htmltokentype {
     T_end_br = 258,
     T_end_img = 259,
     T_row = 260,
     T_end_row = 261,
     T_html = 262,
     T_end_html = 263,
     T_end_table = 264,
     T_end_cell = 265,
     T_end_font = 266,
     T_string = 267,
     T_error = 268,
     T_n_italic = 269,
     T_n_bold = 270,
     T_n_underline = 271,
     T_n_overline = 272,
     T_n_sup = 273,
     T_n_sub = 274,
     T_n_s = 275,
     T_HR = 276,
     T_hr = 277,
     T_end_hr = 278,
     T_VR = 279,
     T_vr = 280,
     T_end_vr = 281,
     T_BR = 282,
     T_br = 283,
     T_IMG = 284,
     T_img = 285,
     T_table = 286,
     T_cell = 287,
     T_font = 288,
     T_italic = 289,
     T_bold = 290,
     T_underline = 291,
     T_overline = 292,
     T_sup = 293,
     T_sub = 294,
     T_s = 295
   };
#endif
/* Tokens.  */
#define T_end_br 258
#define T_end_img 259
#define T_row 260
#define T_end_row 261
#define T_html 262
#define T_end_html 263
#define T_end_table 264
#define T_end_cell 265
#define T_end_font 266
#define T_string 267
#define T_error 268
#define T_n_italic 269
#define T_n_bold 270
#define T_n_underline 271
#define T_n_overline 272
#define T_n_sup 273
#define T_n_sub 274
#define T_n_s 275
#define T_HR 276
#define T_hr 277
#define T_end_hr 278
#define T_VR 279
#define T_vr 280
#define T_end_vr 281
#define T_BR 282
#define T_br 283
#define T_IMG 284
#define T_img 285
#define T_table 286
#define T_cell 287
#define T_font 288
#define T_italic 289
#define T_bold 290
#define T_underline 291
#define T_overline 292
#define T_sup 293
#define T_sub 294
#define T_s 295



#if ! defined YYSTYPE && ! defined YYSTYPE_IS_DECLARED
typedef union YYSTYPE
{
/* Line 387 of yacc.c  */
#line 415 "../../lib/common/htmlparse.y"

  int    i;
  htmltxt_t*  txt;
  htmlcell_t*  cell;
  htmltbl_t*   tbl;
  textfont_t*  font;
  htmlimg_t*   img;
  pitem*       p;


/* Line 387 of yacc.c  */
#line 603 "y.tab.c"
} YYSTYPE;
# define YYSTYPE_IS_TRIVIAL 1
# define htmlstype YYSTYPE /* obsolescent; will be withdrawn */
# define YYSTYPE_IS_DECLARED 1
#endif

extern YYSTYPE htmllval;

#ifdef YYPARSE_PARAM
#if defined __STDC__ || defined __cplusplus
int htmlparse (void *YYPARSE_PARAM);
#else
int htmlparse ();
#endif
#else /* ! YYPARSE_PARAM */
#if defined __STDC__ || defined __cplusplus
int htmlparse (void);
#else
int htmlparse ();
#endif
#endif /* ! YYPARSE_PARAM */

#endif /* !YY_YY_Y_TAB_H_INCLUDED  */

/* Copy the second part of user declarations.  */

/* Line 390 of yacc.c  */
#line 631 "y.tab.c"

#ifdef short
# undef short
#endif

#ifdef YYTYPE_UINT8
typedef YYTYPE_UINT8 htmltype_uint8;
#else
typedef unsigned char htmltype_uint8;
#endif

#ifdef YYTYPE_INT8
typedef YYTYPE_INT8 htmltype_int8;
#elif (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
typedef signed char htmltype_int8;
#else
typedef short int htmltype_int8;
#endif

#ifdef YYTYPE_UINT16
typedef YYTYPE_UINT16 htmltype_uint16;
#else
typedef unsigned short int htmltype_uint16;
#endif

#ifdef YYTYPE_INT16
typedef YYTYPE_INT16 htmltype_int16;
#else
typedef short int htmltype_int16;
#endif

#ifndef YYSIZE_T
# ifdef __SIZE_TYPE__
#  define YYSIZE_T __SIZE_TYPE__
# elif defined size_t
#  define YYSIZE_T size_t
# elif ! defined YYSIZE_T && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
#  include <stddef.h> /* INFRINGES ON USER NAME SPACE */
#  define YYSIZE_T size_t
# else
#  define YYSIZE_T unsigned int
# endif
#endif

#define YYSIZE_MAXIMUM ((YYSIZE_T) -1)

#ifndef YY_
# if defined YYENABLE_NLS && YYENABLE_NLS
#  if ENABLE_NLS
#   include <libintl.h> /* INFRINGES ON USER NAME SPACE */
#   define YY_(Msgid) dgettext ("bison-runtime", Msgid)
#  endif
# endif
# ifndef YY_
#  define YY_(Msgid) Msgid
# endif
#endif

/* Suppress unused-variable warnings by "using" E.  */
#if ! defined lint || defined __GNUC__
# define YYUSE(E) ((void) (E))
#else
# define YYUSE(E) /* empty */
#endif

/* Identity function, used to suppress warnings about constant conditions.  */
#ifndef lint
# define YYID(N) (N)
#else
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static int
YYID (int htmli)
#else
static int
YYID (htmli)
    int htmli;
#endif
{
  return htmli;
}
#endif

#if ! defined htmloverflow || YYERROR_VERBOSE

/* The parser invokes alloca or malloc; define the necessary symbols.  */

# ifdef YYSTACK_USE_ALLOCA
#  if YYSTACK_USE_ALLOCA
#   ifdef __GNUC__
#    define YYSTACK_ALLOC __builtin_alloca
#   elif defined __BUILTIN_VA_ARG_INCR
#    include <alloca.h> /* INFRINGES ON USER NAME SPACE */
#   elif defined _AIX
#    define YYSTACK_ALLOC __alloca
#   elif defined _MSC_VER
#    include <malloc.h> /* INFRINGES ON USER NAME SPACE */
#    define alloca _alloca
#   else
#    define YYSTACK_ALLOC alloca
#    if ! defined _ALLOCA_H && ! defined EXIT_SUCCESS && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
#     include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
      /* Use EXIT_SUCCESS as a witness for stdlib.h.  */
#     ifndef EXIT_SUCCESS
#      define EXIT_SUCCESS 0
#     endif
#    endif
#   endif
#  endif
# endif

# ifdef YYSTACK_ALLOC
   /* Pacify GCC's `empty if-body' warning.  */
#  define YYSTACK_FREE(Ptr) do { /* empty */; } while (YYID (0))
#  ifndef YYSTACK_ALLOC_MAXIMUM
    /* The OS might guarantee only one guard page at the bottom of the stack,
       and a page size can be as small as 4096 bytes.  So we cannot safely
       invoke alloca (N) if N exceeds 4096.  Use a slightly smaller number
       to allow for a few compiler-allocated temporary stack slots.  */
#   define YYSTACK_ALLOC_MAXIMUM 4032 /* reasonable circa 2006 */
#  endif
# else
#  define YYSTACK_ALLOC YYMALLOC
#  define YYSTACK_FREE YYFREE
#  ifndef YYSTACK_ALLOC_MAXIMUM
#   define YYSTACK_ALLOC_MAXIMUM YYSIZE_MAXIMUM
#  endif
#  if (defined __cplusplus && ! defined EXIT_SUCCESS \
       && ! ((defined YYMALLOC || defined malloc) \
	     && (defined YYFREE || defined free)))
#   include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
#   ifndef EXIT_SUCCESS
#    define EXIT_SUCCESS 0
#   endif
#  endif
#  ifndef YYMALLOC
#   define YYMALLOC malloc
#   if ! defined malloc && ! defined EXIT_SUCCESS && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
void *malloc (YYSIZE_T); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
#  ifndef YYFREE
#   define YYFREE free
#   if ! defined free && ! defined EXIT_SUCCESS && (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
void free (void *); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
# endif
#endif /* ! defined htmloverflow || YYERROR_VERBOSE */


#if (! defined htmloverflow \
     && (! defined __cplusplus \
	 || (defined YYSTYPE_IS_TRIVIAL && YYSTYPE_IS_TRIVIAL)))

/* A type that is properly aligned for any stack member.  */
union htmlalloc
{
  htmltype_int16 htmlss_alloc;
  YYSTYPE htmlvs_alloc;
};

/* The size of the maximum gap between one aligned stack and the next.  */
# define YYSTACK_GAP_MAXIMUM (sizeof (union htmlalloc) - 1)

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */
# define YYSTACK_BYTES(N) \
     ((N) * (sizeof (htmltype_int16) + sizeof (YYSTYPE)) \
      + YYSTACK_GAP_MAXIMUM)

# define YYCOPY_NEEDED 1

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */
# define YYSTACK_RELOCATE(Stack_alloc, Stack)				\
    do									\
      {									\
	YYSIZE_T htmlnewbytes;						\
	YYCOPY (&htmlptr->Stack_alloc, Stack, htmlsize);			\
	Stack = &htmlptr->Stack_alloc;					\
	htmlnewbytes = htmlstacksize * sizeof (*Stack) + YYSTACK_GAP_MAXIMUM; \
	htmlptr += htmlnewbytes / sizeof (*htmlptr);				\
      }									\
    while (YYID (0))

#endif

#if defined YYCOPY_NEEDED && YYCOPY_NEEDED
/* Copy COUNT objects from SRC to DST.  The source and destination do
   not overlap.  */
# ifndef YYCOPY
#  if defined __GNUC__ && 1 < __GNUC__
#   define YYCOPY(Dst, Src, Count) \
      __builtin_memcpy (Dst, Src, (Count) * sizeof (*(Src)))
#  else
#   define YYCOPY(Dst, Src, Count)              \
      do                                        \
        {                                       \
          YYSIZE_T htmli;                         \
          for (htmli = 0; htmli < (Count); htmli++)   \
            (Dst)[htmli] = (Src)[htmli];            \
        }                                       \
      while (YYID (0))
#  endif
# endif
#endif /* !YYCOPY_NEEDED */

/* YYFINAL -- State number of the termination state.  */
#define YYFINAL  31
/* YYLAST -- Last index in YYTABLE.  */
#define YYLAST   271

/* YYNTOKENS -- Number of terminals.  */
#define YYNTOKENS  41
/* YYNNTS -- Number of nonterminals.  */
#define YYNNTS  39
/* YYNRULES -- Number of rules.  */
#define YYNRULES  69
/* YYNRULES -- Number of states.  */
#define YYNSTATES  116

/* YYTRANSLATE(YYLEX) -- Bison symbol number corresponding to YYLEX.  */
#define YYUNDEFTOK  2
#define YYMAXUTOK   295

#define YYTRANSLATE(YYX)						\
  ((unsigned int) (YYX) <= YYMAXUTOK ? htmltranslate[YYX] : YYUNDEFTOK)

/* YYTRANSLATE[YYLEX] -- Bison symbol number corresponding to YYLEX.  */
static const htmltype_uint8 htmltranslate[] =
{
       0,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     1,     2,     3,     4,
       5,     6,     7,     8,     9,    10,    11,    12,    13,    14,
      15,    16,    17,    18,    19,    20,    21,    22,    23,    24,
      25,    26,    27,    28,    29,    30,    31,    32,    33,    34,
      35,    36,    37,    38,    39,    40
};

#if YYDEBUG
/* YYPRHS[YYN] -- Index of the first RHS symbol of rule number YYN in
   YYRHS.  */
static const htmltype_uint8 htmlprhs[] =
{
       0,     0,     3,     7,    11,    13,    15,    18,    20,    22,
      24,    28,    32,    36,    40,    44,    48,    52,    56,    58,
      60,    62,    64,    66,    68,    70,    72,    74,    76,    78,
      80,    82,    84,    86,    88,    91,    93,    95,    98,    99,
     106,   108,   112,   116,   120,   124,   128,   130,   131,   133,
     136,   140,   141,   146,   148,   151,   155,   156,   161,   162,
     167,   168,   173,   174,   178,   181,   183,   186,   188,   191
};

/* YYRHS -- A `-1'-separated list of the rules' RHS.  */
static const htmltype_int8 htmlrhs[] =
{
      42,     0,    -1,     7,    43,     8,    -1,     7,    66,     8,
      -1,     1,    -1,    44,    -1,    44,    45,    -1,    45,    -1,
      63,    -1,    62,    -1,    46,    44,    47,    -1,    48,    44,
      49,    -1,    54,    44,    55,    -1,    56,    44,    57,    -1,
      50,    44,    51,    -1,    58,    44,    59,    -1,    60,    44,
      61,    -1,    52,    44,    53,    -1,    33,    -1,    11,    -1,
      34,    -1,    14,    -1,    35,    -1,    15,    -1,    40,    -1,
      20,    -1,    36,    -1,    16,    -1,    37,    -1,    17,    -1,
      38,    -1,    18,    -1,    39,    -1,    19,    -1,    28,     3,
      -1,    27,    -1,    12,    -1,    63,    12,    -1,    -1,    67,
      31,    65,    68,     9,    67,    -1,    64,    -1,    46,    64,
      47,    -1,    48,    64,    49,    -1,    54,    64,    55,    -1,
      56,    64,    57,    -1,    50,    64,    51,    -1,    63,    -1,
      -1,    69,    -1,    68,    69,    -1,    68,    78,    69,    -1,
      -1,     5,    70,    71,     6,    -1,    72,    -1,    71,    72,
      -1,    71,    79,    72,    -1,    -1,    32,    66,    73,    10,
      -1,    -1,    32,    43,    74,    10,    -1,    -1,    32,    77,
      75,    10,    -1,    -1,    32,    76,    10,    -1,    30,     4,
      -1,    29,    -1,    22,    23,    -1,    21,    -1,    25,    26,
      -1,    24,    -1
};

/* YYRLINE[YYN] -- source line where rule number YYN was defined.  */
static const htmltype_uint16 htmlrline[] =
{
       0,   447,   447,   448,   449,   452,   455,   456,   459,   460,
     461,   462,   463,   464,   465,   466,   467,   468,   471,   474,
     477,   480,   483,   486,   489,   492,   495,   498,   501,   504,
     507,   510,   513,   516,   519,   520,   523,   524,   527,   527,
     548,   549,   550,   551,   552,   553,   556,   557,   560,   561,
     562,   565,   565,   568,   569,   570,   573,   573,   574,   574,
     575,   575,   576,   576,   579,   580,   583,   584,   587,   588
};
#endif

#if YYDEBUG || YYERROR_VERBOSE || 0
/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */
static const char *const htmltname[] =
{
  "$end", "error", "$undefined", "T_end_br", "T_end_img", "T_row",
  "T_end_row", "T_html", "T_end_html", "T_end_table", "T_end_cell",
  "T_end_font", "T_string", "T_error", "T_n_italic", "T_n_bold",
  "T_n_underline", "T_n_overline", "T_n_sup", "T_n_sub", "T_n_s", "T_HR",
  "T_hr", "T_end_hr", "T_VR", "T_vr", "T_end_vr", "T_BR", "T_br", "T_IMG",
  "T_img", "T_table", "T_cell", "T_font", "T_italic", "T_bold",
  "T_underline", "T_overline", "T_sup", "T_sub", "T_s", "$accept", "html",
  "fonttext", "text", "textitem", "font", "n_font", "italic", "n_italic",
  "bold", "n_bold", "strike", "n_strike", "underline", "n_underline",
  "overline", "n_overline", "sup", "n_sup", "sub", "n_sub", "br", "string",
  "table", "@1", "fonttable", "opt_space", "rows", "row", "$@2", "cells",
  "cell", "$@3", "$@4", "$@5", "$@6", "image", "HR", "VR", YY_NULL
};
#endif

# ifdef YYPRINT
/* YYTOKNUM[YYLEX-NUM] -- Internal token number corresponding to
   token YYLEX-NUM.  */
static const htmltype_uint16 htmltoknum[] =
{
       0,   256,   257,   258,   259,   260,   261,   262,   263,   264,
     265,   266,   267,   268,   269,   270,   271,   272,   273,   274,
     275,   276,   277,   278,   279,   280,   281,   282,   283,   284,
     285,   286,   287,   288,   289,   290,   291,   292,   293,   294,
     295
};
# endif

/* YYR1[YYN] -- Symbol number of symbol that rule YYN derives.  */
static const htmltype_uint8 htmlr1[] =
{
       0,    41,    42,    42,    42,    43,    44,    44,    45,    45,
      45,    45,    45,    45,    45,    45,    45,    45,    46,    47,
      48,    49,    50,    51,    52,    53,    54,    55,    56,    57,
      58,    59,    60,    61,    62,    62,    63,    63,    65,    64,
      66,    66,    66,    66,    66,    66,    67,    67,    68,    68,
      68,    70,    69,    71,    71,    71,    73,    72,    74,    72,
      75,    72,    76,    72,    77,    77,    78,    78,    79,    79
};

/* YYR2[YYN] -- Number of symbols composing right hand side of rule YYN.  */
static const htmltype_uint8 htmlr2[] =
{
       0,     2,     3,     3,     1,     1,     2,     1,     1,     1,
       3,     3,     3,     3,     3,     3,     3,     3,     1,     1,
       1,     1,     1,     1,     1,     1,     1,     1,     1,     1,
       1,     1,     1,     1,     2,     1,     1,     2,     0,     6,
       1,     3,     3,     3,     3,     3,     1,     0,     1,     2,
       3,     0,     4,     1,     2,     3,     0,     4,     0,     4,
       0,     4,     0,     3,     2,     1,     2,     1,     2,     1
};

/* YYDEFACT[STATE-NAME] -- Default reduction number in state STATE-NUM.
   Performed when YYTABLE doesn't specify something else to do.  Zero
   means the default is an error.  */
static const htmltype_uint8 htmldefact[] =
{
       0,     4,    47,     0,    36,    35,     0,    18,    20,    22,
      26,    28,    30,    32,    24,     0,     5,     7,    47,    47,
      47,     0,    47,    47,     0,     0,     9,     8,    40,     0,
       0,     1,    34,     2,     6,     0,     0,     0,     0,     0,
       8,     0,     0,     0,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,    37,     3,    38,    19,    10,    41,
      21,    11,    42,    23,    14,    45,    25,    17,    27,    12,
      43,    29,    13,    44,    31,    15,    33,    16,     0,    51,
       0,    48,     0,    47,    67,     0,    49,     0,    47,     0,
      53,    46,    39,    66,    50,    65,     0,    58,    56,     0,
      60,    52,    69,     0,    54,     0,    64,     0,     0,    63,
       0,    68,    55,    59,    57,    61
};

/* YYDEFGOTO[NTERM-NUM].  */
static const htmltype_int8 htmldefgoto[] =
{
      -1,     3,    15,    16,    17,    35,    58,    36,    61,    37,
      64,    21,    67,    38,    69,    39,    72,    24,    75,    25,
      77,    26,    40,    28,    78,    29,    30,    80,    81,    82,
      89,    90,   108,   107,   110,    99,   100,    87,   105
};

/* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
   STATE-NUM.  */
#define YYPACT_NINF -82
static const htmltype_int16 htmlpact[] =
{
       8,   -82,   209,    10,   -82,   -82,    11,   -82,   -82,   -82,
     -82,   -82,   -82,   -82,   -82,     5,   209,   -82,   209,   209,
     209,   209,   209,   209,   209,   209,   -82,    -5,   -82,    14,
     -20,   -82,   -82,   -82,   -82,   209,   209,   209,   209,   209,
      13,    37,    12,    66,    16,    80,    19,   109,   123,    20,
     152,    15,   166,   195,   -82,   -82,   -82,   -82,   -82,   -82,
     -82,   -82,   -82,   -82,   -82,   -82,   -82,   -82,   -82,   -82,
     -82,   -82,   -82,   -82,   -82,   -82,   -82,   -82,    23,   -82,
     119,   -82,     7,    46,   -82,    38,   -82,    23,    17,    35,
     -82,    13,   -82,   -82,   -82,   -82,    58,   -82,   -82,    53,
     -82,   -82,   -82,    40,   -82,     7,   -82,    59,    69,   -82,
      72,   -82,   -82,   -82,   -82,   -82
};

/* YYPGOTO[NTERM-NUM].  */
static const htmltype_int16 htmlpgoto[] =
{
     -82,   -82,    -4,   232,   -10,    -1,    26,     0,    39,     1,
      50,   -82,   -82,     2,    36,     3,    47,   -82,   -82,   -82,
     -82,   -82,    -2,   148,   -82,     9,    27,   -82,   -68,   -82,
     -82,   -81,   -82,   -82,   -82,   -82,   -82,   -82,   -82
};

/* YYTABLE[YYPACT[STATE-NUM]].  What to do in state STATE-NUM.  If
   positive, shift that token.  If negative, reduce the rule which
   number is the opposite.  If YYTABLE_NINF, syntax error.  */
#define YYTABLE_NINF -63
static const htmltype_int8 htmltable[] =
{
      27,    18,    19,    20,    22,    23,    34,    54,   104,     1,
      31,    56,    86,    33,    32,     2,    27,    27,    27,    94,
      27,    27,    55,    57,   112,    54,   -46,   -62,    79,     4,
      60,    34,    71,    34,    63,    34,    68,    34,    34,    88,
      34,   101,    34,    34,     5,     6,    95,    96,    57,     4,
       7,     8,     9,    10,    11,    12,    13,    14,     4,   102,
     103,    93,   106,   109,     5,     6,   111,    88,    59,   113,
       7,     8,     9,    10,    11,    12,    13,    14,     4,   114,
      60,    91,   115,    62,    97,    70,    27,    18,    19,    20,
      22,    23,     4,     5,     6,    63,    65,    98,    73,     7,
       8,     9,    10,    11,    12,    13,    14,     5,     6,     0,
      92,     0,     0,     7,     8,     9,    10,    11,    12,    13,
      14,     4,     0,     0,    79,     0,     0,     0,    83,    66,
       0,     0,     0,     0,     0,     4,     5,     6,     0,    68,
      84,    85,     7,     8,     9,    10,    11,    12,    13,    14,
       5,     6,     0,     0,     0,     0,     7,     8,     9,    10,
      11,    12,    13,    14,     4,     0,    42,    44,    46,    71,
      49,    51,     0,     0,     0,     0,     0,     0,     4,     5,
       6,     0,     0,     0,    74,     7,     8,     9,    10,    11,
      12,    13,    14,     5,     6,     0,     0,     0,     0,     7,
       8,     9,    10,    11,    12,    13,    14,     4,     0,     0,
       0,     0,     0,     0,    76,     0,     0,     0,     0,     0,
       0,     4,     5,     6,     0,     0,     0,     0,     7,     8,
       9,    10,    11,    12,    13,    14,     5,     6,     0,     0,
       0,     0,     7,     8,     9,    10,    11,    12,    13,    14,
      41,    43,    45,    47,    48,    50,    52,    53,     0,     0,
       0,     0,     0,     0,     0,     0,     0,    41,    43,    45,
      48,    50
};

#define htmlpact_value_is_default(Yystate) \
  (!!((Yystate) == (-82)))

#define htmltable_value_is_error(Yytable_value) \
  YYID (0)

static const htmltype_int8 htmlcheck[] =
{
       2,     2,     2,     2,     2,     2,    16,    12,    89,     1,
       0,    31,    80,     8,     3,     7,    18,    19,    20,    87,
      22,    23,     8,    11,   105,    12,    31,    10,     5,    12,
      14,    41,    17,    43,    15,    45,    16,    47,    48,    32,
      50,     6,    52,    53,    27,    28,    29,    30,    11,    12,
      33,    34,    35,    36,    37,    38,    39,    40,    12,    24,
      25,    23,     4,    10,    27,    28,    26,    32,    42,    10,
      33,    34,    35,    36,    37,    38,    39,    40,    12,    10,
      14,    83,    10,    44,    88,    49,    88,    88,    88,    88,
      88,    88,    12,    27,    28,    15,    46,    88,    51,    33,
      34,    35,    36,    37,    38,    39,    40,    27,    28,    -1,
      83,    -1,    -1,    33,    34,    35,    36,    37,    38,    39,
      40,    12,    -1,    -1,     5,    -1,    -1,    -1,     9,    20,
      -1,    -1,    -1,    -1,    -1,    12,    27,    28,    -1,    16,
      21,    22,    33,    34,    35,    36,    37,    38,    39,    40,
      27,    28,    -1,    -1,    -1,    -1,    33,    34,    35,    36,
      37,    38,    39,    40,    12,    -1,    18,    19,    20,    17,
      22,    23,    -1,    -1,    -1,    -1,    -1,    -1,    12,    27,
      28,    -1,    -1,    -1,    18,    33,    34,    35,    36,    37,
      38,    39,    40,    27,    28,    -1,    -1,    -1,    -1,    33,
      34,    35,    36,    37,    38,    39,    40,    12,    -1,    -1,
      -1,    -1,    -1,    -1,    19,    -1,    -1,    -1,    -1,    -1,
      -1,    12,    27,    28,    -1,    -1,    -1,    -1,    33,    34,
      35,    36,    37,    38,    39,    40,    27,    28,    -1,    -1,
      -1,    -1,    33,    34,    35,    36,    37,    38,    39,    40,
      18,    19,    20,    21,    22,    23,    24,    25,    -1,    -1,
      -1,    -1,    -1,    -1,    -1,    -1,    -1,    35,    36,    37,
      38,    39
};

/* YYSTOS[STATE-NUM] -- The (internal number of the) accessing
   symbol of state STATE-NUM.  */
static const htmltype_uint8 htmlstos[] =
{
       0,     1,     7,    42,    12,    27,    28,    33,    34,    35,
      36,    37,    38,    39,    40,    43,    44,    45,    46,    48,
      50,    52,    54,    56,    58,    60,    62,    63,    64,    66,
      67,     0,     3,     8,    45,    46,    48,    50,    54,    56,
      63,    44,    64,    44,    64,    44,    64,    44,    44,    64,
      44,    64,    44,    44,    12,     8,    31,    11,    47,    47,
      14,    49,    49,    15,    51,    51,    20,    53,    16,    55,
      55,    17,    57,    57,    18,    59,    19,    61,    65,     5,
      68,    69,    70,     9,    21,    22,    69,    78,    32,    71,
      72,    63,    67,    23,    69,    29,    30,    43,    66,    76,
      77,     6,    24,    25,    72,    79,     4,    74,    73,    10,
      75,    26,    72,    10,    10,    10
};

#define htmlerrok		(htmlerrstatus = 0)
#define htmlclearin	(htmlchar = YYEMPTY)
#define YYEMPTY		(-2)
#define YYEOF		0

#define YYACCEPT	goto htmlacceptlab
#define YYABORT		goto htmlabortlab
#define YYERROR		goto htmlerrorlab


/* Like YYERROR except do call htmlerror.  This remains here temporarily
   to ease the transition to the new meaning of YYERROR, for GCC.
   Once GCC version 2 has supplanted version 1, this can go.  However,
   YYFAIL appears to be in use.  Nevertheless, it is formally deprecated
   in Bison 2.4.2's NEWS entry, where a plan to phase it out is
   discussed.  */

#define YYFAIL		goto htmlerrlab
#if defined YYFAIL
  /* This is here to suppress warnings from the GCC cpp's
     -Wunused-macros.  Normally we don't worry about that warning, but
     some users do, and we want to make it easy for users to remove
     YYFAIL uses, which will produce warnings from Bison 2.5.  */
#endif

#define YYRECOVERING()  (!!htmlerrstatus)

#define YYBACKUP(Token, Value)                                  \
do                                                              \
  if (htmlchar == YYEMPTY)                                        \
    {                                                           \
      htmlchar = (Token);                                         \
      htmllval = (Value);                                         \
      YYPOPSTACK (htmllen);                                       \
      htmlstate = *htmlssp;                                         \
      goto htmlbackup;                                            \
    }                                                           \
  else                                                          \
    {                                                           \
      htmlerror (YY_("syntax error: cannot back up")); \
      YYERROR;							\
    }								\
while (YYID (0))

/* Error token number */
#define YYTERROR	1
#define YYERRCODE	256


/* This macro is provided for backward compatibility. */
#ifndef YY_LOCATION_PRINT
# define YY_LOCATION_PRINT(File, Loc) ((void) 0)
#endif


/* YYLEX -- calling `htmllex' with the right arguments.  */
#ifdef YYLEX_PARAM
# define YYLEX htmllex (YYLEX_PARAM)
#else
# define YYLEX htmllex ()
#endif

/* Enable debugging if requested.  */
#if YYDEBUG

# ifndef YYFPRINTF
#  include <stdio.h> /* INFRINGES ON USER NAME SPACE */
#  define YYFPRINTF fprintf
# endif

# define YYDPRINTF(Args)			\
do {						\
  if (htmldebug)					\
    YYFPRINTF Args;				\
} while (YYID (0))

# define YY_SYMBOL_PRINT(Title, Type, Value, Location)			  \
do {									  \
  if (htmldebug)								  \
    {									  \
      YYFPRINTF (stderr, "%s ", Title);					  \
      html_symbol_print (stderr,						  \
		  Type, Value); \
      YYFPRINTF (stderr, "\n");						  \
    }									  \
} while (YYID (0))


/*--------------------------------.
| Print this symbol on YYOUTPUT.  |
`--------------------------------*/

/*ARGSUSED*/
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
html_symbol_value_print (FILE *htmloutput, int htmltype, YYSTYPE const * const htmlvaluep)
#else
static void
html_symbol_value_print (htmloutput, htmltype, htmlvaluep)
    FILE *htmloutput;
    int htmltype;
    YYSTYPE const * const htmlvaluep;
#endif
{
  FILE *htmlo = htmloutput;
  YYUSE (htmlo);
  if (!htmlvaluep)
    return;
# ifdef YYPRINT
  if (htmltype < YYNTOKENS)
    YYPRINT (htmloutput, htmltoknum[htmltype], *htmlvaluep);
# else
  YYUSE (htmloutput);
# endif
  switch (htmltype)
    {
      default:
        break;
    }
}


/*--------------------------------.
| Print this symbol on YYOUTPUT.  |
`--------------------------------*/

#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
html_symbol_print (FILE *htmloutput, int htmltype, YYSTYPE const * const htmlvaluep)
#else
static void
html_symbol_print (htmloutput, htmltype, htmlvaluep)
    FILE *htmloutput;
    int htmltype;
    YYSTYPE const * const htmlvaluep;
#endif
{
  if (htmltype < YYNTOKENS)
    YYFPRINTF (htmloutput, "token %s (", htmltname[htmltype]);
  else
    YYFPRINTF (htmloutput, "nterm %s (", htmltname[htmltype]);

  html_symbol_value_print (htmloutput, htmltype, htmlvaluep);
  YYFPRINTF (htmloutput, ")");
}

/*------------------------------------------------------------------.
| html_stack_print -- Print the state stack from its BOTTOM up to its |
| TOP (included).                                                   |
`------------------------------------------------------------------*/

#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
html_stack_print (htmltype_int16 *htmlbottom, htmltype_int16 *htmltop)
#else
static void
html_stack_print (htmlbottom, htmltop)
    htmltype_int16 *htmlbottom;
    htmltype_int16 *htmltop;
#endif
{
  YYFPRINTF (stderr, "Stack now");
  for (; htmlbottom <= htmltop; htmlbottom++)
    {
      int htmlbot = *htmlbottom;
      YYFPRINTF (stderr, " %d", htmlbot);
    }
  YYFPRINTF (stderr, "\n");
}

# define YY_STACK_PRINT(Bottom, Top)				\
do {								\
  if (htmldebug)							\
    html_stack_print ((Bottom), (Top));				\
} while (YYID (0))


/*------------------------------------------------.
| Report that the YYRULE is going to be reduced.  |
`------------------------------------------------*/

#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
html_reduce_print (YYSTYPE *htmlvsp, int htmlrule)
#else
static void
html_reduce_print (htmlvsp, htmlrule)
    YYSTYPE *htmlvsp;
    int htmlrule;
#endif
{
  int htmlnrhs = htmlr2[htmlrule];
  int htmli;
  unsigned long int htmllno = htmlrline[htmlrule];
  YYFPRINTF (stderr, "Reducing stack by rule %d (line %lu):\n",
	     htmlrule - 1, htmllno);
  /* The symbols being reduced.  */
  for (htmli = 0; htmli < htmlnrhs; htmli++)
    {
      YYFPRINTF (stderr, "   $%d = ", htmli + 1);
      html_symbol_print (stderr, htmlrhs[htmlprhs[htmlrule] + htmli],
		       &(htmlvsp[(htmli + 1) - (htmlnrhs)])
		       		       );
      YYFPRINTF (stderr, "\n");
    }
}

# define YY_REDUCE_PRINT(Rule)		\
do {					\
  if (htmldebug)				\
    html_reduce_print (htmlvsp, Rule); \
} while (YYID (0))

/* Nonzero means print parse trace.  It is left uninitialized so that
   multiple parsers can coexist.  */
int htmldebug;
#else /* !YYDEBUG */
# define YYDPRINTF(Args)
# define YY_SYMBOL_PRINT(Title, Type, Value, Location)
# define YY_STACK_PRINT(Bottom, Top)
# define YY_REDUCE_PRINT(Rule)
#endif /* !YYDEBUG */


/* YYINITDEPTH -- initial size of the parser's stacks.  */
#ifndef	YYINITDEPTH
# define YYINITDEPTH 200
#endif

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

#ifndef YYMAXDEPTH
# define YYMAXDEPTH 10000
#endif


#if YYERROR_VERBOSE

# ifndef htmlstrlen
#  if defined __GLIBC__ && defined _STRING_H
#   define htmlstrlen strlen
#  else
/* Return the length of YYSTR.  */
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static YYSIZE_T
htmlstrlen (const char *htmlstr)
#else
static YYSIZE_T
htmlstrlen (htmlstr)
    const char *htmlstr;
#endif
{
  YYSIZE_T htmllen;
  for (htmllen = 0; htmlstr[htmllen]; htmllen++)
    continue;
  return htmllen;
}
#  endif
# endif

# ifndef htmlstpcpy
#  if defined __GLIBC__ && defined _STRING_H && defined _GNU_SOURCE
#   define htmlstpcpy stpcpy
#  else
/* Copy YYSRC to YYDEST, returning the address of the terminating '\0' in
   YYDEST.  */
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static char *
htmlstpcpy (char *htmldest, const char *htmlsrc)
#else
static char *
htmlstpcpy (htmldest, htmlsrc)
    char *htmldest;
    const char *htmlsrc;
#endif
{
  char *htmld = htmldest;
  const char *htmls = htmlsrc;

  while ((*htmld++ = *htmls++) != '\0')
    continue;

  return htmld - 1;
}
#  endif
# endif

# ifndef htmltnamerr
/* Copy to YYRES the contents of YYSTR after stripping away unnecessary
   quotes and backslashes, so that it's suitable for htmlerror.  The
   heuristic is that double-quoting is unnecessary unless the string
   contains an apostrophe, a comma, or backslash (other than
   backslash-backslash).  YYSTR is taken from htmltname.  If YYRES is
   null, do not copy; instead, return the length of what the result
   would have been.  */
static YYSIZE_T
htmltnamerr (char *htmlres, const char *htmlstr)
{
  if (*htmlstr == '"')
    {
      YYSIZE_T htmln = 0;
      char const *htmlp = htmlstr;

      for (;;)
	switch (*++htmlp)
	  {
	  case '\'':
	  case ',':
	    goto do_not_strip_quotes;

	  case '\\':
	    if (*++htmlp != '\\')
	      goto do_not_strip_quotes;
	    /* Fall through.  */
	  default:
	    if (htmlres)
	      htmlres[htmln] = *htmlp;
	    htmln++;
	    break;

	  case '"':
	    if (htmlres)
	      htmlres[htmln] = '\0';
	    return htmln;
	  }
    do_not_strip_quotes: ;
    }

  if (! htmlres)
    return htmlstrlen (htmlstr);

  return htmlstpcpy (htmlres, htmlstr) - htmlres;
}
# endif

/* Copy into *YYMSG, which is of size *YYMSG_ALLOC, an error message
   about the unexpected token YYTOKEN for the state stack whose top is
   YYSSP.

   Return 0 if *YYMSG was successfully written.  Return 1 if *YYMSG is
   not large enough to hold the message.  In that case, also set
   *YYMSG_ALLOC to the required number of bytes.  Return 2 if the
   required number of bytes is too large to store.  */
static int
htmlsyntax_error (YYSIZE_T *htmlmsg_alloc, char **htmlmsg,
                htmltype_int16 *htmlssp, int htmltoken)
{
  YYSIZE_T htmlsize0 = htmltnamerr (YY_NULL, htmltname[htmltoken]);
  YYSIZE_T htmlsize = htmlsize0;
  enum { YYERROR_VERBOSE_ARGS_MAXIMUM = 5 };
  /* Internationalized format string. */
  const char *htmlformat = YY_NULL;
  /* Arguments of htmlformat. */
  char const *htmlarg[YYERROR_VERBOSE_ARGS_MAXIMUM];
  /* Number of reported tokens (one for the "unexpected", one per
     "expected"). */
  int htmlcount = 0;

  /* There are many possibilities here to consider:
     - Assume YYFAIL is not used.  It's too flawed to consider.  See
       <http://lists.gnu.org/archive/html/bison-patches/2009-12/msg00024.html>
       for details.  YYERROR is fine as it does not invoke this
       function.
     - If this state is a consistent state with a default action, then
       the only way this function was invoked is if the default action
       is an error action.  In that case, don't check for expected
       tokens because there are none.
     - The only way there can be no lookahead present (in htmlchar) is if
       this state is a consistent state with a default action.  Thus,
       detecting the absence of a lookahead is sufficient to determine
       that there is no unexpected or expected token to report.  In that
       case, just report a simple "syntax error".
     - Don't assume there isn't a lookahead just because this state is a
       consistent state with a default action.  There might have been a
       previous inconsistent state, consistent state with a non-default
       action, or user semantic action that manipulated htmlchar.
     - Of course, the expected token list depends on states to have
       correct lookahead information, and it depends on the parser not
       to perform extra reductions after fetching a lookahead from the
       scanner and before detecting a syntax error.  Thus, state merging
       (from LALR or IELR) and default reductions corrupt the expected
       token list.  However, the list is correct for canonical LR with
       one exception: it will still contain any token that will not be
       accepted due to an error action in a later state.
  */
  if (htmltoken != YYEMPTY)
    {
      int htmln = htmlpact[*htmlssp];
      htmlarg[htmlcount++] = htmltname[htmltoken];
      if (!htmlpact_value_is_default (htmln))
        {
          /* Start YYX at -YYN if negative to avoid negative indexes in
             YYCHECK.  In other words, skip the first -YYN actions for
             this state because they are default actions.  */
          int htmlxbegin = htmln < 0 ? -htmln : 0;
          /* Stay within bounds of both htmlcheck and htmltname.  */
          int htmlchecklim = YYLAST - htmln + 1;
          int htmlxend = htmlchecklim < YYNTOKENS ? htmlchecklim : YYNTOKENS;
          int htmlx;

          for (htmlx = htmlxbegin; htmlx < htmlxend; ++htmlx)
            if (htmlcheck[htmlx + htmln] == htmlx && htmlx != YYTERROR
                && !htmltable_value_is_error (htmltable[htmlx + htmln]))
              {
                if (htmlcount == YYERROR_VERBOSE_ARGS_MAXIMUM)
                  {
                    htmlcount = 1;
                    htmlsize = htmlsize0;
                    break;
                  }
                htmlarg[htmlcount++] = htmltname[htmlx];
                {
                  YYSIZE_T htmlsize1 = htmlsize + htmltnamerr (YY_NULL, htmltname[htmlx]);
                  if (! (htmlsize <= htmlsize1
                         && htmlsize1 <= YYSTACK_ALLOC_MAXIMUM))
                    return 2;
                  htmlsize = htmlsize1;
                }
              }
        }
    }

  switch (htmlcount)
    {
# define YYCASE_(N, S)                      \
      case N:                               \
        htmlformat = S;                       \
      break
      YYCASE_(0, YY_("syntax error"));
      YYCASE_(1, YY_("syntax error, unexpected %s"));
      YYCASE_(2, YY_("syntax error, unexpected %s, expecting %s"));
      YYCASE_(3, YY_("syntax error, unexpected %s, expecting %s or %s"));
      YYCASE_(4, YY_("syntax error, unexpected %s, expecting %s or %s or %s"));
      YYCASE_(5, YY_("syntax error, unexpected %s, expecting %s or %s or %s or %s"));
# undef YYCASE_
    }

  {
    YYSIZE_T htmlsize1 = htmlsize + htmlstrlen (htmlformat);
    if (! (htmlsize <= htmlsize1 && htmlsize1 <= YYSTACK_ALLOC_MAXIMUM))
      return 2;
    htmlsize = htmlsize1;
  }

  if (*htmlmsg_alloc < htmlsize)
    {
      *htmlmsg_alloc = 2 * htmlsize;
      if (! (htmlsize <= *htmlmsg_alloc
             && *htmlmsg_alloc <= YYSTACK_ALLOC_MAXIMUM))
        *htmlmsg_alloc = YYSTACK_ALLOC_MAXIMUM;
      return 1;
    }

  /* Avoid sprintf, as that infringes on the user's name space.
     Don't have undefined behavior even if the translation
     produced a string with the wrong number of "%s"s.  */
  {
    char *htmlp = *htmlmsg;
    int htmli = 0;
    while ((*htmlp = *htmlformat) != '\0')
      if (*htmlp == '%' && htmlformat[1] == 's' && htmli < htmlcount)
        {
          htmlp += htmltnamerr (htmlp, htmlarg[htmli++]);
          htmlformat += 2;
        }
      else
        {
          htmlp++;
          htmlformat++;
        }
  }
  return 0;
}
#endif /* YYERROR_VERBOSE */

/*-----------------------------------------------.
| Release the memory associated to this symbol.  |
`-----------------------------------------------*/

/*ARGSUSED*/
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
static void
htmldestruct (const char *htmlmsg, int htmltype, YYSTYPE *htmlvaluep)
#else
static void
htmldestruct (htmlmsg, htmltype, htmlvaluep)
    const char *htmlmsg;
    int htmltype;
    YYSTYPE *htmlvaluep;
#endif
{
  YYUSE (htmlvaluep);

  if (!htmlmsg)
    htmlmsg = "Deleting";
  YY_SYMBOL_PRINT (htmlmsg, htmltype, htmlvaluep, htmllocationp);

  switch (htmltype)
    {

      default:
        break;
    }
}




/* The lookahead symbol.  */
int htmlchar;


#ifndef YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
# define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
# define YY_IGNORE_MAYBE_UNINITIALIZED_END
#endif
#ifndef YY_INITIAL_VALUE
# define YY_INITIAL_VALUE(Value) /* Nothing. */
#endif

/* The semantic value of the lookahead symbol.  */
YYSTYPE htmllval YY_INITIAL_VALUE(htmlval_default);

/* Number of syntax errors so far.  */
int htmlnerrs;


/*----------.
| htmlparse.  |
`----------*/

#ifdef YYPARSE_PARAM
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
int
htmlparse (void *YYPARSE_PARAM)
#else
int
htmlparse (YYPARSE_PARAM)
    void *YYPARSE_PARAM;
#endif
#else /* ! YYPARSE_PARAM */
#if (defined __STDC__ || defined __C99__FUNC__ \
     || defined __cplusplus || defined _MSC_VER)
int
htmlparse (void)
#else
int
htmlparse ()

#endif
#endif
{
    int htmlstate;
    /* Number of tokens to shift before error messages enabled.  */
    int htmlerrstatus;

    /* The stacks and their tools:
       `htmlss': related to states.
       `htmlvs': related to semantic values.

       Refer to the stacks through separate pointers, to allow htmloverflow
       to reallocate them elsewhere.  */

    /* The state stack.  */
    htmltype_int16 htmlssa[YYINITDEPTH];
    htmltype_int16 *htmlss;
    htmltype_int16 *htmlssp;

    /* The semantic value stack.  */
    YYSTYPE htmlvsa[YYINITDEPTH];
    YYSTYPE *htmlvs;
    YYSTYPE *htmlvsp;

    YYSIZE_T htmlstacksize;

  int htmln;
  int htmlresult;
  /* Lookahead token as an internal (translated) token number.  */
  int htmltoken = 0;
  /* The variables used to return semantic value and location from the
     action routines.  */
  YYSTYPE htmlval;

#if YYERROR_VERBOSE
  /* Buffer for error messages, and its allocated size.  */
  char htmlmsgbuf[128];
  char *htmlmsg = htmlmsgbuf;
  YYSIZE_T htmlmsg_alloc = sizeof htmlmsgbuf;
#endif

#define YYPOPSTACK(N)   (htmlvsp -= (N), htmlssp -= (N))

  /* The number of symbols on the RHS of the reduced rule.
     Keep to zero when no symbol should be popped.  */
  int htmllen = 0;

  htmlssp = htmlss = htmlssa;
  htmlvsp = htmlvs = htmlvsa;
  htmlstacksize = YYINITDEPTH;

  YYDPRINTF ((stderr, "Starting parse\n"));

  htmlstate = 0;
  htmlerrstatus = 0;
  htmlnerrs = 0;
  htmlchar = YYEMPTY; /* Cause a token to be read.  */
  goto htmlsetstate;

/*------------------------------------------------------------.
| htmlnewstate -- Push a new state, which is found in htmlstate.  |
`------------------------------------------------------------*/
 htmlnewstate:
  /* In all cases, when you get here, the value and location stacks
     have just been pushed.  So pushing a state here evens the stacks.  */
  htmlssp++;

 htmlsetstate:
  *htmlssp = htmlstate;

  if (htmlss + htmlstacksize - 1 <= htmlssp)
    {
      /* Get the current used size of the three stacks, in elements.  */
      YYSIZE_T htmlsize = htmlssp - htmlss + 1;

#ifdef htmloverflow
      {
	/* Give user a chance to reallocate the stack.  Use copies of
	   these so that the &'s don't force the real ones into
	   memory.  */
	YYSTYPE *htmlvs1 = htmlvs;
	htmltype_int16 *htmlss1 = htmlss;

	/* Each stack pointer address is followed by the size of the
	   data in use in that stack, in bytes.  This used to be a
	   conditional around just the two extra args, but that might
	   be undefined if htmloverflow is a macro.  */
	htmloverflow (YY_("memory exhausted"),
		    &htmlss1, htmlsize * sizeof (*htmlssp),
		    &htmlvs1, htmlsize * sizeof (*htmlvsp),
		    &htmlstacksize);

	htmlss = htmlss1;
	htmlvs = htmlvs1;
      }
#else /* no htmloverflow */
# ifndef YYSTACK_RELOCATE
      goto htmlexhaustedlab;
# else
      /* Extend the stack our own way.  */
      if (YYMAXDEPTH <= htmlstacksize)
	goto htmlexhaustedlab;
      htmlstacksize *= 2;
      if (YYMAXDEPTH < htmlstacksize)
	htmlstacksize = YYMAXDEPTH;

      {
	htmltype_int16 *htmlss1 = htmlss;
	union htmlalloc *htmlptr =
	  (union htmlalloc *) YYSTACK_ALLOC (YYSTACK_BYTES (htmlstacksize));
	if (! htmlptr)
	  goto htmlexhaustedlab;
	YYSTACK_RELOCATE (htmlss_alloc, htmlss);
	YYSTACK_RELOCATE (htmlvs_alloc, htmlvs);
#  undef YYSTACK_RELOCATE
	if (htmlss1 != htmlssa)
	  YYSTACK_FREE (htmlss1);
      }
# endif
#endif /* no htmloverflow */

      htmlssp = htmlss + htmlsize - 1;
      htmlvsp = htmlvs + htmlsize - 1;

      YYDPRINTF ((stderr, "Stack size increased to %lu\n",
		  (unsigned long int) htmlstacksize));

      if (htmlss + htmlstacksize - 1 <= htmlssp)
	YYABORT;
    }

  YYDPRINTF ((stderr, "Entering state %d\n", htmlstate));

  if (htmlstate == YYFINAL)
    YYACCEPT;

  goto htmlbackup;

/*-----------.
| htmlbackup.  |
`-----------*/
htmlbackup:

  /* Do appropriate processing given the current state.  Read a
     lookahead token if we need one and don't already have one.  */

  /* First try to decide what to do without reference to lookahead token.  */
  htmln = htmlpact[htmlstate];
  if (htmlpact_value_is_default (htmln))
    goto htmldefault;

  /* Not known => get a lookahead token if don't already have one.  */

  /* YYCHAR is either YYEMPTY or YYEOF or a valid lookahead symbol.  */
  if (htmlchar == YYEMPTY)
    {
      YYDPRINTF ((stderr, "Reading a token: "));
      htmlchar = YYLEX;
    }

  if (htmlchar <= YYEOF)
    {
      htmlchar = htmltoken = YYEOF;
      YYDPRINTF ((stderr, "Now at end of input.\n"));
    }
  else
    {
      htmltoken = YYTRANSLATE (htmlchar);
      YY_SYMBOL_PRINT ("Next token is", htmltoken, &htmllval, &htmllloc);
    }

  /* If the proper action on seeing token YYTOKEN is to reduce or to
     detect an error, take that action.  */
  htmln += htmltoken;
  if (htmln < 0 || YYLAST < htmln || htmlcheck[htmln] != htmltoken)
    goto htmldefault;
  htmln = htmltable[htmln];
  if (htmln <= 0)
    {
      if (htmltable_value_is_error (htmln))
        goto htmlerrlab;
      htmln = -htmln;
      goto htmlreduce;
    }

  /* Count tokens shifted since error; after three, turn off error
     status.  */
  if (htmlerrstatus)
    htmlerrstatus--;

  /* Shift the lookahead token.  */
  YY_SYMBOL_PRINT ("Shifting", htmltoken, &htmllval, &htmllloc);

  /* Discard the shifted token.  */
  htmlchar = YYEMPTY;

  htmlstate = htmln;
  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  *++htmlvsp = htmllval;
  YY_IGNORE_MAYBE_UNINITIALIZED_END

  goto htmlnewstate;


/*-----------------------------------------------------------.
| htmldefault -- do the default action for the current state.  |
`-----------------------------------------------------------*/
htmldefault:
  htmln = htmldefact[htmlstate];
  if (htmln == 0)
    goto htmlerrlab;
  goto htmlreduce;


/*-----------------------------.
| htmlreduce -- Do a reduction.  |
`-----------------------------*/
htmlreduce:
  /* htmln is the number of a rule to reduce with.  */
  htmllen = htmlr2[htmln];

  /* If YYLEN is nonzero, implement the default value of the action:
     `$$ = $1'.

     Otherwise, the following line sets YYVAL to garbage.
     This behavior is undocumented and Bison
     users should not rely upon it.  Assigning to YYVAL
     unconditionally makes the parser a bit smaller, and it avoids a
     GCC warning that YYVAL may be used uninitialized.  */
  htmlval = htmlvsp[1-htmllen];


  YY_REDUCE_PRINT (htmln);
  switch (htmln)
    {
        case 2:
/* Line 1792 of yacc.c  */
#line 447 "../../lib/common/htmlparse.y"
    { HTMLstate.lbl = mkLabel((htmlvsp[(2) - (3)].txt),HTML_TEXT); }
    break;

  case 3:
/* Line 1792 of yacc.c  */
#line 448 "../../lib/common/htmlparse.y"
    { HTMLstate.lbl = mkLabel((htmlvsp[(2) - (3)].tbl),HTML_TBL); }
    break;

  case 4:
/* Line 1792 of yacc.c  */
#line 449 "../../lib/common/htmlparse.y"
    { cleanup(); YYABORT; }
    break;

  case 5:
/* Line 1792 of yacc.c  */
#line 452 "../../lib/common/htmlparse.y"
    { (htmlval.txt) = mkText(); }
    break;

  case 8:
/* Line 1792 of yacc.c  */
#line 459 "../../lib/common/htmlparse.y"
    { appendFItemList(HTMLstate.str);}
    break;

  case 9:
/* Line 1792 of yacc.c  */
#line 460 "../../lib/common/htmlparse.y"
    {appendFLineList((htmlvsp[(1) - (1)].i));}
    break;

  case 18:
/* Line 1792 of yacc.c  */
#line 471 "../../lib/common/htmlparse.y"
    { pushFont ((htmlvsp[(1) - (1)].font)); }
    break;

  case 19:
/* Line 1792 of yacc.c  */
#line 474 "../../lib/common/htmlparse.y"
    { popFont (); }
    break;

  case 20:
/* Line 1792 of yacc.c  */
#line 477 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 21:
/* Line 1792 of yacc.c  */
#line 480 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 22:
/* Line 1792 of yacc.c  */
#line 483 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 23:
/* Line 1792 of yacc.c  */
#line 486 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 24:
/* Line 1792 of yacc.c  */
#line 489 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 25:
/* Line 1792 of yacc.c  */
#line 492 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 26:
/* Line 1792 of yacc.c  */
#line 495 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 27:
/* Line 1792 of yacc.c  */
#line 498 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 28:
/* Line 1792 of yacc.c  */
#line 501 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 29:
/* Line 1792 of yacc.c  */
#line 504 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 30:
/* Line 1792 of yacc.c  */
#line 507 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 31:
/* Line 1792 of yacc.c  */
#line 510 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 32:
/* Line 1792 of yacc.c  */
#line 513 "../../lib/common/htmlparse.y"
    {pushFont((htmlvsp[(1) - (1)].font));}
    break;

  case 33:
/* Line 1792 of yacc.c  */
#line 516 "../../lib/common/htmlparse.y"
    {popFont();}
    break;

  case 34:
/* Line 1792 of yacc.c  */
#line 519 "../../lib/common/htmlparse.y"
    { (htmlval.i) = (htmlvsp[(1) - (2)].i); }
    break;

  case 35:
/* Line 1792 of yacc.c  */
#line 520 "../../lib/common/htmlparse.y"
    { (htmlval.i) = (htmlvsp[(1) - (1)].i); }
    break;

  case 38:
/* Line 1792 of yacc.c  */
#line 527 "../../lib/common/htmlparse.y"
    { 
          if (nonSpace(agxbuse(HTMLstate.str))) {
            htmlerror ("Syntax error: non-space string used before <TABLE>");
            cleanup(); YYABORT;
          }
          (htmlvsp[(2) - (2)].tbl)->u.p.prev = HTMLstate.tblstack;
          (htmlvsp[(2) - (2)].tbl)->u.p.rows = dtopen(&rowDisc, Dtqueue);
          HTMLstate.tblstack = (htmlvsp[(2) - (2)].tbl);
          (htmlvsp[(2) - (2)].tbl)->font = HTMLstate.fontstack->cfont;
          (htmlval.tbl) = (htmlvsp[(2) - (2)].tbl);
        }
    break;

  case 39:
/* Line 1792 of yacc.c  */
#line 538 "../../lib/common/htmlparse.y"
    {
          if (nonSpace(agxbuse(HTMLstate.str))) {
            htmlerror ("Syntax error: non-space string used after </TABLE>");
            cleanup(); YYABORT;
          }
          (htmlval.tbl) = HTMLstate.tblstack;
          HTMLstate.tblstack = HTMLstate.tblstack->u.p.prev;
        }
    break;

  case 40:
/* Line 1792 of yacc.c  */
#line 548 "../../lib/common/htmlparse.y"
    { (htmlval.tbl) = (htmlvsp[(1) - (1)].tbl); }
    break;

  case 41:
/* Line 1792 of yacc.c  */
#line 549 "../../lib/common/htmlparse.y"
    { (htmlval.tbl)=(htmlvsp[(2) - (3)].tbl); }
    break;

  case 42:
/* Line 1792 of yacc.c  */
#line 550 "../../lib/common/htmlparse.y"
    { (htmlval.tbl)=(htmlvsp[(2) - (3)].tbl); }
    break;

  case 43:
/* Line 1792 of yacc.c  */
#line 551 "../../lib/common/htmlparse.y"
    { (htmlval.tbl)=(htmlvsp[(2) - (3)].tbl); }
    break;

  case 44:
/* Line 1792 of yacc.c  */
#line 552 "../../lib/common/htmlparse.y"
    { (htmlval.tbl)=(htmlvsp[(2) - (3)].tbl); }
    break;

  case 45:
/* Line 1792 of yacc.c  */
#line 553 "../../lib/common/htmlparse.y"
    { (htmlval.tbl)=(htmlvsp[(2) - (3)].tbl); }
    break;

  case 48:
/* Line 1792 of yacc.c  */
#line 560 "../../lib/common/htmlparse.y"
    { (htmlval.p) = (htmlvsp[(1) - (1)].p); }
    break;

  case 49:
/* Line 1792 of yacc.c  */
#line 561 "../../lib/common/htmlparse.y"
    { (htmlval.p) = (htmlvsp[(2) - (2)].p); }
    break;

  case 50:
/* Line 1792 of yacc.c  */
#line 562 "../../lib/common/htmlparse.y"
    { (htmlvsp[(1) - (3)].p)->ruled = 1; (htmlval.p) = (htmlvsp[(3) - (3)].p); }
    break;

  case 51:
/* Line 1792 of yacc.c  */
#line 565 "../../lib/common/htmlparse.y"
    { addRow (); }
    break;

  case 52:
/* Line 1792 of yacc.c  */
#line 565 "../../lib/common/htmlparse.y"
    { (htmlval.p) = lastRow(); }
    break;

  case 53:
/* Line 1792 of yacc.c  */
#line 568 "../../lib/common/htmlparse.y"
    { (htmlval.cell) = (htmlvsp[(1) - (1)].cell); }
    break;

  case 54:
/* Line 1792 of yacc.c  */
#line 569 "../../lib/common/htmlparse.y"
    { (htmlval.cell) = (htmlvsp[(2) - (2)].cell); }
    break;

  case 55:
/* Line 1792 of yacc.c  */
#line 570 "../../lib/common/htmlparse.y"
    { (htmlvsp[(1) - (3)].cell)->ruled |= HTML_VRULE; (htmlval.cell) = (htmlvsp[(3) - (3)].cell); }
    break;

  case 56:
/* Line 1792 of yacc.c  */
#line 573 "../../lib/common/htmlparse.y"
    { setCell((htmlvsp[(1) - (2)].cell),(htmlvsp[(2) - (2)].tbl),HTML_TBL); }
    break;

  case 57:
/* Line 1792 of yacc.c  */
#line 573 "../../lib/common/htmlparse.y"
    { (htmlval.cell) = (htmlvsp[(1) - (4)].cell); }
    break;

  case 58:
/* Line 1792 of yacc.c  */
#line 574 "../../lib/common/htmlparse.y"
    { setCell((htmlvsp[(1) - (2)].cell),(htmlvsp[(2) - (2)].txt),HTML_TEXT); }
    break;

  case 59:
/* Line 1792 of yacc.c  */
#line 574 "../../lib/common/htmlparse.y"
    { (htmlval.cell) = (htmlvsp[(1) - (4)].cell); }
    break;

  case 60:
/* Line 1792 of yacc.c  */
#line 575 "../../lib/common/htmlparse.y"
    { setCell((htmlvsp[(1) - (2)].cell),(htmlvsp[(2) - (2)].img),HTML_IMAGE); }
    break;

  case 61:
/* Line 1792 of yacc.c  */
#line 575 "../../lib/common/htmlparse.y"
    { (htmlval.cell) = (htmlvsp[(1) - (4)].cell); }
    break;

  case 62:
/* Line 1792 of yacc.c  */
#line 576 "../../lib/common/htmlparse.y"
    { setCell((htmlvsp[(1) - (1)].cell),mkText(),HTML_TEXT); }
    break;

  case 63:
/* Line 1792 of yacc.c  */
#line 576 "../../lib/common/htmlparse.y"
    { (htmlval.cell) = (htmlvsp[(1) - (3)].cell); }
    break;

  case 64:
/* Line 1792 of yacc.c  */
#line 579 "../../lib/common/htmlparse.y"
    { (htmlval.img) = (htmlvsp[(1) - (2)].img); }
    break;

  case 65:
/* Line 1792 of yacc.c  */
#line 580 "../../lib/common/htmlparse.y"
    { (htmlval.img) = (htmlvsp[(1) - (1)].img); }
    break;


/* Line 1792 of yacc.c  */
#line 2277 "y.tab.c"
      default: break;
    }
  /* User semantic actions sometimes alter htmlchar, and that requires
     that htmltoken be updated with the new translation.  We take the
     approach of translating immediately before every use of htmltoken.
     One alternative is translating here after every semantic action,
     but that translation would be missed if the semantic action invokes
     YYABORT, YYACCEPT, or YYERROR immediately after altering htmlchar or
     if it invokes YYBACKUP.  In the case of YYABORT or YYACCEPT, an
     incorrect destructor might then be invoked immediately.  In the
     case of YYERROR or YYBACKUP, subsequent parser actions might lead
     to an incorrect destructor call or verbose syntax error message
     before the lookahead is translated.  */
  YY_SYMBOL_PRINT ("-> $$ =", htmlr1[htmln], &htmlval, &htmlloc);

  YYPOPSTACK (htmllen);
  htmllen = 0;
  YY_STACK_PRINT (htmlss, htmlssp);

  *++htmlvsp = htmlval;

  /* Now `shift' the result of the reduction.  Determine what state
     that goes to, based on the state we popped back to and the rule
     number reduced by.  */

  htmln = htmlr1[htmln];

  htmlstate = htmlpgoto[htmln - YYNTOKENS] + *htmlssp;
  if (0 <= htmlstate && htmlstate <= YYLAST && htmlcheck[htmlstate] == *htmlssp)
    htmlstate = htmltable[htmlstate];
  else
    htmlstate = htmldefgoto[htmln - YYNTOKENS];

  goto htmlnewstate;


/*------------------------------------.
| htmlerrlab -- here on detecting error |
`------------------------------------*/
htmlerrlab:
  /* Make sure we have latest lookahead translation.  See comments at
     user semantic actions for why this is necessary.  */
  htmltoken = htmlchar == YYEMPTY ? YYEMPTY : YYTRANSLATE (htmlchar);

  /* If not already recovering from an error, report this error.  */
  if (!htmlerrstatus)
    {
      ++htmlnerrs;
#if ! YYERROR_VERBOSE
      htmlerror (YY_("syntax error"));
#else
# define YYSYNTAX_ERROR htmlsyntax_error (&htmlmsg_alloc, &htmlmsg, \
                                        htmlssp, htmltoken)
      {
        char const *htmlmsgp = YY_("syntax error");
        int htmlsyntax_error_status;
        htmlsyntax_error_status = YYSYNTAX_ERROR;
        if (htmlsyntax_error_status == 0)
          htmlmsgp = htmlmsg;
        else if (htmlsyntax_error_status == 1)
          {
            if (htmlmsg != htmlmsgbuf)
              YYSTACK_FREE (htmlmsg);
            htmlmsg = (char *) YYSTACK_ALLOC (htmlmsg_alloc);
            if (!htmlmsg)
              {
                htmlmsg = htmlmsgbuf;
                htmlmsg_alloc = sizeof htmlmsgbuf;
                htmlsyntax_error_status = 2;
              }
            else
              {
                htmlsyntax_error_status = YYSYNTAX_ERROR;
                htmlmsgp = htmlmsg;
              }
          }
        htmlerror (htmlmsgp);
        if (htmlsyntax_error_status == 2)
          goto htmlexhaustedlab;
      }
# undef YYSYNTAX_ERROR
#endif
    }



  if (htmlerrstatus == 3)
    {
      /* If just tried and failed to reuse lookahead token after an
	 error, discard it.  */

      if (htmlchar <= YYEOF)
	{
	  /* Return failure if at end of input.  */
	  if (htmlchar == YYEOF)
	    YYABORT;
	}
      else
	{
	  htmldestruct ("Error: discarding",
		      htmltoken, &htmllval);
	  htmlchar = YYEMPTY;
	}
    }

  /* Else will try to reuse lookahead token after shifting the error
     token.  */
  goto htmlerrlab1;


/*---------------------------------------------------.
| htmlerrorlab -- error raised explicitly by YYERROR.  |
`---------------------------------------------------*/
htmlerrorlab:

  /* Pacify compilers like GCC when the user code never invokes
     YYERROR and the label htmlerrorlab therefore never appears in user
     code.  */
  if (/*CONSTCOND*/ 0)
     goto htmlerrorlab;

  /* Do not reclaim the symbols of the rule which action triggered
     this YYERROR.  */
  YYPOPSTACK (htmllen);
  htmllen = 0;
  YY_STACK_PRINT (htmlss, htmlssp);
  htmlstate = *htmlssp;
  goto htmlerrlab1;


/*-------------------------------------------------------------.
| htmlerrlab1 -- common code for both syntax error and YYERROR.  |
`-------------------------------------------------------------*/
htmlerrlab1:
  htmlerrstatus = 3;	/* Each real token shifted decrements this.  */

  for (;;)
    {
      htmln = htmlpact[htmlstate];
      if (!htmlpact_value_is_default (htmln))
	{
	  htmln += YYTERROR;
	  if (0 <= htmln && htmln <= YYLAST && htmlcheck[htmln] == YYTERROR)
	    {
	      htmln = htmltable[htmln];
	      if (0 < htmln)
		break;
	    }
	}

      /* Pop the current state because it cannot handle the error token.  */
      if (htmlssp == htmlss)
	YYABORT;


      htmldestruct ("Error: popping",
		  htmlstos[htmlstate], htmlvsp);
      YYPOPSTACK (1);
      htmlstate = *htmlssp;
      YY_STACK_PRINT (htmlss, htmlssp);
    }

  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  *++htmlvsp = htmllval;
  YY_IGNORE_MAYBE_UNINITIALIZED_END


  /* Shift the error token.  */
  YY_SYMBOL_PRINT ("Shifting", htmlstos[htmln], htmlvsp, htmllsp);

  htmlstate = htmln;
  goto htmlnewstate;


/*-------------------------------------.
| htmlacceptlab -- YYACCEPT comes here.  |
`-------------------------------------*/
htmlacceptlab:
  htmlresult = 0;
  goto htmlreturn;

/*-----------------------------------.
| htmlabortlab -- YYABORT comes here.  |
`-----------------------------------*/
htmlabortlab:
  htmlresult = 1;
  goto htmlreturn;

#if !defined htmloverflow || YYERROR_VERBOSE
/*-------------------------------------------------.
| htmlexhaustedlab -- memory exhaustion comes here.  |
`-------------------------------------------------*/
htmlexhaustedlab:
  htmlerror (YY_("memory exhausted"));
  htmlresult = 2;
  /* Fall through.  */
#endif

htmlreturn:
  if (htmlchar != YYEMPTY)
    {
      /* Make sure we have latest lookahead translation.  See comments at
         user semantic actions for why this is necessary.  */
      htmltoken = YYTRANSLATE (htmlchar);
      htmldestruct ("Cleanup: discarding lookahead",
                  htmltoken, &htmllval);
    }
  /* Do not reclaim the symbols of the rule which action triggered
     this YYABORT or YYACCEPT.  */
  YYPOPSTACK (htmllen);
  YY_STACK_PRINT (htmlss, htmlssp);
  while (htmlssp != htmlss)
    {
      htmldestruct ("Cleanup: popping",
		  htmlstos[*htmlssp], htmlvsp);
      YYPOPSTACK (1);
    }
#ifndef htmloverflow
  if (htmlss != htmlssa)
    YYSTACK_FREE (htmlss);
#endif
#if YYERROR_VERBOSE
  if (htmlmsg != htmlmsgbuf)
    YYSTACK_FREE (htmlmsg);
#endif
  /* Make sure YYID is used.  */
  return YYID (htmlresult);
}


/* Line 2055 of yacc.c  */
#line 592 "../../lib/common/htmlparse.y"


/* parseHTML:
 * Return parsed label or NULL if failure.
 * Set warn to 0 on success; 1 for warning message; 2 if no expat.
 */
htmllabel_t*
parseHTML (char* txt, int* warn, htmlenv_t *env)
{
  unsigned char buf[SMALLBUF];
  agxbuf        str;
  htmllabel_t*  l;
  sfont_t       dfltf;

  dfltf.cfont = NULL;
  dfltf.pfont = NULL;
  HTMLstate.fontstack = &dfltf;
  HTMLstate.tblstack = 0;
  HTMLstate.lbl = 0;
  HTMLstate.gvc = GD_gvc(env->g);
  HTMLstate.fitemList = dtopen(&fstrDisc, Dtqueue);
  HTMLstate.fspanList = dtopen(&fspanDisc, Dtqueue);

  agxbinit (&str, SMALLBUF, buf);
  HTMLstate.str = &str;
  
  if (initHTMLlexer (txt, &str, env)) {/* failed: no libexpat - give up */
    *warn = 2;
    l = NULL;
  }
  else {
    htmlparse();
    *warn = clearHTMLlexer ();
    l = HTMLstate.lbl;
  }

  dtclose (HTMLstate.fitemList);
  dtclose (HTMLstate.fspanList);
  
  HTMLstate.fitemList = NULL;
  HTMLstate.fspanList = NULL;
  HTMLstate.fontstack = NULL;
  
  agxbfree (&str);

  return l;
}

