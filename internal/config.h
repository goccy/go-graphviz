/* config.h.  Generated from config.h.in by configure.  */
/* config.h.in.  Generated from configure.ac by autoheader.  */

/* Command to open a browser on a URL */
#define BROWSER "open"

/* Define for any Darwin-based OS. */
#define DARWIN 1

/* Define for Darwin-style shared library names. */
#define DARWIN_DYLIB ""

/* Default DPI. */
#define DEFAULT_DPI 96

/* Path to TrueType fonts. */
#define DEFAULT_FONTPATH "~/Library/Fonts:/Library/Fonts:/Network/Library/Fonts:/System/Library/Fonts"

/* Define if you want DIGCOLA */
#define DIGCOLA 1

/* Define if you want on-demand plugin loading */
//#define ENABLE_LTDL 1

/* Define for DLLs on Windows. */
/* #undef GVDLL */

/* Filename for plugin configuration file. */
#define GVPLUGIN_CONFIG_FILE "config6"

/* Compatibility version number for plugins. */
#define GVPLUGIN_VERSION 6

/* Define if you have the ann library */
/* #undef HAVE_ANN */

/* Define to 1 if you have the `argz_add' function. */
/* #undef HAVE_ARGZ_ADD */

/* Define to 1 if you have the `argz_append' function. */
/* #undef HAVE_ARGZ_APPEND */

/* Define to 1 if you have the `argz_count' function. */
/* #undef HAVE_ARGZ_COUNT */

/* Define to 1 if you have the `argz_create_sep' function. */
/* #undef HAVE_ARGZ_CREATE_SEP */

/* Define to 1 if you have the <argz.h> header file. */
/* #undef HAVE_ARGZ_H */

/* Define to 1 if you have the `argz_insert' function. */
/* #undef HAVE_ARGZ_INSERT */

/* Define to 1 if you have the `argz_next' function. */
/* #undef HAVE_ARGZ_NEXT */

/* Define to 1 if you have the `argz_stringify' function. */
/* #undef HAVE_ARGZ_STRINGIFY */

/* Define to 1 if you have the `cbrt' function. */
#define HAVE_CBRT 1

/* Define to 1 if you have the `closedir' function. */
#define HAVE_CLOSEDIR 1

/* Criterion unit testing framework is installed and available for use. */
/* #undef HAVE_CRITERION */

/* Define to 1 if you have the <crt_externs.h> header file. */
#define HAVE_CRT_EXTERNS_H 1

/* Define to 1 if you have the declaration of `cygwin_conv_path', and to 0 if
   you don't. */
/* #undef HAVE_DECL_CYGWIN_CONV_PATH */

/* Define to 1 if you have the `deflateBound' function. */
#define HAVE_DEFLATEBOUND 1

/* Define if you have the DevIL library */
/* #undef HAVE_DEVIL */

/* Define to 1 if you have the <dirent.h> header file. */
#define HAVE_DIRENT_H 1

/* Define if you have the GNU dld library. */
/* #undef HAVE_DLD */

/* Define to 1 if you have the <dld.h> header file. */
/* #undef HAVE_DLD_H */

/* Define to 1 if you have the `dlerror' function. */
#define HAVE_DLERROR 1

/* Define to 1 if you have the <dlfcn.h> header file. */
#define HAVE_DLFCN_H 1

/* Define to 1 if you have the <dl.h> header file. */
/* #undef HAVE_DL_H */

#ifndef WIN32
/* Define to 1 if you have the `drand48' function. */
#define HAVE_DRAND48 1
#endif

/* Define if you have the _dyld_func_lookup function. */
/* #undef HAVE_DYLD */

/* Define if errno externs are declared */
#define HAVE_ERRNO_DECL 1

/* Define to 1 if you have the <errno.h> header file. */
#define HAVE_ERRNO_H 1

/* Define to 1 if the system has the type `error_t'. */
/* #undef HAVE_ERROR_T */

/* Define if you have the expat library */
#define HAVE_EXPAT 1

/* Define to 1 if you have the <expat.h> header file. */
#define HAVE_EXPAT_H 1

/* Define to 1 if you have the <fcntl.h> header file. */
#define HAVE_FCNTL_H 1

/* Define if FILE structure provides _cnt */
/* #undef HAVE_FILE_CNT */

/* Define if FILE structure provides _IO_read_end */
/* #undef HAVE_FILE_IO_READ_END */

/* Define if FILE structure provides _next */
/* #undef HAVE_FILE_NEXT */

