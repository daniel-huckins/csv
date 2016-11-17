#include <gtk/gtk.h>
#include <webkit2/webkit2.h>

static gboolean closeWebViewCb(WebKitWebView *webView, GtkWidget *window) {
  gtk_widget_destroy(window);
  return TRUE;
}

static void destroyWindowCb(GtkWidget *widget, GtkWidget *window) {
  gtk_main_quit();
}

static WebKitWebView *newWebView() {
  return WEBKIT_WEB_VIEW(webkit_web_view_new());
}

static GtkWidget *newGtkWindow(WebKitWebView *webView, int width, int height) {
  GtkWidget *main_window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
  gtk_window_set_default_size(GTK_WINDOW(main_window), (gint)width,
                              (gint)height);
  gtk_container_add(GTK_CONTAINER(main_window), GTK_WIDGET(webView));

  // callbacks to close the window when user closes it
  g_signal_connect(main_window, "destroy", G_CALLBACK(destroyWindowCb), NULL);
  g_signal_connect(webView, "close", G_CALLBACK(closeWebViewCb), main_window);
  return main_window;
}

static void showWindow(WebKitWebView *webView, GtkWidget *widget) {
  gtk_widget_grab_focus(GTK_WIDGET(webView));
  gtk_widget_show_all(widget);
}

static void init_gtk(int argc) { gtk_init(&argc, NULL); }
