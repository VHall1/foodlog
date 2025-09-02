"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import Link from "next/link";
import { useActionState } from "react";
import { register } from "../actions";
import { registerRequestSchema } from "../schema";

export default function Register() {
  const [lastResult, action] = useActionState(register, null);
  const [form, fields] = useForm({
    lastResult,
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: registerRequestSchema });
    },
    shouldValidate: "onBlur",
    shouldRevalidate: "onInput",
  });

  return (
    <div className="bg-muted min-h-svh flex items-center justify-center p-6 md:p-10">
      <Card className="w-full max-w-sm">
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Sign Up</CardTitle>
        </CardHeader>
        <CardContent>
          <form
            id={form.id}
            onSubmit={form.onSubmit}
            action={action}
            noValidate
          >
            <div className="grid gap-6">
              <div className="grid gap-6">
                <div className="grid gap-3">
                  <Label htmlFor={fields.username.id}>Username</Label>
                  <Input
                    id={fields.username.id}
                    key={fields.username.key}
                    name={fields.username.name}
                    defaultValue={fields.username.initialValue}
                  />
                  {fields.username.errors && (
                    <div className="text-sm text-destructive">
                      {fields.username.errors}
                    </div>
                  )}
                </div>

                <Button type="submit" className="w-full">
                  Register
                </Button>

                {form.errors && (
                  <div className="text-sm text-destructive">{form.errors}</div>
                )}
              </div>

              <div className="text-center text-sm">
                Already have an account?{" "}
                <Link href="/login" className="underline underline-offset-4">
                  Sign in
                </Link>
              </div>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