/* Define if FILE structure provides _r */
#define HAVE_FILE_R 1

/* Define if you have the fontconfig library */
#define HAVE_FONTCONFIG 1

/* Define if you have the freetype2 library */
#define HAVE_FREETYPE2 1

/* Define if you have the GDI+ framework for Windows */
/* #undef HAVE_GDIPLUS */

/* Define if you have the gdk library */
/* #undef HAVE_GDK */

/* Define if you have the gdk_pixbuf library */
/* #undef HAVE_GDK_PIXBUF */

/* Define if you have the gdlib library */
#define HAVE_GDLIB 1

/* Define if the GD library supports FONTCONFIG */
#define HAVE_GD_FONTCONFIG 1

/* Define if the GD library supports FREETYPE */
#define HAVE_GD_FREETYPE 1

/* Define if the GD library supports GIF */
#define HAVE_GD_GIF 1

/* Define if the GD library supports GIFANIM */
#define HAVE_GD_GIFANIM 1

/* Define if the GD library supports JPEG */
#define HAVE_GD_JPEG 1

/* Define if the GD library supports OPENPOLYGON */
#define HAVE_GD_OPENPOLYGON 1

/* Define if the GD library supports PNG */
#define HAVE_GD_PNG 1

/* Define if the GD library supports XPM */
#define HAVE_GD_XPM 1

/* Define to 1 if you have the `getenv' function. */
#define HAVE_GETENV 1

/* Define if you have the glade library */
/* #undef HAVE_GLADE */

/* Define if you have the glitz library */
/* #undef HAVE_GLITZ */

/* Define if you have the GLUT library */
/* #undef HAVE_GLUT */

/* Define if you have the gs library */
/* #undef HAVE_GS */

/* Define if you have the gtk library */
/* #undef HAVE_GTK */

/* Define if you have the gtkgl library */
/* #undef HAVE_GTKGL */

/* Define if you have the gtkglext library */
/* #undef HAVE_GTKGLEXT */

/* Define if you have the gts library */
#define HAVE_GTS 1

/* Define to 1 if you have the `g_object_unref' function. */
/* #undef HAVE_G_OBJECT_UNREF */

/* Define to 1 if you have the `g_type_init' function. */
/* #undef HAVE_G_TYPE_INIT */

/* Define to 1 if you have the <IL/il.h> header file. */
/* #undef HAVE_IL_IL_H */

/* Define if intptr_t is declared */
#define HAVE_INTPTR_T 1

/* Define to 1 if you have the <inttypes.h> header file. */
#define HAVE_INTTYPES_H 1

/* Define if you have the lasi library */
/* #undef HAVE_LASI */

/* Define if you have the libdl library or equivalent. */
#define HAVE_LIBDL 1

/* Define if libdlloader will be built on this platform */
#define HAVE_LIBDLLOADER 1

/* Define if you have the GD library */
#define HAVE_LIBGD 1

/* Define if the LIBGEN library has the basename feature */
/* #undef HAVE_LIBGEN */

/* Define to 1 if you have the `ltdl' library (-lltdl). */
//#define HAVE_LIBLTDL 1

/* Define if you have the XPM library */
/* #undef HAVE_LIBXPMFORLEFTY */

/* Define if you have the Z library */
//#define HAVE_LIBZ 1

/* Define to 1 if you have the <limits.h> header file. */
#define HAVE_LIMITS_H 1

/* Define to 1 if you have the `log2' function. */
#define HAVE_LOG2 1

/* Define to 1 if you have the `lrand48' function. */
#define HAVE_LRAND48 1

/* Define this if a modern libltdl is already installed */
//#define HAVE_LTDL 1

/* Define to 1 if you have the <mach-o/dyld.h> header file. */
#define HAVE_MACH_O_DYLD_H 1

/* Define to 1 if you have the <malloc.h> header file. */
/* #undef HAVE_MALLOC_H */

/* Define to 1 if you have the <memory.h> header file. */
#define HAVE_MEMORY_H 1

/* Define if you have the ming library for SWF support */
/* #undef HAVE_MING */

/* Define to 1 if you have the <ndir.h> header file, and it defines `DIR'. */
/* #undef HAVE_NDIR_H */

/* Define to 1 if you have the `opendir' function. */
#define HAVE_OPENDIR 1

/* Define if you have the pangocairo library */
/* #undef HAVE_PANGOCAIRO */

/* Define to 1 if you have the `pango_fc_font_lock_face' function. */
/* #undef HAVE_PANGO_FC_FONT_LOCK_FACE */

