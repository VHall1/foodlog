import { getUserId } from "@/lib/auth";

type CalorieProgressProps = {
  value: number;
  max?: number;
  size?: number;
};

function CalorieProgress({
  value,
  max = 1800,
  size = 120,
}: CalorieProgressProps) {
  const percent = Math.min(value / max, 1);
  const radius = size / 2 - 8;
  const cx = size / 2;
  const cy = size / 2;
  const strokeWidth = 8;
  const circumference = Math.PI * radius; // top half only
  const dashArray = circumference;
  const dashOffset = dashArray * (1 - percent);

  return (
    <svg width={size} height={size / 2}>
      {/* Background top half circle */}
      <circle
        cx={cx}
        cy={cy}
        r={radius}
        stroke="#eee"
        strokeWidth={strokeWidth}
        fill="none"
        strokeDasharray={dashArray}
        strokeDashoffset={0}
      />
      {/* Progress top half circle */}
      <circle
        cx={cx}
        cy={cy}
        r={radius}
        stroke="#4f8cff"
        strokeWidth={strokeWidth}
        fill="none"
        strokeDasharray={dashArray}
        strokeDashoffset={dashOffset}
      />
      {/* Center text */}
      <text
        x={cx}
        y={cy - 10}
        textAnchor="middle"
        fontSize="18"
        fontWeight="bold"
        fill="#333"
      >
        {value} / {max}
      </text>
    </svg>
  );
}

export default async function Home() {
  const userId = await getUserId();

  return (
    <>
      <div>Hello {userId}</div>
      <CalorieProgress value={800} />
    </>
  );
}
