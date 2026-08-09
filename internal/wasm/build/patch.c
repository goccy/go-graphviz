#include <stdio.h>
#include <stdlib.h>
#include "gvplugin.h"
#include "gvplugin_render.h"

static const char *tmpfilename = "tmpfile";

FILE *tmpfile(void)
{
  return fopen(tmpfilename, "w+");
}

/* flockfile/funlockfile are referenced by the GV_INFO/GV_DEBUG macros in
 * lib/util/debug.h (added in graphviz 13.0.0) and by lib/util/lockfile.h
 * (used by common/emit.c emit_once since graphviz 14.0.0). They are not
 * provided by wasi-libc's stdio.h for user code: the declarations there are
 * guarded by __wasilibc_unmodified_upstream/_REENTRANT, so graphviz callers
 * compile them as implicit-int (i32)->i32. The stubs must therefore return
 * int, matching the call sites, or wasm-ld renames the int variant to
 * "signature_mismatch:flockfile" with an unreachable (trap) body and the
 * emit_once path traps at runtime. The wasm build is single-threaded and
 * Verbose is never set, so no-op stubs returning 0 are correct. */
int flockfile(FILE *file)
{
  (void)file;
  return 0;
}

int funlockfile(FILE *file)
{
  (void)file;
  return 0;
}

extern gvplugin_library_t gvplugin_dot_layout_LTX_library;
extern gvplugin_library_t gvplugin_neato_layout_LTX_library;
extern gvplugin_library_t gvplugin_core_LTX_library;

lt_symlist_t lt_preloaded_symbols[] = {
    { "gvplugin_dot_layout_LTX_library", (void *)(&gvplugin_dot_layout_LTX_library) },
    { "gvplugin_neato_layout_LTX_library", (void*)(&gvplugin_neato_layout_LTX_library) },
    { "gvplugin_core_LTX_library", (void*)(&gvplugin_core_LTX_library) },
};

static gvplugin_api_t api_zero = {(api_t)0, 0};
static gvplugin_installed_t installed_zero = {0, NULL, 0, NULL, NULL};
static lt_symlist_t symlist_zero = {NULL, NULL};

typedef struct { int len; void *data; } GoSlice;

void wasm_bridge_PluginAPI_zero(void **ret) {
  *ret = &api_zero;
}

void wasm_bridge_PluginInstalled_zero(void **ret) {
  *ret = &installed_zero;
}

void wasm_bridge_SymList_zero(void **ret) {
  *ret = &symlist_zero;
}

void wasm_bridge_SymList_default(GoSlice **ret) {
  GoSlice *v = (GoSlice *)malloc(sizeof(GoSlice));
  size_t len = sizeof(lt_preloaded_symbols) / sizeof(lt_preloaded_symbols[0]);
  v->len = len;
  void **data = malloc(8 * len);
  v->data = data;
  for (int i = 0; i < len; i++) {
    lt_symlist_t *elem = (lt_symlist_t *)malloc(sizeof(lt_symlist_t));
    memcpy(elem, &lt_preloaded_symbols[i], sizeof(lt_symlist_t));
    *data = elem;
    data += 2;
  }
  *ret = v;
}

int main() { return 0; }
