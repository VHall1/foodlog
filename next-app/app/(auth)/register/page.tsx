export default function Register() {
  return (
    <div>
      <h1>Register</h1>
      <form action="/api/auth/register" method="post">
        <label>
          user:
          <input type="text" name="username" />
        </label>
        <button type="submit">Register</button>
      </form>
    </div>
  );
}
