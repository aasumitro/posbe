const ALLOWED_REDIRECT_PATHS = [
  "/",
  "/orders",
  "/stores",
];

export function getSafeRedirect(search: string): string {
  const searchParams = new URLSearchParams(search);
  const candidate = searchParams.get("redirect");
  // Allow only relative, internal paths beginning with a single '/'
  if (candidate && ALLOWED_REDIRECT_PATHS.includes(candidate)) {
    return candidate;
  }
  return "/";
}
