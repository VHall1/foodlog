import { prisma } from "@/lib/prisma";

export default function Login() {
  async function handleSubmit(formData: FormData) {
    "use server";

    const username = formData.get("user");

    // validate
    if (typeof username !== "string") {
      throw new Error("Invalid input");
    }

    // authenticate
    const user = await prisma.user.findFirst({
      where: { name: username },
    });

    if (!user) {
      throw new Error("User not found");
    }
  }

  return (
    <div>
      <h1>Login</h1>
      <form action={handleSubmit}>
        <label>
          user:
          <input type="text" name="user" />
        </label>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}
