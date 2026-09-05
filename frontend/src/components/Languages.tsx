type LanguagesProps = {
  languages: {
    [lang: string]: number;
  };
  totalRepositories: number;
};

const languageColors: { [language: string]: string } = {
  Go: "#00ADD8",
  TypeScript: "#3178C6",
  JavaScript: "#F7DF1E",
  Python: "#3776AB",
  Java: "#ED8B00",
  C: "#A8B9CC",
  "C++": "#00599C",
  "C#": "#68217A",
  Rust: "#DEA584",
  PHP: "#777BB4",
  Ruby: "#CC342D",
  Swift: "#F05138",
  Kotlin: "#A97BFF",
  HTML: "#E34F26",
  CSS: "#1572B6",
  Shell: "#89E051",
};

function Languages({ languages, totalRepositories }: LanguagesProps) {
  const entries = Object.entries(languages);

  return (
    <section className="languages">
      <div className="languages-header">
        <h2>Language distribution</h2>
        <p>Across all repositories</p>
      </div>

      <div className="language-bar">
        {entries.map(([lang, count]) => {
          const percent = (count / totalRepositories) * 100;
          const color = languageColors[lang] ?? "var(--primary)";

          return (
            <span
              key={lang}
              style={{
                width: `${percent}%`,
                backgroundColor: color,
              }}
            />
          );
        })}
      </div>

      <div className="language-list">
        {entries.map(([lang, count]) => {
          const percent = (count / totalRepositories) * 100;
          const color = languageColors[lang] ?? "var(--primary)";

          return (
            <div className="language" key={lang}>
              <span className="language-name">
                <span
                  className="language-dot"
                  style={{ backgroundColor: color }}
                />
                {lang}
              </span>

              <span className="language-percent">
                {percent.toFixed(1)}%
              </span>
            </div>
          );
        })}
      </div>
    </section>
  );
}

export default Languages;