/* Define to 1 if you have the `pango_fc_font_unlock_face' function. */
/* #undef HAVE_PANGO_FC_FONT_UNLOCK_FACE */

/* Define to 1 if you have the `pango_font_map_create_context' function. */
/* #undef HAVE_PANGO_FONT_MAP_CREATE_CONTEXT */

/* Define if you have the poppler library */
/* #undef HAVE_POPPLER */

/* Define if libtool can extract symbol lists from object files. */
#define HAVE_PRELOADED_SYMBOLS 1

/* Define if you have the Quartz framework for Mac OS X */
/* #undef HAVE_QUARTZ */

/* Define to 1 if you have the `readdir' function. */
#define HAVE_READDIR 1

/* Define if you have the rsvg library */
/* #undef HAVE_RSVG */

/* Define to 1 if you have the <search.h> header file. */
#define HAVE_SEARCH_H 1

#ifndef WIN32
/* Define to 1 if you have the `setenv' function. */
#define HAVE_SETENV 1
#endif

/* Define to 1 if you have the `setmode' function. */
#define HAVE_SETMODE 1

/* Define if you have the shl_load function. */
/* #undef HAVE_SHL_LOAD */

/* Define to 1 if you have the `sincos' function. */
/* #undef HAVE_SINCOS */

/* Define to 1 if you have the `srand48' function. */
#ifndef WIN32
#define HAVE_SRAND48 1
#endif

/* Define to 1 if stdbool.h conforms to C99. */
#define HAVE_STDBOOL_H 1

/* Define to 1 if you have the <stdint.h> header file. */
#define HAVE_STDINT_H 1

/* Define to 1 if you have the <stdlib.h> header file. */
#define HAVE_STDLIB_H 1

/* Define to 1 if you have the `strcasecmp' function. */
#define HAVE_STRCASECMP 1

/* Define to 1 if you have the `strcasestr' function. */
#define HAVE_STRCASESTR 1

/* Define to 1 if you have the `strerror' function. */
#define HAVE_STRERROR 1

/* Define to 1 if you have the <strings.h> header file. */
#define HAVE_STRINGS_H 1

/* Define to 1 if you have the <string.h> header file. */
#define HAVE_STRING_H 1

/* Define to 1 if you have the `strlcat' function. */
#define HAVE_STRLCAT 1

/* Define to 1 if you have the `strlcpy' function. */
#define HAVE_STRLCPY 1

/* Define to 1 if you have the `strncasecmp' function. */
#define HAVE_STRNCASECMP 1

/* Have librsvg >= 2.36 */
/* #undef HAVE_SVG_2_36 */

/* Define to 1 if you have the <sys/dir.h> header file, and it defines `DIR'.
   */
/* #undef HAVE_SYS_DIR_H */

/* Define to 1 if you have the <sys/dl.h> header file. */
/* #undef HAVE_SYS_DL_H */

/* Define to 1 if you have the <sys/inotify.h> header file. */
/* #undef HAVE_SYS_INOTIFY_H */

/* Define to 1 if you have the <sys/ioctl.h> header file. */
#define HAVE_SYS_IOCTL_H 1

/* Define to 1 if you have the <sys/mman.h> header file. */
#ifndef WIN32
#define HAVE_SYS_MMAN_H 1
#endif

/* Define to 1 if you have the <sys/ndir.h> header file, and it defines `DIR'.
   */
/* #undef HAVE_SYS_NDIR_H */

/* Define to 1 if you have the <sys/select.h> header file. */
#define HAVE_SYS_SELECT_H 1

/* Define to 1 if you have the <sys/stat.h> header file. */
#define HAVE_SYS_STAT_H 1

/* Define to 1 if you have the <sys/time.h> header file. */
#define HAVE_SYS_TIME_H 1

/* Define to 1 if you have the <sys/types.h> header file. */
#define HAVE_SYS_TYPES_H 1

/* Define if you have the tcl library */
/* #undef HAVE_TCL */

/* Define to 1 if you have the <termios.h> header file. */
#define HAVE_TERMIOS_H 1

/* Define if triangle.[ch] are available. */
/* #undef HAVE_TRIANGLE */

/* Define to 1 if you have the <unistd.h> header file. */
#define HAVE_UNISTD_H 1

/* Define to 1 if you have the <values.h> header file. */
/* #undef HAVE_VALUES_H */

/* Define if you have the visio library */
/* #undef HAVE_VISIO */

/* Define to 1 if you have the `vsnprintf' function. */
#define HAVE_VSNPRINTF 1

