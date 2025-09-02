import { getUserId } from "@/lib/auth";

export default async function Home() {
  const userId = await getUserId();

  return <div>Hello {userId}</div>;
}
