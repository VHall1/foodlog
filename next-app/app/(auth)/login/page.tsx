"use client";

import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import { useActionState } from "react";
import { login } from "../actions";
import { loginRequestSchema } from "../schema";

export default function Login() {
  const [lastResult, action] = useActionState(login, null);
  const [form, fields] = useForm({
    lastResult,
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: loginRequestSchema });
    },
    shouldValidate: "onBlur",
    shouldRevalidate: "onInput",
  });

  return (
    <div>
      <h1>Login</h1>
      <form id={form.id} onSubmit={form.onSubmit} action={action} noValidate>
        <label>Username</label>
        <input
          key={fields.username.key}
          name={fields.username.name}
          defaultValue={fields.username.initialValue}
        />
        <div className="text-red-500">{fields.username.errors}</div>
        <button>Login</button>
        <div className="text-red-500">{form.errors}</div>
      </form>
    </div>
  );
}
