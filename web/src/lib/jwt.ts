import dayjs from "dayjs";

interface JwtDecodeOptions {
  header?: boolean;
}

interface JwtHeader {
  typ?: string;
  alg?: string;
  kid?: string;
}

interface JwtPayload {
  iss?: string;
  sub?: string;
  aud?: string[] | string;
  exp?: number;
  nbf?: number;
  iat?: number;
  jti?: string;
}

class InvalidTokenError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "InvalidTokenError";
  }
}

function base64UrlDecode(input: string): string {
  const output = input.replace(/-/g, "+").replace(/_/g, "/");
  const padded = output.padEnd(output.length + (4 - (output.length % 4)) % 4, "=");
  try {
    return decodeURIComponent(
      atob(padded)
        .split("")
        .map((char) => `%${char.charCodeAt(0).toString(16).padStart(2, "0").toUpperCase()}`)
        .join("")
    );
  } catch {
    throw new Error("Invalid base64 string");
  }
}

function jwtDecode<T = JwtHeader | JwtPayload>(
  token: string,
  options?: JwtDecodeOptions
): T {
  const isHeader = options?.header === true;
  const part = token.split(".")[isHeader ? 0 : 1];

  if (!part) {
    throw new InvalidTokenError(`Invalid token: missing part #${isHeader ? 1 : 2}`);
  }

  try {
    const decoded = base64UrlDecode(part);
    return JSON.parse(decoded) as T;
  } catch (e) {
    if (e instanceof SyntaxError) {
      throw new InvalidTokenError(`Invalid token: invalid JSON for part #${isHeader ? 1 : 2}`);
    } else if (e instanceof Error) {
      throw new InvalidTokenError(`Invalid token: ${e.message}`);
    } else {
      throw new InvalidTokenError(`Invalid token: unknown error`);
    }
  }
}

export function isJWTExpired(token: string): boolean {
  if (!token) return true;
  try {
    const t = jwtDecode<JwtPayload>(token);
    return !t.exp || dayjs.unix(t.exp).diff(dayjs()) < 1;
  } catch {
    return true;
  }
}