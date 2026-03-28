const ALLOWED_REDIRECT_PATHS = [
  "/",
  "/orders",
  "/orders/floors",
  "/orders/menus",
  "/stores",
  "/stores/floors",
  "/stores/attributes",
  "/stores/catalogs",
  "/stores/teams",
  "/stores/customers",
  "/stores/transactions",
];

export function getSafeRedirect(search: string): string {
  const searchParams = new URLSearchParams(search);
  const candidate = searchParams.get("redirect");
  // Allow only relative, internal paths beginning with a single '/'
  if (candidate && ALLOWED_REDIRECT_PATHS.includes(candidate)) return candidate;
  return "/";
}
