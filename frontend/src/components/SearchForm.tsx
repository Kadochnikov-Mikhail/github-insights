type SearchFormProps = {
  username: string;
  setUsername: (value: string) => void;
  onAnalyze: () => void;
  loading: boolean;
};

function SearchForm({
  username,
  setUsername,
  onAnalyze,
  loading,
}: SearchFormProps) {
  return (
    <div className="search-container">
      <div className="search">
        <div className="search-input-wrapper">
          <span className="search-icon">
            <svg
              viewBox="0 0 24 24"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <circle
                cx="11"
                cy="11"
                r="6.5"
                stroke="currentColor"
                strokeWidth="1.8"
              />
              <path
                d="M16 16L21 21"
                stroke="currentColor"
                strokeWidth="1.8"
                strokeLinecap="round"
              />
            </svg>
          </span>

          <input
            type="text"
            placeholder="Enter GitHub username"
            value={username}
            onChange={(event) => setUsername(event.target.value)}
            onKeyDown={(event) => {
              if (event.key === 'Enter') {
                onAnalyze()
              }
            }}
          />
        </div>

        <button
          type="button"
          onClick={onAnalyze}
          disabled={loading || !username.trim()}
        >
          {loading ? (
            <>
              <span className="spinner" />
              Analyzing...
            </>
          ) : (
            <>
              Analyze GitHub
              <span>→</span>
            </>
          )}
        </button>
      </div>

      <p className="search-hint">
        Try{' '}
        <button type="button" onClick={() => setUsername('octocat')}>
          octocat
        </button>{' '}
        to see a sample analysis
      </p>
    </div>
  )
}

export default SearchForm