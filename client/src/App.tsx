import React, { useState } from "react";
import logo from "./logo.svg";
import "./App.css";

function App() {
  const [authenticated, setAuthenticated] = useState(false);

  const getAuthURL = () => {
    return fetch("http://localhost:8080/auth/url")
      .then((response) => response.json())
      .then((data) => {
        window.location.replace(data.url);
      });
  };

  const isAuthenticated = () => {
    fetch("http://localhost:8080/auth/spotify/check")
      .then((response) => response.json())
      .then((data) => {
        setAuthenticated(data.status);
        if (!data.status) {
          getAuthURL();
        }
      });
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <button onClick={() => isAuthenticated()}>Login with Spotify</button>
      </header>
    </div>
  );
}

export default App;
