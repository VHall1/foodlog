export default function Login() {
  return (
    <div>
      <h1>Login</h1>
      <form action="/api/auth/login" method="post">
        <label>
          user:
          <input type="text" name="username" />
        </label>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}
