export default function AuthBackground() {
  return (
    <div className="svg-background">
      <svg viewBox="0 0 1920 1080" preserveAspectRatio="xMidYMid slice" xmlns="http://www.w3.org/2000/svg">
        <defs>
          <radialGradient id="gradient1" cx="20%" cy="30%">
            <stop offset="0%" stopColor="#4a4a4a" stopOpacity="1">
              <animate attributeName="stop-color" values="#4a4a4a;#5a5a5a;#4a4a4a" dur="8s" repeatCount="indefinite" />
            </stop>
            <stop offset="100%" stopColor="#1a1a1a" stopOpacity="1" />
          </radialGradient>
          <radialGradient id="gradient2" cx="80%" cy="70%">
            <stop offset="0%" stopColor="#383838" stopOpacity="1">
              <animate attributeName="stop-color" values="#383838;#484848;#383838" dur="10s" repeatCount="indefinite" />
            </stop>
            <stop offset="100%" stopColor="#151515" stopOpacity="1" />
          </radialGradient>
          <filter id="glow">
            <feGaussianBlur stdDeviation="40" result="coloredBlur" />
            <feMerge>
              <feMergeNode in="coloredBlur" />
              <feMergeNode in="SourceGraphic" />
            </feMerge>
          </filter>
        </defs>

        <rect width="100%" height="100%" fill="#1a1a1a" />
        <rect width="100%" height="100%" fill="url(#gradient1)" />
        <rect width="100%" height="100%" fill="url(#gradient2)" opacity="0.6" />

        <circle cx="30%" cy="40%" r="300" fill="#4a4a4a" opacity="0.4" filter="url(#glow)">
          <animate attributeName="cx" values="30%;35%;30%" dur="15s" repeatCount="indefinite" />
          <animate attributeName="cy" values="40%;35%;40%" dur="12s" repeatCount="indefinite" />
          <animate attributeName="r" values="300;350;300" dur="10s" repeatCount="indefinite" />
        </circle>

        <circle cx="70%" cy="60%" r="250" fill="#3a3a3a" opacity="0.5" filter="url(#glow)">
          <animate attributeName="cx" values="70%;65%;70%" dur="18s" repeatCount="indefinite" />
          <animate attributeName="cy" values="60%;65%;60%" dur="14s" repeatCount="indefinite" />
          <animate attributeName="r" values="250;300;250" dur="12s" repeatCount="indefinite" />
        </circle>

        <circle cx="50%" cy="50%" r="200" fill="#5a5a5a" opacity="0.3" filter="url(#glow)">
          <animate attributeName="r" values="200;280;200" dur="16s" repeatCount="indefinite" />
          <animate attributeName="opacity" values="0.2;0.3;0.2" dur="8s" repeatCount="indefinite" />
        </circle>
      </svg>
    </div>
  );
}