/* Define if you have the webp library */
/* #undef HAVE_WEBP */

/* This value is set to 1 to indicate that the system argz facility works */
/* #undef HAVE_WORKING_ARGZ */

/* Define to 1 if you have the <X11/Intrinsic.h> header file. */
/* #undef HAVE_X11_INTRINSIC_H */

/* Define to 1 if you have the <X11/Xaw/Text.h> header file. */
/* #undef HAVE_X11_XAW_TEXT_H */

/* Define to 1 if the system has the type `_Bool'. */
#define HAVE__BOOL 1

/* Define to 1 if you have the `_NSGetEnviron' function. */
#define HAVE__NSGETENVIRON 1

/* Define if you want IPSEPCOLA */
/* #undef IPSEPCOLA */

/* Define if the OS needs help to load dependent libraries for dlopen(). */
/* #undef LTDL_DLOPEN_DEPLIBS */

/* Define to the system default library search path. */
#define LT_DLSEARCH_PATH "/usr/local/lib:/lib:/usr/lib"

/* The archive extension */
#define LT_LIBEXT "a"

/* The archive prefix */
#define LT_LIBPREFIX "lib"

/* Define to the extension used for runtime loadable modules, say, ".so". */
#define LT_MODULE_EXT ".so"

/* Define to the name of the environment variable that determines the run-time
   module search path. */
#define LT_MODULE_PATH_VAR "DYLD_LIBRARY_PATH"

/* Define to the sub-directory in which libtool stores uninstalled libraries.
   */
#define LT_OBJDIR ".libs/"

/* Define to the shared library suffix, say, ".dylib". */
#define LT_SHARED_EXT ".dylib"

/* Define if dlsym() requires a leading underscore in symbol names. */
/* #undef NEED_USCORE */

/* Define if no fpu error exception handling is required. */
/* #undef NO_FPERR */

/* Define to 1 if your C compiler doesn't accept -c and -o together. */
/* #undef NO_MINUS_C_MINUS_O */

/* Postscript fontnames. */
#define NO_POSTSCRIPT_ALIAS 1

/* Define if you want ORTHO */
#define ORTHO 1

/* Define to the address where bug reports for this package should be sent. */
#define PACKAGE_BUGREPORT "http://www.graphviz.org/"

/* Define to the full name of this package. */
#define PACKAGE_NAME "graphviz"

/* Define to the full name and version of this package. */
#define PACKAGE_STRING "graphviz 2.40.1"

/* Define to the one symbol short name of this package. */
#define PACKAGE_TARNAME "graphviz"

/* Define to the home page for this package. */
#define PACKAGE_URL ""

/* Define to the version of this package. */
#define PACKAGE_VERSION "2.40.1"

/* Path separator character. */
#define PATHSEPARATOR ":"

/* Define if you want SFDP */
#define SFDP 1

/* The size of `int', as computed by sizeof. */
#define SIZEOF_INT 4

/* The size of `long long', as computed by sizeof. */
#define SIZEOF_LONG_LONG 8

/* Define if you want SMYRNA */
/* #undef SMYRNA */

/* Define to 1 if you have the ANSI C header files. */
#define STDC_HEADERS 1

/* Define to 1 if you can safely include both <sys/time.h> and <time.h>. */
#define TIME_WITH_SYS_TIME 1

/* Historical artifact - always true */
#define WITH_CGRAPH 1

/* Define to 1 if the X Window System is missing or not being used. */
#define X_DISPLAY_MISSING 1

/* Define to 1 if `lex' declares `yytext' as a `char *' by default, not a
   `char[]'. */
#define YYTEXT_POINTER 1

/* Define so that glibc/gnulib argp.h does not typedef error_t. */
#define __error_t_defined 1

/* Define to a type to use for `error_t' if it is not otherwise available. */
#define error_t int

#define GVC_EXPORTS 1
#define PATHPLAN_EXPORTS 1

#define PRLOADED_SYMBOL_N 5

/* Define to `int' if <sys/types.h> doesn't define. */
/* #undef gid_t */

/* Define to `__inline__' or `__inline' if that's what the C compiler
   calls it, or to nothing if 'inline' is not supported under any name.  */
#ifndef __cplusplus
/* #undef inline */
#endif

/* Define to `int' if <sys/types.h> does not define. */
/* #undef pid_t */

/* Define to `int' if <sys/types.h> does not define. */
/* #undef ssize_t */

/* Define to `int' if <sys/types.h> doesn't define. */
/* #undef uid_t */
