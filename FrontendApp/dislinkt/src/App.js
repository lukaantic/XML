import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";


// Components
import HomePage from "./pages/home/Home.js";
import LoginPage from "./pages/auth/login/Login.js"
import RegistrationPage from "./pages/auth/registration/Registration.js"
import ClickedUserProfile from "./pages/profile/ClickedUserProfile.js";
import UserProfile from "./pages/profile/UserProfile.js";

function App() {

  return (
    <div>
      <Router>
        <Routes>
          <Route exact path="/" element={<LoginPage />} />
          <Route exact path="/home" element={<HomePage />} />
          <Route exact path="/registration" element={<RegistrationPage />} />
          <Route exact path="/account:id" element={<ClickedUserProfile />} />
          <Route exact path="/profile" element={<UserProfile />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
