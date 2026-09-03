type SearchFormProps = {
  username: string
  setUsername: (value: string) => void
  onAnalyze: () => void
}

function SearchForm({
  username,
  setUsername,
  onAnalyze,
}: SearchFormProps) {
  return (
    <div className="search">
      <input
        type="text"
        placeholder="GitHub username"
        value={username}
        onChange={(event) => setUsername(event.target.value)}
      />

      <button type="button" onClick={onAnalyze}>
        Analyze
      </button>
    </div>
  )
}

export default SearchForm