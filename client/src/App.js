import React, { useState } from "react";
import "./App.css";

const baseURL = "http://localhost:10000/";

function App() {
  const [results, setResults] = useState([]);
  const [searchInput, setSearchInput] = useState("");
  const [previewContent, setPreviewContent] = useState(false);

  function showPreview(content) {
    setPreviewContent(content);
  }
  async function fetchResults(e) {
    e.preventDefault();
    const { data } = await fetch(
      `${baseURL}/search?q=${searchInput}`
      // const { data } = await fetch(
      // `http://localhost:10000/search?q=wordpress shortcode`
    ).then((res) => res.json());
    console.log(data);
    setResults(data);
  }
  return (
    <div className="App">
      <header className="App-header">
        <form onSubmit={fetchResults}>
          <label htmlFor="search">
            <p className="">Search</p>
            <input
              type="text"
              id="search"
              name="search"
              value={searchInput}
              onChange={(e) => setSearchInput(e.target.value)}
            />
          </label>
        </form>
      </header>
      <ul>
        {results.length ? (
          results.map((result) => {
            const { path, content, heading, id } = result;
            return (
              <li key={id}>
                <button onClick={() => showPreview(content)}>
                  {heading.replace("title: ", "")}
                </button>
              </li>
            );
          })
        ) : (
          <p>No Results</p>
        )}
      </ul>
      {previewContent ? (
        <div>
          <button onClick={() => setPreviewContent(null)}>close</button>
          <pre>{previewContent}</pre>
        </div>
      ) : null}
    </div>
  );
}

export default App;
