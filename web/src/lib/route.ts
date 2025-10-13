export function getSafeRedirect(search: string): string {
  const searchParams = new URLSearchParams(search);
  const candidate = searchParams.get("redirect");
  // Allow only relative, internal paths beginning with a single '/'
  if (candidate && /^\/(?!\/)/.test(candidate)) {
    return candidate;
  }
  return "/";
}
