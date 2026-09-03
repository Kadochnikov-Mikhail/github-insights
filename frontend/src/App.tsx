import "./App.css";
import Header from "./components/Header";
import SearchForm from "./components/SearchForm";
import Stats from "./components/Stats";
import Languages from "./components/Languages";
import { useState } from "react";

function App() {
  const [username, setUsername] = useState("");

  function handleAnalyze() {
    console.log(username);
  }
  return (
    <div className="app">
      <Header />

      <main className="main">
        <h2 className="main-title">Analyze GitHub Profile</h2>

        <SearchForm
          username={username}
          setUsername={setUsername}
          onAnalyze={handleAnalyze}
        />
      </main>
      <Stats />
      <Languages />
    </div>
  );
}

export default App;
