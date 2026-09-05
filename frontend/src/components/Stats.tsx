type StatsProps = {
  stats: {
    repositories: number;
    stars: number;
    languages: number;
  };
};

function Stats({ stats }: StatsProps) {
  return (
    <section className="stats">
      <div className="stat-card">
        <div>
          <p className="stat-label">Repositories</p>
          <p className="stat-value">{stats.repositories.toLocaleString()}</p>
        </div>

        <span className="stat-icon">
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M6 3V21"
              stroke="currentColor"
              strokeWidth="1.8"
              strokeLinecap="round"
            />
            <path
              d="M6 7H17L20 10L17 13H6"
              stroke="currentColor"
              strokeWidth="1.8"
              strokeLinejoin="round"
            />
          </svg>
        </span>
      </div>

      <div className="stat-card">
        <div>
          <p className="stat-label">Total stars</p>
          <p className="stat-value">{stats.stars.toLocaleString()}</p>
        </div>

        <span className="stat-icon">
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M12 3L14.8 8.7L21 9.6L16.5 14L17.6 20.2L12 17.3L6.4 20.2L7.5 14L3 9.6L9.2 8.7L12 3Z"
              stroke="currentColor"
              strokeWidth="1.8"
              strokeLinejoin="round"
            />
          </svg>
        </span>
      </div>

      <div className="stat-card">
        <div>
          <p className="stat-label">Languages</p>
          <p className="stat-value">{stats.languages}</p>
        </div>

        <span className="stat-icon">
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M8 7L3 12L8 17"
              stroke="currentColor"
              strokeWidth="1.8"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M16 7L21 12L16 17"
              stroke="currentColor"
              strokeWidth="1.8"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M14 4L10 20"
              stroke="currentColor"
              strokeWidth="1.8"
              strokeLinecap="round"
            />
          </svg>
        </span>
      </div>
    </section>
  );
}

export default Stats;
