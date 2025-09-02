"use server";

import { Prisma, User } from "@/generated/prisma";
import { setUserToken } from "@/lib/auth";
import { prisma } from "@/lib/prisma";
import { parseWithZod } from "@conform-to/zod";
import { redirect } from "next/navigation";
import { loginRequestSchema, registerRequestSchema } from "./schema";

export async function login(_: unknown, formData: FormData) {
  const submission = parseWithZod(formData, { schema: loginRequestSchema });

  if (submission.status !== "success") {
    return submission.reply();
  }

  const { username } = submission.value;
  const user = await prisma.user.findFirst({
    where: { name: username },
  });

  if (!user) {
    return submission.reply({
      fieldErrors: { username: ["User not found"] },
    });
  }

  try {
    await setUserToken(user.id);
  } catch (error) {
    console.error(error);
    return submission.reply({
      formErrors: ["Failed to set user token"],
    });
  }

  redirect("/");
}

export async function register(_: unknown, formData: FormData) {
  const submission = parseWithZod(formData, { schema: registerRequestSchema });

  if (submission.status !== "success") {
    return submission.reply();
  }

  let user: User;
  try {
    const { username } = submission.value;
    user = await prisma.user.create({
      data: { name: username },
    });
  } catch (error) {
    if (
      error instanceof Prisma.PrismaClientKnownRequestError &&
      error.code === "P2002"
    ) {
      return submission.reply({
        fieldErrors: { username: ["Username already taken"] },
      });
    }

    console.error("failed to create user:", error);
    return submission.reply({
      formErrors: ["Failed to create user"],
    });
  }

  try {
    await setUserToken(user.id);
  } catch (error) {
    console.error("failed to set user token:", error);
    return submission.reply({
      formErrors: ["Failed to set user token"],
    });
  }

  redirect("/");
}
