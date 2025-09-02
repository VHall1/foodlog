import z from "zod";

export const loginRequestSchema = z.object({
  username: z.string(),
  // TODO: implement passwords
  // password: z.string(),
});
