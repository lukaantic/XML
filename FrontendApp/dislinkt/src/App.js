import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";


// Components
import HomePage from "./pages/home/Home.js";
import LoginPage from "./pages/auth/login/Login.js"
import RegistrationPage from "./pages/auth/registration/Registration.js"

function App() {

  return (
    <div>
      <Router>
        <Routes>
          <Route exact path="/" element={<LoginPage />} />
          <Route exact path="/home" element={<HomePage />} />
          <Route exact path="/registration" element={<RegistrationPage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
