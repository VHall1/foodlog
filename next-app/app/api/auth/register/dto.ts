import z from "zod";

export const registerRequestSchema = z.object({
  username: z.string(),
  // TODO: implement passwords
  // password: z.string(),
});
