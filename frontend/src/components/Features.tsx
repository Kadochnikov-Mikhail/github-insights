function Features() {
  return (
    <section className="features" id="features">
      <div className="features-heading">
        <p className="section-label">What you get</p>

        <h2>Everything you need to understand your GitHub</h2>

        <p>
          Get a clear picture of your repositories, technology stack,
          and development activity.
        </p>
      </div>

      <div className="features-grid">
        <div className="feature-card">
          <div className="feature-icon">
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
          </div>

          <h3>Repository analysis</h3>

          <p>
            See your total repositories and understand the scale of
            your GitHub activity.
          </p>
        </div>

        <div className="feature-card">
          <div className="feature-icon">
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
          </div>

          <h3>Language statistics</h3>

          <p>
            Discover which programming languages make up your
            development portfolio.
          </p>
        </div>

        <div className="feature-card">
          <div className="feature-icon">
            <svg
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M4 19L9 13L13 16L20 7"
                stroke="currentColor"
                strokeWidth="1.8"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M16 7H20V11"
                stroke="currentColor"
                strokeWidth="1.8"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
          </div>

          <h3>Historical statistics</h3>

          <p>
            Track how your GitHub activity changes over time as your
            project grows.
          </p>
        </div>
      </div>
    </section>
  );
}

export default Features;