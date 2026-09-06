import { useState } from "react";
import "./App.css";

import Header from "./components/Header";
import SearchForm from "./components/SearchForm";
import ProfileHeader from "./components/ProfileHeader";
import Stats from "./components/Stats";
import Languages from "./components/Languages";
import Features from "./components/Features";
import Footer from "./components/Footer";

type Insights = {
  username: string;
  repositories: number;
  total_stars: number;
  languages: {
    [language: string]: number;
  };
  created_at: string;
};

function App() {
  const [username, setUsername] = useState("");
  const [insights, setInsights] = useState<Insights | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function handleAnalyze() {
    if (!username.trim()) {
      setError("username is required");
      return;
    }

    const trimmedUsername = username.trim();

    setError(null);
    setInsights(null);
    setLoading(true);

    try {
      const response = await fetch(
        `/api/github/insights?user=${encodeURIComponent(trimmedUsername)}`,
      );

      if (!response.ok) {
        setError("Failed to fetch GitHub data");
        return;
      }

      const data = await response.json();

      setInsights(data);
    } catch {
      setError("Something went wrong");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="app">
      <Header />

      <section className="hero">
        <div className="hero-badge">
          <span></span>
          Live repository analytics
        </div>

        <h2 className="hero-title">
          Understand your GitHub activity <span>at a glance</span>
        </h2>

        <p className="hero-description">
          Enter a GitHub username to explore repositories, stars, languages, and
          the momentum behind your work.
        </p>

        {error && <p className="hero-error">{error}</p>}

        <SearchForm
          username={username}
          setUsername={setUsername}
          onAnalyze={handleAnalyze}
          loading={loading}
        />
      </section>

      {insights && (
        <div className="results">
          <ProfileHeader username={insights.username} />

          <Stats
            stats={{
              repositories: insights.repositories,
              stars: insights.total_stars,
              languages: Object.keys(insights.languages).length,
            }}
          />

          <Languages
            languages={insights.languages}
            totalRepositories={insights.repositories}
          />
        </div>
      )}

      <Features />
      <Footer />
    </div>
  );
}

export default App;
