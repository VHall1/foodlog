import { Prisma, User } from "@/generated/prisma";
import { setUserToken } from "@/lib/auth";
import { prisma } from "@/lib/prisma";
import { redirect } from "next/navigation";
import { registerRequestSchema } from "./dto";

export async function POST(request: Request) {
  const payload = await request.json();
  const { success, data } = registerRequestSchema.safeParse(payload);

  if (!success) {
    return new Response("Invalid request", { status: 400 });
  }

  let user: User;
  try {
    const { username } = data;
    user = await prisma.user.create({
      data: { name: username },
    });
  } catch (error) {
    if (
      error instanceof Prisma.PrismaClientKnownRequestError &&
      error.code === "P2002"
    ) {
      return new Response("User already exists", { status: 422 });
    }

    console.error("failed to create user:", error);
    return new Response("Internal server error", { status: 500 });
  }

  await setUserToken(user.id);

  // TODO: implement "redirect back" param?
  return redirect("/");
}
