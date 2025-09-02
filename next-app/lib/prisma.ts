import { PrismaClient } from "@/generated/prisma";
import "server-only";

const prisma = new PrismaClient();

export { prisma };
