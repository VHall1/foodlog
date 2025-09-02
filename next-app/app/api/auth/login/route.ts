import { setUserToken } from "@/lib/auth";
import { prisma } from "@/lib/prisma";
import { redirect } from "next/navigation";
import { loginRequestSchema } from "./dto";

export async function POST(request: Request) {
  const payload = await request.json();
  const { success, data } = loginRequestSchema.safeParse(payload);

  if (!success) {
    return new Response("Invalid request", { status: 400 });
  }

  const { username } = data;
  const user = await prisma.user.findFirst({
    where: { name: username },
  });

  if (!user) {
    return new Response("User not found", { status: 404 });
  }

  await setUserToken(user.id);

  // TODO: implement "redirect back" param?
  return redirect("/");
}
