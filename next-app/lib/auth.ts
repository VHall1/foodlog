import * as jwt from "jsonwebtoken";
import { cookies } from "next/headers";
import invariant from "tiny-invariant";

import "server-only";

const JWT_SECRET = process.env.JWT_SECRET;
invariant(JWT_SECRET, "JWT_SECRET is not defined");

// 7 days
const COOKIE_TTL = 7 * 24 * 60 * 60;
const COOKIE_NAME = "foodlog__session";

export const setUserToken = async (userId: string) => {
  const token = jwt.sign({ userId }, JWT_SECRET, {
    expiresIn: COOKIE_TTL,
  });

  const cookieStore = await cookies();
  cookieStore.set(COOKIE_NAME, token, {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    maxAge: COOKIE_TTL,
    path: "/",
  });
};

export const getUserId = async () => {
  const cookieStore = await cookies();
  const token = cookieStore.get(COOKIE_NAME);

  if (!token) return null;

  try {
    const payload = jwt.verify(token.value, JWT_SECRET);
    if (typeof payload !== "string" && typeof payload?.userId === "string") {
      return payload.userId;
    }
    return null;
  } catch (error) {
    console.error("failed to verify token:", error);
    return null;
  }
};
