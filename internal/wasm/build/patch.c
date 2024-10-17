#include <stdio.h>
#include <stdlib.h>
#include "gvplugin.h"
#include "gvplugin_render.h"

static const char *tmpfilename = "tmpfile";

FILE *tmpfile(void)
{
  return fopen(tmpfilename, "w+");
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
