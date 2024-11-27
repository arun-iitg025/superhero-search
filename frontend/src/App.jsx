import { useState } from "react";
import "./App.css";
import SearchIcon from "./search.svg";
import MovieCard from "./MovieCard";
import SuperheroCard from "./SuperheroCard";

const App = () => {
  const [results, setResults] = useState([]);
  const [searchTerm, setSearchTerm] = useState("");

  const search = async () => {
    try {
      // Call backend API
      const response = await fetch(`http://localhost:8082/search?query=${searchTerm}`);
      const data = await response.json();
      setResults(data);
    } catch (error) {
      console.error("Error fetching data:", error);
      setResults([]);
    }
  };

  return (
    <div className="app">
      <h1>Superhero Search Engine IMDb</h1>
      <div className="search">
        <input
          placeholder="Search for superheroes or movies (e.g., Batman type:movie)"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
        <img src={SearchIcon} alt="search" onClick={search} />
      </div>
      {results?.length > 0 ? (
        <div className="container">
          {results.map((item, index) =>
            item.title ? (
              <MovieCard key={index} item={item} />
            ) : (
              <SuperheroCard key={index} superhero={item} />
            )
          )}
        </div>
      ) : (
        <div className="empty">
          <h2>No results found</h2>
        </div>
      )}
    </div>
  );
};

export default App;
