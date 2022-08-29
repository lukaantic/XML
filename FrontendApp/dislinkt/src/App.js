import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";


// Components
import HomePage from "./pages/home/Home.js";
import LoginPage from "./pages/auth/login/Login.js"

function App () {

  return (
    <div>
      <Router>
        <Routes>
          <Route exact path="/" element={<LoginPage />} />
          <Route exact path="/home" element={<HomePage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
